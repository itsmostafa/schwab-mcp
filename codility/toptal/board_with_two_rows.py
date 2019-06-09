'''
There is a board with 2 rows and N columns, represented by a matrix M. Rows are
numbered form 0 to 1 from top to bottom and columns are numbered from 0 to N-1
from left to right. Each cell contains either a 0 or a 1. You know that:
    - the sum of integers in the 0th (upper) row is equal to U
    - the sum of integers in the 1st (lower) row is equal to L
    - the sum of integers in the kth column is equal to C[K]

Your job is to recover M based on this information.

Write a function:

    def solution(U, L, C)

that, given two integers U, L, and an array C of N integers, as described above,
returns a string describing the matrix M in the following format. The first part
of the string shouldbe the description of the upper row (N characters: 0 or 1),
then there should comma (,), and finally there should be the description of the lower row
(N characters: 0 or 1)
The output string should not contain any whitespace.

If there exists multiple valid Ms, your function may return any one of them. If no valid M
exists, your function should return the word IMPOSSIBLE.

Write an efficient algorithm for the following assumptions:
    - U and L are integers within the range [0.. 100,000];
    - N is an integer within the range [1.. 100,000];
    - each element of array C is an integer within the range [0.. 2];

For examples, see board_with_two_rows_examples.png
'''