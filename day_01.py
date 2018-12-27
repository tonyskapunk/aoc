#!/usr/bin/env python

import sys

if len(sys.argv) != 2:
    print("Error: Just one argument with a filename is expected")
    sys.exit(1)

filename = sys.argv[1]

with open(filename, "r") as f:
    lines = f.read().splitlines()
    f.close()


def one(lines):
    total = 0
    for line in lines:
        total += int(line)
    return total


def two(lines):
    total = 0
    frequency = [total]
    index = 0
    while True:
        for line in lines:
            total += int(line)
            if total in frequency:
                return total
            frequency.append(total)


print(one(lines))
print(two(lines))
