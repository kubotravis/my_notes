Python Basics
---
Minimal Info can found here.

### Intro
Get it from internet or some where.

**Note**
Be alert on INDENTATION

### Basic Math

```python
>>> print 1 + 1  # This is the reason its called loosely type
2
>>> print 1 * 3 (Mul)
>>> print 4 / 2 (Divide)
>>> print 2**16 (Exponentional - 2 Power 16)
>>> print 10 - 5 (Sub)
```

### Strings
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
- `<` less than
- `<=` less than or equal to
- `>` (Greater than)
- `>=` (Greater than equal)
- `==`  (Equals to)
- `!=` or `<>`  (Not Equal to)

***Note:*** I just came to know about declaring the var/values in the same line

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

**File I/O I**

- In python if you open a file, the file handlers gonna run IN-MEMORY
- Just any other language its does all things with python

- Modes that are available
  `r` - read
  `rb` - read binary
  `w` - write
  `wb` - write binary
  `r+` - read and write

- below code, doesn't have try/error handling mechanism
- Check the document page, basically each mode has its own may function

- Ex: 
  - `r` - has `readline()` - Where used to read the line by line (in a string)
  - `readline()` returns the very first line ("first char" to breaking of the line "\n") & stored in string
  - `read()` reads entire file stored as one string
  - `read(12)` only reads the till 1st to 12 character in the file (count start from 1 not zero)
  - `readlines()` its is literal, its prints out everything (including the line breaks)

***Learned:***
- If readline() & read() are in the same script, which left by readline() returned in the read() function.
i.e if file has 10 line, 1 st line goes by readline(), and the read 9 lines returned by read( )

- Same goes for every other method, how to avoid ? - before using any other read method you should close the file

***Note:***
read() - bit dangerous if your file is very large, chances it may exhaust your memory

```python
#!/usr/bin/python

# Opening a file in READ mode
handler1 = open("myfile.txt", "r")

# Just prints the file descriptor/handler information
print handler1

# Or even you can assign into variable
#print handler1.readline()

firstline = handler1.readline()
print type(firstline)
print firstline

# Reading the entire file
whole_content = handler1.read()
print type(whole_content)
print whole_content

# Reading the specific line
print "Reading the 3rd line only"
spec_line = handler1.read(3)
print "3rd line: ", spec_line


# Using the readlines()
everything = handler1.readlines()
print "Here we are printing everything by using the readlines()", everything
```

**File I/O II**

- Here we will be using 2 handlers for both "Read" & "Write" operation
- Just came to know Write line function doesn't insert the NEW line ("\n"). Yes, its so if you are copying any content to an another file it gonna be same as original
- when ever do any operation on the file you have to "CLOSE", if its missed sometimes changes wont take effect
- Reading from one file & writing it to another files then use readlines(), cuz we required trailing new lines to print it as a original so the combination is readlines() & writelines()

```python
#!/usr/bin/python

# Here file names can be variablized, if you require
hand1 = open("data1", "r")
hand2 = open("data2", "w")

val1 = hand1.readlines()

#for i in val1:
#  hand2.write(i)

# Avoid the loop to write the content
# It takes care of the looping the content
hand2.writelines(val1)

# Close the file handlers
hand1.close()
hand2.close()
```

**File I/O III**
Here we are going to prompt the user to enter the files name/location - Nothing fancy

- some "append" based operation
  - `a` Append mode - incase the target file is not available then it creates, If its available just does the append

If you run the blow multiple time you can see that "data3" file would have been grown by append operation like by concatinations

```python
#!/usr/bin/python

hand1 = open("data1", "r")
hand2 = open("data3", "a")

val1 = hand1.readlines()
hand2.writelines(val1)

hand1.close()
hand2.close()
```

- Working with "r+" (aka Read and Write)

  - If you var/val to APPEND/write there wont be any newline char
  - If you readlines() there is proper "Trailing" new line character

```python
#!/usr/bin/python

# Reading and Writing it to the same file

hand1 = open("rdwrite.txt", "r+")

val1 = hand1.readlines()
print val1

#val2 = "Yo Yo"
val2 = hand1.readlines()

hand1.writelines(val2)

hand1.close()
```

**File I/O IV**

- Basically content handled in FILE I/O was `String`, Especially in `write()` operation
- So variable has integer values `(val = 3)` & you are trying to write it into another file by using `write()` it will fail.
- Because write consider the input as `STRING`, but we actually passed the `INT`
- Also it consider it has a `Tuple`, so Tuples are immutable

- Refer the below code, if you tries execute it will fail !

```python
#!/usr/bin/python
hand1 = ("rdwrite.txt", "w")

myvar = 34
mystr = "String"
myint = 56

hand1.write(mystr,myvar,myint)

hand1.close()
```

- So in-order to fix the above code we have to use the `Formatter` by adding the `%`
  - Types of formatter
    - `%s` - string
    - `%d` - integer
    - `%f` - float

So refer the below code fixed code

```python
#!/usr/bin/python
hand1 = ("rdwrite.txt", "w")

mystr = "String"
myvar = 34
myint = 56

hand1.write("%s %d %d" % (mystr,myvar,myint))

hand1.close()
```

- Above fixcode is not working in mycase (with python 2.7)

- Another sample code to increent and store it

```python
#!/usr/bin/python
hand1 = ("rdwrite.txt", "w")

myvar = 34
mystr = "String"
myint = 56

while myint < 100:
  hand1.write("%s %d %d" % (mystr,myvar,myint))
  myvar = myvar + 1
  myint = myint + 1

hand1.close()
```

- Another example with FLOAT (print only 2more num after the dot)

```python
#!/usr/bin/python
hand1 = ("rdwrite.txt", "w")

myvar = 34.50
mystr = "String"
myint = 56

while myint <= 100:
# %2f say that print only 2 digit after the dot (like 33.20, not 33.234334)
  hand1.write("%s %.2f %d" % (mystr,myvar,myint))
  myvar = myvar + 1
  myint = myint + 1

hand1.close()
```

**Exceptions I & II**

- It can used to print pretty decent error info when failure, instead of Generic Python error msg
- Basic demo on try & except
- After try: and except "indentation" is must

Simple script to ask the file name

```python
#!/usr/bin/python

f1 = raw_input("Enter the File location: ")

try:
  hand1 = open(f1, "r")
except:
  print "Problem opening file name called: ", f1

print "Will be moving on..."
print hand1.readline()
```

Example with the never ending loop (for me below notbreaking after entering the right filename also)

```python
#!/usr/bin/python
while 1:
  f1 = raw_input("Enter the File location: ")
  try:
    hand1 = open(f1, "r")
  except:
    print "Problem opening file name called: ", f1


print "Will be moving on..."
print hand1.readline()
```

Here is the example of `BREAK` with the never ending script

```python
#!/usr/bin/python
while 1:
  f1 = raw_input("Enter the File location: ")
  try:
    hand1 = open(f1, "r")
    break
  except:
    print "Problem opening file name called: ", f1


print "Will be moving on..."
print hand1.readline()
```

Instead of infinite lets give based on Count, with support of `if` conditional

```python
#!/usr/bin/python
import sys

count = 1

while 1:
# Staring the if block
  if count == 3:
    answer = raw_input("Reached max chances, would you like to Continue (yes/No): ")
    if answer == "yes":
      sys.exeit()
    else:
      count = 0

  f1 = raw_input("Enter the File location: ")
  try:
    hand1 = open(f1, "r")
    break
  except:
    print "Problem opening file name called: ", f1
  count = count + 1

print "Will be moving on..."
print hand1.readline()
```

There is an another example with `string.lower`, converting user input to lower case the perform the file based operation

```python
#!/usr/bin/python
import sys
import string

count = 1

while 1:
# Staring the if block
  if count == 3:
    answer = raw_input("Reached max chances, would you like to Continue (yes/No): ")
    if answer == "yes":
      sys.exeit()
    else:
      count = 0

  f1 = raw_input("Enter the File location: ")
  try:
    f1 = string.lower(f1)
    print f1
    hand1 = open(f1, "r")
    break
  except:
    print "Problem opening file name called: ", f1
  count = count + 1

print "Will be moving on..."
print hand1.readline()
```

**Functions**
- Function does kind of mask/encapsulation to the code
- Use decent name for the function which is not supposed to be clash with build-in calls/functions
- Its ultimately based on the INTENDATION
- Anything declared in-side the function block is only accessed by `function` not by out-side of the program as mentioned in the others
- Just like in BASH you have call the function some where in the main program to execute, else it returns NOTHING
- Once after the function declaration only we have to call the function, else you it get error. So define function & call it

```bash
$ python
>>> help()
help > keywords
```

Below code will not print anything

```python
#!/usr/bin/python

def hword():
  print "Hi There"

hword()
```

Author functioned the life expectency program, lets do some small math

```python
#!/usr/bin/python

print "Lets do some math"

num1 = input("Enter the 1st Number: ")
num2 = input("Enter the 2nd Number: ")

def mymath(n1, n2):
  added = n1 + n2
  return added

value = mymath(num1, num2)

answer = input("What could be the addition of them ?: ")

if answer != value:
  print "You Dump as b**th"
else:
  print "Alright, its very easy anyway"
```

**Modules**
- How to load module? by using `import <module_name>`
- How get anyother functions of loaded modules, by using `dir (mod_name)`
- multiple values of any function stored in list, no TUPLEs `(>>> sys.path)`

- From the modules directory, you get the *.py ascii code of the modules
  - *.pyc - is the bit code compiles version of *.py files, when you say "import.sys" this is files been used for the faster execution (not sure check the internet)

- You can `import` only required function in the module, instead importing the complete modules
- like `from sys import path`

```bash
$ python
>>> import sys
>>> dir(sys)
>>> print sys.platform
>>> print sys.path
```

  - From the above example `sys.path` returns the all available path where python library is store
  - You can add you custom path also by update the `PYTHONPATH` shell environment variable, so which is included in `path`

***IMPORTANT NOTICE***
- When you `import` a module it uses totally different namespace from actual main program that you wrote
- For exmple, after `import sys`  the `sys.path` & `path` are totally different, so you no need to worry about the overlapping

**SHUTIL**

Shell Utilities

- This is further discussion about the FileIO utilization
- why we shoud be shutils for copy a file? because its cross platform

Copying the a file

```python
#!/usr/bin/python

import shutil

#You can do this also
# s = shutil
# s.copy("file1", "file2")


# It created new file with NEW timestamp (current time stamp)
shutil.copy("file1", "files2")

# Below crates the NEW file with same timestamp as original file
shutil.copy2("file1", "file3")

# This is moving the file
shutil.move("file2", "/tmp/")

# Copy tree function
# It copies the content in recursively
src = "temp"
dst = "temp2"

shutil.copytree(src, dst)
```

- There is pretty cool feature about `copytree` is that it accepts the `BOLEAN` values as `0` or `1` in arg area
  - If it 1, then it copies the SYM link from the source to dst
  - if its 0, then it copies the actual file where SYM link associated (but for the naming convention it keeps the SYM link file name for the actual file)

- In order prove the above refer the below, make sure you have symlink in the "temp" dir, & remove if there any "temp2" directory

```python
#!/usr/bin/python

src = "temp"
dst = "temp2"

# 0 - Copies the actual file where the symlink pointed
shutil.copytree(src, dst, 0)
```

- In order to remove the directory tree you can use `rmtree`
  - It has the option `ignore_error` - where it accepts the values as BOOLEAN & ignore in case any error while removing the directory tree
  - `onerror` - is an another option. Not sure for what ?

```python
#!/usr/bin/python
import shutil

dst = "temp2"

shutil.rmtree(dst)
```

- Below code for the move the file
- likewise the is `movetree` you can check the reference page for more information

```python
#!/usr/bin/python
import shutil

src = "file3"
dst = "temp"

shutil.move(src, dst)
```


**Regular Expressions I**

- By default regex not inclueded in the Python core, so we have to import like `import re`
- From interactive shell you can do the following,

  ```bash
  $ python
  >>> import re
  >>> dir (re)
  ```

- Worst case of "realized" - If you want to use any function from the imported module, you have to use the module name as the prefix so,

  ```python
  import re
  re.blah(blah)
  ```

- In order to perform the match in Python,
  - 1st you have compile the regrex (im bit doubtful)
  - 2nd you have run the "Match" against the complied item
  - Refer the below interactive shell code

  ```python
      >>> import re
      >>> reg1 = re.compile('itsamatch')
      >>> reg1.match('itsamatch')
  ```
  - `group()` sub method, which is used to return the match
      
      ```python
      >>> print reg1.match('itsamatch').group()
      ```

  - Try for non match
      ```python
      >>> print reg1.match('it')
      None
      >>> print reg1.match('itamatc')
      None
      ```

- Its Regex, ofcourse it's case sensitive
- `\` used for escaping the special character in the Regex
- Instead of variable in the `.match()` function, refer the below code for variable substitute
      
      ```python
      >>> content = re.compile('matchme')
      >>> regex1 = content.match('matchme')
      >>> print regex1
      <_sre.SRE_Match object at 0x100ba1a58>
      >>> print regex1.group()
      matchme
      ```

Made small script from the above understandings

```python
#!/usr/bin/python

import re

# This is verbatim string based search
reg1 = re.compile('Hola')
match1 = reg1.match('Hola')
print "From example-1"
print match1.group()

# This example based on the VARIABLE used for storing the search string
reg2 = re.compile('Gracias')
searchstring = "Gracias"
match2 = reg2.match(searchstring)
print "From example-2"
print match2.group()
```

**Regular Expressions II & III**

- Will see how to "Ignore" the case sensitive, You can ignore as mentioned in the below, [dont mess around "re.IGNORECASE" keep as it is]
  ```python
  >>> reg1 = re.compile('itsamatch', re.IGNORECASE)
  ```

- Lets try below example to demonstrate
  ```python
  >>> import re
  >>> reg1 = re.compile('complicated', re.IGNORECASE)
  >>> match1= reg1.match('CoMpLiCaTeD')
  >>> print match1.group()
  CoMpLiCaTeD
  ```

- With `^`
  - This carrot sign which allows us to search & returns the thing which is `STARTS` with a specific char/word
  - Refer the below code
    ```python
    >>> import re
    >>> searchstring = "This is a test"
    >>> reg1 = re.compile('^T')
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    ```

- In order to get the char ending with you can use the `$` symbol, please refer the below code [I Dont understand this shit]
  ```python
  >>> import re
  >>> searchstring = "This is a test"
  >>> reg1 = re.compile('T$')
  >>> match1 = reg1.match(searchstring)
  >>> print match1.group()
  ```

- TIL `.` used to print the any character except the new-line `\n` character separator (it depends on how many times dots are prensented)

- If you want to print/match everything, you have use the match both `.` & `*`
  - `.*`
    - `.` prints very first char, regardless of whatever it is!
    - `*` print anything matches (else it repeats the . for inifite times - im not sure at this point) (it literally means)

***Notes:*** This combination doesn't work for multiple lines, you have to use the `re.DOTALL` sub-function with it

- For the above scenario please refer the below sample code

  ```python
  >>> import re
  >>> searchstring = "This is a test"
  >>> reg1 = re.compile('.*')
  >>> match1 = reg1.match(searchstring)
  >>> print match1.group()
  ```

- For the multiple line & `re.DOTALL` combination
  ```python
  >>> import re
  >>> searchstring = "This is a test\nAnd this another test line"
  >>> reg1 = re.compile('.*', re.DOTALL)
  >>> match1 = reg1.match(searchstring)
  >>> print match1.group()
  ```

- There are special meta character `^, ., *, +, ?`
  - `^` -> Starting with
  - `.` -> Match any char at only once
  - `*` -> Infinite match for any char
  - `+` -> Incremental match for specific char, breaks when NEW char found
  - `?` -> search the match ONE or zero time

- Also you can print the `RANGEs` as mentioned in the below

  `[a-z]` - Prints only 1 char, which matched the range

  `[a-z]+` = Prints 1 to Many the supplied char matched, breaks if its didn't find anything in range
  - Refer the below example: just to print the range
    
    ```python
    >>> import re
    >>> searchstring = "this is a test"
    >>> reg1 = re.compile('[a-z]')         # This is just range
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    t
    >>> reg1 = re.compile('[a-z]+')        # Range with continue (breaks when it finds the "SPACE" character)
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this
    >>> reg1 = re.compile('[a-z]+.*')      # We are accepting even the SPACW with "." + printing everything with "*"
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a test
    ```

- In order to get/evaulate the get space you can try the following
    ```python
    >>> reg1 = re.compile('[a-z]+\s')      # \s goes for the space
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this
    ```

- `\s` is goes to find the space and give resul accordingly. But when you want to keep `\` as also a part regex you have mention it has `raw` (you may required this). simply prefix -ing the "r", which stands for the "raw notation"

    ```python
    >>> reg1 = re.compile(r'[a-z]+\s')
    ```

- One more example to print only one more character next to the space
    
    ```python
    >>> reg1 = re.compile('[a-z]+\s[a-z]')      # \s goes for the space with 1 char
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this i

    >>> searchstring = "awesome right you"
    >>> reg1 = re.compile('[a-z]+\s[a-z]+')      # \s goes for the space with 1 or more char
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    awesome right

    >>> searchstring = "this is a test"                         # In this example we are trying to print everything
    >>> reg1 = re.compile('[a-z]+\s[a-z]+\s[a-z]\s[a-z]+')      # Adjusting as per out search string
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a test

***Note:*** regex are case sensitive default, unless you dont specify the IGNORE the case sensitive

- Going through with the CAPS letter
    ```python  
    >>> searchstring = "this is a TEST"
    >>> reg1 = re.compile('[a-z]+\s[a-z]+\s[a-z]\s[A-Z]+')      # [A-Z] did for the CAPS
    # Also you can use the IGNORECASE
    # reg1 = re.compile('[a-z]+\s[a-z]+\s[a-z]\s[a-z]+', re.IGNORECASE)
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a TEST
    ```

- Going with through the NUMERIC
    ```python
    >>> searchstring = "this is a 2017"
    >>> reg1 = re.compile('[a-z]+\s[a-z]+\s[a-z]\s[0-9]+')      # [0-9] did for the NUMERIC part
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a 2017
    ```

***Note:***
We can use the extended notation to get this things done as seen previous example - matching all words & numeric.
Just like space notation `\s`, we cause for word as `\w` & digit has `\d` and more.

- Refer the below code for "\d" example
    ```python
    >>> searchstring = "this is a 2017"
    >>> reg1 = re.compile('[a-z]+\s[a-z]+\s[a-z]\s\d+')      # \d+ did for the NUMERIC part
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a 2017
    ```

**Regular Expressions IV**

- There is something called `literal` search (refer the below code)
  ```python
  >>> reg1 = re.compile('[a-z] pid \d+')
  ```
  - In the above `pid` and `Space` bar are literal search we did user `[a-z]` & `\s` to denote them

- Here we are re-writing the word search by using the `\w` (`\w+` 1 or more possibility it finds)
  - `\w` (word) covers the following in its search,
    `a-z`, `A-Z`, `_`, `0-9`

- Lets rewrite the above examples bit,
    ```python
    >>> searchstring = "this is a test"
    >>> reg1 = re.compile('\w+\s\w+\s\w?\s\w+')
    >>> match1 = reg1.match(searchstring)
    >>> print match1.group()
    this is a 2017
    ```

**Regular Expressions V**

- Writing the script for the above scenario(s) we come across

```python
#!/usr/bin/python
import re

reg1 = re.compile('hello', re.IGNORECASE) # Matching Verbatim
searchstring = raw_input("Wassup Dog: ")
match1 = reg1.match(searchstring)

if match1:
  print "Yo we got match: ", match1.group()
else:
  print "Oops, No Match"
```

- Another code use the same above except "reg1"

```python
reg1 = re.compile('\d+', re.IGNORECASE)
```

***Output:**
```bash
~/tmp/Python-Tut/RegEx>_ ./regex-3.py
Going through DIGIT based search
Wassup Dog: 123Dog
Yo we got match:  123

~/tmp/Python-Tut/RegEx>_ ./regex-3.py
Going through DIGIT based search
Wassup Dog: 123 Dog
Yo we got match:  123
```

***Note:** As per regex we have speified int he "reg1", ot  goes and finds 1 or more digits and breaks if its no Digits in provided input

- Another example for "Digit" "Space" & "Words" as well

```python
reg1 = re.compile('\d+\s+\w+', re.IGNORECASE)
```

- Another script with "Chances" (just by giving the 3 chances)

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\d+\s+\w+', re.IGNORECASE)

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Waddup, Dog: ")
  match1 = reg1.match(searchstring)

  if match1:
    print "You have got matches: ", match1.group()
    print "So exiting out"
    sys.exit()
  else:
    print c, "chance is gone"
    c = c + 1
```

**Regular Expressions VI & VII**

- Scaling out the above script
- Since all case/example we have seen searching BGINNING of the string ! There is no examples shown to search middle of string (right?) all beacuse we have used the `match` function call `re.match`

- In order to search the pattern anywhere we have to use the `search` call `re.search`

- Refer the simple example
  
  ```python
  >>> import python
  >>> searchstring = " This 2017 yo"
  >>> reg1 = re.compile('\d+')
  >>> search1 = reg.search(searchstring)
  >>> print search1
  >>> print search1.group()
  ```

Comparison between match/search

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\d+\s+\w+', re.IGNORECASE)
reg2 = re.compile('\d+')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Waddup, Dog: ")
  match1 = reg1.match(searchstring)
  search1 = reg2.search(searchstring)

  if match1:
    print "You have got matches: ", match1.group()
  else:
    print "No Match found"
  if search1:
    print "you have got SEARCH: ", search1.group()
  else:
    print "No SEARCH found"
  c = c + 1
```

- Another example with findall
- Reason is `re.search` goes find sequence of digit, if there no-digit values it breaks immediately & then there is DIGIT it can't pickup
- For instance if the search string is `You its 2017 and 2132`, its only picks the `2017`, but can't pickup the `2132`

***Note:*** Find all doesn't required group, cuz by nature its output format LIST

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\d+\s+\w+', re.IGNORECASE)
reg2 = re.compile('\d+')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Waddup, Dog: ")
  match1 = reg1.match(searchstring)
  search1 = reg2.search(searchstring)
  findall1 = reg2.findall(searchstring)

  if match1:
    print "You have got matches: ", match1.group()
  else:
    print "No Match found"
  if search1:
    print "you have got SEARCH: ", search1.group()
  else:
    print "No SEARCH found"
  if findall1:
    print "from Findall: ", findall1  # Find all doesn't required group, cuz by nature its output format LIST
  else:
    print "Nothing from FINDALL"
  c = c + 1
```

- Will going through an another example with `re.finditer`
- `finditer` items should looped (These especially designed looping the items), without loopoing the "fintiter" values are non-useable

```
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\d+\s+\w+', re.IGNORECASE)
reg2 = re.compile('\d+')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Waddup, Dog: ")
  match1 = reg1.match(searchstring)
  search1 = reg2.search(searchstring)
  findall1 = reg2.findall(searchstring)
  finditer1 = reg2.finditer(searchstring)

  if match1:
    print "You have got matches: ", match1.group()
  else:
    print "No Match found"
  if search1:
    print "you have got SEARCH: ", search1.group()
  else:
    print "No SEARCH found"
  if findall1:
    print "from Findall: ", findall1
  else:
    print "Nothing from FINDALL"
  if finditer1:
    print "Printing the item in FINDITER"
    for i in finditer1:
      print i.group()
  else:
    print "Nothing on FindIter"
  c = c + 1
```

**Regular Expression VIII**

- Will sending new compile format to search the e-mail address
- Following can be used
  ```python
  re.compile(.*@.*) # This is simplest form
  re.compile(\w+#.*) # Another form
  ```
- Refer the sample code the example
***Note:*** This code doesn't pickup the `first.secondname@domain.net` like address, since `.` is not evaluated in the regex
- A simple mail-id will work, like `myname@mydomain.com`

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\w+@.*')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Wuddup, Dog: ")
  match1 = reg1.match(searchstring)

  if match1:
    print "Got the mail address: ", match1.group()
  else:
    print "No E-Mail id"

  c = c + 1
```

***Note:*** If you want to search for literal char `.` use `\.` in the regex

- So in-order to get the pattern like `first.secondname@domain.net` use the below regex
- Still there is an issue with this script
  - Incase if you input as `my email id is name@domain.com` - it will output as `my email id is name@domain.com`
  - Which we don't want, check for next example for the fix

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('.*@.*')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Wuddup, Dog: ")
  match1 = reg1.match(searchstring)

  if match1:
    print "Got the mail address: ", match1.group()
  else:
    print "No E-Mail id"

  c = c + 1
```

- So if you want to avoid output like = `my email id is name@domain.com`
- Below code is not working (i dont know why)

```python
#!/usr/bin/env python
import re
import sys

reg1 = re.compile('\s+\w+@.*')

c = 1
t = 3
print "You got only", t,"chances :"
while c<=t:
  searchstring = raw_input("Wuddup, Dog: ")
  match1 = reg1.match(searchstring)

  if match1:
    print "Got the mail address: ", match1.group()
  else:
    print "No E-Mail id"

  c = c + 1
```

**Regular Expressions IX & X**

- Will be working with the files
- Find the below code simply match/search/findall/finditem for given condition

```python
#!/usr/bin/python

import re

reg1 = re.compile('\d+')

hand1 = open("datafile", "r")

for searchstring in hand1.readlines():
  match1 = reg1.match(searchstring)
  search1 = reg1.search(searchstring)
  findall1 = reg1.findall(searchstring)
  finditer1 = reg1.finditer(searchstring)

  if match1:
    print "from MATCH: ", match1.group()
  else:
    print "Nothing from MATCH"
  if search1:
    print "from SEARCH: ", search1.group()
  else:
    print "Nothing from SEARCH"
  if findall1:
    print "from FINDALL: ", findall1
  else:
    print "Nothing from FINDALL: "
  if finditer1:
    print "from FINDITER: "
    for i in finditer1:
      print i.group()
  else:
    print "Nothing from FINDITER"
```

- Using the above script with FILE handlers

```python
#!/usr/bin/python

import re

filename = "/var/log/system.log"

reg1 = re.compile('\d+')

hand1 = open(filename, "r")

for searchstring in hand1.readlines():
  match1 = reg1.match(searchstring)
  search1 = reg1.search(searchstring)
  findall1 = reg1.findall(searchstring)
  finditer1 = reg1.finditer(searchstring)

  if match1:
    print "from MATCH: ", match1.group()
  else:
    print "Nothing from MATCH"
  if search1:
    print "from SEARCH: ", search1.group()
  else:
    print "Nothing from SEARCH"
  if findall1:
    print "from FINDALL: ", findall1
  else:
    print "Nothing from FINDALL: "
  if finditer1:
    print "from FINDITER: "
    for i in finditer1:
      print i.group()
  else:
    print "Nothing from FINDITER"
```

- As of now we are doing just search/match/find(all/iter), How about search and repalce ?
- There is sub-function call `re.sub()` for that
  - The feature with `sub()`, is `count` which is allows you to perform the limits of substitute in HORIZONDALLY
  - Which means that in each if the actuall finds 4~5 occurance, if count set to `1` then it replaces only one OCCURANCE (as per count=1)
  - So If there is NO count specified, it changes every single occurance

```python
#!/usr/bin/python
import re

reg1 = re.compile('\d+')

filename = "data1"

hand1 = open(filename, "r")

for searchstring in hand1.readlines():
  print reg1.sub("2016", searchstring)
  # IF you want to limit the SUB
  # print reg1.sub("2016", searchstring, count=1)

hand1.close()
```

- Above code so far Displays the SUBS change but didn't write to anyfiles.
  - Let open another file handlers & write the content to it
  - Refer the below code for the same

```python
#!/usr/bin/python
import re

reg1 = re.compile('\d+')

filename = "data1"
filename2 = "sub_data"

hand1 = open(filename, "r")
hand2 = open(filename2, "w")
for searchstring in hand1.readlines():
  subs = reg1.sub("2016", searchstring, count=1)
  hand2.writelines(subs)

hand1.close()
hand2.close()
```

- Here is another example to mask the email ids

```python
#!/usr/bin/python
import re

reg1 = re.compile('\w+@.*')

org = "email-id.org"
dup = "email-id.mask"

hand1 = open(org, "r")
hand2 = open(dup, "w")

for searchstring in hand1.readlines():
  mask_is = reg1.sub("dummy@domain.local", searchstring)
  hand2.writelines(mask_is)
```

**Syslog Integration I**

- Python has capability of logging info into `SYSLOG` (from level 7 to 0)
- To create any log object you have to initiate the `logger` object with `getLogger` sub method `logger = logging.getLogger()`
- Additionally this module has the `FileHandler` sub method, so we no need to user any other Filehandler from the another module

A simple model works like this

```
Model
(Object)Logger --> Handler --> Destination
```

- We are creating the Object for logger.
  - Logger sends the content/log to handler.
    - Handler sends loginfo/msg to the Destination
      (Destination can be Syslog/TCP/UDP/File & Etc)

- Please refer the below sampl code

```python
#!/usr/bin/python
import logging

# Definition & Instantiation of logger object
logger = logging.getLogger()

# Definition of the Handler
# Here the default is "Append" mode, which means the below code same as un-commented
# hand1 = logging.FileHandler('my_log.log', 'a')
hand1 = logging.FileHandler('my_log.log')

# Definition of Log content format
format = logging.Formatter('%(asctime)s %(levelname)s %(message)s')
hand1.setFormatter(format)

# Add handler to logger
logger.addHandler(hand1)

# Invocation of logger
logmessage = "This is Test log message"
logger.error(logmessage)
```

- By default `logger` default mode is `append`, which means that if you run the above scipt multiple times the log content get appeneded
- If you dont want "append" as defult, then specify required method

**Syslog Integration II**

- Here we are modifying the script to log content to syslog file itself
- `import logging` support only 2 handlers, which is `Stream/File`, but in-order to log the information to syslog file you have import some other handlers as well
- Refer the below sample code for more information
- There is no limit for using file related handlers, we can append the content into many files by nature. Means that one logger and multiple handlers

- From the below example code we did specify the any log file name to log the message for `SysLogHandler()`
  - which is `hand2 = logging.handlers.SysLogHandler()`

- We can configure the `SYSLOG` daemon incase of any log info from the user to `SysLogHandler()` then appent to an any specfic file (/var/log/custom.log)
  - For that you have to update following line in `/etc/syslog.conf` file
  ```
  "user.*         /var/log/custom.log"
  ```

    - `user` - means from any user from the system
    - `*`    - all kind of severity (like ERROR, INFO, EMER, DEBUG and others)
    - `/var/log/custom.log` - where all msg from the user should be stored
  - restart the SYSLOG daemon `service syslog restart`

```python
#!/usr/bin/python

import logging
import logging.handlers

logger = logging.getLogger()

hand1 = logging.FileHandler('my_log.log')
hand2 = logging.handlers.SysLogHandler()

format = logging.Formatter('%(asctime)s %(levelname)s %(message)s')
hand1.setFormatter(format)
hand2.setFormatter(format)

logger.addHandler(hand1)
logger.addHandler(hand2)

logmessage = "This is Test log message"
logger.error(logmessage)
```

***Notes:***
- Syslog opens the UDP port `514`, so we can use that port send message to Syslog
- If there is no proper ROUTE in the `/etc/syslog.conf` by all messages sends/route to the `/var/log/messsage`
- If we used `SysLogHandler()` by default it prefix `Date/Time Stamp, hostname` like things

In latest system, it gonna be `rsyslog`, so thinking that steps on the OS gonna be different.

**CGI I**

[Common Gateway Interface]

- Typically all content stays in `/var/www/cgi-bin`
- If you want more configuration detail `grep -i -ne cgi /etc/httpd/conf/httpd.conf`

- Writing the simple "helloworld.py" cgi script

```python
#!/usr/bin/python

print "Content-type: text/html"
print

print "<html>"
print "<title>From Py</title>"
print "<body>"
print "<h1><strong><font face='Monaco' size='+2'>Hi There !!!</font></strong></h1>"
print "</body>"
print "</html>"
```
Make it executable and call form the browser


**CGI II**
------------
- Now we trying to setup a simple form
- create file name called "form.py"

```python
#!/usr/bin/python

print "Content-type: text/html"
print

print "<html>"
print "<title>Form from Py</title>"
print "<body>"
print "<form action='action1.py' method='post'>"
print "Name: <input type='text' name='name' size='50'><br>"
print "Title: <input type='text' name='title' size='50'><br>"
print "E-Mail: <input type='text' name='email' size='50'><br>"
print "<input type='submit' value='submit'>"
print "</form>"
print "</body>"
print "</html>"
```

- let's create the "action1.py" file

```python
#!/usr/bin/python
import cgi

print "Content-type: text/html"
print

form = cgi.FieldStorage()

name = form["name"].value
title = form["title"].value
email = form["email"].value

print "Yo, ", name, "Wassup !<br>"
print "Hmm,", title, "ah....<br>"
print "Is this your mail > ", email, "<br>"
```

**Globbing**

- Kind of quick and dirty
- Search against for file/directory search (I think so)
    
  ```python
  >>> import glob
  >>> search1 = glob.glob('*.py') # list of files has the extension of "py"
  >>> print search1  # By nature output will be list "Mutable"
  >>> len(search1)
  ```
  - You can loop them for doing anything with FileIO if required
  - By default "glob" search under the "Current" path

  - ex: Will be using some other directory

    ```python
    >>> search = glob.glob('/home/ramdasssadha01/*.yaml')
    ### search = glob.glob('./test/*.yaml')
    >>> print search
    ['/home/ramdasssadha01/lineinfile.yaml', '/home/ramdasssadha01/rev-proxy.yaml']
    >>> len(search)
    2
    ```

  - Another kind of search
    
    ```python
    >>> search1 = glob.glob('fileio?.py')
    >>> search1 = glob.glob('fileio[1-4].py')
    ```

- Another task to put all together in the script

```python
#!/usr/bin/python
import glob

query1 = raw_input("Enter search format with inclusion of directory if required: ")

search1 = glob.glob(query1)

if search1:
  print "So far you have", len(search1), "matches, and they are below"
  print search1
else:
  print "Boss, No Match"
```

- Simply, Here we are going to loop the search results

```python
#!/usr/bin/python
import glob

query1 = raw_input("Enter search format with inclusion of directory if required: ")

search1 = glob.glob(query1)

if search1:
  print "So far you have", len(search1), "matches"
  print "You have got matches, And will gonna be in LOOP"
  for i in search1:
    print i
else:
  print "Boss, No Match"
```

