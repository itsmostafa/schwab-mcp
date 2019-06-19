# `bytes` – an immutable sequence of bytes 

`bytes` type is similar to the `str` type, except:
- rather than each instance being a sequence of Unicode code points, each instance is a sequence of bytes.
- bytes objects are used for raw binary data and fixed-width, single-byte character encodings, such as ASCII. 

#### Literal `bytes`

As with strings they have a simple, literal form delimited by either single or double quotes.

**Literal** `bytes` the opening quote must be preceded by a lower-case b

for example:

```
>>> b'data'
b'data'
>>> b"data"
b'data' 
```

There is also a bytes constructor, but it has fairly complex behavior. 
For now, it's sufficient to recognize bytes literals and understand that they support many of the same operations as
`str`, such as indexing and splitting: 

```
>>> d = b 'some bytes'
>>> d.split()
[ b 'some' , b 'bytes' ] 
```

#### Converting `bytes` and `str`

- Know theencoding of the byte sequence used to represent the string’s Unicode code points as bytes. 
- Python supports a wide-variety of codecs 
    - UTF-8
    - UTF-16
    - ASCII
    - Latin-1
    - Windows-1251
    - [List of Codecs supported](https://docs.python.org/3/library/codecs.html#standard-encodings)

In Python we can *encode* a Unicode `str` into a `bytes` object and vice versa, we can *decode* a `bytes` object into a Unicode `str`.
