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

How to find out the Netwirk address in case of any class of IPv4 address ? ip is `192.34.56.67`

Answer:
Since its Class C IPv4. so first three bit goes for NW bit & rest goes for HOST
Ip `192.34.56.0` -> `0` is the starting of the HOST IP here, So this is Network IP
it must belongs to any of network devices (like Router/Switch)

Commonly used IPv4 Private Address

```
Class A: 10.x.x.x
Class B: 172.16-31.x.x
Class C: 192.168.x.x
```

What is the Purpose of the Private Address ?

Imagine any company setup network with `8.x.x.x` & some of other device got the IP from it.
When ever you make query to Google public DNS `8.8.8.8`, it will NEVER match or get response.
Because the PC thinks it local & it never go for DEFAULT GATEWAY to resolve the `8.8.8.8`
So in-order to avoid such collision we had to have the PRIVATE IPs as listed above.

### 4 Beautiful Binary

#### Questions:
1. What is the binary values of 4th position (right to left) ?
2. Are there only 10 types of people regarding binary ?
3. How many place holders do we need to be concerned with ?

#### Answers:
1. 8
2. Gotcha - this is Binary joke
3. I didn't get the question at all !

**Notes**

Base10 Numbering  ->  Decimal
   - Here numbering system goes like below, (from left to right)

```
ex:
45,1234 ->
... 10000  <- 1000 <- 100 <- 10 <- 1
```

Base2 numbering - 1 & 0 -> Binary
  - here numbering system goes like below, (from left to right)

```
... 128 <- 64 <- 32 <- 16 <- 8 <- 4 <- 2 <- 1
```

### 5 Decimal to Binary Conversion

#### Questions:
1. What is the Decimal equivalent of 01011010 ?
2. What is the binary equivalent of 130 ?
3. What is the binary equivalent of 255 ?

#### Answer:
Simple, do it yourself !

**Notes**
Octet = group of 8 bits
8bits = 1 Byte

#### "Does it fit" game

If you want convert any number into binary, write down the below

`128 | 64 | 32 | 16 | 8 | 4 | 2 | 1`  & ask that number does it fit in ?

 ex: take a number 130 
  - It will fit in in 128 & take reminder 2 & 2 only will fit in 2.
  - Now mark as 1 which ever FIT in & make rest as 0
  - This is how you will be the Decimal to Binary conversion
 so,
 
 ```
 130 -> 10000010
 193 -> 11000001
 17 -> 00010001
 95 -> 01011111
 136 -> 10001000
```

#### Binary to Decimal conversion
- This is reverse of the above function
- Get the binary number & map it as below

```
 1      0     1     1    1    0    0    0
128 <- 64 <- 32 <- 16 <- 8 <- 4 <- 2 <- 1
```

Add numbers which  got all ONEs

so above binary becomes,
`10111000 - 184`

### 6 The Mask Unveiled

#### Questions:
1. PC has the IPv4 15.69.35.28 with the mask of 255.255.0.0
   What is the network this computer is connected to ?
2. How many bits from the IP are being used to represent the "network" ?
3. What is the significance of a bit being "on" in a mask ?
4. Do the Mask bits, going from left to right, have to be "on" in a order ?
5. What is the purpose of Mask ?

**Notes:**
Purpose Mask: Actually it identifies IPv4 address as which portion is part of "network" & which part is "Host"

  - IP - 32 bits of address
  - Mask - 32 bits of address

Mask/Network Mask/Subnet Mask - all are same simply naming convention differs

How they are saying its 24 bit network (something like this 192.168.1.0/24)

**Ex:**

The IPv4 is    : `192.168.1.15`
and its Binary for each decimal:   `11000000 10101000 00000001 00001111`

Mask is : `255.255.255.0`
Binary: `11111111 11111111 11111111 00000000`

If you MULTIPLY the IPv4 binary with Mask you get the follow

```
IPv4 11000000 10101000 00000001 00001111
Mask 11111111 11111111 11111111 00000000
----------------------------------------
     11000000 10101000 00000001 00000000   -> 192.168.1.0
```
   - Since each deciam has 8bits so 3*8=24, since last bit considered as HOST so you cant include it for calculating the netwowrk
     (Almost 3 decimal same as IP). so it becomes 192.168.1.0/24

 - Here is the another question
  ```
  Ip : 192.168.1.15
  Mask: 255.255.0.0
  ```
  
- Now what is the NetworkIp or network address ?
  Its gonna be : `192.168.0.0` 
  ( Possibly from given mask first 2 field goes for Network & 2 for HOST, so it becomes `192.168.0.0/16`)

- One more question
```
 Ip : 192.168.1.15
 Mask: 255.0.0.0
```
  - Now what is the Network IP or network address ?
    Its gonna be : `192.0.0.0` ( Possibly from given mask first 3 field goes for Network bit & 1 for HOST, so it becomes `192.0.0.0/8`)

Vid-7 Stealing the Host Bits
---

#### Questions:
1. How many host bits would be sacrificed to network addressing to create exactly 16 new subnetworks ? 4

2. In the finger game, When you start with a 2 and verbally double the number with each new digit raised, does the number you speak represents how many new subnets you can create, and the digits in the air represent the number of host bits sacrificed ?
3. What is the new mask, in the octet where 2 previous host bits are being converted over to network addressing ?


#### Answers:
1. 4 host bits required for 16 new subnets
2. Raised fingers are - HOST bit sacrifies count

Number are - New Subnet

3. 192 ( refer the Ref section)


**Notes:**

**Task:**
 - There is given network topology (Topology contains 3 different subnet & 6 routers - As per the diagram in the give excercise)
 - Now you boss want to have 192.168.5.0/24 network have access to over all topology

Things are required:-
 1. Identify No.of required subnet
 2. Use the Finger Game
 3. Identify no.of Host bits needs to be sacrificed
 4. Let Everyone know (Meant devices in the network)

**Answers:**
1. Identify No.of required subnet -> is 2 (As per the diagram in the give excercise)
2. Finger Game: (how many host bit required to create 3 subnets)

mark `2` in the THUMB finger.

Release other finger as how many subnet you want. For example if you are releasing POINTER & MIDDLE finger next.

`2 * 2 = 4` mark INDEX finder values (Thumb finger values * Verbally same number)
`4 * 2 = 8` mark MIDDLE finger vales (Previous finger values * Thumb finger values)
... (If you the same as above for rest other 2 fingers)

`8 * 2 = 16` RING finger
`16 * 2 = 32` LIITLE finger

Which means that:-

- So all FINGERs got values (They are UP), -> If you took 5 bits from HOST & Allocate it to NETWORK bit then you can have 32 brand new SUBNETs
- In that case we have to SACRIFICE 2 bits from HOST bits to create 4 subnets ( actually we are going to create only 3)
- So those sacrificed bits from HOST & can NOT be re-used in host bits at all

3. From the above 2 bits needs to be sacrificed from the host bit

4. This only done by modifying the MASK, so the given -> `192.168.5.0/24`

  - After `2` bits goes from HOST to Network, i gonna be -> `192.168.5.0/26` (so the host bit gonna only have `6` bits but we got `2` more subnet [Network address space])

  - since from the above we took 2 bits from HOST, so -> `192.168.5.0/26`

  - it becomes
  ```
  ip -> 192.168.5.0
  mask -> 255.255.255.192 (This calculated based on below Ref section)
  ```

 **Ref:**

 ```
 Mask Values: 128 192 224 240 248 252 254 255
 Weights:     128 64  32  16  8   4   2   1
```

 **Explanation:**
  a) write down the weights first
  b) since 2 bits were taken from host & so mark the first 2 decimal as 1
  c) And then add the all marked ones

```
a) ->  128 64  32  16  8   4   2   1
b) ->  1    1   -   -  -   -   -   -
c) ->  128+64 = 192  (so then mask is 255.255.255.192) - (same value mentioned in the MASK values)
```

**Tasks**
How many host bit you going to steal to create the 42 subnets?

  - Use the finger game:-
  - As per the finger game we required `6` (`6` finger UP) host bits to create the `42` subnet (Actually in `6` finger maximum we can create the `64` subnets)
  - So if the given is -> `192.168.0.1/24`

  - Ans goes like this
      ```
      192.168.0.1/30
      mask is 255.255.255.255.252
      ```