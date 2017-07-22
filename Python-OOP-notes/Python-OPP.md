# Python _ OOP

### Chapter 01: Introduction

##### What I have got here ?

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