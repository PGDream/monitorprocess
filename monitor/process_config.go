package monitor

type ProcessConfigInfo struct {
	Process string `yaml: process`
	Metadata []struct {
		ProcessName string `yaml: processName`
		ProcessPort string `yaml: processPort`
		StartCmd    string `yaml: startCmd`
		StartUser   string `yaml: startUser`
	}
}
