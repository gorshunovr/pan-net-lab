# Ansible

## Roles and playbooks

Roles:

* `docker` installs and sets up Docker
* `getweather` builds *getweather* container
* `rsyslog` sets up syslog
* `scanner` installs dependencies of *scanner* application

The `setup.yml` playbook applies roles onto the host.

### Conventions

Roles and playbooks should be linted by `ansible-lint` and/or `ansible-review`.
