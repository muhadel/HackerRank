def reverseArray(a):
    # print(*arr[::-1])
    for i in a[::-1]:
        print(i, end=' ')
    return a[::-1]

if __name__ == '__main__':
    arr = [1, 4, 3, 2]
    reverseArray(arr)