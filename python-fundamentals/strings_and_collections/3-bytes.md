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

There is also a bytes constructor, but it has fairly complex behavior and we
defer coverage of it to the second book in this series, The Python Journeyman .
At this point in our journey, it’s sufficient for us to recognize bytes literals and understand that they support many of the same operations as
str , such as indexing and splitting: 

```
>>> d = b 'some bytes'
>>> d.split()
[ b 'some' , b 'bytes' ] 
```