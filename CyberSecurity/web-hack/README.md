# web-hack

### Objective

This project assingment is to deploy a [DVWA(Damn Vulnerable Web Application)](https://github.com/digininja/DVWA) and find atleast 3 vulnerabilities in the application and exploit them. In additsion will have to develop a C99/R57 Type PHP Shell that allows you to add a file, delete a file and execute a shell commands.

In this md file I will be documenting the vulnerabilities I have found and how I exploited them.

### Damn Vulnerable Web Application (DVWA)

**What is DVWA?**
 
  Damn Vulnerable Web Application (DVWA) is a PHP/MySQL web application that is damn vulnerable. Its main goal is to be an aid for security professionals to test their skills and tools in a legal environment, help web developers better understand the processes of securing web applications and to aid both students & teachers to learn about web application security in a controlled class room environment. 


You can calibrate the security level to your preference, and test your skills in a safe enviorment (low, medium, high or impossible). The security level changes the vulnerability level of DVWA:
- Low - This security level is completely vulnerable and has no security measures at all. It's use is to be as an example of how web application vulnerabilities manifest through bad coding practices and to serve as a platform to teach or learn basic exploitation techniques.

- Medium - This setting is mainly to give an example to the user of bad security practices, where the developer has tried but failed to secure an application. It also acts as a challenge to users to refine their exploitation techniques.

- High - This option is an extension to the medium difficulty, with a mixture of harder or alternative bad practices to attempt to secure the code. The vulnerability may not allow the same extent of the exploitation, similar in various Capture The Flags (CTFs) competitions.

- Impossible - This level should be secure against all vulnerabilities. It is used to compare the vulnerable source code to the secure source code.
   Prior to DVWA v1.9, this level was known as 'high'.


**Deploying DVWA**

*It's reccomended to deploy DVWA in a virtual environment. For this project I have used VirtualBox and Kali Linux VM which is set to NAT networking mode.*

DVWA can be deployed in many ways:
1. [Using Docker](https://hub.docker.com/r/vulnerables/web-dvwa/)
2. [Using Kali Linux in VirtualBox](https://www.youtube.com/watch?v=WkyDxNJkgQ4)
3. [Windows + XAMPP](https://www.youtube.com/watch?v=Yzksa_WjnY0)

For more detailed information on deploying DVWA, trobleshooting and more please refer to the [official documentation](https://github.com/digininja/DVWA)


## Vulnerabilities

### 1. Finding and Exploiting Reflected XSS Vulnerability

**What is Cross-Site Scripting(XSS)?**

Cross-Site Scripting (XSS) attacks are a type of injection, in which malicious scripts are injected into otherwise benign and trusted websites. XSS attacks occur when an attacker uses a web application to send malicious code, generally in the form of a browser side script, to a different end user. Flaws that allow these attacks to succeed are quite widespread and occur anywhere a web application uses input from a user within the output it generates without validating or encoding it.

An attacker can use XSS to send a malicious script to an unsuspecting user. The end user’s browser has no way to know that the script should not be trusted, and will execute the script. Because it thinks the script came from a trusted source, the malicious script can access any cookies, session tokens, or other sensitive information retained by the browser and used with that site. These scripts can even rewrite the content of the HTML page.

**Cross-Site Scripting (XSS) attacks occur when**:

Data enters a Web application through an untrusted source, most frequently a web request. The data is included in dynamic content that is sent to a web user without being validated for malicious content.

The malicious content sent to the web browser often takes the form of a segment of JavaScript, but may also include HTML, Flash, or any other type of code that the browser may execute. The variety of attacks based on XSS is almost limitless, but they commonly include transmitting private data, like cookies or other session information, to the attacker, redirecting the victim to web content controlled by the attacker, or performing other malicious operations on the user’s machine under the guise of the vulnerable site.

**How to Protect Yourself?**

The primary defenses against XSS are described in the OWASP XSS Prevention Cheat Sheet.

Also, it’s crucial that you turn off HTTP TRACE support on all web servers. An attacker can steal cookie data via Javascript even when document.cookie is disabled or not supported by the client. This attack is mounted when a user posts a malicious script to a forum so when another user clicks the link, an asynchronous HTTP Trace call is triggered which collects the user’s cookie information from the server, and then sends it over to another malicious server that collects the cookie information so the attacker can mount a session hijack attack. This is easily mitigated by removing support for HTTP TRACE on all web servers.

### Exploiting Reflected XSS Vulnerability:

Visiting the *Vulnerability: Reflected Cross Site Scripting (XSS)* page in DVWA you are presented with a form that takes a user input and echos it back to the user *Hello <user_input>*.

![form](/pics/form.png)

Viewing the source code of the form we can see that it is using a *GET* method. This also the reason why it is shown on the url bar. This should be avoided by using *POST* method.

![get](/pics/get.png)

This also means that everything is directly inserted to the html.

- **Low security level:** The user input is not sanitized and is directly inserted into the html. This means that you can insert a script tag and it will be executed.

In the user field insert the following code:

```<script> alert("You've been hacked!") </script>```

- **Medium security level:** The user input is sanitized poorly and the script tags are removed if it's a direct tag match. However, you can still insert a script tag by changeing tags to all caps.

```<SCRIPT> alert("You've been hacked!") </SCRIPT>```

- **High security level:** The user input is sanitized and the script tags are removed, but you can still use HTML event attributes to manipulate the HTML.

```<img src="x" onerror="alert('You\'ve been hacked!')">```

- **Impossible security level:** This level hacking needs some social engineering to be successful. You send a link to the user that contains the malicious script tag and the user clicks on the link. You will get the session and user token.

Examples how different levels of security are implimented:

![xsslvl](/pics/xsslvl.png)

### 2. Brute Force Attack

*For this exploit I will be using [BurpSuite](https://www.kali.org/tools/burpsuite/) and [FoxyProxy](https://getfoxyproxy.org/) to perform the brute force attack.*

**What is Brute Force Attack?**

A brute force attack can manifest itself in many different ways, but primarily consists in an attacker configuring predetermined values, making requests to a server using those values, and then analyzing the response. For the sake of efficiency, an attacker may use a dictionary attack (with or without mutations) or a traditional brute-force attack (with given classes of characters e.g.: alphanumeric, special, case (in)sensitive). Considering a given method, number of tries, efficiency of the system which conducts the attack, and estimated efficiency of the system which is attacked the attacker is able to calculate approximately how long it will take to submit all chosen predetermined values.

Brute-force attacks are often used for attacking authentication and discovering hidden content/pages within a web application. These attacks are usually sent via GET and POST requests to the server. In regards to authentication, brute force attacks are often mounted when an account lockout policy is not in place.

To start we will need an baseline of successful and unsuccessful login attempts. 

- Successful login attempt:

![login](/pics/login.png)

- Unsuccessful login attempt:

![loginfail](/pics/loginfail.png)


*Side note: If you left click on the succesful login picture you can see the url bar contains /users/<user_name>. If you insert only  /users/ you can see all the users in the database. What may become handy when choosing who to hack.*

![users](/pics/users.png)

**Exploiting Brute Force Attack Using Burp Suite:**

After you have successfully configured your browser and Burp Suite to work together, you can start the brute force attack.

- **Low security level:** 

The login has no sleep function and you can spam the login with different passwords. The successful login will have a longer response than the unsuccessful login.

1. Open Burp Suite and go to the *Proxy* tab and click on *Intercept is on*.

2. Go to the login page and insert the username and random password. Click on the *Login* button.

3. The request will be intercepted by Burp Suite.

![randompw](/pics/randompw.png)

4. Left click on Intruder
Go to the *Intruder* tab and select the *Positions* tab. Click on *Clear* to remove all the positions.


7. Click on the *Add* button to add a position. Click on the *Request* tab and select the username and password field. Click on *Add*.

8. Choose your Attack type. In this case we will be using *Cluster Bomb* attack type.

9. Go to the *Payloads* tab and choose *Runtime file* as payload type and insert the payload list. I used the one in metasploit library. Do this to both payload sets if you dont know the username you would like to attack. Payload set 1 is for the username and payload set 2 is for the password. You can also use username from list provided in the users page.

10. Click on the *Start attack* button. The attack will start and you will see the progress in the *Intruder* tab.

![bruteattack](/pics/bruteattack.png)

11. After the attack is finished you can see the results in the *Intruder* tab. You cant determen a successful login by the length field of the response. The successful login will have a longer response than the unsuccessful login and we can determen that the  correct username is **admin** and the correct password is **password**.

![brutesuc](/pics/brutesuc.png)

- **Medium security level:**

The login has a sleep function that will delay the response if the login is unsuccessful. This will make the brute force attack slower.

Only difference between the low and medium security level is that the response time is longer for the unsuccessful login.(200ms)

- **High security level:**

If we view the source code of the login page we can see that web server will first check for Anti-CSRF token. If the token is not present the login will fail. After each failed login attempt the server will sleep for 3 seconds.

**What is Anti-CSRF token?**

An Anti-CSRF token (also known as a CSRF token, XSRF token, or one-time token) is a security mechanism used to prevent Cross-Site Request Forgery (CSRF) attacks. CSRF is an attack where a malicious website tricks a user's browser into making unintended requests to a different site where the user is authenticated, potentially leading to unauthorized actions.

**How CSRF Works?**

CSRF attacks exploit the trust that a site has in the user's browser. If the user is logged in to a site (e.g., a bank or email account), and the browser has an active session or authentication cookie, a malicious site can make requests to the legitimate site without the user’s knowledge.

Only difference between the medium and high security is that you have to capture the user_token save and use it when attacking. Also the wait time will be second longer.

- **Impossible security level:**

    Requires social engineering to be successful or will take unreasonable amount of time using brute force. This is because it has implimented time out after 3 failed login attempts user will have to wait 15 min to try to login again and the Anti-CSRF token is required.

### 3. Command Injection

**What is command injection?**

Command injection is an attack in which the goal is execution of arbitrary commands on the host operating system via a vulnerable application. Command injection attacks are possible when an application passes unsafe user supplied data (forms, cookies, HTTP headers etc.) to a system shell. In this attack, the attacker-supplied operating system commands are usually executed with the privileges of the vulnerable application. Command injection attacks are possible largely due to insufficient input validation.

This attack differs from Code Injection, in that code injection allows the attacker to add their own code that is then executed by the application. In Command Injection, the attacker extends the default functionality of the application, which execute system commands, without the necessity of injecting code.

### Command injection:

- **Low security level:**

In this exploit we will be using the *Vulnerability: Command Injection* page in DVWA. The page takes a IP address and pings it. The IP address is inserted directly into the command and is not sanitized.

![ij](/pics/ij.png)

After little bit of google search I found out that I can use the following command to get identical response in bash.

```ping -c4 <IP-address>```

Because it's a plain command directly to shell I can use bash syntax to execute commands.

For example:

```127.0.0.1 && echo "Get pwned" && ls;```

In the result we will that it's printing out the IP pings, echos out the message and lists the files in the directory.


![getpwnd](/pics/getpwnd.png)

You can also use netcat in your terminal to get a reverse shell.

```nc -nlvp <port_num>```

- **Medium security level:**

In Medium security level the user input is sanitized and the command is passed through a function that removes the semicolon and the double ampersand characters. This means that you can't chain commands.

Solution is to background the ping command and then execute the command.

```127.0.0.1 & echo "Get pwned"```

- **High security level:**

  The developer has blacklisted substitiution characters and the command is passed through a function that removes the semicolon and the double ampersand characters. This means that you can't chain commands. But he made a mistake with "| " having a trailing space. This means that you can use the following command to execute commands.

```127.0.0.1 |echo "Get pwned"```

![highij](/pics/highij.png)

### 4. File Inclusion - Php Shell

This is a bonus exploit that I have added to the project was covered in previous *["local"](https://01.kood.tech/git/IngvarLeerimaa/local)* project.

It's even easier here to use it:
- Go to the *Vulnerability: File Inclusion* page in DVWA
- Click upload
- Choose your php script
- Click upload
- Open terminal
- Use netcat to get a reverse shell ```nc -nlvp <port_num>```
- Open http://localhost/dvwa/hackable/uploads/<your_php_script>.php in your browser.

Congratulations you have a shell access!

**How to protect against file inclusion?**
- Limit Allowed File Types: Only allow file uploads of specific types (e.g., images or documents) that are necessary for your application. Use a whitelist approach to permit only certain extensions (e.g., .jpg, .png, .pdf).

- Validate File Contents: Check the contents of uploaded files to ensure they match their file type. For example, an image file should be validated to ensure it contains image data rather than executable code.

- Rename Uploaded Files: Rename uploaded files to a safe format and remove any original file extensions or names that could potentially be dangerous.

- Use a Secure Directory for Uploads: Store uploaded files in a directory outside the web root so that they cannot be directly accessed via a URL.