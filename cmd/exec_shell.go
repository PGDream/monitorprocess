package cmd

import (
	"log"
	"os/exec"
)

func init() {
	log.SetPrefix("【exec-shell】")
}

func ExecShell(cmdInfo string) ([]byte, error) {
	if cmdInfo == "" {
		return nil, nil
	}
	cmdObject := exec.Command("/bin/bash", "-c", cmdInfo)
	outInfo, outErr := cmdObject.CombinedOutput()
	/**err := cmdObject.Run()
	if err != nil {
		log.Fatalln("exec shell fialed, expection info: ", err)
	}
	*/
	if outErr != nil {
		log.Println("get exec shell out info failed, expection info: ", outErr)
		return nil, outErr
	}
	return outInfo, nil
}

func ExecCmd(cmdInfo string) string {
	if cmdInfo == "" {
		return ""
	}
	cmdObject := exec.Command("/bin/bash", "-c", cmdInfo)
	outInfo, outErr := cmdObject.Output()
	if outErr != nil {
		log.Println("get exec shell out info failed, expection info: ", outErr)
		return ""
	}
	return string(outInfo)
}
