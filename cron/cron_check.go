package cron

import (
	"github.com/PGDream/monitorprocess/monitor"
	"log"
	"os/exec"
	"fmt"
	"strconv"
	"time"
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
处理process
 */
func handleProcess(processConfigInfo *monitor.ProcessConfigInfo) {
	pidProcessMap := catchPIDProcess(processConfigInfo)
	if pidProcessMap != nil {
		for _, pidProcess := range (pidProcessMap) {
			if pidProcess.PID == 0 {
				info := exeCmd(fmt.Sprintf(startProcess, pidProcess.StartUser, pidProcess.StartCmd))
				log.Println(pidProcess.ProcessName + "----start info----" + info)
			}
		}
	}
}

/**
 catch PID，返回值为uint，默认值为0
 */
func catchPIDProcess(processConfigInfo *monitor.ProcessConfigInfo) (processInfo map[string]*monitor.ProcessMetadata) {
	if len(processConfigInfo.Process) <= 0 {
		return nil
	}
	pids := make(map[string]*monitor.ProcessMetadata, len(processConfigInfo.Process))
	for _, process := range (processConfigInfo.Process) {
		process.PID = getPID(&process)
		pids[process.ProcessName] = &process
	}
	return pids
}

/**
获取pid
 */
func getPID(metadata *monitor.ProcessMetadata) int {
	pid := exeCmd(fmt.Sprintf(checkPIDCmd, metadata.ProcessPort, metadata.StartUser))
	if pid == "" {
		return 0
	}
	num, _ := strconv.Atoi(pid)
	return num
}

/**
exec cmd
 */
func exeCmd(cmdInfo string) string {
	log.Println("command info:", cmdInfo)
	cmd := exec.Command("/bin/bash", "-c", cmdInfo)

	out, err := cmd.Output()
	if err != nil {
		log.Println("command exec failer:", err)
		return ""
	}

	return string(out)
}
