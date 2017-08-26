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

```python
#!/usr/bin/python

class InstanceCounter(object):
  count = 0

  def __init__(self, val):
    self.val = val
    InstanceCounter.count += 1   # Here we are accessing the class data in "object.attribute" format. Is that mean Class also an object - Refer below explanation
  def set_val(self, newval):
    self.val = newval
  def get_val(self):
    return self.val

  def get_count(self):
    return InstanceCounter.count

a = InstanceCounter(5)
b = InstanceCounter(13)
c = InstanceCounter(17)

for obj in (a, b, c):
  print "val of obj: %s" % (obj.get_val())
  print "count is: %s" % (obj.get_count())
```
(Reminder: Everything in Python is an object)

- As we you can the above commented item, "Is class also is an object?"
  Let's go by example:

```python
    >>> class MyClass(object):
    ...     pass
    ...
    >>> print MyClass
    <class '__main__.MyClass'>

    # If this is object, does it have attribute ?
    >>> print dir(MyClass)
    ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__']
    # Seems there are PRIVATE/MAGIC attribute

    # Ok, lets set some value/attribute (attribute name is "val")
    >>> MyClass.var = 5

    # Interestingly, the aattribute which is setup 'ed by us also included in the DIRECTORY list (you can see the "var" at the last in the list)
    >>> print dir(MyClass)
    ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'var']
```

- Another question
  
  If we are able to setup some values, then why are accessing the class value/attribute with class name
    - From the previous exampe (here - instanceCounter.count += 1)
    Reason:
      Because of global variable. If you simply call "var" it going to access the global variable instead of class attribute
      So in-order to avoid that we have to use the "object.attribute" format which is "instanceCounter.count"

- Previous code output explanation:

```bash
    val of obj: 5
    count is: 3
    val of obj: 13
    count is: 3
    val of obj: 17
    count is: 3
```
- from above output you might be asking why we are getting "3" as count ?
  Because we are running/creating the instance (all 3 via loop), so every single time __init__ got incremented & that's why we are getting the "3" value

- As of above code example we getting the count "3" via class method attrbute. As we already aware instance lookup for the attribute 1st instance itself.
  So can we access the count via instance itself? answer is "yes" please refer the below code

```python
#!/usr/bin/python

class InstanceCounter(object):
  count = 0

  def __init__(self, val):
    self.val = val
    InstanceCounter.count += 1
  def set_val(self, newval):
    self.val = newval
  def get_val(self):
    return self.val

  def get_count(self):
    return InstanceCounter.count

a = InstanceCounter(5)
b = InstanceCounter(13)
c = InstanceCounter(17)

for obj in (a, b, c):
  print "val of obj: %s" % (obj.get_val())
  print "count (from instance): %s" % (obj.count)
```

**Assignment**:
Rough Definition:
  - You have max_number & empty list
  - Append some data to the list which should <= max_number, If the list len bigger that max_number then drop the last element from the list
  - Write a Class and process the same
  - Here is the code

```python
#!/usr/bin/python

class MaxSizeList(object):

  def __init__(self, max):
    self.max_size = max
    self.innerlist = []

  def push(self, obj):
    self.innerlist.append(obj)
    if len(self.innerlist) > self.max_size:
      self.innerlist.pop(0)

  def get_list(self):
    return self.innerlist

a = MaxSizeList(3)
b = MaxSizeList(1)

a.push("one")
a.push("two")
a.push("three")
a.push("four")

b.push("five")
b.push("six")
b.push("seven")
b.push("eight")

print(a.get_list())
print(b.get_list())
```

**Note for myself:**
- Check out import the custom class to module & using those module in your script

**Inheritance and Polymorphism**
**Inheritance Attributes**

**Recap:**
**Encapsulation**
  - Its a facility store the Data and Objects
  - Provides data Integrity via Setter and mGetter methods

**Inheritance:**
  - Ability that, one Class inherits the attributes of another Class

- Lets write some code example:

```python
#!/usr/bin/python

class Date(object):       # Inherts from the 'object' class (not explained as of now - its there below)
  def get_date(self):
    return '2020-03-03'

class Time(Date):         # Have look at this declaration, This is inherits from the 'Date' class
  def get_time(self):
    return '01:02:03'

dt = Date()
print(dt.get_date())

tm = Time()
print(tm.get_time())
print(tm.get_date())    # Found this method in the 'Date' class, Since we declared "Date" object as ARG in TIME we are able to call it with tm.xx()
```

- So class `Time` inherits from `Date` & `Date` inherits from the `CLASS` called object
 (=> "object" built-in class provided by python)

- So we have created the object `Date()` with `dt`
- Also we have created the object `Time()` with `tm`

- now `dt` calling its class method - `get_date()`
- also `tm` calling its Object method - `get_time()` adding its also calling `get_date()` method, because the we pahe placed the Date Class name in the Argument list

- If you remove DATE from this line `class Time(Date):` and run the script it will fail

- In general Python will follow the `LOOKUP` inorder to find the things
- Earlier we were discussing about the ATTRIBUTE lookup, now we are going to the `object.lookup` (since we are Inheriting the another class in to the class)
- From the above code `print(tm.get_time())` is the `object.attribute()` lookup
- Python doesn't care wether the its `Arrtibute, Class.attribute, object.attribute()` by default its considered as the ATTRIBUTE lookup
- Every single case python will follow the these place inorder to do the Attribute lookup
  - Instance itself
  - The Class (Instance belongs to)
  - The Class from which Class this inherits (basically looking for the parent class)

**Some Inheritance Terms**
- An Inheriting class known as follows: `class MyClass(YourClass)`
  - Child class
  - Derived class
  - Subclass

- An Inherited class known as follows: `Class YourClass(object)`
(Also an child class if "object" but in this case)
  - Parent class
  - Base class
  - Super class

**Inheritance**

Inheritance: The 2nd pillar of OOP

One class can inherit from the another

  - The class attributes are inherited
  - In particular, its methods are inherited
  - This means that instances of an inheriting (child) class an access attributes of the inheritited (parent) class
  - This is simply another level of attribute lookup: Instance, then class, then inherited class

**Inheritance examples**

- Lets start with the code

```python
#!/usr/bin/python

class Animal(object):
  def __init__(self, name):
    self.name = name
  def eat(self, food):
    print '%s is eating %s.' % (self.name, food)

class Dog(Animal):
  def fetch(self, thing):
    print '%s goes after the %s' % (self.name, thing)

class Cat(Animal):
  def swatstring(self):
    print '%s shreds the string' % (self.name)

r = Dog('Rover')
f = Cat('Fluffy')

r.fetch('Paper')    # Rover goes after the Paper
f.swatstring()       # Fluff shreds the string
r.eat('Dog food')   # Rover eating Dog food
f.eat('CAt food')   # Fluffy eating Cat food

# Below tto simply through an error
# AttributeError: 'Dog' Object has  no attribute 'swatstring'
r.swatstring()
```

Above script is self explanatory
- Parent class is `Animal`
- `Animal` class access from child class `Dog` & `Cat`
- Both Child class invoking the Parent class methods `eat`
- When Dog child class trying to access the `Cat` child's method you get error, since attribute couldn't find by python. Cus we are invoking this no-where in the code

**Polymorphism**
("Many Shapes")

**Definition:**
  - 3rd pillar of OOP
  - Two class with the same interface
  - The methods() are often different, but conceptually similar
    refer the below code for reference

- Lets go by example

```python
#!/usr/bin/python

class Animal(object):

  def __init__(self, name):
    self.name = name

  def eat(self, food):
    print '{0} eats {1}'.format(self.name, food)

class Dog(Animal):

  def fetch(self, thing):
    print '{0} goes after the {1}!'.format(self.name, thing)

  def show_affection(self):
    print '{0} wags tail'.format(self.name)

class Cat(Animal):

  def swatstring(self):
    print '{0} shreds the string !'.format(self.name)

  def show_affection(self):
    print '{0} purs'.format(self.name)

for a in (Dog('Rover'), Cat('Fluffy'), Cat('Precious'), Dog('Scout')):
  a.show_affection()

# Above line is the Polymorphism call
```

**Some Explanation:**
- So both two class Dog & Cat are inherits from the Animal(Super/Parent) class
- from the above code both Dog/Cat were inherits `eat()` method from the Parent class

- But `show_affection()` is not inherit method
  - They have same `NAME`, well functionally same but conceptually different & also they stick with their class.
  - They independent method with teir class

This is POLYMORPHISM - these method is more of a Technic

***Adding some more information for notes***
 - Allows for expressiveness in Design: we can say that this group of related classes implement the same action
 - Duck typing refers to reading an object's attribute to decide whether it is of proper type, rather than checking the type itself


- Python itself has classes & Polymorphic
  This is happens pretty much for almost most of the operations

- For example `len()` function can be used for multiple objects, means that it can play along with string, list, tuples & dict

```python
    >>> len('Yolo')
    4
    >>> len([1, 2, 3, 4])
    4
    >>> len(('a', 'b'))
    2
    >>> len({'a': 1, 'c': 3})
    2
```

- You may ask `len()` is build-in function, how it can be polymorphic when itsn't even a method.
- Here is magic/truth

  `len()` is built-in function, but basically when we use they translate the method call on the object being passed to them

```python
    >>> var = "Hi"
    >>> len(var)
    2
    # This is method() of "STRING" object
    >>> var.__len__()
    2
```
  - Checking the availble method() for the object "var"
  - Basically we should call Attribute, not necessarily moethod here !

```python
    >>> dir(var)
    ['__add__', '__class__', '__contains__', '__delattr__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getnewargs__', '__getslice__', '__gt__', '__hash__', '__init__', '__le__', '__len__', '__lt__', '__mod__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__rmod__', '__rmul__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '_formatter_field_name_split', '_formatter_parser', 'capitalize', 'center', 'count', 'decode', 'encode', 'endswith', 'expandtabs', 'find', 'format', 'index', 'isalnum', 'isalpha', 'isdigit', 'islower', 'isspace', 'istitle', 'isupper', 'join', 'ljust', 'lower', 'lstrip', 'partition', 'replace', 'rfind', 'rindex', 'rjust', 'rpartition', 'rsplit', 'rstrip', 'split', 'splitlines', 'startswith', 'strip', 'swapcase', 'title', 'translate', 'upper', 'zfill']
```

- Basically we are using them INDIRECTLY without knowing them. i mean `len()` against the string/tuples & others

- Ok, let try with List/Tuples that how many associated method/attributes are availabel

```python
    >>> dir([1, 2, 3, 4])
    ['__add__', '__class__', '__contains__', '__delattr__', '__delitem__', '__delslice__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getslice__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__setslice__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']
    >>> dir(('a', 'b'))
    ['__add__', '__class__', '__contains__', '__delattr__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getnewargs__', '__getslice__', '__gt__', '__hash__', '__init__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__rmul__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', 'count', 'index']
```

**Inheriting the Constructor**

- lets start eith code example

```python
#!/usr/bin/python

class Animal(object):

  def __init__(self, name):
    self.name = name

  def eat(self, food):
    print '{0} eats {1}'.format(self.name, food)

class Dog(Animal):

  def fetch(self, thing):
    print '{0} goes after the {1}!'.format(self.name, thing)

d = Dog('dogname')

print d.name
```

- Another example

```python
Code ===================
#!/usr/bin/python
import random

class Animal(object):

  def __init__(self, name):
    self.name = name

  def eat(self, food):
    print '{0} eats {1}'.format(self.name, food)

class Dog(Animal):

  def __init__(self, name):
    super(Dog, self).__init__(name)
    self.breed = random.choice(['Mutt', 'Shih Tzu', 'Beagle', 'Huskey'])

  def fetch(self, thing):
    print '{0} goes after the {1}!'.format(self.name, thing)

d = Dog('dogname')

print d.name
print d.breed
```