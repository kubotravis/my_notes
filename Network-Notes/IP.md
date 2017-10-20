IPv4 SubNetting
---

### 1 Introduction
Refer the Wikipedia

### 2 Fun with IPv4 Basics

#### Questions
  1. How many periods (dots) are in the dotted decimal IPv4 address ?
  2. What are the two parts of IPv4 address
  3. What determines the  dividing the  line between the two parts of an IPv4 address ?
  4. How do you see the current IPv4 address you are using it right now ?

#### Answer
3. MASK does(determines) the divide.
  ```
  ex:
  ip: 192.168.1.5
  mask: 255.255.255.0
  ```
so the first 3 bit are belongs to NEWORK & rest goes for HOST

   - Detailed explanation will be there later section
   - This guy says only function of MASK to divide the network/host bits from Actual ip - new stuff

**Other Note**
   - Getting Dynamic IP - this thing is done by "Dynamic Host Protocol"

### 3 Classes, Masks and Private IPs

#### Questions
1. What the class of `191.192.127.99` ?
2. What is the default mask of the `175.32.66.123` ?
3. What is the range for Class B private address ?
4. Is it really matter knowing the class of IP ?

#### Answers
4. Because of Default mask (sort of identifier) - which tells (or controls) that IP address going front and back
  - for Class A, the very first bit is NW bit since the default Mask is `255.0.0.0`
  - for Class B, the first two bit is NW bit since the default Mask is `255.255.0.0`
  - for Class A, the first three bit is NW bit since the default Mask is `255.255.255.0`

#### IP Range of Class
  - Class A (1-126)
  - Class B (128-191)
  - Class C (192-223)

- `127` belongs to no class (Basically goes to localhost), Its reserved for loopback interface

How to find out the Netwirk address in case of any class of IPv4 address ?
IP is  `192.34.56.67`

Ans:

Since its Class C IPv4. so first three bit goes for NW bit & rest goes for HOST
Ip 192.34.56.0 -> 0 is the starting of the HOST IP here, So this is Network IP
it must belongs to any of network devices (like Router/Switch)

Commonly used IPv4 Private Address
Class A: 10.x.x.x
Class B: 172.16-31.x.x
Class C: 192.168.x.x

What is the Purpose of the Private Address ?
Imagine any company setup network with `8.x.x.x` & some of other device got the IP from it.
When ever you make query to Google public DNS `8.8.8.8`, it will NEVER match or get response.
Because the PC thinks it local & it never go for DEFAULT GATEWAY to resolve the `8.8.8.8`
So in-order to avoid such collision we had to have the PRIVATE IPs as listed above