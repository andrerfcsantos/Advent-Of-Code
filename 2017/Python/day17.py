# -*- coding: utf-8 -*-

DAY = 17

ntimes = int(open(f'../in/day{DAY:02}.txt').readline().strip())

vals = [0]
pos = 0
currval = 1

while currval != 2017+1:
    size = len(vals)

    pos = 1 + (pos + ntimes) % size
    vals.insert(pos,currval)

    currval+=1

part1 = vals[pos+1]

while currval != 50000000+1:
    size = len(vals)

    pos = 1 + (pos + ntimes) % size
    vals.insert(pos,currval)

    currval+=1

part2 = vals[vals.index(0)+1]

print(f'Part 1: {part1}')
print(f'Part 2: {part2}')
