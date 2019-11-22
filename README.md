# Pan-Net lab

## Environment

Environment is expected launched via `vagrant up` command. Vagrant would bring up
2 Ubuntu 16.04 virtual machines:

* *n0* - this is the main VM:

  * Docker
  * Ansible (VM manages itself trough Ansible)
  * Containerized `getweather` application
  * `scanner` application

* *n1* - this VM is used for testing purposes only

## Exercise 1

### Programming (Exercise 1)

Please, see [README](getweather/README.md) for details on *getweather*
application.

### Ansible (Exercise 1)

Ansible is installed on *n0* virtual machine via `bootstrap.sh` shell script,
which is called by Vagrant during provisioning process. After Ansible is
installed, Vagrant runs Ansible playbook `setup.yml`:

```bash
cd /vagrant/ansible/ && ansible-playbook -i inventory.txt setup.yml
```

**NOTE:** Inventory refers only to this local node (*n0*) via `local` connection
driver.

The playbook performs the following actions in order:

* sets up `rsyslog` service via role for the **Exercise 3**
* installs and sets up Docker via role
* builds *getweather* application via role

Roles have simplified structure; non-needed parts have been excluded for
simplicity (e.g. ansible-galaxy info).

```code
ansible/
├── inventory.txt
├── roles
│   ├── docker
│   │   ├── files
│   │   │   └── daemon.json
│   │   ├── handlers
│   │   │   └── main.yml
│   │   └── tasks
│   │       └── main.yml
│   ├── getweather
│   │   └── tasks
│   │       └── main.yml
│   └── rsyslog
│       ├── defaults
│       │   └── main.yml
│       ├── handlers
│       │   └── main.yml
│       ├── tasks
│       │   └── main.yml
│       └── templates
│           ├── 50-default.conf.j2
│           └── rsyslog.conf.j2
└── setup.yml
```

### Docker

Ansible playbook builds `getweather:1.0` container with Go application packaged
into it.

## Exercise 2

### Programming (Exercise 2)

Please, see [README](scanner/README.md) for details on *scanner* application.

## Exercise 3

### Syslog configuration

Syslog is configured with the following parameters:

* default logging to `/var/log/*`
* custom log files - per-host logging

### Ansible (Exercise 3)

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
account (assuming `192.168.122.10` is an IP address of *n0*):

```bash
logger --server 192.168.122.10 "Test message from n1 VM"
```

And checking results on *n0* VM:

```bash
cat /var/log/192.168.122.10/system.log
```
