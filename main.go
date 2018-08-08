package main

import (
	"errors"
	"os"
	"log"
	"syscall"

	"github.com/urfave/cli"
	"time"
	"runtime"
	"strings"
)

func init() {
	log.SetPrefix("【main】")
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "请输入配置文件目录",
		},
	}

	app.Action = func(c *cli.Context) error {
		configFile := c.String("c")
		if configFile == "" {
			return errors.New("配置文件不能为空")
		}
		_, err := os.Open(configFile)
		if err != nil && os.IsNotExist(err) {
			return errors.New("配置文件不存在")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("获取参数异常,信息: ", err)
	}
	// create daemon process
	daemon(0, 1)
	for {
		time.Sleep(time.Second * 1)
	}
}

func daemon(nochdir, noclose int) int {
	var ret, ret2 uintptr
	var err syscall.Errno
	// platform macos linux unix
	platform := []string{"darwin", "linux", "unix"}
	runPlatform := runtime.GOOS
	darwin := false
	for _, value := range (platform) {
		if strings.EqualFold(runPlatform, value) {
			darwin = true
		}
	}
	// already a daemon
	if syscall.Getppid() == 1 {
		return 0
	}
	// fork off the parent process
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return -1
	}
	// failure
	if ret2 < 0 {
		os.Exit(-1)
	}
	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}
	// if we got a good PID, then we call exit the parent process.
	if ret > 0 {
		os.Exit(0)
	}
	/* Change the file mode mask */
	_ = syscall.Umask(0)

	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		return -1
	}
	if nochdir == 0 {
		os.Chdir("/")
	}
	if noclose == 0 {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}
	return 0
}
