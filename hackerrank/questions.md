### 1. Solve me first

```
function solveMeFirst(a, b) {
  // Hint: Type return a+b below
    return a + b;
}
```

### 2. Simple array sum

```
function simpleArraySum(ar) {
    /*
     * Write your code here.
     */
    var total = 0;
    for (var i = 0; i < ar.length; i++) {
        total += ar[i];
    }
    return total;
}
```

```
def simpleArraySum(ar):
    #
    # Write your code here.
    #
    total = 0
    for i in ar:
        total += i
    return total
```

```
def simpleArraySum(ar):
    #
    # Write your code here.
    #
    return sum(ar)
```

### 3. Compare the Triplets

```
def compareTriplets(a, b):
    length = len(a)
    alice = 0
    bob = 0
    for i in range(0, length):
        if a[i] > b[i]:
            alice += 1
        elif a[i] < b[i]:
            bob += 1
    return [alice, bob]
```

```
function compareTriplets(a, b) {
    var alice = 0;
    var bob = 0;
    for (var i = 0; i < a.length; i++) {
        if (a[i] > b[i]) {
            alice++;
        } else if (a[i] < b[i]) {
            bob++;
        }
    }
    return [alice, bob]
}
```

```
def compareTriplets(a, b):
    ret = [0, 0]
    for x, y in zip(a, b):
        if x > y:
            ret[0] += 1
        elif x < y:
            ret[1] += 1
    return ret
```

### 4. A Very Big Sum
```
function aVeryBigSum(ar) {
    /*
     * Write your code here.
     */
    var total = 0;
    for (var i = 0; i < ar.length; i++) {
        total += ar[i];
    }
    return total;
}
```

### 5. Diagonal Difference
```
def diagonalDifference(arr):
    length = len(arr[0])
    left = []
    right = []
    for i in range(length):
        left.append(arr[i][i])
        x = length-i-1
        right.append(arr[i][x])
    
    left_total = sum(left)
    right_total = sum(right)

    return abs(left_total - right_total)
```

```
function diagonalDifference(arr) {
    var sum = function(list) {
        var total = 0;
        for (var i = 0; i < list.length; i++) {
            total += list[i];
        }
        return total;
    }

    var left = [];
    var right = [];
    for (var i = 0; i < arr.length; i++) {
        left.push(arr[i][i]);
        right.push(arr[i][arr.length - i - 1])
    }

    return Math.abs(sum(left) - sum(right))
}
```

```
def diagonalDifference(arr):
    length = len(arr[0])
    left, right = [], []
    for i, _list in enumerate(arr):
        left.append(_list[i])
        right.append(_list[length-i-1])
    return abs(sum(left) - sum(right))
```

```
def diagonalDifference(arr):
    length = len(arr[0])
    left, right = 0, 0
    for i, _list in enumerate(arr):
        left += _list[i]
        right += _list[length-i-1]
    return abs(left - right)
```

### 6. Plus Minus

```
def plusMinus(arr):
    length = len(arr)
    minus = round(float(len([i for i in arr if i < 0])/length), 6)
    plus =  round(float(len([i for i in arr if i > 0])/length), 6)
    zero = round(float(1 - plus - minus), 6)
    print(plus)
    print(minus)
    print(zero)
```

```
function plusMinus(arr) {
    var length = arr.length;
    var plus = 0;
    var minus = 0;
    for (var i = 0; i < arr.length; i++) {
        if (arr[i] > 0) {
            plus++;
        } else if (arr[i] < 0) {
            minus++;
        }
    }
    var lenPlus = (plus / length).toFixed(6);
    var lenMinus = (minus / length).toFixed(6);
    console.log(lenPlus);
    console.log(lenMinus);
    console.log((1 - lenPlus - lenMinus).toFixed(6));
}
```

### 7. Staircase

```
def staircase(n):
    ret = '#'*n
    for i in range(1, n+1):
        print(' '*(n-i) + ret[:i])
```

```
function staircase(n) {
    var ret = "#".repeat(n);
    for (var i = 1; i <= n; i++){
        console.log(' '.repeat(n-i) + ret.substring(0, i));
    }
}
```

### 8. Mini-Max Sum
