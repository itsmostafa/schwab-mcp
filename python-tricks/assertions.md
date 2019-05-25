# Assertions

Python's `assert` statement is a debugging aid that tests a condition.
- If the `assert` statement is `True`:
  - Nothing happens.
  - your program continues to execute as normal
- If the `assert` statement is `False`
  - `AssertionError` exception is raised
  - The exception comes with an optional error message

Example: 
```
def apply_discount(product, discount):
    price = int(product['price'] * (1.0 - discount))
    assert 0 <= price <= product['price']
    return price
```

The `assert` statement will guarantee that the discounted prices calculated by this function cannot be lower than $0 and they cannot be higher than product’s original price.

### Assert Syntax
```
assert_stmt ::= "assert" expression1 ["," expression2]
```
- `expression1` is the condition we test
- `expression2` is an error message that's displayed if the assertion fails

At execution, the interpreter transforms each assert statement into roughly the following sequence of statements:

```
if __debug__:
    if not expression1:
        raise AssertionError(expression2)
```
- Before the assert condition is checked, there's an additonal check for the `__debug__` global variable
  - it's built under a boolean flag that's true under normal circumstances
  - false if optimizations are requested
- `expresion2` is used to pass an optional error message that will be displayed with the `AssertionError` in the traceback
  - This can simplify debugging even further

Example:
```
if cond == 'x':
    do_x()
elif cond == 'y':
    do_y()
else:
    assert False, (
        'This should never happen, but it does
        occasionally. We are currently trying to
        figure out why. Email me if you encounter
        this in the wild. Thanks!')
```

### Common pitfalls of using Asserts
- Asserts for Data Validation
- Asserts that never fail

**Don't use Asserts for Data Validation**
- Assertions can be globally disabled with
  - the `-o` and `--o` command line switches
  - `PYTHONOPTIMIZE` environment variable in CPython
  - This is an intentional design

**Incorrect** way of using assertions example:
```
def delete_product(prod_id, user):
    assert user.is_admin(), 'Must be admin'
    assert store.has_product(prod_id), 'Unknown product'
    store.get_product(prod_id).delete()
```
- Checking for admin previleges with an `assert` statement is dangerous
  - if assertions are disabled, *any* user can delete products.
- The `has_product` check is skipped when assertions are disabled
  - `get_product` can now be called with invalid product IDs

**Correct** way of doing data validations by **not** using assertions example:
```
def delete_product(product_id, user):
    if not user.is_admin():
        raise AuthError('Must be admin to delete')
    if not store.has_product(product_id):
        raise ValueError('Unknown product id')
    store.get_product(product_id).delete()
```
- This code now raises semantically correct `AuthError` or `ValueError`

**Asserts that never fail**

An assertion that will never fail example:
```
assert(1 == 2, 'This should fail')
```
- If you pass a tuple to an assert statement, it will lead to an assert condition always being true
  - This will lead to the above `assert` statement being *useless* becuase it will never fail and trigger an exception

Bad multi-line assertion example:
```
assert (
    counter == 10,
    'It should have counted all the items'
)
```
- This assertion will always evaluate as to True
  - This is because it asserts the `True` value of a tuple object

Tips:
- Use a Code linter to prevent this syntax quirk
- Newer versions of Python 3 will show syntax warnings for these dubious asserts

### Summary
- Assertions speed up debugging efforts considerably
- Make programs more maintainable in the long run
- Assertions are meant to be *internal self checks* for your program
- Not a mechanism for handling runtime errors
- An assertion error should **never** be raised unless there's a bug in your program
- Asserts can be globally disabled with an interpreter setting