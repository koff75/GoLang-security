# GoLang-security
Golang for Security. Shodan | Search scraper | Metasploit

## Shodan
It uses the Shodan API like this : https://api.shodan.io/shodan/host/search?key={YOUR_API_KEY}&query={query}&facets={facets}

usage: SHODAN_API_KEY=YOUR-KEY go run main.go tomcat

results: IPs

## Search engines

usage:go run main.go sitename.com docx

results: links + document creator / modification date / software / version / ...
