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


### 03 -  Creating an EC2 instances - AMI Selection

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

### 04 - Creating an EC2 instances - Pricing

**Notes:**
1. On-Demand Instance pricing - No Commit
   -> Just walkin into AWS and grab the Xtra large server and run  for hour and shutdown so will get paid for that hour only - No commit - just like that - most people used to do.
   -> Note that per hour charge will change by region - it will vary
   -> Also you may get charged if you are using other than ephimiral storage (S3/EBS will cost you)

2. Reserved Instance Princing - Better than the On-Demand
   -> You have to pay some down payment, so that you will end up with paying less money fo the AMI's
   -> Down payment depends on the size(micro-small-medium-large-xtralarge)/utilization(light-medium-heavy)/region of AMI
   -> Refer the AWS pricing - it has more info
   -> A suggested model for the Horizontal scalling
   -> Basically in here pricing model based on, how long youa are running in the AMI

3. Spot Instance Pricing
   -> Its a almost like gambling
   -> The left over CPU/MEM out bit by you for some dollars, if some one bit more than you then your AMI's goes down and theirs gets up
   -> This is suitable for the Video streaming service , TV's and others
   -> When you AMI goes due to your low bit, then what about date then ?  Still persistance
   -> Basically one demand you scaling your anv based on Bid
   -> Saying that great idea, but explore it more - From AWS site

Winning model - Reserved Instances
