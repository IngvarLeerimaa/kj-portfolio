# mal-track

### Project Overview:

The goal of this project is to understand the basic operation of a computer virus in a Windows environment and explore simple methods to eradicate them.

During the audit official Windows virtual machine is required.
## Setup

1. **Install VirtualBox:** Download VirtualBox and install it on your host machine.

2. **Download Windows 10 ISO:** Obtain the Windows 10 ISO from Microsoft.

3. **Create a Virtual Machine:** Set up a new virtual machine using the downloaded Windows 10 ISO.

*Note: You may need to disable the Floppy drive and remove the floppy controller in the VM settings to successfully install Windows.*

4. **Install Python:** Once Windows 10 is installed on your VM, install Python.

*Note: This can be ommited if you convert the script to an executable file and transfer it to your VM.*

5. **Download Malware:** Boot into the Windows 10 VM and download the malware sample from [here.](https://01.kood.tech/git/root/public/src/branch/master/subjects/cybersecurity/mal-track/resources/mal-track(Fynloski%20sample%2C%20ON%20VM%20ONLY).zip)

*Note: You may need to make an exception in your antivirus to download, extract and use the malware.*

6. **Enable Host-VM File Communication:** To transfer files between your host and VM, go to the VM window, select Devices -> Insert Guest Additions CD Image, navigate to the CD drive, and run the setup file. Restart the VM afterward. On Virtual box navigate to settings -> Shared Folders -> Add a new shared folder. Select the folder you want to share. This will be visable in the VM under network.

7. **Clone script.py:** Transfer the script.py file from your repository to the VM.

## Objective

**Create script that will:**
1. Check for the malware process and kill it.
2. Remove the malware's execution from the startup programs.
3. Stop and remove the malware from the virtual machine.
4. Display the IP address of the attacker.


## Usage

1. **Run the Malware:** Launch mal-track.exe on your VM.
2. **Verify Process:** Check the malware process in the Task Manager.
3. **Verify Startup Execution:** Check for the malware in the Documents folder.
4. **Run the Script:** Execute script.py on your VM.
5. **Verify Removal:** Check if the malware has been successfully removed.

## Audit Questions

### How we can manage the startup programs in windows?

1. Open Task Manager:
- Press **Ctrl + Shift + Esc**
- Or right-click on the taskbar and select **Task Manager**
- Or press **Ctrl + Alt + Del** and select **Task Manager**
2. Click on **More details** if needed, then go to the Startup tab.
3. Right-click on the program you want to disable and select **Disable**.
### How to Extract the IP Address of the Attacker?

Extracting the IP address of an attacker from a malware file typically involves analyzing the malware's binary or code. Here's a simplified method to do this:

**Steps to Extract the IP Address:**
1. **Read the Binary Data:** Open the binary file (e.g., malware.exe) in binary mode ('rb').
2. **Search for IP Patterns:** Use a regular expression (regex) to find patterns in the binary data that match the structure of an IP address (e.g., 192.168.0.1).
3. **Validate and Output:** Ensure the extracted sequences are valid IP addresses and print them out.

### Example Code from script.py:
```python
import re

def extract_ip_address(binary_path):
    ip_addresses = []
    
    with open(binary_path, 'rb') as file:
        binary_data = file.read()

        # Use regular expressions to find IP addresses in the binary data
        ip_regex = r'\b(?:\d{1,3}\.){3}\d{1,3}\b'
        matches = re.findall(ip_regex, str(binary_data))
        
        # Add unique IP addresses to the list
        for match in matches:
            if match not in ip_addresses:
                ip_addresses.append(match)
    
    return ip_addresses
```
### Why This Might Work:
- **Malware Network Information:** Many types of malware need to communicate with external servers, and these addresses might be hardcoded into the binary.
- **Binary Inspection:** By analyzing the binary file, you can extract useful information without needing to decompile or execute the malware.
### Limitations:
- **False Positives:** The regex might match sequences that are not IP addresses.
- **Obfuscation:** Advanced malware might obfuscate its communication protocols, making this method ineffective.
- **Dynamic Resolution:** Some malware may resolve IP addresses dynamically, which static analysis won't capture.
### Advanced Techniques:
- **Dynamic Analysis:** Running the malware in a sandboxed environment to observe its behavior can reveal the IP addresses it contacts.
- **Decompilation:** Disassembling the malware to inspect its code logic for IP address retrieval mechanisms.
## mal-track task description

Task description and audit questions can be found  [here](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/mal-track).