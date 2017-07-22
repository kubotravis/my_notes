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

- Defining the Class

```
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