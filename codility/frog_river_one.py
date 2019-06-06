'''
A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.

For example, you are given integer X = 5 and array A such that:

  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4
In second 6, a leaf falls into position 5. This is the earliest time when leaves appear in every position across the river.

Write a function:

def solution(X, A)

that, given a non-empty array A consisting of N integers and integer X, returns the earliest time when the frog can jump to the other side of the river.

If the frog is never able to jump to the other side of the river, the function should return −1.

For example, given X = 5 and array A such that:

  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4
the function should return 6, as explained above.

Write an efficient algorithm for the following assumptions:

N and X are integers within the range [1..100,000];
each element of array A is an integer within the range [1..X].
'''

def solution(X, A):
	new_list = [0]*X
	for i, v in enumerate(A):
		new_list[i]+=1
		import ipdb; ipdb.set_trace()
		if sorted(set(new_list))==list(range(1, max(A)+1)) and X in A:
			return i
			break
	return -1


# def solution(X, A):

#     count = [0]*X
#     tot_sum = X*(X+1)/2
#     each_sum = 0
#     for i in range(len(A)):
#         if count[A[i]-1]==0:
#             count[A[i]-1]+=1
#             each_sum +=A[i]
#             if each_sum == tot_sum:
#                 return i
#     return -1

# Tests
list_1 =  [1, 3, 1, 4, 2, 3, 5, 4]
x_1 = 5
print(solution(x_1, list_1)) # returns 6

list_2 =  [2, 2, 2, 2, 2]
x_2 = 2
print(solution(x_2, list_2))

list_3 =  [1, 2, 3, 5, 3, 1]
x_3 = 5
print(solution(x_3, list_3))

list_4 =  [1, 3, 1, 3, 2, 1, 3]
x_4 = 3
print(solution(x_4, list_4)) # returns 4

list_5 =  [1]
x_5 = 1
print(solution(x_5, list_5)) # returns 0