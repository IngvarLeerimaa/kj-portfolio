
# [passive](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive)

## Table of Contents
- ### [General Information](#general-information)
- ### [Osint](#osint)
- ### [Setup](#setup)
- ### [Usage](#usage)
- ### [Technologies used](#technologies-used)
- ### [Audit Questions](#audit-questions)


### General Information
For this project, you will create a program that will allow you to retrieve information about a person using Full name, an IP address, or a username from public sources. I used the Sherlock project to search for usernames on social networks, and ipapi to search for information about an IP address. The program will save the results of each command in a result.txt file. If the result.txt file already exists, a new file will be created. The program will be able to search for a full name, an IP address, and a username. The program will display the address, and the telephone number for the full name entered. The program will display the ISP, and position for the entered IP address. The program will check if the user entered is on 5 popular social networks. I limited it to 5 random popular sites to limit the run time. If needed its possible to check against all the popular sites. The program will retrieve this information from a public source. The program will save the result of each command in a result.txt file. If the result.txt file already exists, a new file will be created.

### Osint
Open Source Intelligence (OSINT) involves collecting and analyzing publicly available information to produce actionable intelligence. OSINT leverages various open sources like websites, social media, public records, news articles, and other publicly accessible data. Here are the key aspects:

Sources of OSINT:

Internet: Websites, blogs, forums, online databases.
Social Media: Platforms like Twitter, Facebook, LinkedIn.
Public Records: Government publications, company filings, legal documents.
News Media: Newspapers, TV broadcasts, online news sites.
Academic Publications: Research papers, academic journals.
Technical Sources: WHOIS information, domain registrations, IP address data.
Techniques:

Web Scraping: Extracting data from websites.
Social Media Monitoring: Analyzing social media activities.
Search Engines: Using advanced search operators to find specific information.
Data Aggregation: Compiling information from multiple sources.
Applications:

Cybersecurity: Identifying potential threats and vulnerabilities.
Investigations: Supporting law enforcement and private investigations.
Business Intelligence: Gathering competitive intelligence.
Due Diligence: Conducting background checks.
Tools and Software:

Maltego: For graphical link analysis.
Shodan: Search engine for Internet-connected devices.
SpiderFoot: Automated OSINT tool.
Sherlock: Finds usernames across social networks.
Ethical and Legal Considerations:

Ensure compliance with laws and regulations regarding data privacy and usage.
Ethical responsibility to avoid harm and respect privacy.
### Setup
1. Clone the repository:
    ```sh
    git clone https://github.com/IngvarLeerimaa/passive.git
    ```
2. Navigate to the project directory:
    ```sh
    cd passive
    ```
3. Install the required dependencies:
    ```sh
    pip install -r requirements.txt
    ```
### Usage

OPTIONS:

    `-fn`         Search with full-name
    `-ip`         Search with ip address
    `-u`          Search with username

EXAMPLES:

For name, phone number, address:

```python3 passive.py -fn "Jean Dupont"```

For ISP, position etc..:

```python3 passive.py -ip 127.0.0.1```

For social networks (Github, Reddit, Instagram, Discord, Youtube):

```python3 passive.py -u "@user01"```


Visual example:
![](https://01.kood.tech/git/IngvarLeerimaa/passive/audit.gif)


### Technologies used
- [Beautiful Soup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/)
- [ipapi](https://ipapi.co/)
- [Sherlock-project](https://github.com/sherlock-project)

and other libraries.



### [Audit Questions](https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive/audit)

#### General

###### Is the student able to explain clearly the used investigative methods? 

###### Is the student able to explain clearly what OSINT means?

###### Is the student able to explain clearly how his program works?

##### Check the Repo content

Files that must be inside your repository:

- Your program source code.

- A README.md file, which clearly explains how to use the program.

###### Are the required files present?

##### Ask the student to present his program to you by doing 3 tests

###### Is the information entered as an argument a full name, an IP address, and a username?

##### Try flag "-fn" with the following command `passive -fn "Jean Dupont"`

###### Does the program display the address, and the telephone number for the full name entered?

##### Try flag "-ip" with the following command `passive -ip 127.0.0.1`

###### Does the program display the ISP, and position for the entered IP address?

##### Try flag "-u" with the following command `passive -u "@user01"`

###### Does the program check if the user entered is present in is present in at least 5 social networks?

###### Does the program retrieve this information from a public source?

###### Does the program save the result of each command in a result.txt file?

###### If the result.txt file already exists is a new file created?