# Ansible

> **NOTE:** Inventory `inventory.txt` refers only to this local node (*n0*) via
  `local` connection driver.

The `setup.yml` playbook performs the following actions in order:

* sets up `rsyslog` service via role for the
  [Exercise 3](../README.md#Exercise-3)
* installs and sets up Docker via `docker` role
* builds *getweather* application via `getweather` role
* installs dependencies for the *scanner* application via role for the
  [Exercise 2](../README.md#Exercise-2)

Roles have simplified structure; non-needed parts have been excluded for
simplicity (e.g. `ansible-galaxy` meta information). Each role has corresponding
tag assigned (see `ansible-playbook --list-tags ...`). It is possible to run
specific role by specifying corresponding tag
(e.g. `ansible-playbook --tags scanner ...`).

```code
ansible/
├── roles/
│   ├── docker/
│   │   ├── files/
│   │   │   └── daemon.json
│   │   ├── handlers/
│   │   │   └── main.yml
│   │   └── tasks/
│   │       └── main.yml
│   ├── getweather/
│   │   └── tasks/
│   │       └── main.yml
│   ├── rsyslog/
│   │   ├── defaults/
│   │   │   └── main.yml
│   │   ├── handlers/
│   │   │   └── main.yml
│   │   ├── tasks/
│   │   │   └── main.yml
│   │   └── templates/
│   │       ├── 50-default.conf.j2
│   │       └── rsyslog.conf.j2
│   └── scanner/
│       └── tasks/
│           └── main.yml
├── inventory.txt
├── README.md
└── setup.yml
```

## Roles and playbooks

Roles:

* `docker` installs and sets up Docker
* `getweather` builds *getweather* container
* `rsyslog` sets up syslog
* `scanner` installs dependencies of *scanner* application

The `setup.yml` playbook applies roles onto the host.

### Conventions

Roles and playbooks should be linted by `ansible-lint` and/or `ansible-review`.
