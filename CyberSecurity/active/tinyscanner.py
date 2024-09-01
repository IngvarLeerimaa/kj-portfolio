import socket
import argparse
from concurrent.futures import ThreadPoolExecutor

def scan_tcp(ip, port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
        sock.settimeout(1)
        result = sock.connect_ex((ip, port))
        if result == 0:
            return port, 'open'
        else:
            return port, 'closed'

def scan_udp(ip, port):
    with socket.socket(socket.AF_INET, socket.SOCK_DGRAM) as sock:
        sock.settimeout(1)
        try:
            sock.sendto(b'', (ip, port))
            sock.recvfrom(1024)
            return port, 'open'
        except socket.timeout:
            return port, 'closed | timeout'
        except Exception:
            return port, 'closed'

def port_scan(ip, ports, scan_type):
    results = []
    scan_func = scan_udp if scan_type == 'udp' else scan_tcp
    
    with ThreadPoolExecutor(max_workers=100) as executor:
        futures = [executor.submit(scan_func, ip, port) for port in ports]
        for future in futures:
            results.append(future.result())
    
    return results

def main():
    parser = argparse.ArgumentParser(description='Tiny Port Scanner')
    parser.add_argument('host', type=str, help='Host to scan')
    parser.add_argument('-p', '--ports', type=str, required=True, metavar='', help='Specific or range of ports to scan')
    parser.add_argument('-u', '--udp', action='store_true', help='UDP scan')
    parser.add_argument('-t', '--tcp', action='store_true', help='TCP scan')
    
    args = parser.parse_args()
    
    if not args.tcp and not args.udp:
        print('Please specify at least one of -u or -t for UDP or TCP scan')
        return
    
    ports = []
    if '-' in args.ports:
        start_port, end_port = map(int, args.ports.split('-'))
        ports = range(start_port, end_port + 1)
    else:
        ports = [int(args.ports)]
    
    if args.udp:
        results = port_scan(args.host, ports, 'udp')
        for port, status in results:
            print(f"UDP Port {port} is {status}")
    
    if args.tcp:
        results = port_scan(args.host, ports, 'tcp')
        for port, status in results:
            print(f"TCP Port {port} is {status}")

if __name__ == "__main__":
    main()
