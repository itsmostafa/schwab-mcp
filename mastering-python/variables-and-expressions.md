# Python Variables and Expressions

Variables do not contain objects. They are pointers to objects. Look at this code:

```
a = [2, 4, 6]
b = a
a.append(8)
print(b)
[2, 4, 6, 8]
```
- A variable has been created and references a list object
- Another variable has been created and references the same object
- both `a` and `b` reflect the change

#### Python is a **dynamically typed language**. 

This means:
- We can bind variable names to different **types** and **values**
- Whenever a **variable** is initialized in python, we *don't* need to declare a type
- The object the variable points to can also change its **type**

#### Variable Scope

It is important to understand the scope surrounding a variable in a function
- New local **namespaces** are created whenever a function is executed.
  - **namespaces** represent environments in which you'll find the names of each parameter and each variable that the *function* assigned
  - When a function is called, the namespace needs to be resolved. To do this, the python interpreter will search the function, if it can't find a match, it'll then search the module in which the function is defined. This is the **global namespace**
    - Otherwise the interpreter will raise a **NameError** exception
  - Unless specifically changed using `global` within a function, variables will make not reflections of local variables within the global scope

Look at this code:

```
a=20; b=30
def my_function():
    global a
    a=21; b=31
my_function()
print(a) #prints 21 print(b) #prints 30
```

#### Iteration and Flow Control

Python programs are made up of sequences of statements.
- Each statement is executed until all are done
- All statements have an equal status, regardless whether they are:
  - function definitions
  - variables assignments
  - module imports
  - class definitions
- Program exaction flow can be controlled in two ways:
  - Loops
  - Conditional Statements
  
Conditional Statements are:
- if
- elif
- else
  
Loops are:
- while
- for
