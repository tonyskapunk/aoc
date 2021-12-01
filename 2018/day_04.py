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
# sort logs
logs = lines
logs.sort()


def get_data(logs):
    """Creates a dict by parsing the logs

    sleepy:
        'g_id1': { # Guard ID
            'date1': [asleep, wake-up, asleep, wake-up, ..]
              # Date is MM-DD
              # asleep is int on non-pair positions
              # wake-up is int on pair positions
            ..
        ..
    """
    sleepy = {}
    for log in logs:
        if '#' in log:
            d, g = log.split("#")
            # Get guard ID
            m = re.search('\d+\s', g)
            g_id = m.group(0).strip()
            # Get date
            m = re.search('\d\d-\d\d\s', d)
            date = m.group(0).strip()
            # Verify the guard ID has not been added already
            # otherwise create an empty entry for this date
            if not sleepy.get(g_id, None):
                sleepy[g_id] = {}
                sleepy[g_id][date] = []
            else:
                # Guard ID already exists, simply add the date
                # if it has not been already added
                if not sleepy[g_id].get(date, None):
                    sleepy[g_id][date] = []
        if 'falls' in log:
            d, _ = log.split("]")
            # Get minute falling asleep
            m = re.search("\s00:\d{2}$", d)
            _, sleep = m.group(0).split(":")
            # Get date, is possible that shift started prior falling sleep
            # (a day before)
            m = re.search('\d\d-\d\d\s', d)
            new_date = m.group(0).strip()
            # If falling sleep is not when shift started, then use the new date
            if date != new_date:
                sleepy[g_id][new_date] = sleepy[g_id][date]
                sleepy[g_id].pop(date, None)
                date = new_date
            sleepy[g_id][date].append(int(sleep))
        if 'wakes' in log:
            d, _ = log.split("]")
            # Get minute awaking
            m = re.search("\s00:\d{2}$", d)
            _, awake = m.group(0).split(":")
            sleepy[g_id][date].append(int(awake))
    return sleepy


def one(sleepy):
    """Get the guard facts
    - Guard has the most minutes asleep
    - Minute the guard spend asleep the most
    """
    sleep_time = {}
    # Count the sleep time per guard ID
    for g_id in sleepy.keys():
        tot_sleep = 0
        # Count the sleep time per date
        for date in sleepy[g_id].keys():
            # Count the sleep time per diff of time between asleep and awake
            # (sleep time is define by pair index, while awake time is non-pair
            # index)
            entries = int(len(sleepy[g_id][date])/2)
            for i in range(entries):
                index = i*2
                tot_sleep += (sleepy[g_id][date][index+1] -
                              sleepy[g_id][date][index])

        sleep_time[g_id] = tot_sleep
    # Order the time slept by guard into a list of tuples
    sort_by_asleep = sorted(sleep_time.items(), key=lambda kv: kv[1])
    # Obtain the guard ID and the time slept of the guard that slept the most
    sleepy_guard, slept_time = sort_by_asleep[-1]
    print(f'Guard that slept the most is {sleepy_guard}, slept {slept_time}')

    # Print the minutes the guard slept by day
    d = []
    dates = sleepy[sleepy_guard]
    # Loop through each date
    for date in dates.keys():
        # (sleep time is define by pair index, while awake time is non-pair
        # index)
        entries = int(len(dates[date])/2)
        # Loop through each range slept
        print(f'{date}',  end='')
        t = ''
        l_i = 0
        for i in range(entries):
            index = i*2
            s = dates[date][index]
            a = dates[date][index+1]
            for j in range(l_i, s):
                t += '.'
            for j in range(s, a):
                t += '#'
            l_i = a
        for j in range(l_i, 60):
            t += '.'
        print(f' {t}')
        d.append(t)

    # Obtain the minute the guard slept the most
    zzz = 0
    for i in range(60):
        z = 0
        for l in d:
            if l[i] == '#':
                z += 1
        if z > zzz:
            zzz = z
            min_zzz = i
    print(f'On minute {min_zzz} guard slept the most with {zzz} times')
    print()
    print(f'Answer: {int(sleepy_guard)*int(min_zzz)}')


def two(sleepy):
    """
    """
    most_min_per_guard = {}
    # Loop through each guard and each day

    for g_id in sleepy.keys():

        d = []
        # Loop through each date
        for date in sleepy[g_id].keys():
            # (sleep time is define by pair index, while awake time is non-pair
            # index)
            entries = int(len(sleepy[g_id][date])/2)
            # Loop through each range slept
            t = ''
            l_i = 0
            for i in range(entries):
                index = i*2
                s = sleepy[g_id][date][index]
                a = sleepy[g_id][date][index+1]
                for j in range(l_i, s):
                    t += '.'
                for j in range(s, a):
                    t += '#'
                l_i = a
            for j in range(l_i, 60):
                t += '.'
            d.append(t)

        # Obtain the minute the guard slept the most
        zzz = 0
        for i in range(60):
            z = 0
            for l in d:
                if l[i] == '#':
                    z += 1
            if z > zzz:
                zzz = z
                min_zzz = i
        most_min_per_guard[f'{g_id}:{min_zzz}'] = zzz

    # Order the minute most slept by guard into a list of tuples
    sort_by_most = sorted(most_min_per_guard.items(), key=lambda kv: kv[1])
    gid_and_min, tot_min_slept = sort_by_most[-1]
    gid_most_slept, min_most_slept = gid_and_min.split(":")
    print(f'Guard {gid_most_slept} slept the most {tot_min_slept} minutes '
          f'on minute {min_most_slept}')

    print(f'Answer: {int(gid_most_slept)*int(min_most_slept)}')


sleepy = get_data(logs)

one(sleepy)
two(sleepy)
