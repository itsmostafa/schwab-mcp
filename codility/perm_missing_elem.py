'''
An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

def solution(A)

that, given an array A, returns the value of the missing element.

For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)]
'''

def solution(A):
    length = len(A) + 1
    return int(length*(1+length)/2 - sum(A))

# Tests
list_1 = [1, 2, 2, 3, 5, 6, 6]
print(solution(list_1))

list_2 = [7, 9, 8, 5, 7, 4]
print(solution(list_2))

list_3 = [2]
print(solution(list_3))

list_4 = []
print(solution(list_4))

list_5 = [2, 3, 1, 5]
print(solution(list_5))
