# -*- coding: utf-8 -*-

DAY = 17

ntimes = int(open(f'../in/day{DAY:02}.txt').readline().strip())

vals = [0]
pos = 0
currval = 1

while currval != 2017+1:

    pos = 1 + (pos + ntimes) % currval
    vals.insert(pos,currval)

    currval+=1

part1 = vals[pos+1]

pos = 0
pos_zero = 0
after_zero = None

for i in range(1, 50000001):
    pos = (pos + ntimes) % i + 1

    if pos < pos_zero:
        pos_zero += 1
    if pos == pos_zero + 1:
        after_zero = i

part2 = after_zero

print(f'Part 1: {part1}')
print(f'Part 2: {part2}')
