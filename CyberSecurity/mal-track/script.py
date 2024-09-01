import os
import subprocess
import shutil
import winreg
import re
import logging

# Setup logging
logging.basicConfig(filename='script.log', level=logging.INFO, 
                    format='%(asctime)s:%(levelname)s:%(message)s')

def stop_process(process_name):
    try:
        subprocess.run(['taskkill', '/F', '/IM', process_name], check=True)
        logging.info(f'Stopped Process {process_name}')
    except subprocess.CalledProcessError as e:
        logging.error(f"Error occurred while stopping process '{process_name}': {e}")

def delete_file(filename):
    for root, dirs, files in os.walk('C:\\'):
        if filename in files:
            file_path = os.path.join(root, filename)
            try:
                os.remove(file_path)
                logging.info(f'Deleted file: {file_path}')
            except OSError as e:
                logging.error(f"Error occurred while deleting file '{file_path}': {e}")

def delete_registry_entry(key_path, entry_name):
    try:
        key = winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE, key_path, 0, winreg.KEY_ALL_ACCESS)
        try:
            winreg.DeleteValue(key, entry_name)
            logging.info(f'Registry entry {entry_name} deleted from {key_path}')
        except FileNotFoundError:
            logging.warning(f"Registry entry '{entry_name}' not found in '{key_path}'.")
        winreg.CloseKey(key)
    except OSError as e:
        logging.error(f"Error occurred while deleting registry entry '{entry_name}' from '{key_path}': {e}")

def extract_ip_address(binary_path):
    ip_addresses = []

    if not os.path.exists(binary_path):
        logging.error(f"File not found: {binary_path}")
        return ip_addresses

    try:
        with open(binary_path, 'rb') as file:
            binary_data = file.read()
        logging.info(f"Successfully read binary file: {binary_path}")
    except IOError as e:
        logging.error(f"Error reading file '{binary_path}': {e}")
        return ip_addresses

    ip_regex = r'\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b'
    matches = re.findall(ip_regex, str(binary_data))

    for match in matches:
        if match not in ip_addresses:
            ip_addresses.append(match)

    logging.info(f"Extracted IP addresses: {ip_addresses}")
    return ip_addresses

if __name__ == "__main__":
    stop_process("maltrack.exe")
    delete_registry_entry(r"SOFTWARE\Microsoft\Windows\CurrentVersion\Run", "mal-track")
    delete_registry_entry(r"SOFTWARE\Microsoft\Windows\CurrentVersion\Run", "maltrack.exe")
        
    ip_addresses = extract_ip_address('C:\\Users\\user\\Documents\\maltrack\\maltrack.exe')
    if ip_addresses:
        print("Possible attacker IP addresses found:")
        for ip in ip_addresses:
            print(ip)
    else:
        print("No IP addresses found.")

    delete_file("maltrack.exe")
    logging.info("System scan completed.")
    
    user_input = input("Press Enter to exit...")
    if user_input == "":
        logging.info("Script ended by user.")
