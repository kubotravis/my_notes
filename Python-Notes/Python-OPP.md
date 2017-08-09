# Python _ OOP

### Chapter 01: Introduction

**What I have got here ?**

**Procedure Programming**
- You have data (variables & its with its values)
- Then you will be using some standard function to process them

**OOP**
- Your data's are called as OBJECTS
- Your OBJECTS are processed by METHODS
- All of you OBJECT & METHODS maintained under the CLASS

**Three Pillars of OOP**
- Encapsulation
- Inheritance
- Polymorphism

**TIL**
- Python also has the "None" as one of the type
- "Boolean" I just came to know, for function I never tried

```python
#!/usr/bin/env python

mybool = True
mynone = None
mylist = ['a', 'b', 'c']

def myfunc():
  print "Hi"

print type(mybool)
print type(mynone)
print type(mylist)
print type(myfunc)

this_type = type(mylist)
print type(this_type)
```

- This is interesting

```
var=5
dir(var)
```

- From the above, the "attributes" associated with `int` data type are listed here
- The attribute name begins with "--" is called `PRIVATE` or `MAGIC` attributes
- Checking with some attributes

```
 var.real
 var.numeric
 var.bit_length()
```

**Modules**
- A simple file that contains the Python code, So it can executed directly or can be imported from another python script

**Class**
- Simply a Python code

### Chapter 02: Simply Prerequisites
Basic knowledge on Python

### Chapter 03: Classes
- Every callable methods is an Attribute, Not every attributes are Method.
  - [As i already knew, Method is a FUNCTION inside the CLASS]

- 2 Examples so far the clarification
  - Example 1: 
    - Consider the "Car" factory is a class
    - Where you can pass your own data (data) to Methods which produces the car, as per your wish
    - As result there is a NEW car (instance), with your custom spec
    - And its object will be run(), stop(), halt() and so on

- So Each different set of DATA produces a NEW instances
  - Example 2:
  - ATM Machine also can be an another example. Dont have time to write the explanation, better try match as above

**Defining the Class**

1. Code to for a class, its a very simple class does nothing

```python
#!/usr/bin/python

class MyClass(object):
  pass
this_obj = MyClass()
print this_obj

```

Excution of the above code

```
~/Python-Tut/OOP/Classes>_ ./class-1.py
<__main__.MyClass object at 0x10ec882d0>
```
- Above output confirms that we have instance of MyClass()
    - "__main__" is a NameSpace of the script that currently being executed
    - hexcode - Refers the address that "object" being stored

- About "pass"
  - Since we are NOT doing anything inside the class, may Python will get confuse. So inorder to indicate that its just a block & pass on

2. Adding to VAR into Class

```python
#!/usr/bin/python

class MyClass(object):
  var = 10

this_obj = MyClass()
that_obj = MyClass()

print this_obj.var
print that_obj.var
```

- How the Instances get the var ? Since instances are the COPY/RUNNING from the class, so it can access the Class Variables

**6 Points to Understanding the Class**
(So far we covered the 2 points, so i'm gonna write those 2 of them)
  -  "Instance" of "class" knows what class its from
  - Varibales (vars) defined in the class are available to the instance

**Instance Method**
- Just a recall, a variable declared in the Class is accessible by its Instance(s)
- Why we recalling is ? - Anything that declared in the Class accessible by its Instance
  In the same we are going to access the Method (a function) which is defined inside the Class

- Let me write some code later, will go by explanation

```python
#!/usr/bin/python

class Joe(object):
  def callme(self):           # Just like variable declarion we defined the Method (Function)
    print('calling the "callme" method with instance: ')

thisjoe = Joe()

print thisjoe.callme()        # Here we called the method, by an Object
```
When you execute the above code, it will return "None"
  - Because you are not Supposed call a method(), with "print"
  - Also when the method/function doesnt return anything, "print" will return the "None"
    Refer the clean example of the code

```python
#!/usr/bin/python

class Joe(object):
  def callme(self):           # Just like variable declarion we defined the Method (Function)
    print('calling the "callme" method with instance: ')

thisjoe = Joe()

thisjoe.callme()        # Here we called the method, by an Object
```

- Above will return the proper output (What we wanted)

**Now the explanation about "self" in the declared method ?**

- In general if you declare a function with ARG, when call that function you have to supply required ARG in the CALL method.
    Else it will return a error
- But in the above scenario we didn't passwd any value to the method while calling it such as "thisjoe.callme()", but still passes
- What if we remove the "self" label & execute the code, it will retuen the error. as follows
  - callme() takes no arguments (1 given)
  - Wait a second, we didnt supply any args in the "thisjoe.callme()" call !!?
  - This is pretty clear that "thisjoe.callme()" still passing to arguments to the Class Method still nothing in it.
- Here is the MAGIC
  - When we call an METHOD on a an INSTANCE, the INSTANCE gets passed as the first argument to the METHOD automatically
  - This is the default IMPLICIT behaviour that we shoud aware
  - So now you understand what is "self" in there ! Its the instance ! Lets try to print the "self" then.

```python
#!/usr/bin/python

class Joe(object):
  def callme(self):
    print('calling the "callme" method with instance: ')
    print(self)

thisjoe = Joe()

thisjoe.callme()
```

If you execute the above code will get the below response

```
~/Python-Tut/OOP/Classes>_ ./3B-Class-with-Method-self-print.py
This is calling the "callme" method from instance:
<__main__.Joe object at 0x10cbd22d0>
```
- Output decribes, its "Joe" instance & where it stored in the memory

- Now actually we can prove that this is same instance on we called, by printing "thisjoe"; refer the below example

```python
#!/usr/bin/python

class Joe(object):
  def callme(self):
    print('calling the "callme" method with instance: ')
    print(self)

thisjoe = Joe()

thisjoe.callme()

# below line added additionally, to prove that we calling the same in the above declared whatever in the method

print thisjoe
```

if you execute the above code will get the below response

```bash
~/Python-Tut/OOP/Classes>_ ./3C-Class-with-Method-proving.py
This is calling the "callme" method from instance:
<__main__.Joe object at 0x106de72d0>
<__main__.Joe object at 0x106de72d0>
```

- From the above we can identify that, those 2 objects are ONE object; they just printed into defferent places

**Summary from this lessons:**
- Instance methods are variable defined in the class
- They are acces through the instances: instance.method()
- when called through the instance, instance automatically passed as 1st argument to the class
- Because of this automatic passing of the instance,
    - Instance method are known as "bound" methods. i.e, bound to the instance upon which it is called

**From 6 point:**
A method on a instance passes instance as the first ARGUMENT to the method (named self in the method)

**What I Learned?**
- Below code snippet is standard on all class declaration
- Since the we know that "INSTANCE" always passed to an 1st ARG to the "self"

```python
  def callme(self):    # Especially this "self" is super standard on the all class definition
    print('calling the "callme" method with instance: ')
```

**Instance Attributes**

What is `self` recall it (The Keyword in the Class declaration).

"self" - is an instance upon which the method was called.

Below example explains accessing the Instance(Method) values from the method. Has we talked earlier car manufacturing with spec in the example.

Here is the example how we can access the Methods objects

```python
#!/usr/bin/python
import random

class MyClass(object):
  def dothis(self):
    # Setting the Instance attribute value
    # Instance.attribute = value
    self.rand_val = random.randint(100, 200)

myinst = MyClass()
myinst.dothis()

print(myinst.rand_val)
```

Covered one of the **6 points** is,
 - Instance have their own data, instance attribute

**Encapsulation**

Let's start with code example

```python
#!/usr/bin/python

class MyClass():
  def set_val(self, val):
    self.value = val

  def get_val(self):
    return self.value

a = MyClass()
b = MyClass()

a.set_val(10)
b.set_val(100)

print(a.get_val())
print(b.get_val())
```

- Above is the clean code, where we have 2 method in the Class (setter & Getter), and we are setting the values in the instance & getting them back

- The question is by using the instance method we setting up the value `[a.set_val(10) & b.set_val(100) to => self.value = val ]`, can we directly
  - the set ATTRIBUTE value in instance itself ?
  - I mean instead of all the way going through the Method? Answer is Yes refer the below code
  - (Its NOT the right way, refer further)

```python
#!/usr/bin/python

class MyClass():
  def set_val(self, val):
    self.value = val

  def get_val(self):
    return self.value

a = MyClass()
b = MyClass()

a.set_val(10)
b.set_val(100)

# Here we are Directly setting the value to the ATTRIBUTE of method without invoking the Setter method
a.value = "Hi There"

print(a.get_val())
print(b.get_val())
```

Excute the above code has goes below,

```bash
./5A-Class-Values-WO-setter.py
Hi There
100
```

**From the above:**
- Originally we have set a to `10` in `[a.set_val(10)]` but then directly w/o interaction of the Method in the class we have setup the values to `"Hi There"`
  - Means that we accessed the ATTRIBUTE and setup the value. So there is `NO` restriction to do this way. Well, then why IF we can do it in main body of the code, then why
  - we have to write the whole method to set the value then ?
    - The reason is, by using such a method provides a Kind of Gateway to do any operation with the instances.
    - In this gateway we can ensure the integrity of the Instance attribute!

- Let's go by another example

```python
#!/usr/bin/python

class MyInteger(object):
  def set_val(self, val):
    try:
      val = int(val)
    except:
      return
    self.val = val

  def get_val(self):
    return self.val

  def increment_val(self):
    self.val = self.val + 1

i = MyInteger()

i.set_val(9)
print i.get_val()   # Will return "9"

i.set_val('Hi')
print i.get_val()   # Will return "9", since "Hi" will through away with "except"
```

**Explanation:** (About the what code doinng)
- It has the `set_val` method where it changing the any value to interger & saving it to `val`. Then later GETs the same & INCREMENT the same
- Please note the "Try & Except" in case if its fails to convert the given value to an integer it simply "return" (Return nothing)
- What if we think smart that chage the ATTRIBUTE value directly in body of the code (we did it in the previous example, here we are trying to that's wrong)

- For chaning the ATTRIBUTE value directly please refer the below code

```python
#!/usr/bin/python

class MyInteger(object):
  def set_val(self, val):
    try:
      val = int(val)
    except:
      return
    self.val = val

  def get_val(self):
    return self.val

  def increment_val(self):
    self.val = self.val + 1

i = MyInteger()

# Here we are DIRECTLY setting the ATTRIBUTE values
i.val = 'Hi'
print i.increment_val()
```
- when you execute the above code you will get an error !

```bash
    ./5c-Class-Increment-attribute-change-directly.py
    Traceback (most recent call last):
      File "./5c-Class-Increment-attribute-change-directly.py", line 21, in <module>
        print i.increment_val()
      File "./5c-Class-Increment-attribute-change-directly.py", line 15, in increment_val
        self.val = self.val + 1
    TypeError: cannot concatenate 'str' and 'int' objects
```

- From the above output you can see, it BYPASS the "try & except" goes for increment. This is what we are taking about the `GATEWAY`. So don't SET the values in body of script by DIRECTLY setting some values for ATTRIBUTE

**Thing that Learned**

- So, by requiring the `set_val()` method we ensure the integrity of the `val` value as INT. This is what we meant by an Encapsulation.
- Setting the value of method attribute, without calling "Setter" is we called "Breaking the Encapsulation". Big `NO` step

**Further**
- We are going to see how we can enforce values should go through even we set DIRECTLY

**Encapsulation**

- 1st of 3 pillars in OOP
- As referred to an safe storage of the data (as attribute) in an instance
- Data should only accessed through instance methods
- Data should validated as correct (depending on the requirement set in the class methods)
- Data should be safe from changes by external processes

Breaking the Encapsulation:
- Although normally set in a setter method, instance attribute values can be seet anywhere
- Encapsulation in Python is a voluntary restriction
- Python does not implement data hiding, as other language does

**Init Constructor**

**The` __init__` constructor**
- Its special method which is allows to initialize the attributes at the time instance is constructed
- The double underscore `__init__` considered as the PRIVATE/MAGIC method
  - PRIVATE => Because its intention to use internally, not called by the user of the class
  - MAGIC => Its called automatically when the particular event happens
  - `__init__` is automatically called when the instances is constructed)

- So "init" is called automatically when the instance is initialized, refer the below example

```python
#!/usr/bin/python

class MyNum(object):
  def __init__(self):
    print "calling the __init__"
    self.val = 0

  def increment(self):
    self.val = self.val + 1

dd=MyNum()
dd.increment()
dd.increment()

print dd.val
```
- Another example with passing some values

```python
#!/usr/bin/python

class MyNum(object):
  def __init__(self, value):
    print "calling the __init__"
    self.val = value

  def increment(self):
    self.val = self.val + 1

dd=MyNum(5)
dd.increment()
dd.increment()

print dd.val
```
- From the above example, you can see we passed the value "5" to `__init__` & then incremented
- There is a chance we can pass the some "STRING" value also, but that not supposed to be. Let's see how we can set the integrity GATE to the init method
- Refer the below code INTEGRITY check

```python
#!/usr/bin/python

class MyNum(object):
  def __init__(self, value):
    try:
      value = int(value)
    except ValueError:
      value = 0
    self.val = value

  def increment(self):
    self.val = self.val + 1

aa = MyNum('Hello')
bb = MyNum(100)

aa.increment()
aa.increment()
bb.increment()

print aa.val
print bb.val
```
 - From the above the, even we passed the "string" value as argument to init but it got ingore due the INTEGRITY check we made via TRY&EXCEPT

 **Notes**
- `__init__` is a keyword variable, It must be named as `init`
- Its a method automatically called when a new instance is constructed
- If its not present, it is not called
- The `self` argument is the 1st appearence of the instance
- `__init__` offers the opprtunity to initialize attributes in the instance at the time of construction

**Class Attributes**

This section describes about - Class valus vs Instance value/attributes

- Below code a very Normal Class

```python
#!/usr/bin/python

class YourClass(object):
  classy = 10

  def set_value(self):
    self.insty = 100

dd = YourClass()

dd.set_value()
print(dd.classy)
print(dd.insty)
```

- Here is an another example to demostrate the ATTRIBUTE preference in the INSTANCE

```python
#!/usr/bin/python

class YourClass(object):
  classy = 'Class Value !'

dd = YourClass()

print(dd.classy)

dd.classy = "Inst Value"

print(dd.classy)

del dd.classy

print(dd.classy)
```

Output

```bash
./7A-Inst-Attribute-Preference.py
Class Value !
Inst Value
Class Value !
```

- From the above, it proves that ATTRIBUTE lookup order in the INSTANCE 1st it will lookup itself ! Then only ot will go for Class value/attribute

**Notes**
Class Attributes Vs. Instance Attributes
  - Attributes / Variables in the class are acccessible through the instance
  - Instance attribute also accessible from instance
  - when we use the syntax object.attribute, we are asking Python to look up the attribute
      -> First in the instance
      -> Then in the class
  - Method calls through the instance follow this lookup

**From the 6 point**
 - Variable defined in the class are "Class Attributes"
 - When we read an attribute, Python looks for it first in the instance, and then the class

**Working with Class and Instance Data**

- We know that Class data can be accessed by as many INSTANCE we create!
  So class data intend to created for "SHARE"
- Will start with an example
