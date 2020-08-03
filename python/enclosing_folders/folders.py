from timeit import timeit


test_list = [str(x) for x in range(0, 100000)]


def run_loop():
    test_list = []
    for x in range(0, 100000):
        test_list.append(str(x))
    return test_list


def run_comprehension():
    test_list = [str(x) for x in range(0, 100000)]
    return test_list


# print("loop ", timeit(run_loop, number=500))
print("comprehension ", timeit(run_comprehension, number=500))

