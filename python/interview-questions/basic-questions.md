# Python Basic Questions

1. ### What is Python and what are its key features?

- Python is an interpretive computer programming language, both object-oriented and interactive, and used for scripting. It is designed as the most readable computer language in the world.
- Its key features are:
  - It is interpretive which means you do not need to compile the program before you run it; it is interpreted at runtime by the interpreter, unlike other languages like C.
  - It is a dynamically-typed language and this means there is no need to define datatypes for the variables you declare, or anything else for that matter.  For example, a variable called x=10 can be declared, followed by x=”Hello World” and, providing there are no errors, the datatype will be defined as per the value.
  - Python functions are first-class objects.
  - Python may be used for many applications, cross-platform, including web applications, big data applications, scientific models, and a whole lot more.
  - It contains objects, modules, threads, exceptions and automatic memory management.

2. ### What is the difference between a List and a Tuple?
- The main difference between lists and a tuples is the fact that lists are mutable whereas tuples are immutable.
- A mutable data type means that a python object of this type can be modified.
- An immutable object can’t be modified.

3. ### What operator types does Python use?

- Arithmetic operators:
    - Addition (+)
    - Subtraction (-)
    - Multiplication (*)
    - Division (/)
    - Exponentiation (**)
    - Floor Division (//)
    - Modulus (%)

- Relational operators:
    - Less than (<)
    - Greater than (>)
    - Less than or equal to (<=)
    - Greater than or equal to (>=)
    - Equal to (==)
    - Not equal to (!=)

- Assignment operators:
    - Assign (=)
    - Add and Assign (+=)
    - Subtract and Assign (-=)
    - Divide and Assign (/=)
    - Multiply and Assign (*=)
    - Modulus and Assign (%=)
    - Exponent and Assign (**=)
    - Floor-Divide and Assign (/=)

- Logical operators:
    - and
    - or
    - not

- Bitwise operators

4. ### What is PEP-8

- PEP 8 is a coding convention, a set of recommendation, about how to write your Python code more readable

5. ### What is pickling and unpickling?

- Pickle module accepts any Python object and converts it into a string representation and dumps it into a file by using dump function, this process is called pickling.
- The process of retrieving original Python objects from the stored string representation is called unpickling.

6. ### How is Python Interpreted?

- Python program runs directly from the source code. It converts the source code that is written by the programmer into an intermediate language, which is again translated into machine language that has to be executed.

7. ### How is memory managed in Python?

- Python memory is managed by Python private heap space. All Python objects and data structures are located in a private heap. The programmer does not have an access to this private heap and interpreter takes care of this Python private heap.
- The allocation of Python heap space for Python objects is done by Python memory manager. The core API gives access to some tools for the programmer to code.
- Python also have an inbuilt garbage collector, which recycle all the unused memory and frees the memory and makes it available to the heap space.

8. ### How are arguments passed by value or by reference?

- Everything in Python is an object and all variables hold references to the objects. The references values are according to the functions; as a result you cannot change the value of the references. However, you can change the objects if it is mutable.
