package monitor

import (
	"encoding/json"
	"log"
)

type ProcessConfigInfo struct {
	Process []ProcessMetadata
}

type ProcessMetadata struct {
	ProcessName string `yaml:"process_name"`
	ProcessPort string `yaml:"process_port"`
	StartCmd    string `yaml:"start_cmd"`
	StartUser   string `startUser:"start_user"`
}

func (processConfigInfo *ProcessConfigInfo) Tostring() string {
	info, err := json.Marshal(processConfigInfo)
	if err != nil {
		log.Println("to string failed, info :", err)
		return ""
	}
	return string(info)
}
