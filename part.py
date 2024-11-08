def partition(arr, low, high):
    pivot = arr[low]
    i = low - 1
    j = high + 1

    while (True):

        # Find leftmost element greater than
        # or equal to pivot
        i += 1
        while (arr[i] < pivot):
            i += 1

        # Find rightmost element smaller than
        # or equal to pivot
        j -= 1
        while (arr[j] > pivot):
            j -= 1

        # If two pointers met.
        if (i >= j):
            return j

        arr[i], arr[j] = arr[j], arr[i]

arr = [14, 11, 1, 9, 15, 12, 17, 11]
partition(arr, 0, len(arr)-1)
print(arr)
