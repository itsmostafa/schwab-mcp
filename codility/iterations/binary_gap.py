'''
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

def solution(N)

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
'''

def solution(N):
    binary = f'{N:0b}'
    binary_gap = False
    bin_max = 0
    bin_counter = 0
    for i in binary:
        if i == '1':
            if bin_max < bin_counter:
                bin_max = bin_counter
            binary_gap = True
            bin_counter = 0
        elif binary_gap:
            bin_counter += 1
    return binary, bin_max


## Tests
fifteen = solution(15)
print(fifteen)
thirty_two = solution(32)
print(thirty_two)
six_four_seven = solution(647)
print(six_four_seven)
one_zero_four_one = solution(1041)
print(one_zero_four_one)
'''
Alternate solution:

def solution(N):
  return len(max(bin(N)[2:].strip('0').strip('1').split('1')))
'''