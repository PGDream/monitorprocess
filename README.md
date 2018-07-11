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