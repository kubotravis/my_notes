Introduction to Flask
---

#### 1. Set Up

Setting up the virtual environment using `Python 2.7 & 3.4`, since the course also refers the same.

If you are using the Python 2.7

`$ virtualenv venv`

If you are using the Python3

`$ python3 -m venv venv`

Activate the above

```bash
$ source venv/bin/activate
$ which python # It should be pointing to the venv's python location
```

**Ref:**
- [Intro-to-flask-video](http://github.com/miguelgrinberg/oreilly-intro-to-flask-video.git)

#### 2. Simple Hello World Application

Please the File folder the Hello world code sample. Will paste here as well for vidsibility

```python
from flask import Flask                   # Importing the class "Flask" from package "flask"
app = Flask(__name__)                     # Creating the Instance of 'Flask' main class and storing it in the global variable "app" & also passing the arg __name__  its a
                                          # python variable. This tells the location of application. Here its search for templated & static file. (Its defacto like)

app.route('/')                            # @ - denotes its a decorator and its registers the url '/', which will run the function
def index():                              # index(): & returns the below std output
  return "<h1>Hello, World!</h1>"         # This is content diplayed in the browser

if __name__ == '__main__':                # If body
  app.run(debug=true)                     # Running the web server with debug option
```

Install the flask & run the application (Its under venv enabled)

```bash
$ pip install flask

$ python hello.py
```

Refer the comment section in the code block for the explanation.

Another example but addition of one more route, refer the below

```python
@app.route('/user/<name>')                       # <name> is arg & its dynamic. It can be anything
def user(name):                                  # Here is the Arg value used from the route which is given by client
    return '<h1>Hello, {0}!</h1>'.format(name)   # So we are storing & returing the what ever suplied in the URI as a greeting
```

#### 3. Introduction to Templates