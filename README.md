# monitorprocess
`monitorprocess` 暂时只支持linux系统,`monitorprocess` 实现原理比较简单，通过配置的进程端口，使用系统`lsof`命令获取`PID`
是否大于*0*如果为空或者小于*0*就根据配置的启动命令在用户下启动，再进行判断是否成功
## 安装要求
- 系统软件必须安装`lsof`
```
centos:
yum install lsof

ubuntu:
apt-get update
apt install lsof
```

## 功能完成进度

|序号|功能描述|完成进度|
|-|-|-|
|1|执行系统命令模块|已完成|
|2|解析yaml配置文件|已完成|
|3|守护进程功能|未开发|
|4|定时器,定时检查进程|未开发|
|5|日志输出|未开发|
