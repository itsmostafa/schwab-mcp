'''
Write a function:

def solution(A)

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
'''

def solution(A):
    final_list = []
    for i, x in enumerate(list(set(A)) + [None], 1):
        if i != x:
            final_list.append(i)
    if len(final_list) == 0:
        return 1
    else:
        return min(final_list)

## Tests
list_1 = [1, 2, 9, 3, 9, 7, 9]
print(1, solution(list_1))

list_2 = [11, 7, 11, 9, 4, 5, 1]
print(2, solution(list_2))

list_3 = [-3, -5, -7, -8, -9, -1]
print(3, solution(list_3))

list_4 = []
print(4, solution(list_4))

list_5 = [1, 1, 1, 1, 1]
print(5, solution(list_5))

list_6 =[0, 0, 0]
print(6, solution(list_6))

list_7 = [1]
print(7, solution(list_7))

list_8 = [0, 0, 2]
print(8, solution(list_8))

list_9 = [93, 1, 85, 1, 103, 166]
print(9, solution(list_9))

list_10 = [1, 3, 6, 4, 1, 2]
print(10, solution(list_10))