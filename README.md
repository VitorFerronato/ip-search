# IP Search CLI

<div style="display:flex;">
  <img align="center" alt="go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
</div>

<br/>
This is a Go project that provides a command line application (CLI) to search for IP addresses and servers of a given host on the internet.

## 📌 Features

- Find IP addresses of a host

- Look up a host's nameservers (NS)

## 🚀 Installation

1. Clone
```bash
git clone https://github.com/seu-usuario/ip-search.git
cd ip-search
```
2. Install the dependecies

Make sure you have Go installed and run:
```bash
go mod tidy
```

## 🏃‍♂️ How to use

Run the following command to run the application:
```bash
go run main.go
```

🔍 Search IPs from a host
```bash
go run main.go ip --host=example.com
```

🔍 Look up a host's nameservers (NS)
```bash
go run main.go server --host=example.com
```

## 🛠 Technologies Used
- Go
- urfave/cli for CLI Building
