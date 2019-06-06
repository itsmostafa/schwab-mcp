'''
A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

For example, consider array A such that:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
We can split this tape in four places:

P = 1, difference = |3 − 10| = 7 
P = 2, difference = |4 − 9| = 5 
P = 3, difference = |6 − 7| = 1 
P = 4, difference = |10 − 3| = 7 
Write a function:

def solution(A)

that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

For example, given:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1,000..1,000]
'''


def solution(A):
    sum_of_A = sum(A)
    sum_of_left_side = 0
    differences = []
    result = 0
    for i in A:
        sum_of_left_side += i
        differences.append(abs(sum_of_left_side-(sum_of_A-sum_of_left_side)))

    result = min(differences[:-1])

    return result


## Tests
list_1 = [3, 1, 2, 4, 3]
print(solution(list_1))

list_2 = [11, 7, 11, 9, 4, 5, 1]
print(solution(list_2))

list_3 = [-3, -5, -7, -8, -9, -1]
print(solution(list_3))

list_4 = [5, 6, 30]
print(solution(list_4))

list_5 = [1, 1, 1, 1, 1]
print(solution(list_5))

list_6 =[0, 0, 0]
print(solution(list_6))

list_7 = [1, 2]
print(solution(list_7))

list_8 = [0, 0, 2]
print(solution(list_8))

list_9 = [93, 1, 85, 1, 103, 166]
print(solution(list_9))

list_10 = [1, 3, 6, 4, 1, 2]
print(solution(list_10))
