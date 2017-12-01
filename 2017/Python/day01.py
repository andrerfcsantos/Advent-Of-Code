# -*- coding: utf-8 -*-

DAY = 1
numbers = open(f'../in/day{DAY:02}.txt').readline().strip()

sum, sum2 = 0,0
step = int(len(numbers)/2)

for i,number in enumerate(numbers):
    if number == numbers[(i+1)%len(numbers)]:
        sum += int(number)
    if number == numbers[(i+step)%len(numbers)]:
        sum2 += int(number)


part1, part2 = sum, sum2

print(f'Part 1: {part1}')
print(f'Part 2: {part2}')

