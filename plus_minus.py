#!/usr/bin/python3


def plusMinus(arr):
    denom = len(arr)
    positive_count = 0
    negative_count = 0
    zero_count = 0
    for number in arr:
        if number > 0:
            positive_count += 1
        elif number < 0:
            negative_count += 1
        else:
            zero_count += 1
    print(positive_count / denom)
    print(negative_count / denom)
    print(zero_count / denom)
    return [positive_count / denom, negative_count / denom, zero_count / denom]


plusMinus([1, 2, 0, -1, -2, 4])
