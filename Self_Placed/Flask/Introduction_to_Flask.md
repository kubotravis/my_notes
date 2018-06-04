Introduction to Flask
---

Table of Content
----------------

[1. Set Up](https://github.com/kubotravis/my_notes/blob/master/Self_Placed/Flask/Introduction_to_Flask.md#1-set-up)
[2. Simple Hello World Application](https://github.com/kubotravis/my_notes/blob/master/Self_Placed/Flask/Introduction_to_Flask.md#2-simple-hello-world-application)
[3. Introduction to Templates](https://github.com/kubotravis/my_notes/blob/master/Self_Placed/Flask/Introduction_to_Flask.md#3-introduction-to-templates)
[4. Using Jinja2 Templates](https://github.com/kubotravis/my_notes/blob/master/Self_Placed/Flask/Introduction_to_Flask.md#3-using-jinja2-templates)
----

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

Taka a look at the 2a exampled code, you may noticed that there is template dir with 2 files

```bash
2a
├── hello.py
└── templates
    ├── index.html
    └── user.html

1 directory, 3 files
```

If notice there new return statement as below

```python
    return render_template('index.html')            # render_template is the function which is imported from Flask, So its basically grabs the ARG values which is 
...                                                 # the actual template files
    return render_template('user.html', name=name)  # Since its dynamic, here we are passing the ARG to the template to render at runtime
```

So inorder to render the dynamic ARG in the template we should have the placeholder in the template file, please refer the below

```html
<h1>Hello, {{ name }}!</h1>
```

- Anything with in the `{{ }}` will be replaced with variable that we provide to the function in the code

#### 4. Using Jinja2 Templates