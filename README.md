## About The Project

Gori is a tiny and rapid config file scanner. Gori allows to do concurrents scans regarding .git, .env, .travis.yml, config.yml files within GET request.

## Getting Started

### Prerequisites
* go > 1.13
Optional:
  ```sh
  go mod init
  ```

### Installation

1. Clone the repo
     ```sh
     git clone https://github.com/GetSource1234/gori.git
     ```
2. Setup go modules (optional if you see build errors)
     ```sh
     go mod init
     ```

## Usage
```sh
go run main.go --urlPath=<PATH_TO_DOMAIN_LIST>
```
Supported flags: urlPath, verbose, tor.
Domain list: simple .text file with whitespace as a separator.

Run tests:
```sh
go test ./...
```
#Note!!!
Some false positive results may occur, double check manually if needed.

#Note!!!
Make sure if tor socks proxy is running if you use *--tor* flag
*Tor Security note:*
https://2019.www.torproject.org/docs/faq.html.en#TBBSocksPort

## License

Distributed under the MIT License. See `LICENSE` for more information.


## DISCLAIMER

Usage of Gori for attacking targets without prior mutual consent is illegal. It is the end user's responsibility to obey all applicable local, state and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program.
