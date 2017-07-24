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
  - This is pretty clear that "thisjoe.callme()" still passing to arguments to the Class Method still nothing in it !

- Here is the MAGIC !
    When we call an METHOD on a an INSTANCE, the INSTANCE gets passed as the first argument to the METHOD automatically !
    This is the default IMPLICIT behaviour that we shoud aware !
    So now you understand what is "self" in there ! Its the instance ! Lets try to print the "self" then !

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
