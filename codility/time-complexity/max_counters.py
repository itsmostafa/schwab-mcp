def solution(N, A):
    current_max = 0
    ret = [0] * N
    
    for X in A:
        if 1 <= X <= N:
            current_index = X-1
            ret[current_index] += 1
            
            current_value = ret[current_index]
            if current_value > current_max:
                current_max = current_value
        else:
            ret = [current_max] * N
    
    return ret