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
>>> print 2**16 (Exponentiona - 2Power16)
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

```python
msg="abcdefghijklmnopqrstuvwxyz"

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