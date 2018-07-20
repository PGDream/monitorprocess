package cmd

import (
	"testing"

	"log"
)

func TestExecShell(t *testing.T) {
	contents := ExecShell("ls ~")
	log.Println(string(contents))

}
