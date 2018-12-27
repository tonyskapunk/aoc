#!/usr/bin/env python

import sys
from collections import namedtuple

if len(sys.argv) != 2:
    print("Error: Just one argument with a filename is expected")
    sys.exit(1)

filename = sys.argv[1]

with open(filename, "r") as f:
    lines = f.read().splitlines()
    f.close()
# Parsing input
# Builds a dict with: claim_id: { b: [x,y], s: [x,y] }
# b -> begin
# s -> size
claims = {}
for line in lines:
    _, claim = line.split("#")
    claim_id, claim_info = claim.split("@")
    b, s = claim_info.split(":")
    claims[int(claim_id.strip())] = {
        "b": [int(b.strip().split(",")[0]),
              int(b.strip().split(",")[1])],
        "s": [int(s.strip().split("x")[0]),
              int(s.strip().split("x")[1])],
    }

# Named tuple
# P -> Position, s -> start, e -> end
P = namedtuple('P', 's e')


def get_position(claim):
    """Return a set of two pair of named tuples
    Args:
        claim: 'b': [30, 168], 's': [20, 14]}

    Returns:
        set with two named tuples, e.g (P, P)
    """
    s_x = claim['b'][0]
    e_x = claim['b'][0] + claim['s'][0]
    s_y = claim['b'][1]
    e_y = claim['b'][1] + claim['s'][1]
    return (P(s_x, e_x), P(s_y, e_y))


def one(claims):
    """ Get overlapped claims"""
    # Empty fabric of 1000x1000
    fabric = []
    for i in range(1000):
        fabric.append([0 for j in range(1000)])
    overlap = 0
    for c_id in claims.keys():
        pos = get_position(claims[c_id])
        x_range = pos[0]
        y_range = pos[1]
        for x in range(x_range.s, x_range.e):
            for y in range(y_range.s, y_range.e):
                # If the coordinate is a string, then is a overlap
                # we count this already, continue with the next iteration
                if isinstance(fabric[x][y], str):
                    continue
                # when there is a 0 it means is empty, but when there is
                # anything else, that's a coordinate already claimed, set a
                # overlap: "#"
                if fabric[x][y]:
                    fabric[x][y] = "#"
                    overlap += 1
                else:
                    fabric[x][y] = c_id
    return overlap


def two(claims):
    """ Get non overlapped claim"""
    # Empty fabric of 1000x1000
    fabric = []
    for i in range(1000):
        fabric.append([0 for j in range(1000)])
    o = list(claims.keys())
    for c_id in claims.keys():
        pos = get_position(claims[c_id])
        x_range = pos[0]
        y_range = pos[1]
        for x in range(x_range.s, x_range.e):
            for y in range(y_range.s, y_range.e):
                # If the coordinate is a string, then is a overlap, continue
                # remove the c_id from o if that exists
                if isinstance(fabric[x][y], str):
                    if c_id in o:
                        o.remove(c_id)
                    continue
                # when there is a non 0 value that's a coordinate already
                # claimed, set an overlap: "#"
                if fabric[x][y]:
                    # if the coordinate exists in o, remove it
                    if fabric[x][y] in o:
                        o.remove(fabric[x][y])
                    # if the current claim id exists in o, remove it
                    if c_id in o:
                        o.remove(c_id)
                    fabric[x][y] = "#"
                else:
                    fabric[x][y] = c_id
    return o


print(one(claims))
print(two(claims))
