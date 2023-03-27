// Auther: winnielovegood@gmail.com
// Date: 20230323
// 用途: 采集icdncacher server 日志中的 n_recv, n_succ, n_done, n_hit
// use for: Collect icdncacher  logs n_recv, n_succ, n_done, n_hit

package collector

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"os/exec"
	"regexp"
	"strconv"
)

// errLogSubsystem 定义系统名称
const (
	icdncacherLogSubsystem = "icdncacherlog"
)

// icdncacherLogCollector 定义一个结构体
type icdncacherLogCollector struct {
	logger log.Logger
}

// 写一个 init方法调用 registerCollector 注册自己
func init() {
	registerCollector(icdncacherLogSubsystem, defaultEnabled, NewicdncacherLogCollector)
}

// NewicdncacherLogCollector 工厂函数返回 icdnLogCollector
func NewicdncacherLogCollector(logger log.Logger) (Collector, error) {
	return &icdncacherLogCollector{logger}, nil
}

// Update 完成 Update 方法,更新指标数据到 chan 中
func (c *icdncacherLogCollector) Update(ch chan<- prometheus.Metric) error {
	// 定义一个 Gauge 类型的 变量
	var metricType prometheus.ValueType
	metricType = prometheus.GaugeValue
	// 获取服务器上错误日志
	level.Debug(c.logger).Log("msg", "start Get icdncacherLog")
	output, err := GetLog()
	if err != nil {
		level.Debug(c.logger).Log("msg", "Set icdncacheLog", "err", err)
		return err
	}
	/*
		map[broken:0 done:15773575 fail:0 giveup:0 hit:3732595 recv:15773575 succ:15773575 timeout:0 unhealthy:0]
	*/
	// 按行切割错误日志
	level.Debug(c.logger).Log("msg", "start cut icdncacherLog")
	for name, value := range output {

		// 输出到日志
		level.Debug(c.logger).Log("msg", "Set icdncacherLog", "name", name, "value", value)

		// 最关键一步: 写入 ch 中
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, icdncacherLogSubsystem, name),
				fmt.Sprintf("from icdncacher server.log  n_%s", name),
				nil, nil,
			),
			metricType, float64(value),
		)
	}
	return nil
}

// GetLog  调用 grep 命令从 icdncacher 中过滤出  n_recv, n_succ, n_done, n_hit
func GetLog() (map[string]int, error) {
	// 在 os 上执行命令
	errLogCmd := `ls -t /usr/local/release/icdncacher/log/server*.log | head -1 | sort -n | xargs grep 'profile' | tail -1`
	cmd := exec.Command("sh", "-c", errLogCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	// 将命令结果转成 string
	logLine := string(output)

	// 使用 正则 功能 过滤 n_recv, n_succ, n_done, n_hit

	/*
		过滤后的效果：去掉了日志中的 n_
		map[broken:0 done:15773575 fail:0 giveup:0 hit:3732595 recv:15773575 succ:15773575 timeout:0 unhealthy:0]
	*/

	re := regexp.MustCompile(`n_(\w+):\s*(\d+)`)
	matches := re.FindAllStringSubmatch(logLine, -1)

	// 定义一个 map 用来存过滤后的值
	values := map[string]int{}

	for _, match := range matches {
		key := match[1]
		value, _ := strconv.Atoi(match[2])
		values[key] = value
	}

	return values, nil

}
