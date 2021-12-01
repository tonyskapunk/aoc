#!/usr/bin/env python

import re
import sys

if len(sys.argv) != 2:
    print("Error: Just one argument with a filename is expected")
    sys.exit(1)

filename = sys.argv[1]

with open(filename, "r") as f:
    lines = f.read().splitlines()
    f.close()


def one(lines):
    two = 0
    three = 0
    for line in lines:
        l = list(line)
        s = set(line)
        if len(s) == len(l):
            continue
        found_two = False
        found_three = False
        for c in s:
            freq = line.count(c)
            if freq == 3 and not found_three:
                #print(f'found (three) {freq} {c} in: {line}')
                found_three = True
            if freq == 2 and not found_two:
                #print(f'found (two)   {freq} {c} in: {line}')
                found_two = True
            if found_two and found_three:
                break
        if found_three:
            three += 1
        if found_two:
            two += 1
    #print(f'Two founds: {two}')
    #print(f'Three founds: {three}')
    return two * three


def two(lines):
    boxes = list(lines)
    rest_of_boxes = list(lines)

    # Loop through boxes, box will be compared to the rest of boxes
    for box in boxes:
        # Generate the rest of boxes
        rest_of_boxes.remove(box)
        # Loop through the rest of boxes
        for xob in rest_of_boxes:
            # Get the chars that are diff on each box ID
            in_box = [x for x in box if x not in xob]
            in_xob = [x for x in xob if x not in box]
            # Go to the next iteration if there is more than one char diff.
            if ((len(in_box) > 1 or len(in_xob) > 1) or
               not(in_box or in_xob)):
                continue
            # Construct lists with for both box IDs
            l_box = list(box)
            l_xob = list(xob)

            # When there is only one char diff in box
            # Remove the diff char from box
            # Remove the char in the same index on xob
            if not in_box:
                del l_box[xob.index(in_xob[0])]
                l_xob.remove(in_xob[0])

            # When there is only one char diff in xob
            # Remove the diff char from xob
            # Remove the char in the same index on box
            if not in_xob:
                del l_xob[box.index(in_box[0])]
                l_box.remove(in_box[0])

            # When one char is diff on either box ID
            # Check if those are on the same index
            if in_box and in_xob:
                # If the diff chars are not in the same index go to the next
                # iteration.
                if box.index(in_box[0]) != xob.index(in_xob[0]):
                    continue
                # Remove the chars on the same index
                l_box.remove(in_box[0])
                l_xob.remove(in_xob[0])

            # Recreate the strings from the lists
            s_box = "".join(l_box)
            s_xob = "".join(l_xob)
            # Compare the new box IDs
            if s_box == s_xob:
                return s_box


print(one(lines))
print(two(lines))
