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
    if not A:
        return 1
    for index, value in enumerate(A):
        if len(A) < value <= 0:
            continue
        while index + 1 != A[index] and 0 < A[index] <= len(A):
            v = A[index]
            A[index], A[v-1] = A[v-1], A[index]
            A[v-1] = v

            if A[index] == A[v-1]:
                break

    for index, value in enumerate(A, 1):
        if value != index:
            return index
    return len(A) + 1


## Tests
list_1 = [1, 2, 9, 3, 9, 7, 9]
print(solution(list_1))

list_2 = [11, 7, 11, 9, 4, 5, 1]
print(solution(list_2))

list_3 = [-3, -5, -7, -8, -9, 1]
print(solution(list_3))

list_4 = []
print(solution(list_4))