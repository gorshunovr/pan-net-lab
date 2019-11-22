# The scanner application

## Programming

This program provides a way to run repetitive network scans displaying
differences between subsequent scans.

* Target of the scan must be provided as CLI argument
* Target can be single IP address as well as network range

Since program uses `nmap` internally, you can provide different `nmap` options
as you need by setting `NMAP_OPTS` variable. Multiple hosts as a scan target are
supported in `nmap` *target specification* format, as long as it's one parameter
(e.g. `192.168.56.0/24` or `192.168.56.10-20`).

By default, `nmap` would scan 1,000 popular ports for each protocol, and we scan
with `TCP CONNECT` scan, as it's a fast scan option and does not require root
privileges.

Parsing of the scan log files is done by `awk` script `parser.awk`. Full scan
logs are available in `logs/` subdirectory.

```code
scanner/
├── logs/
├── parser.awk
├── README.md
└── scanner.sh*
```

## Running program and output example

Run initial scan against *n1*:

```bash
vagrant@n0:/vagrant/scanner$ ./scanner.sh 192.168.56.11
*Target - 192.168.56.11: Full scan results:*
Host: 192.168.56.11 Ports: 22/open/tcp//ssh///
vagrant@n0:/vagrant/scanner$
```

Run repetitive scan with no changes on target host:

```bash
vagrant@n0:/vagrant/scanner$ ./scanner.sh 192.168.56.11
*Target - 192.168.56.11: No new records found in the last scan.*
vagrant@n0:/vagrant/scanner$
```

Install web-service onto *n1* (in Ubuntu services start running right after
installation):

```bash
vagrant@n1:~$ sudo apt-get update && sudo apt-get install -y lighttpd
...
Setting up lighttpd (1.4.35-4ubuntu2.1) ...
...
vagrant@n1:~$
```

Run repetitive scan with changes on target host:

```bash
vagrant@n0:/vagrant/scanner$ ./scanner.sh 192.168.56.11
*Target - 192.168.56.11: Full scan results:*
Host: 192.168.56.11 Ports: 22/open/tcp//ssh///
Host: 192.168.56.11 Ports: 80/open/tcp//http///
vagrant@n0:/vagrant/scanner$
```

Run scan against /24 network:

```bash
vagrant@n0:/vagrant/scanner$ ./scanner.sh 127.0.0.1/24
*Target - 127.0.0.100: Full scan results:*
Host: 127.0.0.100 Ports: 22/open/tcp//ssh///
...
*Target - 127.0.0.9: Full scan results:*
Host: 127.0.0.9 Ports: 22/open/tcp//ssh///
vagrant@n0:/vagrant/scanner$
```
