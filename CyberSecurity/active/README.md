# active
In following the subject of cybersecurity, the goal of this project is to create a program that will be able to build a simple port scanner, wich will tell you if the port is open or closed. 

I've also created a simple go server that can be used to test the port scanner.

### Basics
What is a port?
- A port is a communication endpoint that identifies a specific process or a type of service. Ports are identified for each protocol and address combination by 16-bit unsigned numbers, commonly known as the port number. The port number is included in the data sent in the header of the packet.

What is a port scanner?
- A port scanner is a software application or online tool designed to probe a network host for open ports. This is often used by administrators to verify security policies of their networks and by attackers to identify running services on a host with the view to compromise it.

Why is port scanning important in pentesting?
- Port scanning is an essential technique used by pentesters to discover open ports and services running on a target machine. This information is crucial for identifying potential vulnerabilities and misconfigurations that could be exploited by attackers.

How does my port scanner work?
- Please refer to the code for a detailed explanation of how the port scanner works. Usage is down below.

### Usage

You can either type in the commands in the terminal or utilise the makefile to run the program.

```bash
$>python3 tinyscanner.py [-h] -p  [-u] [-t] host

Tiny Port Scanner

positional arguments:
  host           Host to scan

options:
  -h, --help     show this help message and exit
  -p , --ports   Specific or range of ports to scan
  -u, --udp      UDP scan
  -t, --tcp      TCP scan
$>  python3 tinyscanner.py -u 127.0.0.1 -p 8081
UDP Port 8081 is open
$> python3 tinyscanner.py -t 127.0.0.1 -p 8080
TCP Port 8080 is open
$> python3 tinyscanner.py -t 10.53.224.5 -p 80-83
TCP Port 80 is closed
TCP Port 81 is closed
TCP Port 82 is closed
TCP Port 83 is closed
```

**Using the Makefile commands:**

_I recommend using two terminals, one for the server and the other for the port scanner._

```make server``` to create a simple go server

```make u``` to run the UDP scan

```make t``` to run the TCP scan

```make all``` to run both the UDP and TCP scans



### Technologies used
- Go
- Python

[Task Description](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/active)

[Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/active/audit)