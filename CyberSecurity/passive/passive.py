import os
import re
from collections import Counter
import subprocess
import json
from bs4 import BeautifulSoup
import click
import requests
import pyap


@click.command(context_settings={"help_option_names": ["-h", "--help"]})
@click.option("-fn", "--fn", help="Search with full-name", show_default=False)
@click.option("-ip", "--ip", help="Search with IP address", show_default=False)
@click.option("-u", "--u", help="Search with username", show_default=False)
def passive(fn: str, ip: str, u: str) -> None:
    if fn:
        print(f"Searching with full name: {fn}")
        name_finder(fn)
        return
    if ip:
        print(f"Searching with IP address: {ip}")
        ip_finder(ip)
        return
    if u:
        print(f"Searching with username: {u}")
        user_finder(u)
        return
    print("No search option provided. Use --help for more information.")


def name_finder(name: str) -> None:
    try:
        first_name, last_name = name.split(" ", 1)
    except ValueError:
        print(
            f"Error: Full name must contain first and last name separated by a space."
        )
        return

    phone_number = phone_finder(first_name, last_name)
    address = address_finder(first_name, last_name)
    print(
        f"First name: {first_name}\nLast name: {last_name}\nAddress: {address}\nNumber: {phone_number}"
    )
    save_data(
        {
            "First name": first_name,
            "Last name": last_name,
            "Phone number": phone_number,
            "Address": address,
        }
    )


def phone_finder(first_name: str, last_name: str) -> str:
    data = search(first_name, last_name, "phone number")
    all_phone_numbers = filter_numbers(data)
    most_common_phone_number = find_most_common_item(all_phone_numbers)
    return most_common_phone_number


def address_finder(first_name: str, last_name: str) -> str:
    data = search(first_name, last_name, "address")
    all_addresses = filter_addresses(data)
    most_common_address = find_most_common_item(all_addresses)
    return most_common_address


def search(first_name: str, last_name: str, type: str) -> list:
    url = f"https://www.google.com/search?q={first_name}+{last_name}+{type}"
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
    }
    response = requests.get(url, headers=headers, timeout=10)
    html_content = response.text
    soup = BeautifulSoup(html_content, "html.parser")
    span_elements = soup.find_all("span")
    div_elements = soup.find_all("div")

    text = [element.text.strip() for element in span_elements + div_elements]
    return text


def filter_addresses(data):
    addresses = []
    address_patterns = [
        r"\b[\w\s\.]+\b,\s*\d{5}\s+\w+",
        r"\b[\w\s\.]+\b\s+\d+-\w+,\s*\d{5}",
        r"^[A-Z][a-zA-Z\s]+\d{1,3}(?:-[A-Za-z]?\d{1,3})?, \d{5}$",
        r"^[A-Z][a-zA-Z\s]+\d{1,3}(?:-[A-Za-z]?\d{1,3})?,\s*",
        r"\b[\w\s\.]+\b \d{1,3}, \d{5} [A-Za-z]+",
        r"\b[\w\s\.]+\b \d+/\d+-\d+, \d{5}",
    ]

    for item in data:
        parsed_addresses = pyap.parse(item, country="US")
        for address in parsed_addresses:
            addresses.append(str(address))

    if not addresses:
        addresses = find_matches(address_patterns, data)

    return addresses


def filter_numbers(data):
    patterns = [
        r"\b5\d{6}\b",
        r"\b5\d{7}\b",
        r"\b5\d{6,10}\b",
    ]

    phone_numbers = find_matches(patterns, data)

    if not phone_numbers:
        patterns = [
            r"\(\d{3}\) \d{3}-\d{4}",
            r"\d{3}-\d{3}-\d{4}",
            r"\d{3}.\d{3}.\d{4}",
            r"\(\d{3}\)\d{3}-\d{4}",
            r"\d{10}",
            r"\d{3} \d{3} \d{4}",
        ]
        phone_numbers = find_matches(patterns, data)

    return phone_numbers


def find_matches(patterns, text_content):
    all_matches = []

    for pattern in patterns:
        regex = re.compile(pattern)

        for line in text_content:
            matches = regex.findall(line)
            if matches:
                all_matches.extend(matches)
    return all_matches


def find_most_common_item(items):
    most_common_item = Counter(items).most_common(1)
    return most_common_item[0][0] if most_common_item else None


def ip_finder(ip: str) -> None:
    print(f"Searching for {ip} from the World Wide Web")

    url = f"https://ipapi.co/{ip}/json/"
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
    }

    try:
        resp = requests.get(url, headers=headers, timeout=10)

        if resp.status_code == 200:
            data = resp.json()
            print(json.dumps(data, indent=4))
            save_data(data)
        else:
            print(f"Error: Received status code {resp.status_code}")

    except requests.exceptions.RequestException as e:
        print(f"An error occurred: {e}")


def user_finder(u: str) -> None:
    result = ""
    try:
        command = [
            "sherlock",
            u,
            "--site",
            "Github",
            "--site",
            "Reddit",
            "--site",
            "Instagram",
            "--site",
            "Discord",
            "--site",
            "Youtube",
        ]
        subprocess.run(command, check=True)

        filename = f"{u}.txt"
        while not os.path.exists(filename):
            pass

        with open(filename, "r") as file:
            result = file.read()

            os.remove(filename)

    finally:
        save_data(result)


def save_data(data: dict) -> None:
    base_filename = "result"
    extension = ".txt"
    filename = f"{base_filename}{extension}"
    i = 2
    if not data:
        print("No data to save")
    else:
        while os.path.exists(filename):
            filename = f"{base_filename}{i}{extension}"
            i += 1

        with open(filename, "w", encoding="utf-8") as file:
            file.write(json.dumps(data, indent=4))
            print(f"Data saved to {filename}")


def main() -> None:
    passive()


if __name__ == "__main__":
    main()
