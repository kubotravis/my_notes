Core AWS Administration
-----------------------

Warning: Content based on the 2012 

### 01 - Thinking about your server Presence on AWS

Key things what you should pick
  1. Pick your Region - Asia, Europe & Etc
  2. Pick Availability Zone - ua-west-1c (something like that - depends on Region)
  3. Provision you server - Select the Operation system and configuration -medium, large & etc
  4. Configure your service - web, db, media streaming & etc
  5. Expand to other availability zone
     - (i) Purpose: serve the same service to different region
     - (ii) Advantage : Reduce the latency & high avail too.

  6. Expand to other region - Giving the service to different regions

| Vertical scaling  | Horizondal Scaling |
| ------------- | ------------- |
| Traditional | AWS style  |
| Increasing the HW | Increase the instance |
| Increasing the Capacity | Shared Capacity |
| Easy to do | Typically requires planning |

In simple way - Instead of increasing the hw, simply increase the instances and setup loadbalancer for requests, Whenever there is less usage simply S-OUT some of the server.


### 02 - Getting started with AWS

Below are called core services of AWS - i can say it.

1. Cloudwatch
   - Its a monitoring mechanism.
   - You will get chanrged based on frequency of monitoring yourstuff


2. Simple DB/Dynamo DB
   - Simple DB - as named its small, slow and slowperformance, May it wanished/replace by any other tools by AWS in future
   - Dynamo DB - Good speed, performance(Ninja), you have pay more so..

3. Elastic Compute cloud (EC2)
   - Where you generate the servers. AWS term is Instance (AMI - Amazon Instance)
     - Eg, CentOS, Ubuntu, Win - Server instances
  
4. Simple storage Service (S3)
   - Its storage bucket, by deafult does the Replication of your instaces in the 2 different place/region.Unless its copy is done on the both place, then only it will return the success msg to you. S3 is quite slow.So when you run the instances it will not run from the S3 due to slowness.

   - Actually when you run the instance, it will copy the instance from the S3 to EPHEMERIAL storage, and it will run from there. When shutdown the AMI it will erase the from the EPHEMERIAL Storage and back to S3 as a normal form - so there is no permant storage what ever you worked.To take over this drawback AWS comes up EBS

   - <EPHEMERIAL Storage> - Its a mechanism, copied your VM image to Physical server & runs.

   - Its ninja speed but you will lose date when you shut down the AMI.

5. Elastic Block Storage (EBS)
   - Is Decent performance area, which is saving the changes that whatever you worked before you shutdown

6. Route 53
   - Its straight on DNS
   - You can create name record for your hosts/AMI, AWS will host/manages those.

7. Virtual Private Cloud (VPC)
   - By default AWS sets up the firewall/ip/ip to public & etc.If you are not ok with that you can manage those stuff from here.
   - E.g: you can setup the Site-to-Site VPN, make you feel you own it.

8. Auto Scaling
   - It adds the AMI's based on your threshold value. Saying that in a particaluar time create one AMI for me, or if the CPU reaches 89% add one more AMI for me like that.
   - On another word-it will allows to system to grow with your business

9. Cloud Formation
   - Allows you to define the templates of server
   - By default AWS does the design of AMI, but you create on yourself
   - Also you can make use of other's (AWS's subscriber) template to bring up your very own ENV, like Web,App & DB server AMI by some clicks.

10. Identity and Access Management (IAM)
     - Managing the Access over your AMI & others
     - its an unltimate Access management system with lot of stuff
     - Like defining algorithm, token mechinism, and blah blah blah

11. Elastic Load Balancing (ELB)
    - As named its Load Balancing

The free tier:
  - For more info http://aws.amazon.com/free


### 03 -  Creating an EC2 instances (AMI Selection)

**Key Concepts:**
1. Accessing the AWS Management Console
2. Choosing AMI (Amazon Machine Instance)
3. Understanding the EC2 pricing models

**Key Notes:**
  1. Sing-up and get the account & login to Management console
  2. Select your/nearby region
  3. Check for the availability zone
  4. AMI (Amazon Machine Instance)
     - A running server 
     - Basically build/developed by Amazon and frozened long time as a HDD format, You go and just kick that server, up & run.
  5. When creating an Instance, Amazon takes care of it License.
     - You no need to pay full license amount or you get paid by hours of usage, its nature way of AWS

**Steps-Better do it yourself**

- SignUp ->Management Console
  - -> Select Region->Launch Instance
  - -> Select AMI 'AWS' (AMI from Quick Start/MyAMI/Community) -> Continue
  - -> Enter the No.of Instance
  - -> Instance type (Micro Instance Doesnt charge-falls under Free Tier)
  - -> And Done

**Note** : There is something "Request Spot Instance" which will be convered in next chapter also about Core/CPU/Mem stuff too.

### 04 - Creating an EC2 instances (Pricing)

**Notes:**
1. On-Demand Instance pricing - No Commit
   - -> Just walkin into AWS and grab the Xtra large server and run  for hour and shutdown so will get paid for that hour only - No commit - just like that - most people used to do.
   - -> Note that per hour charge will change by region - it will vary
   - -> Also you may get charged if you are using other than ephimiral storage (S3/EBS will cost you)

2. Reserved Instance Pricing - Better than the On-Demand
   - -> You have to pay some down payment, so that you will end up with paying less money fo the AMI's
   - -> Down payment depends on the size(micro-small-medium-large-xtralarge)/utilization(light-medium-heavy)/region of AMI
   - -> Refer the AWS pricing - it has more info
   - -> A suggested model for the Horizontal scalling
   - -> Basically in here pricing model based on, how long youa are running in the AMI

3. Spot Instance Pricing
   - -> Its a almost like gambling
   - -> The left over CPU/MEM out bit by you for some dollars, if some one bit more than you then your AMI's goes down and theirs gets up
   - -> This is suitable for the Video streaming service , TV's and others
   - -> When you AMI goes due to your low bit, then what about date then ?  Still persistance
   - -> Basically one demand you scaling your anv based on Bid
   - -> Saying that great idea, but explore it more - From AWS site

Winning model - Reserved Instances

### 05 - Creating an EC2 Instances (Instance Type)

**Discuss:**
- Reading Instance type secification
- What's ECU's (EC2 Computing Unit)
- Factoring in I/O performance (NW/Disk Access speed)

Instances is mix of following (Basic Unit):
1. Memory
2. PRocesssing Unit (ECU)
3. Storage (EBS)
4. I/O performance - Basically depends on the 1,2,3

Instance Family 
- -> Micro
- -> Standard
- -> High Memory
- -> High-CPU
- -> Cluster Compute / 6 PU 

Basically Family based on the capacity of Basic  Units

Process Power Ratings
- -> Each Computig module capacity - 1.0 / 1.2 Ghz Xeon processors

How to calculate the you CPU capacity of each processor
- -> When the AWS says like this 

  ```
  8 EC2 Compute Units (4 Virtual core with 2 EC2 Compute Units each)
  So, here is the calculation
  8 * 1.2 Ghz = 9.6 Ghz
  9.6 / 4 = 2.4 Ghz
  ```

- -> Ans: We have quad core processor which each has 2.4 Ghz processing power

Understanding that I/O Rating
- -> Even if hold the high CPU/MEM but if you dont have proper Network bandwidth between your process and disk connnectivity - that is pointless
- -> Above case will worst when you page file are really big while running Windows server
- -> Consider the what kind of ethernet access you have between you EC2 AMI and the EBS HDD storage area - I/O Level
- -> You can also make use of the EBS to add multiple Disk on the single server and make out it has RAID disk

- -> Keep in mind the risk - Do the Benchmark or you can follow already done by other people on the AWS

**Notes:**
- -> Best Practice is start with Small, Benchmark and scale up if necessary

### 07 - Creating an EC2 Instances (Tags and Key Pairs)

Ok, We have up and running instances

**Tags:**
- -> What they are ?
  - Basically this are kay-value pair
  - Are like having the right naming convention.
    ```
    E.g
    Name: CORE-TS01
    Description : MyGateway Server
    Date : 12/12/12
    ```

**Key Pair:**
Basic Concepts
1. Public/Private key cryptography
2. How you can use it Windows AMI's
3. How to use it on Linux AMI's
4. What if I lose these keys

- -> AWS style of Private/Public key pair mechanism
  - While creating the AMI's you define the Key, and you will get download option on the Private key - Means that as you generate the AMI's you will get the Private Key
- -> ON Windows
  - Once you created the Windows AMI, you need to get the default Admin password the downloaded private key (adasd.pem), How ?
  - Right click on the running windows `AMI -> Get password -> export the downloaded key -> RDP to the windows server -> Reset pasword` (That's easy right...)
  - Baically we exported the Private key then you got the User/Pass to login
  -> On Linux
    - In linux soley you can login by using the Private key - Linux has that capability (NO need to export and get user/pass to login), so How ?
      - 1. Download the AWS command line tool and genearate the AMI keypairs
      - 2. Using that key just login
    - Command line (AWS command line tool - which is used to see all you setup instead of using the Web Interface of AWS management tool)
    - `ec2-add-keypair test-keypair `
      - Above command to generate the keypair using the AWS CMD tool
    - `ssh -i id_rsa-test-keypair root@ec2-xxx-xx-xx-x-blah.amazonaws.comn`
      - -> Above command which is used login into the AMI, Actually here we are not import/exporting any keys-just login to AMI
- Lose key ?
  - On windows/Linux - That is it you are locked out of your AMI - Amazon can't to do anything, So you have to terminate the AMI.
  - But you dont want to lose the DATA, better you can create your own AMI image of it, and create new instance of it.
  - If you remember, while creating the AMI, saying that "My AMI"; over there your backed up AMI image will be listed

**FYI :**
How the secure communication works in the realtime ?
Ex:
- When you connect to Bank site through your browser, they will send the Public Key to you.
- On the browser side, It creates the Session Key, This session key gonna be ecnryption/de-cryption key for the whole communication between Bank/Your browser.
- This is session key got encrypted by the Public key which is generated by the bank.
- Basically Public key is the half of the algorithm/math calulation, You can only encrypt the data cant de-cryprt.
- On the bank they decypt using the their Privatkey
- Private key has the only capability to do de-cryption
- This is how almost all secure channel communication works

### 08- Creating an EC3 Instance (Security Groups)
**Key Concepts:**
- Inbound Filtering for your AMI's
- `"Security Groups"`
- Rules - Inbound/outbound/Within the group
- Can't change the Security Group (Rules - Without VPC) - I guess u can now

- On other words, its a building an Firewall for your AWS

**My Understanding so FAR**
- Security Group is the one of the step you have to fill when you are creating an AMI (Last question/step berfore getting in got Review)
- It's a setting which is applied to an group
- What is Group ? - Is multiple no.of AMI's, also one only AMI can be a part of group
- Its not changeable once its applied to an group (Without VPC-w'll talk about later)
- If you want you can create your brand new custom rules, that will appear all the time while creating an AMI's

- The options you will see while creating an AMI at Security Group tab
  - YOUR Custome Security Group (Set of rules) 
  - Default (Only Outbound no Inbound but Each AMI can communicate fully)
  - Linux Default (SHH) 
  - Quick-start-1 ~> This is the default security group for windows server which is assigned by AWS

- Cannot Modify, What ? (IN terms of Security Group)
  - That means you can update rule which already defined, but you can simply delete and add which you want - little pain

- Can't change, what ? (IN terms of AMI)
  - Once you have selected any Security Group to any AMI and it up and running, you can't change the security group for that AMI unless if you are inside the VPC

- What you should aware ! (IN terms of Security Group)
  - If you are creating a Security group for each AMI's, that means you are more than one Security gruoup
  - You know that AMI's inside an one single Security group can communicate each other but not outside, so you have create a seperate rule in each Security group to communicate other security groups
  
- So, How can make security group communicate each other ?
  - Imagine that you have 2 security group, naming that `TEST-1 &TEST-2`
  - In order to make communicate each other, select any one of the group (TEST-1) and got edit
  - Add custom rule forInbound in ALL TCP
  - Select the "Source" as TEST-2 (it will get updated by its groupid)
  - And save/apply the changes to take effect

**Real time eaxmaples, how to design ?**
- Imagine that you are serving an web server for the world.
- Those web servers are connected to DB server for saving the transactions
- So in the above case DB servers are suposed to be in different security group, to be apart from the Web servers
- There will rule on Web/DB sevrer side to allow only inbount connection to the DB, not the outbound
- Main purpose is that, we are not expose the DB server to anybody apart from Webservers

**Notes:**
- "world of Worldcraft" server port configuration - check that out
- You should have the knowledge about Subnetmasking too