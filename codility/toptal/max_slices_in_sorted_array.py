'''
We are given an array A consisting of N distinct integers. We would like to sort array A
into ascending order using a simple algorithm. First, we divide it into one or more slices
(a slice is a contiguous subarray). Then we sort each slice. After that, we join the sorted
slices in the same order. Write a function, solution, that returns the maximum number
of slices for which the algorithm will return a correctly sorted array.

    def solution(A)

Write an efficient algorithm for the following assumptions:
    - N is an integer within the range [1.. 100,000];
    - each element of array A is an integer within the range
    [1.. 1,000,000,000]
    - the elements of A are all distinct

See max_slices_in_sorted_array.png for examples
'''