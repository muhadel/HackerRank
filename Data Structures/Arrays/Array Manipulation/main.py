def arrayManipulation(n, queries):
    arr = [0] * n
    for i in queries:
        for j in range(i[0], i[1] + 1):
            arr[j - 1] += i[2]
    return max(arr)

'''
def arrayManipulation(n, queries):
    arr = [0] * n
    for i in queries:
        arr[i[0] - 1] += i[2]
        arr[i[1]] -= i[2]
        print(arr)
    maxval = 0
    itt = 0
    for q in arr:
        itt += q
        if itt > maxval:
            maxval = itt
    return maxval
'''
if __name__ == '__main__':
    queries=[
        [1, 5, 3],
        [4, 8, 7],
        [6, 9, 1]
        ]
    n = 10
    print(arrayManipulation(n, queries))