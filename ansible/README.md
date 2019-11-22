# Ansible

## Roles and playbooks

Roles:

* `docker` installs and sets up Docker
* `getweather` builds container
* `rsyslog` sets up syslog

The `setup.yml` playbook applies roles onto the host and installs prerequisite
for the *scanner* application.
