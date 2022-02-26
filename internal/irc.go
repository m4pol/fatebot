package lib

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

func Recv(str string, arg int) string {
	for {
		recv := strings.Split(str, " ")
		if len(recv) == arg {
			continue
		}
		return recv[arg]
	}
}

func Find(read, str string) bool {
	return strings.Contains(read, str)
}

func name(alphabet rune) string {
	return string(alphabet + rune(rand.Intn(26)))
}

func (b *Bot) Send(str string) {
	fmt.Fprintf(b.IRC, "%s\r\n", str)
}

func (b *Bot) Report(str string) {
	b.Send("PRIVMSG " + b.Channel + "  :" + str)
}

func (b *Bot) Login() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(10000000)
	group := name('A')

	user := fmt.Sprint("USER [BOT][", group, "][", id, "]", " 8 * :bot")
	nick := fmt.Sprint("NICK [BOT][", group, "][", id, "]")

	b.Send(user)
	b.Send(nick)
}
