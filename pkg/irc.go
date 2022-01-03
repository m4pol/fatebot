package pkg

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

type logFormation struct {
	user, nick string
}

func IRC_Conn(server string) net.Conn {
	conn, err := net.Dial("tcp", server)
	for err != nil {
		time.Sleep(10 * time.Second)
		continue
	}
	return conn
}

func IRC_Find(read, msg string) bool {
	return strings.Contains(read, msg)
}

func IRC_Recv(command string, arg int) string {
	return strings.Split(command, " ")[arg]
}

func (irc *IRC) IRC_Send(data string) {
	fmt.Fprintf(irc.Report, "%s\r\n", data)
}

func (irc *IRC) IRC_Report(msg string) {
	irc.IRC_Send("PRIVMSG " + irc.Channel + " " + " :" + msg)
}

func (irc *IRC) IRC_Login() {
	rand.Seed(time.Now().UnixNano())
	alphabet := 'A' + rune(rand.Intn(26))
	sAlphabet := string(alphabet)
	botID := rand.Intn(1000000)

	//Change bot profile here.
	botformation := logFormation{
		user: fmt.Sprint("USER [FATE][", sAlphabet, "][", botID, "]", " 8 * :bot"),
		nick: fmt.Sprint("NICK [FATE][", sAlphabet, "][", botID, "]"),
	}

	irc.IRC_Send(botformation.user)
	irc.IRC_Send(botformation.nick)
}
