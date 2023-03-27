# single install for consul

## preparation

```shell
# download
wget https://releases.hashicorp.com/consul/1.15.1/consul_1.15.1_linux_amd64.zip

# unzip
unzip consul_1.15.1_linux_amd64.zip -d /r2/monitor/consul/
```

## configuration file

```shell
# config
cat <<EOF > /r2/monitor/consul/single_server.json
{
    "datacenter": "dc1",
    "node_name": "consul-svr-01",
    "server": true,
    "bootstrap_expect": 1,
    "data_dir": "/r2/monitor/consul/data",
    "log_level": "INFO",
    "log_file": "/r2/monitor/consul/logs/",
    "ui": true,
    "bind_addr": "your ip",
    "client_addr": "0.0.0.0",
    "retry_interval": "10s",
    "raft_protocol": 3,
    "enable_debug": false,
    "rejoin_after_leave": true,
    "enable_syslog": false
}
EOF
```

## consul.service

```shell
cat <<EOF > /etc/systemd/system/consul.service
[Unit]
Description=consul server
Wants=network-online.target
After=network-online.target

[Service]
ExecStart=/r2/monitor/consul/consul agent  -config-file=/r2/monitor/consul/single_server.json
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=consul
[Install]
WantedBy=default.target
EOF

```

## start consul

```shell
systemctl daemon-reload

systemctl start consul   

systemctl status consul
```

## test and verify

```shell
http://yourip:8500/
```

