package cron

import (
	"github.com/PGDream/monitorprocess/monitor"
	"log"
	"fmt"
	"strconv"
	"time"
	"github.com/PGDream/monitorprocess/cmd"
)

const checkPIDCmd = "lsof -i:%s|grep %s|awk -F \" \" '{printf $2}'"
const startProcess = "sudo -u %s %s"
const timeInterval = 5 // second

/**
定时检查
 */
func cron(processConfigInfo *monitor.ProcessConfigInfo) {
	timer := time.NewTicker(timeInterval * time.Second)
	for {
		select {
		case <-timer.C:
			handleProcess(processConfigInfo)
		}
	}
}

/**
 catch PID，返回值为uint，默认值为0
 */
func handleProcess(processConfigInfo *monitor.ProcessConfigInfo) {
	if processConfigInfo == nil || len(processConfigInfo.Process) <= 0 {
		log.Println("ProcessConfiginfo object is nil")
		return
	}
	for _, process := range (processConfigInfo.Process) {
		process.PID = getPID(&process)

		func(process *monitor.ProcessMetadata) {
			if process.PID == 0 {
				info := cmd.ExecCmd(fmt.Sprintf(startProcess, process.StartUser, process.StartCmd))
				log.Println(process.ProcessName + "----start info----" + info)

			}
		}(&process)
	}
}

/**
获取pid
 */
func getPID(metadata *monitor.ProcessMetadata) int {
	pid := cmd.ExecCmd(fmt.Sprintf(checkPIDCmd, metadata.ProcessPort, metadata.StartUser))
	if pid == "" {
		return 0
	}
	num, _ := strconv.Atoi(string(pid))
	return num
}
