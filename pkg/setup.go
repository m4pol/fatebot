package pkg

import "net"

/*
	This structure is For report about handler and process
	through the irc server.
*/
type IRC struct {
	Report           net.Conn
	Channel, ChanKey string
}
