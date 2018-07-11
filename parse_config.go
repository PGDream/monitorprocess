package main

import (
	"io/ioutil"
	"os"
	"github.com/prometheus/common/log"
)

type ParseConfig struct {
	FileName string
}

func NewParseConfig(fileName string) *ParseConfig {

	return &ParseConfig{
		FileName: fileName,
	}
}

func (parseConfig *ParseConfig) Parse() interface{} {

}

/**
  return can`t "" charset
 */
func readConfig(parseConfig *ParseConfig) string {
	_, err := os.Open(parseConfig.FileName)
	if err != nil && os.IsNotExist(err) {
		log.Fatalln("read config failed, file not exists or get config for fd failed, expection info:", err)
		return ""
	}
	contents, err := ioutil.ReadFile(parseConfig.FileName)
	if err != nil {
		log.Fatalln("read config file content failed, expection info:", err)
	}
	return string(contents)
}
