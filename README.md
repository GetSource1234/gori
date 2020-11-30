## About The Project

Gori is a tiny and rapid config file scanner. Gori allows to do concurrents scans regarding .git, .env, .travis.yml, config.yml files within GET request.

## Getting Started

### Prerequisites
* go > 1.13
  ```sh
  go mod init
  ```

### Installation

1. Clone the repo
     ```sh
     git clone https://github.com/your_username_/Project-Name.git
     ```
2. Setup go modules (optional if you see build errors)
     ```sh
     go mod init
     ```

## Usage
```sh
go run main.go --urlPath=<PATH_TO_DOMAIN_LIST>
```
Supported flags: urlPath, verbose.
Domain list: simple .text file with whitespace as a separator.

Run tests:
```sh
go test ./...
```

## License

Distributed under the MIT License. See `LICENSE` for more information.


## DISCLAIMER

Usage of Gori for attacking targets without prior mutual consent is illegal. It is the end user's responsibility to obey all applicable local, state and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program.
