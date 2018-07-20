package main

import (
	"io/ioutil"
	"os"
	"log"

	"github.com/PGDream/monitorprocess/monitor"
	"gopkg.in/yaml.v2"
)

type ParseConfig struct {
	FileName string
}

func init() {
	log.SetPrefix("【parse-config】")
}

func (parseConfig *ParseConfig) NewParseConfig(fileName string) *ParseConfig {

	return &ParseConfig{
		FileName: fileName,
	}
}

func (parseConfig *ParseConfig) Parse() *monitor.ProcessConfigInfo {
	contents := readConfig(parseConfig)
	if contents == nil {
		log.Println("config content is nil")
		return nil
	}
	processConfigInfo := new(monitor.ProcessConfigInfo)
	err := yaml.Unmarshal(contents, &processConfigInfo)
	if err != nil {
		log.Fatalln("parse config to object failed, expection info:", err)
		return nil
	}
	return processConfigInfo
}

/**
  return can`t "" charset
 */
func readConfig(parseConfig *ParseConfig) []byte {
	_, err := os.Open(parseConfig.FileName)
	if err != nil && os.IsNotExist(err) {
		log.Fatalln("read config failed, file not exists or get config for fd failed, expection info:", err)
		return nil
	}
	contents, err := ioutil.ReadFile(parseConfig.FileName)
	if err != nil {
		log.Fatalln("read config file content failed, expection info:", err)
	}
	return contents
}
