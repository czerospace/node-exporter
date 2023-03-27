// Auther: winnielovegood@gmail.com
// Date: 20230323
// 用途: 采集系统错误日志
// use for: Collect system error logs

package collector

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"os/exec"
	"strconv"
	"strings"
)

// errLogSubsystem 定义系统名称
const (
	errLogSubsystem = "errlog"
)

// errLogCollector 定义一个结构体
type errLogCollector struct {
	logger log.Logger
}

// 写一个 init方法调用 registerCollector 注册自己
func init() {
	registerCollector(errLogSubsystem, defaultEnabled, NewErrLogCollector)
}

// NewErrLogCollector 工厂函数返回 errLogCollector
func NewErrLogCollector(logger log.Logger) (Collector, error) {
	return &errLogCollector{logger}, nil
}

// Update 完成 Update 方法,更新指标数据到 chan 中
func (c *errLogCollector) Update(ch chan<- prometheus.Metric) error {
	// 定义一个 Gauge 类型的 变量
	var metricType prometheus.ValueType
	metricType = prometheus.GaugeValue
	// 获取服务器上错误日志
	output := errLogGrep()
	/*
		kubelet: 16822
		containerd: 9350
		kernel: 5
		grafana-server: 10
	*/
	// 按行切割错误日志
	for _, line := range strings.Split(output, "\n") {
		l := strings.Split(line, ":")
		if len(l) != 2 {
			continue
		}

		name := strings.TrimSpace(l[0])
		// 按照普罗米修斯的规范 name 中不能出现 -, 使用 _ 代替 -
		// eg: grafana-server --->>> grafana_server
		name = strings.Replace(name, "-", "_", -1)

		value := strings.TrimSpace(l[1])
		v, _ := strconv.Atoi(value)

		// 输出到日志
		level.Debug(c.logger).Log("msg", "Set errLog", "name", name, "value", value)

		// 最关键一步: 写入 ch 中
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, errLogSubsystem, name),
				fmt.Sprintf("/var/log/message err log %s", name),
				nil, nil,
			),
			metricType, float64(v),
		)
	}
	return nil
}

// 调用 grep 命令从 /var/log/messages 中过滤出 error 日志
func errLogGrep() string {
	errLogCmd := `grep -i error /var/log/messages |awk '{a[$5]++}END{for(i in a) print i,a[i]}'`
	cmd := exec.Command("sh", "-c", errLogCmd)
	output, _ := cmd.CombinedOutput()
	return string(output)
}
