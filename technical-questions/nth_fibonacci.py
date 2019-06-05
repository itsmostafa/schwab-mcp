'''
The Fibonacci sequence is defined as follows:
The first number of the sequence is 0, the second number is 1,
and the nth number is the sum of the (n-1)th and (n-2)th numbers.

Write a function that takes integers n and returns the nth fibonacci number.

Sample input: 6
Sample output: 5 (0, 1, 1, 2, 3, 5)
'''


def solution(n):
    result = [0, 1]
    counter = 3
    while counter <= n:
        next_fib = result[0] + result[1]
        result[0] = result[1]
        result[1] = next_fib
        counter += 1
    return result[1]

print(solution(6))