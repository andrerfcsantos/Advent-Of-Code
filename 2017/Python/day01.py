# -*- coding: utf-8 -*-

DAY = 1
numbers = open(f'../in/day{DAY:02}.txt').readline().strip()

sum, sum2 = 0,0
size = len(numbers)
step = size//2

for i,number in enumerate(numbers):
    if number == numbers[(i+1)%size]:
        sum += int(number)
    if number == numbers[(i+step)%size]:
        sum2 += int(number)

print(f'Part 1: {sum}')
print(f'Part 2: {sum2}')

