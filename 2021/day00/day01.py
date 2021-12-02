#!/usr/bin/env python3


def count_increments(measurements):
    count = 0
    if not measurements or len(measurements) == 1:
        return count

    for i, m in enumerate(measurements):
        if i == 0:
            continue
        if measurements[i - 1] < m:
            count += 1
    #            print(f"[{measurements[i-1]} < {m}] | {count} (increased)")
    #        else:
    #            print(f"[{measurements[i-1]} >= {m}] | {count} (decreased)")

    return count


if __name__ == "__main__":
    with open("input_day01", "r") as f:
        measurements = f.read().splitlines()

    # fixed!
    print(count_increments(measurements) + 1)
