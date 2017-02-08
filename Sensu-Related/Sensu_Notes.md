All about Sensu Monitoring
--------------------------

### 01. Architecture
- Runs the Checks (Internal & External - like Nagios)
- It can also schedule service check on periodic time (Its called Standalone method) - (In this case totally independent)
- So sensu client executes the checks & Passes the result to the RabbitMq on the Sensu server
  [Even if the Sensu API dead, all your results stays in the RabbitMQ - cool]
- Now, Sensu takes the result from RMQ & operates them. Based the result runs the Handlers
- Handlers are basically scripts (user defined, like send mail, PD or etc)
- Sensu server can request for check on client, so the request registered on the RMQ. so the request picked up by Sensu client & response back. Above scenarios are called subscription check.
- All component in sensu is stateless, so the persistence data will be stored in Redis internally
- Now SensuAPI
  - Its talks with Redis & RMQ and provides required data in a format of REST-API
- So that Sensu-cli/Dashboard will utilize the API for anypther purpose

### 02. Other Features
- Its not required to maintained (new/old) client configuration on Sensu server
- All the clients can register itself to sensu server ON-THE-FLY
- For instance when the apache package been installed on the server & meantime you can configure the Apache related checks on the server. So that it automatically register with Sensu server.

**How enable the NEW/DYNAMIC checks on the Sensu server side and pass it to client & start Checks with restarting the client component ?**
Which means that there is no such CHECKS configured on the Sensu client but still we can execute the desired checks on the FLY !

How? - by writing new META-CHECK (may be cron job) and sending the response to the Sensu-client itself & then it will forward it to the Sensu Server

[Basicall JSON report sent localhost port 3030 & it gets forward to the Sensu server via Sensu client]


### 03. Installation

**Platform: Ubuntu**
- [Installing the Redis](https://sensuapp.org/docs/0.26/installation/install-redis-on-ubuntu-debian.html)
- [Installation of RabbitMQ](https://sensuapp.org/docs/0.26/installation/install-rabbitmq-on-ubuntu-debian.html)
- [Installation of Sensu API & Server](https://sensuapp.org/docs/0.26/platforms/sensu-on-ubuntu-debian.html#sensu-core)
- config.json
  This file basically referred by Sensu for where to look Redis/RabbitMQ connection detail
- config_memory.json # this is check configuration, leave it for a while
- default_handler.json # this is a handler configuration, just cat 'ing the actual output. will be covered in detail later.

### 04. Installing the Sensu Client
- Theres is no specific package for client (Sensu API/Server/Client all comes together)
- Also all comes together in the omnibus package
- Once after the installation create the `client.json` file under `/etc/sensu/conf.d` directory
  - where it has its `name`, `ipaddress` of the RMQ server `IP/port/user/password/vhost` value
- configure the check_mem.sh file under /etc/sensu/plugins to run the checks
  (we are doing above because we made a subscription while setting up the Sensu server, so this client will respond to that)

### 05. Adding the handlers
- Default handler will simply cat the piped content but you can configure the MAIL handler to send mail
- You can install through gem
```
  $ /opt/sensu/embedded/bin/gem install <plugin_name>
  $ /opt/sensu/embedded/bin/gem install contents <plugin_name> # will display all the content that installed
```
- Fill the required information (in JSON) file pass ARG to handler will take care the rest

### 06. Adding the Real Check
- Here we are kind of migrating the Nagios service checks to the sensu client side
- Installing the nagios pluing
  `$ apt-get install nagios-plugins-basic`
- All plugin will bw available on `/usr/lib/nagios/plugins/check_*`
- write the new configuration file for check_disk (as JSON dictionary)

```json
  {
    "checks": {
      "check_slash": {
        "command": "/usr/lib/nagios/plugins/check_disk -c 20% /",
        "standalone": true,
        "interval": 10
      }
    }
  }
```

- Check the above config syntax as ho NO problem (by using jq)
  `jq . file-name.json`
- Now restart the sensu client & tail the logs

**Note:-**

Every monitoring system has some sort of/ or its own queue system.
Sensu don't want to re-invent the wheel so using awesome `RabbitMQ` for that.

`jq` - tool for 'pretty' view of the JSON

**Client configuration**

- Below is to connect and send information to RabbitMQ queue fron the client node

  `$ vim /etc/sensu/conf.d/rabbitmq.json`

```json
  {
    "rabbitmq": {
      "host": "10.0.0.2",
      "port": 5672,
      "vhost": "/sensu",
      "user": "sensu",
      "password": "secret"
    }
  }
```

- This is actual information about client node

  `$ vim /etc/sensu/conf.d/client.json`

```json
  {
    "client": {
      "name": "exp-client502z",
      "address": "10.0.0.1",
      "environment": "exp",
      "subscriptions": [
        "dev",
        "rhel"
      ]
      }
  }
```
