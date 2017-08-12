Python Basics
---
Minimal Info can found here.

**Intro**
Get it from internet or some where.

**Note**
Be alert on INDENTATION

**Basic Math**

```python
>>> print 1 + 1  # This is the reason its called loosely type
2
>>> print 1 * 3 (Mul)
>>> print 4 / 2 (Divide)
>>> print 2**16 (Exponentional - 2 Power 16)
>>> print 10 - 5 (Sub)
```

**Strings**
(Double quotes)
```python
>>> print "Hello World"
Hello World  # Since python is loosely coupled, so yo no need to specify the new line character (\n) but still it prints the new line
>>> print "Hello World\n"    # Here it prints extra new line

>>>
```

(Single quote) - Python interpreter is pretty smart to handle here
```python
>>> print 'Hello\nWorld'   # It doesn't print the "\n" at all, it evaluates instead.
Hello
World
```

(Help - stuff)
```python
>>> help() # Gives the help information
help> keywords
help> module spam # Searching the documentation for the module
help> quit  # Exit out from the help interpreter
```

**Variables**

```python
>>> print 200 / 4 * 5
160
```

```python
name = "Python"   # This how assigns VAR/VALUE combination
```

 - You don't required any `"$"` like in perl to define the variable value
 - "=" is the assignment operator here
 - We have used String so we have to keep it under `""` much like in all other programming language

**Data Types**

- String, Int, Float, Lists, Tuples, Dictionary
- Python interpreter was pretty smart to understand what is the variable type which is available in the script, will cover more in the type() later
- Dictionary are like Hash in Perl/Ruby (Key/value pairs)
- Tuples & Lists are like arrays
- Int, String, Float are just types of variables

- Just like any other language variable names are case sensitive

```python
>>> name = "Python"
>>> print name
Python
>>> type(name)         # As you can see, Python interpreter was smart
<type 'str'>
```

- If something variable quoted in "", its considered as STRING, refer the below

```python
>>> Price = 666
>>> price = "555"
>>> print price, Price
555 666
>>> type(Price)
<type 'int'>
>>> type (price)      # Since price value ("555") mentioned under double quotes ("") so it become string
<type 'str'>

>>> price = 365.00    # Here it goes for FLOAT
>>> type(price)
<type 'float'>
```

Another useful build function is `id()` - which is useful to get the memory location where the values currently stored in

```python
>>> price = 365.00
>>> id(price)
140653874971696
```

Also when you refer the already store values, it doesn't create the new it simply refers the previous one (Kind of alias)

```python
>>> newprice = price
>>> id(newprice)
140653874971696
```

Also note that, if you update existing variable also it update to NEW location

```python
>>> id(val)
140653874966912
>>> val = 23445
>>> id(val)
140653874997720
```
**Basic Math**

- Python follow standard operation
- Just like all other programming language, python also its own order evaluation

```
(PEMDAS)
 P - Parenthesis
 E - Exponential
 M - Multiple
 D - Division
 A - Addition
 S - Subtraction
```
  - Above are the order python evaluates in the math operation
  - But there are some things, Since D&M considered as same priority it will interchange in case D comes first while evaluating
  Same applies for A & S

examples

```python
>>> print 2**8 * 4 / 4 + 44 - 44  # Normal order
256
>>> print 100 / 2 * 3   # This one doesn't follow as mentioned in the case #2 (since D & M same line but D comes first)
150

```
- In python in case INT divide operation doesn't return FLOAT, it will return the values in INT only, refer the below

```python
>>> print 100 / 6
16
>>> print 100 / 6.0
16.6666666667
```

PARENTHESIS - here is the example

```python
>>> print 100 / 2 * 3
150
>>> print 100 / (2 * 3)
16
```

Modulo - Calculating the reminder in the division

```python
>>> print 100 / 16
6
>>> print 100 % 16
4
```

We can put the above in the single line

```python
>>> print "100 divided 16 is :", 100 / 16, "An the reminder is:", 100 % 16
100 divided 16 is : 6 An the reminder is: 4
```

**STDIN input**

String Concatenation

```python
>>> print "First", "Second"
First Second
>>> print "First" + "Second"    # This is the example of string concatenation, simple "+" will do
FirstSecond
>>> print "First "*3            # Another example (simply use space in string itself to get rid of ugly format)
First First First
```

Some more examples

```python
>>> print "First\n"*3
First
First
First

>>> print "First\t"*3
First First First
```

**How to collect the input from the user ?**
- By using the "input" & "raw_input" built in functions
    - "input" will be used get the INT data type
    - "raw_input" for string & all others as well

```python
>>> message = raw_input("Wassup? ")                 # Just storing the input values from user in "message"
Wassup? Hi There
>>> print message
Hi There

>>> my_num = input("Please enter any INT: ")       # Just storing the input values from user in "my_num"
Please enter any INT: 6
>>> print my_num
6
```

If you tried insert the "String" under "input" function it will return error, refer the below

```python
>>> my_num = input("Please enter any INT: ")
Please enter any INT: one million
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<string>", line 1
    one million
              ^
SyntaxError: unexpected EOF while parsing
```

Sample script

```python
#!/usr/bin/python
# Purpose: Hello World
# Author: Me
# Date : 31-12-2017

name = raw_input("Yo, Whats your name? ")
age = input("BTW, Dude what is your age? ")
exp = input("Well, How long you gonna live? ")
org = exp - age

print "Dear", name, "As per our calculation, you are gonna live only,", org, "more years !"
print "Have FUN !!!!!"

#END
```

**Strings I**

- Calculating the string with len()

```python
>>> import string
>>> msg="Hello"
>>> len(msg)
5
```

- Extracting specific string from given word, the very first character noted as "0"

```python
>>> print msg[0]
H
```
- Some other examples
`msg="abcdefghijklmnopqrstuvwxyz"`

For the below, very first character noted as "0

- `print msg[:]` - return the entire string (just same as "print msg")

- `print msg[:2]` - Print only very first two (0,1) character (boundary value never gets printed)
`'ab'`

- `print msg[2:10]` - Print from 2rd to 9 character only & don't print the rest
`'cdefghij'`

- `print msg[2::10]` - print the from the 2rd value + every 5th occurences of it
`'cmw'`

**Strings II**

In order to use the string related operation you have use the following once after the shebang shell `import string`

```python
#!/usr/bin/python
import string

name="yo yo"

print len(name) # This is doesn't required the "import string"
print string.upper(name)
print string.lower(name)

print string.capwords(name)   # Cap's the every first char of word
print string.capitalize(name)  # very first char of whole sentence

print string.split(name)  # Here delimiter considered as " " since name delimited by the same

```

**String III**

Write the script from above - String II
(Include "for" looping the each character in the string)

```python
#!/usr/bin/python
import string

name="yo yo"
print "The Length of the name is: ", len(name)
print "Convert to Upper: ", string.upper(name)
print "Convert to back to Lower: ", string.lower(name)

print "CAP the very each 1st word only: ", string.capwords(name)
print "CAP the very 1st word: ", string.capitalize(name)

print "Split the name: ", string.split(name)

print "Loop the Name: "
for i in name:
  print i

Name="yo yo"

if name == Name:
  print "They Match"

print "Join the Name: ", string.join(name)
```

**Note**
Interactive Shell doesn't required the "print" to return value but script EXPLICITLY required the "print" return the values

**Lists I**
- Which is used to allow store in the "String,int, whatever" values in comma separated list & gives access to them
- List should declared inside the `[]` with `","` separated values
- anything quoted within `""` are considered as strings

Functions which is available over the List
- `listname.reverse()` - reverse the content
- `listname.append()` - appends to list with new list in the element place (Refer the below code result)
- Basically above makes the nested list
- `listname.pop()` Removes the last # It considers the list as stack & perform LIFO method
- listname.extend()` - this is Extends the listname with new list
- listname.insert()` - inset the element into the specific index value, refer the below code

```
ex:
    listname.insert(index-position, value)
    listname.insert(0,1)
```

**NOTE**
FIFO = used in the queue mechanism, especially in Mailing system

```python
#!/usr/bin/python
import string

numlist=[1,2,3,4,5]
print "Print the Numlist", numlist

numlist.reverse()
print "Reverse it: ", numlist

numlist.reverse()
print "Back to Normal", numlist

numlist2=[6,7,8,9]
# Below does the append (not the extend)
# Appending the numlist2 into numlist1 in last element place
# Does the nested list
numlist.append(numlist2)
print "1St element" ,numlist[0]
print "3rd element", numlist[2]

print "6th element", numlist[5] # It prints full list which nested
print "1th from 6th element", numlist[5][0] # It prints only the 1st element from numlist[5]
print "4th from 6th element", numlist[5][3]

# Delete the last
numlist.pop()
print "After POP", numlist

numlist.extend(numlist2)
print "After extent: ", numlist

numlist.pop()
print "Removing again last item from numlist: ", numlist

# Remoce the very 1st element from the numlist
numlist.pop(0)
print "After removing 1st element by pop: ", numlist

# INserting the values
numlist.insert(0, 1)
print "After Insert: ", numlist
```

**List II**

- `range(10)` => print 0-9 in list format
- `range(1, 10)` => prints from 1 to 9
- `range(1,10,2)` => prints incremental of 2, from 1 through 10, such as below

```python
>>> range(10)
[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
>>> range(1, 10)
[1, 2, 3, 4, 5, 6, 7, 8, 9]
>>> range(1,10,2)
[1, 3, 5, 7, 9]
```

Lets move string based LIST

 - string list needs to be inside the double quotes `""`
 - Almost all functions are same which ever we evaluated in the `INTEGER` based list as above

```python
#!/usr/bin/python

# Simply printing some INT by range
print range(10)
print range(1,10)
print range(1,10,2)

stringlist = ["My", "String", "list", "is", "HERE", "Okay"]

print stringlist

stringlist.reverse()
print "Lets Reverse it", stringlist

stringlist.reverse()
print "Revert the stringlist back to normal, ", stringlist

strlist2 = ["This", "is", "Another", "string", "list"]
print "Lets print string list2: ", strlist2

stringlist.append(strlist2)
print "After Append: ", stringlist

print "1st Elmement: ", stringlist[0]
print "3rd Element: ", stringlist[2]
print "7th Element", stringlist[6]

print "1st of 7th Element: ", stringlist[6][0]

print "Pop the last element from stringlist, ", stringlist.pop()

stringlist.pop(0)

print "After Removing the 1st element: ", stringlist

stringlist.insert(0, "myNew")

print "After insert: ", stringlist
```

**List III**

- Gonna dealing with the String manipulation, by processing the logfile
- `logfile[0:3]` - prints 0th, 1st & 2nd index element not the 3rd one: refer below code

```python
#!/usr/bin/python
import string

logfile = "20170122 1.1.1.1 9000 2.2.2.2 8880 300 test.go"

print "The logfile content: ", logfile

logfile2 = string.split(logfile)

print "The logfile2 is: ", logfile2

print "logfile is: ", type(logfile)
print "logfile2 is: ", type(logfile2)

print "1st element in logfile2", logfile2[0]

print "Only 0st to 2nd", logfile2[0:3]

print "1st to 5th", logfile2[1:6]

print "Form the new LIst from the above output"
logfile3 = logfile2[1:6]

print "Here we go, logfile3: ", logfile3
print "Bring the logfile3 bck to string", string.join(logfile3)

# now it prints as the list, because in previous we didn't save it ! just converted and printed
print logfile3
```

**List IV**

Writes the script for the above (We already did it)

**Dictionaries**

Intro:

- list as mutable => means can change at any time by using the string operation
- Tuple => Are immutable, we cant change once it declared | In other words its readonly

```python
#!/usr/bin/python

distro = ["Centos", "RHEL", "SuSe", "Debian", "Ubuntu", "Arch", "Manjaro"]

print distro
print type(distro)

# Instead deleting the Changin the values

distro[2] = "OpenSuse"
print distro

# So above operation applicable only in List, Not possible in Tuples
# Lets do that also

# This is how you have declare the Tuples by using () parenthesis
distro2 = ("Rpi", "Elementary", "Solus", "Xbmc")
print distro2
print type(distro2)

# When you trying to change the you get an error
# like blow
# distro2[2] = "Fedora"

# We append, prepend to it, but we can change # Serious I'm not sure i have to check
```

That's all abut the Tuples

**Jumping into Dictionary**
- Basically its a BULK in number of Key-Value pair
- Also its an un-ordered list / If you want list, then you have to use SORT function but it by default print in KEY in alphabetical order

```python
#!/usr/bin/python

mydict = {'Year': 2017, 'Month': 1, 'Date': 22}

print mydict
print type(mydict)

# How to extract anyof those
print "getting the Year: ", mydict['Year']
print "getting the Month: ",mydict['Month']
print "getting the Date: ",mydict['Date']

# Adding the new KV
mydict['Time'] = 13.00
print "After adding time: ", mydict

# Changing the values
mydict['Time'] = 13.35
print "After changing the time: ", mydict

print "Printing all the Keys: ", mydict.keys()
print "Printing all the Values: ", mydict.values()

# Deleting the KV pair
del mydict['Time']
print "After deleting the TIME: ", mydict

# Iterate them
# Iteritem in build python capability
for k,v in mydict.iteritems():
  print k, v

# Multiple values for single Key
# Basiaclly 1-to-Many Relationship

values = [1, 2, 3, 4]

mydict['nums'] = values

print "With Multiple Values Dict: ", mydict

# Iterate them again
for k,v in mydict.iteritems():
  print k, v
```

**Conditional I**

Conditional Operations/Operator
  - < (less than)
  - <= (less than or equal to)
  - > (Greater than)
  - >= (Greater than equal)
  - ==  (Equals to)
  - != / <>  (Not Equal to)

Note: I just came to know about declaring the var/values in the same line

```python
#!/usr/bin/python

# This is declaring the 2 var is same line

min, max = 8, 9

print "Min", min
print "Max", max

# If condition, condition is true it gonna print
if min < max:
  print min,"Min is < Max", max

# The condition is NOT true nothing gonna print here !
if min > max:
  print min,"Min is > Max", max

if min == max:
  print "Min eq to Max"
else:
  print "Its not Equal"

if min <= max:
  print "Min less than Max"

if min >= max:
  print "Min is greater than Max"
else:
  print "Max is Max"

if min <> max:
  print "They are NOT equal"
else:
  print "They are equal"
```

**Conditional II & Conditional III**

- Command line arguments were stored in the `sys` module, so if want to play around arguments you have to import the "sys" module
- Basically its stores into `sys.argv`, basically its a list. It will store all the cmdline arguments which you are passed
- `sys.argv` 1st element is the name of the script (just like BASH)
- Get the length of by using `len(sys.argv)`, it includes the script name also in the count (actually the very 1st one)
- `elif` are considered as the `nested if`
- `else` does use any condition parameter, its used to by-pass the all conditions

```python
#!/usr/bin/python

import sys

print "Simply print the SYS.ARGV", sys.argv
# Run script pass some are
# ex: ./basic_cond2.py arg1 arg2 arg3
# Results: Simply print the SYS.ARGV ['./basic_cond2.py', 'arg1', 'arg2', 'arg3']

passed = len(sys.argv)-1
print "Length of the ARGS: ", passed

my_len = 2

if passed == 0:
  print "You are SICK, enter something"
# Simple condition I
#else:
#  print "You got it Bruh !"
elif passed == my_len:
  print "Its Equal!"
elif passed < my_len:
  print "You are insane"
elif passed > my_len and passed <= 10:
  print "Its under 10"
elif passed >= 11:
  print "U r awesome"
```

**For Loops**

- Does the iteration of each of the values from any type of the data

```python
#!/usr/bin/python
import string

name="HellYeah"

print "Looping the Name Char"
for var in name:
  print var

list = range(1,10)
print "Looping the Nums"
for j in list:
  print j

stringlist = ["This", "is", "String", "list"]
print "Looping the Stringlist"
for k in stringlist:
  print k

# Some weird example, i don't that im gonna use it!
logfile = ["2017 1.1.1.1 3333", "2017 2.2.2.2 4444"]
logfile2 = []

for i in logfile:
  logfile2.extend(string.split(i))
  print logfile2
```

**While Loops I & While Loop II**

- Unless the "Condition" becomes TRUE it performs the block of the code
- So the condition is become "TRUE" then it stops the block & control over to the next

```python
#!/usr/bin/python

count = 0
print "Gonna Perform WHILE"
while count <= 5:
  print count
  count = count + 1

print "WHILE is ended"

# Below writing the Infinite loop
# while 1:
#   print "Infinite loop"

# Some dump code to check while

my_num = 9

guess_num = input("Guess the number: ")

while my_num != guess_num:
  guess_num = input("Guess the number: ")
```

Another example

```python
#!/usr/bin/python
import sys

my_num = 9
count = 0

guess_num = 100 - my_num

while my_num != guess_num:
  guess_num = input("Guess the number: ")

  if count == 3:
    print "Would you like to exit the Program? (yes/no): "
    user_pref = raw_input()
    if user_pref == "yes":
      sys.exit()
    else:
      count = 0
  count = count + 1
```

***Note:*** From the above first it will ask for "4" times, but when you say "no" to exit, it will only asks "3" times. This scenario will be explained later
