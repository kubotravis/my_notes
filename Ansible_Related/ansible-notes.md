Ansible Reference Notes
-----------------------

### Table of Contents
[1. Introduction](https://docs.ansible.com/ansible/latest/index.html)

[2. Architecture & Process flow](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#2-architecture--process-flow)

[3. Creating an Environment](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#3-creating-an-environment)

[4 Ansible Inventory and Configuration](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#4-ansible-inventory-and-configuration)

[5. Modules](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#5-modules)

[6. Plays and Playbook](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#6-plays-and-playbook)

[7. Roles](https://github.com/kubotravis/my_notes/blob/master/Ansible_Related/ansible-notes.md#7-roles)

---

### 2. Architecture & Process flow

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
---

### 3. Creating an Environment

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
<192.168.1.4> ESTABLISH SSH CONNECTION FOR USER: centos
<192.168.1.4> SSH: EXEC sshpass -d12 ssh -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r 192.168.1.10 '/bin/sh -c '"'"'( umask 77 && mkdir -p "` echo $HOME/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752 `" && echo ansible-tmp-1467459236.91-162511876456752="` echo $HOME/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752 `" ) && sleep 0'"'"''
<192.168.1.4> PUT /tmp/tmpvdp3jC TO /home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/ping
<192.168.1.4> SSH: EXEC sshpass -d12 sftp -o BatchMode=no -b - -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r '[192.168.1.4]'
<192.168.1.4> ESTABLISH SSH CONNECTION FOR USER: centos
<192.168.1.4> SSH: EXEC sshpass -d12 ssh -C -vvv -o ControlMaster=auto -o ControlPersist=60s -o User=centos -o ConnectTimeout=10 -o ControlPath=/home/centos/.ansible/cp/ansible-ssh-%h-%p-%r -tt 192.168.1.10 '/bin/sh -c '"'"'LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8 LC_MESSAGES=en_US.UTF-8 /usr/bin/python /home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/ping; rm -rf "/home/centos/.ansible/tmp/ansible-tmp-1467459236.91-162511876456752/" > /dev/null 2>&1 && sleep 0'"'"''
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
---
### 4 Ansible Inventory and Configuration

**Inventory fundamentals**
- Basically inventory file can be at anywhere in the files system
- What are things can be done on the inventory files
  - Create an required parameter with values
  - Create Groups
  - Grouping the Groups
  - Assign variables
  - Scaling out of multiple files
  - Creating static/dynamic information

**Sample inventory files**

**Sample 1**

```
[db]  # its a groupname
server1.ex.com ansible_ssh_user=dummy ansible_ssh_pass=123 # ansible related var/val
server2.ex.com ansible_python_interpreter=/usr/bin/python

[datacenter-west:children]   # here dc-west has the 'db' group as child group (means all nodes under db group are target)
db

[datacenter-west]   # here dc-west has the 'db' as NODE (actual target not a GROUP)
db

[datacenter-west:vars]   # here dc-west has the group specific var/values
ansible_ssh_user=testuser
ansible_ssh_pass=khjhj%^$%
```

**Sample 2**

```
exp-elastic501z ansible_ssh_host=192.168.1.209
exp-elastic502z ansible_ssh_host=192.168.1.4
exp-rundeck501z ansible_ssh_host=192.168.1.37
exp-kibana401z ansible_ssh_host=192.168.1.216
exp-elastic601z ansible_ssh_host=192.168.1.169
exp-kibana101z ansible_ssh_host=192.168.1.236
exp-jslave402z ansible_ssh_host=192.168.1.113
exp-jslave401z ansible_ssh_host=192.168.1.115
exp-jmaster401z ansible_ssh_host=192.168.1.116

[efk]
exp-elastic501z
exp-elastic502z
exp-kibana401z
exp-kibana101z

[cicd]
exp-rundeck501z
exp-jmaster401z
exp-jslave401z
exp-jslave402z

[exp:children]
efk
cicd

[exp:vars]
ansible_ssh_user=centos
ansible_ssh_passwd=redhat or [ansible_ssh_private_key_file=~/.ssh/id_rsa]
```

**Scaling out with multiple files**

- Basic directory structure
  - Means that, taking about managing multiple inventory files by ansible which gives more clarity. also which has the directory structure
  - Ansible can merge separate inventory files into 1 file (if you want)

- Variable and values precedence
  - That how variables can be applied while executing any tasks on ansibe

```  
    > All (Group_vars)
    >> GroupName (Group_vars)
    >>> HostName (Host_Vars)
```
  - From the above we can see host specific variables will override anything which is specified in the group & all

**Demo Scaling out with multiple files**

- Here we are going to demo 'How variable precedence flows based on ENV directory files'
- Except the inventory file, all files are used for declaring the variables

Info:
  My directory structure
  Files which are listed here are YAML files [So remember the YAML syntax]

```
├── exp
│   ├── group_vars   # All group related VAR goes here
│   │   ├── all
│   │   └── efk      # Specific group relate VAR goes here
│   ├── host_vars
│   │   └── some_server_name_from_efk     # Host specific VAR goes here
│   └── inventory_exp
└── test
    ├── group_vars
    │   ├── all
    │   └── efk
    ├── host_vars
    │   └── some_server_name_from_efk
    └── inventory_test
```

- `inventory_exp` is the same file from ex-2 inventory file
- Now, we are going to add some user on the target server by using the 'all' file to define the var/values

  `$ vi exp/group_vars/all`

  Above file content goes like below

  ```
  ---
  # This is our test username
  username: testuser
  ```

`$ ansible efk -i exp/inventory_exp -m user -a "name={{username}} password=12345"`

where,

"-a" to pass the ARG values to the module "user", about the parenthesis declaration will se in PLAYBOOK creation

- Watch the error regarding the cicd files since its not YAML you might get the error regarding this
- Also there is chance of failing due to permission, so supply the "--sudo" with it.

`$ ansible efk -i "exp/inventory_exp" -m user -a "name={{username}} passwd=12345" --sudo`

- Now create the "exp/group_vars/efk" file with same content different userame & run the above command.
- This is time it will create the USERNAME which is specified the "group_vars" file

`$ ansible efk -i "exp/inventory_exp" -m user -a "name={{username}} passwd=12345" --sudo`

- Now create the "exp/host_vars/some_server_name_from_efk" with same content but with DIFFERENT username & run the same command as above
- This is time it will create the USERNAME which is mentioned in the HOST specific var file

`$ ansible efk -i "exp/inventory_exp" -m user -a "name={{username}} passwd=12345" --sudo`

- So this how precedence follows, simply it goes has follows
  ```
    Host Precedence
    |
    |_ Group Precendence
        |
        |_ All precedence
  ```

**Ansible Configuration basics**

- This how Ansible will load the configuration file location  [Anyone if them, at time not ALL]

```
 1st ~> $ANSIBLE_CONFIG [variable]
 2nd ~> ./config.cfg [From the $PWD]
 3rd ~> ~/.config [from $HOME as hidden file]
 4th ~> /etc/ansible/ansible.cfg [A global config file] - used at LAST
```

- IMPORTANT - Configuration files are NOT merged, FIRST on always wins [which ever it finds]
- But you can onverride the configuration settings by ENV variables

```
ex:
  $ANSIBLE_<sometthing>=123
  [or]
  export ANSIBLE_<something>=123
  [or]
  Update the things in bashrc/profile files to load as soon as you login
```
- Some of the setting for [under [default] section]
 - **fork** [how many execution at the time, default is 5, you can make your own number here]
 - **host_key_cehcking** [highly recommended on PROD]
 - **logging** [writes info above exec, default it NULL, you can set any files location to log the info]

**Demo on above theory**

- Here you can find the all options that available `http://docs.ansible.com/ansible/intro_configuration.html`

- EX-1
 - delete the content ~/.ssh/known_host
 - create the ansible.cfg file (inside/outside of the exp dir)
 - Append the following conetent

```
[defaults]
host_key_checking=False
```
 - now if you run any ansible command to check the HOST related it will work
 - Also it appends the TARGET fingerprint in to the ~/.ssh/known_host too in background silently
 - EX-2

 - export the following to check the failure case above scenarion
  `$ export $ANSIBLE_HOST_KEY_CHECKING=True`
 - Run any ansible command against the Target servers it will fail now !
 - Because the precedence goes to ENV variables over CFG files

**Working with Python 3 based systems**

 - its pretty simple , install python 2.x version on the target system & update the HOST parameter to pick up the same
```
$ vim invetory
10.0.0.0 ansible_python_interpreter=/usr/bin/python2.7
```

***Note:*** Above talks with base system that which you run ansible commands. There are support to Python3 ont the both the end. Yet to chaeck the Docs page.

---

### 5. Modules

Modules are Building block of you Automation

**Fundamentals**

- It does every single action on the target system for you

**3 types of modules**
- Core [from Ansible]
- Extra [from Community]
- Deprecated [Old one]

You no need to all the time to internet to browse the information about modules,

```bash
$ ansible-doc -l
  - list all modules from Ansible

$ ansible-doc <modulesname>
  - more info about the module

$ ansible-doc -s <name>
  - playbook snippet examples
```

- Ansible has 15 Categeries of playbook it supports as of now (server, config mngmt, n/w device mangment, deploy, cicd & many)

Some of the modules, 

  - copy [local to remote]
  - fetch [remote to local]
  - apt/yum/deb for update
  - service module [start/stop service]

**Demo - Web/DB/iptable configuration**
```bash
$ ansible-doc -l
$ ansible-doc yum [all yum related params and args]
```

- Installing the package using yum module

`ansible cfgmng -i exp/inventory_exp -m yum -a "name=vim state=present" --sudo`


Ex: `$ ansible webservers -i exp/inventory_exp -m yum -a "name=httpd state=present" --sudo`

Above command installs the httpd server if its not installed

- Before starting the service, let have look at "service" module
 
```bash
$ ansible-doc service
$ ansible webservers -i exp/inventory_exp -m service -a "name=httpd enabled=yes state=started" --sudo
``

- Install DB

```bash 
$ ansible dbservers -i exp/inventory_exp -m yum -a "name=mysql-server state=present" --sudo
$ ansible dbservers -i exp/inventory_exp -m service -a "name=mysqld state=started" --sudo
```

- Stopping the Iptables
`$ ansible webservers:dbservers -i exp/inventory_exp -m service -a "name=iptables state=stopped" --sudo`

***Note:*** what's up with `webservers:dbservers` pattern

Target Pattern

```
-OR (both the group)
  webservers:dbservers

-NOT (except dbservers)
  !:dbservers

-Wildcard
  web*.ex.com

-Regex
  (~web[0-9]+)

-AND (pretty complex not technically)
  (group1:&group2) - Here we are talking with server which are in both group NOT all node in the both group
  Make sense when you have server with ROLE as WEB & also in PROD

-Complex one
  webservers:&prod:!python3
  - webserver which are in PROD environment but NOT having python 3
```

**Demo using setup modules**

- Setup modules are called `FACT` modules, which brings the all information about the Operating System

`$ ansible cfgmng -i exp/inventory_exp -m setup`
  which retrieves the ultimate information about the system [like ohai in chef]

`$ ansible cfgmng -i exp/inventory_exp -m setup -a "filter=ansible_eth*"`
  which only returns the informataion about the Network interfaces

`$ ansible cfgmng -i exp/inventory_exp -m setup -a "filter=ansible_mount"`
  which gives the information about HDD & other mounts

`$ ansible all -i exp/inventory_exp -m setup --tree ./setup`
  all node information will dumped under setup directory

---

### 6. Plays and Playbook

Playbooks - Ultimate Automated tasks list

Basics - Kind of Glue brings all things together

`Playbook contains-> Multiple Play contains-> Multiple TASKS ->Target Group/set of Groups`

- Playbook execution happens top to bottom
- If any host fails in the task execution, it will not perform rest of the TASKS on the target server; Basically it will add the `FAILED` target to `RE_TRY` file/list
   - just by using the `--limit` option with ansible will ONLY execute against the FAILED nodes again
   - `MODULE.retry` file will be place on under HOME, this location can be configured in the CFG file.


**What playbook contains**
 1) Target list (may be host/group name)
 2) Actual Tasks
 3) Repeats the same

**Exec:**
```
  $ ansible-playbook playbook.yml
    - assumes that INVEMTORY files is already mentioned in the configuration file, else please pass the inventory file name with '-i' option
  $ ansible-playbook -i inventory playbook.yml
```

**First Playbook Demo**

below is my sample palybook

```
---
- hosts: webservers
  sudo: yes

  tasks:
  - name: Install Apache Server
    yum: name=httpd state=present

  - name: Start And Enable Apache
    service: name=httpd enable=yes state=started

- hosts: db
  sudo: yes

  tasks:
  - name: Install MySQL server
    yum: name=mysql-server state=present

  -name: Start MySQL server
   service: name=mysqld state=started

- hosts: webserver:db
  sudo: yes

  tasks:
  - name: Stop IPTABLE
    service: name=iptables state=stopped

```

**RHEL 7 based**

```
---
- hosts: webservers
  sudo: yes

  tasks:
  - name: Install Apache Server
    yum: name=httpd state=present

  - name: Start And Enable Apache
    service: name=httpd enabled=yes state=started

- hosts: db
  sudo: yes

  tasks:
  - name: Install MySQL server
    yum: name=mariadb-server state=present

  - name: Start MySQL server
    service: name=mariadb state=started

- hosts: webservers:db
  sudo: yes

  tasks:
  - name: Stop IPTABLE
    service: name=firewalld state=stopped
```

```
  $ ansible-playbook web_db.yaml
  $ ansible-playbook web_db.yaml --limit $web_db.retry
```
  - does only execution against the FAILED target servers


**Playbook and Logic and More**

**INCLUDE**
  - There is `include/include_vars` directive you call in playbook to point the some YAML file where it can look for TASK/VAR & etc - which gives the flexibility to breakdown the playbook

**REGISTER**
  - Using the one task output to another task, or you can make decision based on previous play execution by using `REGISTER` class
    - Using the register you can record the play output & you can reuse
    - Also can used to identify the success and failure of the previous task

**DEBUG**
  - `Debug` modules helps more for troubleshooting purpose
    - debug - accepts only two parameters `["msg", "var"]` which goes "full string msg to the user" & "output the result of the variable"
    - Playbook don't have to be static

**PROMPT**
  - `prompt` module will prompt for the any input that you wanna inject in the playbook to run, basically entering fields are become invisible (like password)

**HANDLERS**
  - Some called `handlers` does any changes/actions based on ANOTHER thing on the system that you have defined in the playbook [in chef 'only_if' "not_if"]
    - There are some condition on HANDLER tasks
      - It will executed after all task been executed/completed
      - And it will executed at only ONCE [No matter how many times you call]

**EXE CONDITION**
  - just by using the "when" clause [like in chef]

**COND CLAUSE based on OUTPUT**
  - ignoring the error and executing the playbook further

**TEMPLATED**
   - by using the `jinja2` engine place the you values based exec to target
     "sdfsd.j2 to file.cfg/any" [in chef erb -> to any files]

**Advanced Playbook usage**

- Here we are going to play around with tamplete(file copy j2), notify, handlers
- Will start editing the previously created playbook

- Create a directory called `template` we used this one place out j2 templates

- create a file name called `index.j2` with following cotent
  ```
    <html>
    <title>Ansibel</title>
    <body>
    <head>Hello from Ansible </head>
    <H1> well, Its's working </H1>
    <p> Yo ! {{ username }} !!  Wassup </p>
    </body>
    </html>
  ```

**Content: web_db.yaml**

  ```
  - hosts: webservers
    sudo: yes
    vars:               # -> these are the var/values that we are going to use in the http.conf.j2/index.j2 template files (as attribute to erb template in chef)
      http_port: 80
      doc_dir: /ansible/
      doc_root: /var/www/html
      max_clients: 5
  
    vars_prompt:        # -> Still you can promt anything while running the playbook & that values can used in the j2 template
      - name: username
        prompt: What's your Name, Yo ?
  
    tasks:
    - name: Install Apache Server
      yum: name=httpd state=present
      when ansible_os_family == "Redhat"      # -> Here we are telling that RUN this task only WHEN the os family is REDHAT
  
    - name: Start And Enable Apache
      service: name=httpd enabled=yes state=started
      when: ansible_os_family == "Redhat"
  
    - name: Copy the INDEX data
      template: src=template/index.j2 dest=dest={{ doc_root }}/index.html   # -> in order to copy the local template from the Ansible server to TARGET, also we user the VARS
      notify:                               # -> If this task's "changed=true" then notify the handlers to continue {here handler will restart the httpd}
        - Restart Apache                    # -> Spell correctly & its case sensitive (same as handlers name anyway)
  
    handlers:                           # -> This is the handler module, takes the action based on the above notify modules
      - name: Restart Apache
        service: name=httpd state=restarted
  
  - hosts: db
    sudo: yes
  
    tasks:
    - name: Install MySQL server
      yum: name=mariadb-server state=present
  
    - name: Start MySQL server
      service: name=mariadb state=started
  
  - hosts: webservers:db
    sudo: yes
  
    tasks:
    - name: Stop IPTABLE
      service: name=firewalld state=stopped
  ```

- let run the above ansible playbook
- Expectation it only supposed to copy the files & restart the apache(httpd), since the ENV already has the required packages & service up and running

```$ ansible-playbook web_db.yaml
PLAY RECAP *********************************************************************
exp-jslave401z             : ok=6    changed=2    unreachable=0    failed=0
```

---

### 7. Roles

- You can say "An abstracted function which is SYSTEM go perform - tags" - Confusing ! me tooo.
- Which helps to refactor the manage the PLAYBOOK
[For me looks this for grouping the tasks under ROLE name]
[In chef we have `include <somerecipename>` some thing like that]

**Role Basics**
 - ROLE has pre-defined directory structure, something like below

```
 ── roles
    ├── builders
    ├── server-common
    │   ├── defalts
    │   ├── files
    │   ├── handlers
    │   ├── meta
    │   ├── tasks
    │   ├── templates
    │   └── vars
    └── webservers
```
 - From above you no need to required all the directories, for ex if you dont use "handlers" you can remove & same as all servers
 - Now you no need "REFERNCE" on your play, ex: for copy files yoy no need of absolute path of the SOURCE cuz it automatically picked from the `files` directory
 - Every directory has the `main.yaml` where you have to place the content [in chef "default.rb"], and its not necessary that you have use `main.yaml` only
   - You can use some other files name to keep your date, call then using `include` module/directive [same like in chef]
   - BUT dont create other than `main.yaml` for VARS because its special & unique from the usage perspective

 - `main.yaml` is very go-to file for ansible
 - `site.yaml` file holds the your infra details & again you can have different files like and call them using the "include/directive"
 - `site.yaml` you can use `tag` to execute the playbook against that "TAG" 'ed infra

**PRE/POST tasks**
    You cam make use like, remove web server from SLB and so operation later put them back to the SLB
    Or During any operation you can silence the MONITORING alarm & bring it back once it done

 - Additionally TAG can played with "--limit" to execute against only the specific group

**Creating the ROLES**

- Role are should be non-interactive means that fully automated so there is NO `PROMPT` in here
- As told its required some nice DIR structure as follows, else create one

```
   $ mkdir -pv webserver/{vars,tasks,handlers,templates}
  
  exp
  ├── ansible.cfg
  ├── inventory_exp
  └── roles
      └── webserver
          ├── handlers
          ├── tasks
          ├── templates
          └── vars
```
 - create the main.yaml file in the all directories under "webserver"

```
  $ cd roles/webserver
  $ vim tasks/main.yaml

---
- name: Install Apache Server
  yum: name=httpd state=present

- name: Start And Enable Apache
  service: name=httpd enabled=yes state=started

- name: Copy the INDEX data
  template: src=index.j2 dest={{ doc_root }}/index.html
  notify:
    - Restart Apache
```

`$ vim vars/main.yaml`

```
---
http_port: 80
doc_dir: /ansible/
doc_root: /var/www/html
max_clients: 5
username: DamnOps
```

`$ vim handlers/main.yaml`

```
 ---
- name: Restart Apache
  service: name=httpd state=restarted
```

- now can copy the templates from the previous example to "templates" directory
 - now go the top directory where "roles" directory present & create the swebserver playbook
 
 `$ vim webserver.yaml # this files name can be anything`

  ```
  ---
  - hosts: webserver
    gather_facts: no
    become: yes
    roles:
    - webserver  # this supposed to be the directory name of the all stuff consists [tasks/templates/vars/handlers]
  ```
- Then execute the following command to run the playbook

 `$ ansible-playbook webserver.yaml`

- Now do the same dbserver

 `$ mkdir -pv roles/dbserver/tasks`

 `$ vim roles/dnserver/tasks/main.yaml`

  ```
  ---
   - name: Install MariaDB
     yum: name=mariadb-server state=present
  
   - name: Start MariaDB server
     service: name=mariadb state=started
  ```
- now create the dbserver playbook [where roles folder present]

 `$ vim dbserver.yaml`

```
---
 - hosts: db
   gather_facts: no
   become: yes
   roles:
   - dbserver
```

- Now execute the playbook
 `$ ansible-playbook dbserver`

- Now lets create the server-common playbook

 `$ mkdir -pv roles/server-common/tasks`

 `$ vim server-common/tasks/stuff.yaml`

```
---
- name: Install Stuff-1
  yum: name=vim state=present
```

 `$ vim server-common/tasks/sub.yaml`

```
---
- name: Install Stuff-2
  yum: name=screen state=present

- name: Install screen
  yum: name=mlocate state=present
```

 `$ vim server-common/tasks/main.yaml`

```
---
- include: stuff.yaml
- include: sub.yaml
```

- now create the site.yaml file, where roles directory present
  site.yaml will be apply/encompasses to all server

`$ vim site.yaml`

```
  ---
  #- hosts: all
  - hosts: webserver:db
    gather_facts: no
    become: yes
    roles:
    - server-common
  
  - include: webserver.yaml
  - include: dbserver.yaml
```

- then execute the following command

 `$ ansible-playbook site.yaml`

**Ansible Galaxy Intro**

- community based Playbook developemet
- It works like this -> Download -> Share -> Review
- installing/downloading the roles from Galaxy

`$ ansible-galaxy install username.rolename`

**Demo on the above**
- http://galaxy.ansible.com
- You can also do browse roles & also you can setup the filter to setup the search for the roles

- If you download any roles form Galaxy it stored under `/etc/ansible/roles` directory
  - something it will look like `/etc/ansible/roles/username.git`

**How we can use this roles then?**
 - cd ex-7/exp
   - Update the inventory files if required
   - now create the "something-new.yaml"
   - Add the following content

  ```
  ---
  - hosts: webserver
    sudo yes
    roles:
     - username.git   # -> this how you are calling that Roles here
    vars:
      variable: values  # -> Here we can override the values which already declared in the downloaded role
  ```
  - now you can execute the roles
  `$ ansible-playbook something-new.yaml`