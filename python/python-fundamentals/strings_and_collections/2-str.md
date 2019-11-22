# `str` - an immutable sequence of Unicode points

A **string** is a sequence of Unicode code-points, and for the most part
you can think of code-points as being like characters, although they aren’t
strictly equivalent. 

The sequence of code-points in a Python
string is **immutable**, so once you’ve constructed a string, you can’t modify its contents.

Literal strings in Python are delimited by quotes: 

`'This is a string'`

You can use single quotation marks, as we have above. Or you can use double quotation marks, as shown below: 

`"This is also a string"`

You must, however, be consistent. For example, you can’t use a double quotation mark paired with a single quotation mark otherwise you'll get this error:

```
"inconsistent' File "<stdin>" , line 1 "inconsistent' 

SyntaxError : EOL while scanning string literal
```

#### Concatenation of adjacent strings

```
>>> "first" "second"

'firstsecond' 
```

#### Multiline strings

```
>>> '''So ... is ... this.'''

'So \n is \n this.' 
```

#### Universal Newline Support
Python 3 has a feature called **universal newline support** which translates from the simple
`\n` to the native newline sequence for your platform on input and output. You can read more about Universal Newline Support in [PEP 278](https://www.python.org/dev/peps/pep-0278/).

#### The `str` Constructor

We can use the `str` constructor to create strings representations of
other types, such as integers:
```
>>> str ( 496 )
>>> '496'
```
or floats:
```
>>> str ( 6.02e23 )
'6.02e+23' 
```

#### Strings as sequences

Strings are **sequence types**, which means they
support certain common operations for querying ordered series of elements.

For example, we can access individual characters using square brackets with a zero-based integer index:
```
>>> s = 'parrot'
>>> s [ 4 ] 'o' In contrast 
```

Unlike other programming languages, there isn't a separate character type distinct from the string type. The indexing operation returns a full-blown string that just contains a single code point element, a fact we can demonstrate using Python’s built-in type() function:

```
>>> type ( s [ 4 ])
< class ' str '>
```

#### String methods

String objects also support a wide variety of operations implemented as
methods. We can list those methods by using help() on the string type:

`help(str)`

#### Strings with unicode

Strings are fully Unicode capable, so you can use them
with international characters easily, even in literals

The default source code encoding for Python 3 is **UTF-8**.

You can use the hexadecimal representations of Unicode
code points as an escape sequence prefixed by `\u `

You can use the `\x` escape sequence followed by a 2-character
hexadecimal string to include one-byte Unicode code points in a string literal