All about Docker Cli usages
---

#### Intro
Docker Commands that i understood and remember

Some of the very Basic information Docker Engine that i've missed earlier, let me list it here the

Topic:
Part #1 Hello World in Container
Part #2 Running a simple application

- As of now on CentOS 7 , just install the Docker Engine from the official repository
- And add any user to the group 'docker', so user can run commands over Docker

- `$ docker version`
which displays the client/server information
 - 'docker' is a client side command line utility which calls the Docker Server/API to get the job DONE

- `$ docker run ubuntu /bin/echo "Hello World"`
 which runs the 'ubuntu' image as container, as result it displays the "Hello World" and dies asap

- `$ docker run -t -i ubuntu /bin/bash`
 Run the Ubuntu container, -t assign the psuedo tty, -i give that tty as INTERACTIVE with bash shell

- `$ docker run -d ubuntu:latest /bin/bash -c "while true; do echo Hello World; sleep 1; done` 
-d daemonized it - will be running in the background for all the time

- `$ docker ps`
 All running Docker Container information - that 7 Column info

- `$ docker ps -l`
Give me the recently ran Docker container information, `-l` last container that started.

- `$ docker ps -a`
 Give all running/Stopped Docker container informtion

- `$ docker log <contain_name/id>`
 Which prints the All application log from the app container, ex: hello world on literally prints all echo output

- `$ docker stop <container_name>`
 It just stops the running container, just STOP

- `$ docker start <container_name>`
 It just starts the application which is stopped above

- `$ docker run -d -P training/webapp python app.py`
 `-P` which expose the some random port on Localhost for the container application

- `$ docker run -d -p 80:5000 training/webapp python app.py`
80 - localhost port, 5000 - Container application port

Here we are manually assigning the port number to the container application

- `$ docker port nostalgic_morse 5000`
which will give the more information about the localhost port which mapped by `5000`

- `$ docker logs -f nostalgic_morse`
 `-f` is equivalant to `<tail -f>`, which is literally tails the application log from the Container app

- `$ docker top <container_name>`
 We can see all process which is running inside the application container

- `$ docker inspect <container_name>`
 
 - ex:`$ docker inspect nostalgic_morse`
  which gives the very low level information about the Running container
  still we narrow down some specific required values from the application container
 - `$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nostalgic_morse`

- `$ docker rm <container-name>`
 basically its not possible remove container which is running, so we have to stop it first: still its possible to remove the running container by `docker rm -f'` flag

Notes:
I have to explore the option, that changing the port mapping on demand

#### Playing with Images
- May some of the commands are repeated here

- `$ docker images`
 which displays the Docker image information

- `$ docker search <image_name>`
 which does the search against the DockerHub for the images

- `$ docker pull <image_name>:<tag>`
which downloads the image from the Docker Hub
 ex:

```
$ docker pull ubuntu:latest
$ docker pull ubuntu:14.04
```

- `$ docker run -t -i ubuntu:14:04 /bin/bash`
 Running a specific version of Image as a container

#### Creating your own image
  - 1. With minor change
  - 2. By using Dockerfile

 1~>
    - Pull any image
    - Install something
    - Note the hostname somewhere which os required while committing the changes [root@0b2616b0e5a8]
    - 'Exit' out of the container
    
     - Run the following command then

      `$ docker commit -m "Added json gem" -a "Kate Smith" 0b2616b0e5a8 ouruser/sinatra:v2`
      `$ docker images`

2~>
    - create some directory where do you wanna carry the operation
    - And create the create the 'Dockerfile' and give the  instructions or steps which needs t be performed

    `$ mkdir test && cd test && touch Dockerfile && vim Dockerfile`

    - Append the following content

  ```
    # This is a comment
    FROM ubuntu:14.04
    MAINTAINER Kate Smith <ksmith@example.com>
    RUN export http_proxy=http://pkg.proxy.prod.jp.local:10080/ && apt-get update && apt-get install -y ruby ruby-dev
    RUN export http_proxy=http://pkg.proxy.prod.jp.local:10080/ && gem install sinatra
  ```

    - Run the following command to start the build from it

    `$ docker build -t test/name:v2 .`
        then watch for PASS/FAIL


 As image got created, set up the TAG for it
 
  `$ docker tag 5db5f8471261 ouruser/sinatra:devel`
  `$ docker images`

- There is some talk about 'digests' that i dont understand as of now [content-addressable identifier called a digest]
  `$ docker images --digests | head`

  Pushing the images to Docker Registry - which you don't care about it..!

  Removing the images locally
  `$ docker rmi training/sinatra`
  `$ docker rmi -f training/sinatra` [with force]

#### 3 Docker Networking

`$ docker run -d -P --name web training/webapp python app.py`
- above command used to naming the container
- Container name must be unique

`$ docker ps -l`
`$ docker inspect web`
 - inspecting the container by using the NAME
`$ docker stop web`
 - STOP the container by using NAME
`$ docker rm web`
 - REMOVE the container by NAME

if you want to use the same name for other container you should remove the old one in order to create the NEW one

Basically Docker Networking is Advanced topic

- By default docker provides 3 different networks/drivers
  bridge, host, none

`$ docker network ls`

`$ docker network inspect bridge`
 Inspecting the bridge network drivers

- Run a container & remove it from the network [which is assigned defualt]
`$ docker run -itd --name=networktest ubuntu`
`$ docker network inspect bridge`
`$ docker network disconnect bridge networktest`

- creating new BRIDGE network using bridge drivers
`$ docker network create -d bridge my-bridge-network`

Add containers to a network
- in-order to attach network in Container, container should be running first
`$ docker run -d --net=my-bridge-network --name db training/postgres`
Above command attaches the Container to the "=my-bridge-network" network drive

`$ docker inspect --format='{{json .NetworkSettings.Networks}}'  db`
 Inspecting the "db" container

`$ docker run -d --name web training/webapp python app.py`
`$ docker inspect --format='{{json .NetworkSettings.Networks}}'  web`

`$ docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' web`
 Getting the IP of the container

`$ docker exec -it db bash`
 Now, open a shell to your running db container
from here if you are trying to ping the WEB contner u'll get 100% of packet loss
To avoid the above scenario u can attach the 'my-bridge-network' to 'web' container by using the following command
 `$ docker network connect my-bridge-network web`

#### 4 Managing data in Container

There are two primary ways,
- Data volumes
- Data volume containers

 ~> Data volumes
 These are the special-designed directory
  Simple way
    - A random or specific directory from Docker host been getting mounted inside the docker container [by bypassing the UnionFS]
    - Its persistent & independent storage volumes, regardless of container life cycle
    - Same volumes can be shared via many App container
    - It get initiated when container gets created
    - It cant be removed by any container based operation, even if the container itself deleted
    - By default it mounts with the RW mode


Adding Random volume to container
`$ docker run -d -P --name web -v /webapp training/webapp python app.py`
  - This will create a new volume inside a container at /webapp
  - if the directory /webapp already exists & some data is available then it does overlay & get the HOST content into the CONTAINER
  - It will good to specify the ABSOLUTE path on the Docker engine, else it will create volume for the NAME that you have supplied
  - MUST/Always use ABSOLUTE path for the container specific one [meant container folder]


Inspecting
`$ docker inspect web`  check the MOUNTS area

Mounting the Host directory as data volume
`$ docker run -d -P --name web -v /src/webapp:/opt/webapp training/webapp python app.py`

  - in the both place if the directory doesn't exists, its crates the directory for you
  - By default it mounts the host volume with 'rw' permission
  ~> Use blow command to mounts it has 'ro' only
  `$ docker run -d -P --name web -v /src/webapp:/opt/webapp:ro training/webapp python app.py`

  There is a STORY if you use MAC/WIN32 based you just shared/mount the directory like Linux
  Only /c/User & /Users directory in WIN32 & MAC can be mounted respectively

Shared storages can be mounted [NFS/iSCSI/FC] inside the containers by using the plugins in Docker & Container side [Citation required]

~~> by using the Volume Drivers
 which means mount the Docker host directory with additional driver information
 Its like creating the volume using dd command and format & use it

##### Warning
In Dockerfile, you can't use the HOST directory to mount inside the container as specifies above.
because built images should be portable. A host directory wouldn’t be available on all potential hosts

Mounting some DATA volume (instead of host directory)
 `$ docker run -d -P --name web -v foo:/tmp/foo training/webapp python app.py`
  - foo: represents the volume name

 `$ docker run -d -P --name web -v foo:/tmp/foo:ro training/webapp python app.py`
  - Here we mounted the volume as RO (Read-Only)

Mount a host file as a data volume
 `$ docker run --rm -it -v ~/.bash_history:/.bash_history ubuntu /bin/bash`
  - now type command on the Container that will get stored in the .bash_history & exit out of the container
    You can visualize command that ran in the container form HOST machine

Usage:
Instant code checking if you application reacts weird

Creating and mounting a data volume container

Basically its simple, if you dont specify the ABSOLUTE file/directry from the Docker host while mounting/creating the container, Docker simply created NAME volume that you supplied.

ex:
  `$ docker run -it -d --name test -v foo:/bar busybox`
  - From the above example Docker host side volume name is "foo" (it supposed to /foo), in case docker creatde the VOLUME with that name, verify using the following command
  `$ docker volume ls`
  `$ docerk volume inspect foo`

From Docker man pages

  `$ docker create -v /dbdata --name dbdata training/postgres /bin/true`
  - This command creates the "dbdata" container with some volume name & mounts & dies

  `$ docker run -d --volumes-from dbdata --name db1 training/postgres`
   - this command re-uses the above first created VOLUME fro the old container (dbdat) to new "db1" container
  `$ docker run -d --volumes-from dbdata --name db2 training/postgres`

  `$ docker run -d --volume-from db1 --name db3 training/postgres`
  - this so called "extend the chain of the mount", that referring the db1 creating the db3,
    In this case also, all are container pointing to the same DATA VOLUME from the Docker host

~> So you can use as many as volume to any container (means that "---volumes-from") from any container

~> interesting:-
    If the postgres image contained a directory called /dbdata then mounting the volumes from the dbdata container hides the /dbdata files from the postgres image.
    The result is only the files from the dbdata container are visible

Backup, restore, or migrate data volumes

Backup:-
  `$ docker run --volumes-from dbdata -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar /dbdata`
    1) above command creates the container
    2) attaches the volume from "dbdata" container
    3) mounts the Docker host local directory $(pwd) to the container in "backup" directory
    4) takes the backup of "dbdata" content to the "/backup" directory which is already mounted with Docker host (so you can get the data locally)

Restore:-
  `$ docker run --volumes-from dbstore2 -v $(pwd):/backup ubuntu bash -c "cd /dbdata && tar xvf /backup/backup.tar"`
   1) create a container by attaching the volume from "dbstore2" container
   2) mount the local directory to "/backup" dir in container
   3) get inside the "/dbdata" dir in container & extract the "backup.tar" in it [Commandline section]

So you use "Backup & Restore" for automated backup functionalities

This below part required some study; during the middle of learning i could find this content anywhere in the Docker manual (Might they removed since its required lot of stuff)

 `$ docker volume create -d flocker --name my-named-volume -o size=20GB`
 Which creates the volume on the Docker host

 ```$ docker run -d -P \
  -v my-named-volume:/opt/webapp \
  --name web training/webapp python app.py
  ```
 Above command utilizes the volumes which is created previously

 ```
 $ docker run -d -P \
  --volume-driver=flocker \
  -v my-named-volume:/opt/webapp \
  --name web training/webapp python app.py
  ```
   Above command creates the on the fly over Container creation

Warning:
 - All container writing into same volume might cause the Data corruption
 - Since Data volumes are available at Docker as filesystem, make sure you dont do any operation on it, bcuz container have no chance to AWARE you docker host function
 -
ERROR:
   Intalling the driver 'flocker' fails due saying that there is no flocker plugin installed
   Follow the steps to install 'flocker' plugin

~> Installing the DRIVER based volume become super flop - so please wait... !

