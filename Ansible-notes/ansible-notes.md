Ansible Reference Notes
-----------------------

#### 1. Architecture & Process flow

**Control server**

(Ansible) ------> Target server

Requirements: (Ansible-Control server)
- Python 2.6+
- Platform Unix/Linux NOT Win32

Target servers:
- \*nix
- Python (simpleJSON library) + SSH
(During the course my learning Python 3.x is not supported)

**Components overview**

```
  Inventory
    |
    |
  PlayBook [Consists of play(s)]  ----> Ansible Config ---> Python |------->SSH -------> Target Servers
    ^
    |
  Modules
```

Inventory -> TXT file, that describes your servers & system
[host, host groups, user accounts - which is used to manage the stuff, host related env variables]

Modules -> Programmed unit of work to be done. Its the command centre of the system
[Core modules, extra modules]
[ex: yum modules install some package on the system]

PlayBook -> A single task from a module, executed on a host or a hostgroup
[Glue that brings all together]

Ansible Config -> Global configuration of ansible stays here
[How many parallel exe supposed to run]
[Notify if the target been skipped]
[etc]

**Some variable information**
Three are some variables

1. Host Variables [will be there is Inventory files]
2. Facts - Gathered data from target system before the playbook run, you can use those values to evaluate something, if required !
3. Dynamic Variables - These are the var/values created during the course of playbook run & destroyed after a playbook run

**Process of Execution and Flow**
 - Identify which play/playbooks needs to executed
 - Identify the Target to execute them
 - Establish the SSH & Copy the Play/Play book to the target temp directory
 - Execure the playbook on the target servers
 - Target sends results back to the Control server as a JSON

**Execution Types**
There are TWO execution types
- remote (default): remote execution
- local: local ansible server executing the playbook
[Local execution is pretty typical execution though, since the python package cant be received by default for local server, for that u have to make HTTP/REST call to execute on the local ]
