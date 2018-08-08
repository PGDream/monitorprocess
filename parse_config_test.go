package main

import (
	"testing"
	"log"
	"runtime"
)

func TestNewParseConfig(t *testing.T) {
	fileName := "/Users/pc001/go/src/github.com/PGDream/monitorprocess/config.yaml"
	parseConfig := ParseConfig{
		FileName: fileName,
	}
	processInfo := parseConfig.Parse()
	log.Println(processInfo.Tostring())

	log.Println(runtime.GOOS)
}
