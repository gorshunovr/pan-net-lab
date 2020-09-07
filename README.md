# Pan-Net lab

![CI-linting](https://github.com/gorshunovr/pan-net-lab/workflows/CI-linting/badge.svg?branch=master)

## Environment

Environment is expected to be launched via `vagrant up` command. Vagrant would
bring up 2 Ubuntu 16.04 virtual machines:

* *n0* - this is the main VM:

  * Docker
  * Ansible (VM manages itself trough Ansible)
  * Containerized `getweather` application
  * `scanner` application

* *n1* - this VM is used for testing purposes only

```code
.
├── ansible/
├── getweather/
├── scanner/
├── bootstrap.sh*
├── LICENSE
├── README.md
└── Vagrantfile
```

## Exercise 1

### Programming 1

Please, see [README](getweather/README.md) for details on *getweather*
application.

### Ansible 1

Ansible is installed on *n0* virtual machine via `bootstrap.sh` shell script,
which is called by Vagrant during provisioning process. After Ansible is
installed, Vagrant runs Ansible playbook `setup.yml`:

```bash
cd /vagrant/ansible/ && ansible-playbook -i inventory.txt setup.yml
```

Please, see [README](ansible/README.md) for details on Ansible setup and run.

### Docker

Ansible playbook builds `getweather:1.0` container with Go application packaged
into it.

## Exercise 2

### Programming 2

Please, see [README](scanner/README.md) for details on *scanner* application.

## Exercise 3

### Syslog configuration

Syslog is configured with the following parameters:

* default logging to `/var/log/*`
* custom log files - per-host logging

### Ansible 3

Ansible role `rsyslog` has been created.
Using Jinja2 templates of `rsyslog.conf` and `50-default.conf` files it allows to
parametrize the following:

* logging only default log files
* logging custom files
* selecting external log server to send logs to

Default values of parameters are listed in role defaults:

* `rsyslog_udp_server` boolean and `rsyslog_udp_server_port` number for `rsyslog.conf` file
* `rsyslog_tcp_server` boolean and `rsyslog_tcp_server_port` number for `rsyslog.conf` file
* `rsyslog_custom_rules` list of strings with rules for `50-default.conf.j2` file
* `relay_logs` list of `facility` (e.g. `mail`, `user`, etc.), with logs sent to defined
  `destination_server` via both UDP and TCP (not RELP).

Verification that logs from other servers are being properly delivered to *n0* VM
is done by running the following command on test *n1* VM under standard user
account (assuming `192.168.56.10` is an IP address of *n0*):

```bash
logger --server 192.168.56.10 "Test message from n1 VM"
```

And checking results on *n0* VM:

```bash
vagrant@n0:/vagrant/scanner$ sudo grep Test /var/log/n1/system.log
Nov 22 20:48:26 n1 vagrant Test message from n1 VM
vagrant@n0:/vagrant/scanner$
```

> **NOTE:** Path to log file includes hostname *n1* as requested.

## Conventions

* Ansible roles and playbooks should be linted by `ansible-lint` and/or `ansible-review` tools
* Shell scripts should be linted by `shellcheck` tool
* AWK scripts should be linted by `awk --lint ...`
* Go code should be linted by `go fmt ...`
* Documentation should be in Markdown format and linted by `markdownlint`
