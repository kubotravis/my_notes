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


#### 2. Creating an Environment

LAB - Setup the 3 vm by using the vagrant (acs,web,db)

**Installing the Ansible**

- On one of CENTOS its epel-release based installation

```bash
$ yum install epel-release
$ yum install ansible -y
```
- Validate by using,

`$ ansible`


**Testing the LAB setup**

Creating an inventory [Ansible relay on this inventory to run any playbook]

`$ mkdir -pv ex-1`

Create a file name called 'inventory' and add the client IP address

```bash
$ vim inventory
  192.168.1.10
  192.168.1.20
```

Run the following command to get reposince as PONG [since we are using ping]

`$ ansible 192.168.1.10 -i ./inventory -u vagrant -m ping -k`


where,
```
 192.168.1.10 - Target server
 -i inventory file location
 -u username
 -m which modules needs to run
 -k prompt for the password
```

- Basically you will get the error msg from SSH for fingerprint related, in order to fix that just do a ssh command to each of server which is in inventory file to add the FINGERPRINT to known_hosts file

How to go in verbose/debug mode - simple add -v at the end of the command

```bash
 $ ansible 192.168.1.10 -i ./inventory -u vagrant -m ping -k -v
 $ ansible 192.168.1.10 -i ./inventory -u vagrant -m ping -k -vv
```

Here is an example of more verbose mode

```bash
$ ansible 192.168.1.10 -i ./inventory -u vagrant -m ping -k -vvv

Loaded callback minimal of type stdout, v2.0
<100.67.165.4> ESTABLISH SSH CONNECTION FOR USER: centos
<100.67.165.4> SSH: EXEC sshpass -d12 ssh -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r 192.168.1.10 '/bin/sh -c '"'"'( umask 77 && mkdir -p "` echo $HOME/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752 `" && echo ansible-tmp-1467459236.91-162511876456752="` echo $HOME/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752 `" ) && sleep 0'"'"''
<100.67.165.4> PUT /tmp/tmpvdp3jC TO /home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/ping
<100.67.165.4> SSH: EXEC sshpass -d12 sftp -o BatchMode=no -b - -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r '[100.67.165.4]'
<100.67.165.4> ESTABLISH SSH CONNECTION FOR USER: centos
<100.67.165.4> SSH: EXEC sshpass -d12 ssh -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r -tt 192.168.1.10 '/bin/sh -c '"'"'LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8 LC_MESSAGES=en_US.UTF-8 /usr/bin/python /home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/ping; rm -rf "/home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/" > /dev/null 2>&1 && sleep 0'"'"''
```

**What happens during the execution**

What actually happens when you run the PING modules against the any targer server (as info retrieved from the above -vvv option)
 - creates the ssh session by using the user/pass info
 - Then starts compiling the python module 'ping' [just compiles the package to make it executables]
    - making directory (space) for the executrion
    - then insode the making everything as "Executables"
    - then ECHO backs the DIR path to ansible, now ansible knows where the stuff needs tp be copied
    - now it puts the complied modules that already (/tmp/tmpvdp3jC) created into that NEW directory that created above
    - as last step by using the PYTHON it gonna execute that ping module
    - and removes the modules after the execution

**How to run the ad-hoc commands**
[by using the command modules]

`$ ansible 192.168.1.10 -i inventory -u centos -m command -a "uptime" -k`

Acutally you NO need to use the `-m command` to specify module because its default, so you can use as follows too

`$ ansible 192.168.1.10 -i inventory -u centos -a "uptime" -k`

- Well there is another module call `Shell`, where you can execute the shell commands.
- So, what is the difference in command & shell modules ?
  Shell module support ENVIRONMENT vairables (HOME,USER,PWD & others) where `command` moduls doesn't support.
