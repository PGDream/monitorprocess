package cmd

import (
	"log"
	"os/exec"
)

func init() {
	log.SetPrefix("【exec-shell】")
}

func ExecShell(cmdInfo string) []byte {
	if cmdInfo == "" {
		return nil
	}
	cmdObject := exec.Command("/bin/bash", "-c", cmdInfo)
	outInfo, outErr := cmdObject.CombinedOutput()
	/**err := cmdObject.Run()
	if err != nil {
		log.Fatalln("exec shell fialed, expection info: ", err)
	}
	*/
	if outErr != nil {
		log.Fatalln("get exec shell out info failed, expection info: ", outErr)
	}
	return outInfo
}
