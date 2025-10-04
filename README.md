# godork
A handy companion for effortlessly generating Google dorking links.


## Prerequisite

* Go 

## Installation

```bash
# Install godork
$ go install github.com/bl155x0/godork@v0.1.0

# Get the dorkfile template or create your own
wget https://raw.githubusercontent.com/bl155x0/godork/refs/heads/main/dorks.txt
```

## Usage

Create the dorks for "example.com" and the specified dorks file
```
$ godork -d /opt/dorks.txt -u example.com
```

Specify a dorks file with an environment variable
```
# Specify the dorkfile to use by setting the env variable "GODORK_DORKFILE"
$ export GODORK_DORKFILE="/opt/dorks.txt"
$ godork -u example.com
```


Create dorks for various (sub)domains
```
cat subdomains.txt | godork
```
