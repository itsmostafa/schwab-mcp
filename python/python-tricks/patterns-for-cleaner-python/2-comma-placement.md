# Complacent Comma Placement

When adding and removing items from a list, dict, or set constant in Python: *Just end all of your lines with a comma.* 

For example:
```
>>> names = ['Alice', 'Bob', 'Dilbert']
```

Most source control systems are line-based and have a hard time highlighting multiple changes to a single line.

For example, it’ll be hard to tell what was modified by looking at a `git diff`

A quick fix for that is to adopt a code style where you spread out list, dict, or set constants across multiple lines, like so:
```
>>> names = [
...     'Alice',
...     'Bob',
...     'Dilbert'
... ]
```

That way there’s one item per line, making it perfectly clear which one was added, removed, or modified when you view a diff in your source control system. It will make it easier for you and your teammates to review code changes. 

However you'll need to watch out for **string literal concatenation** behavior:

```
>>> names = [
...     'Alice',
...     'Bob',
...     'Dilbert' # <- Missing comma!
...     'Jane'
... ]
```

will output:

```
>>> names
['Alice', 'Bob', 'DilbertJane']
```

Instead, place a comma after every item, including the last item like this:

```
>>> names = [
...     'Alice',
...     'Bob',
...     'Dilbert',
...     'Jane',]
```

#### String literal concatenation

String literal concatenation is a useful feature in some cases. For example, you can use it to reduce the number of backslashes needed to split long string constants across multiple lines:

```
my_str = ('This is a super long string constant '
          'spread out across multiple lines. '
          'And look, no backslash characters needed!') 
```

#### Key takeaways

- Smart formatting and comma placement can make your list, dict, or set constants easier to maintain.
- Python’s string literal concatenation feature can work to your benefit, or introduce hard-to-catch bugs.
