# -*- coding: utf-8 -*-

DAY = 1
numbers = list(map(int,open(f'../in/day{DAY:02}.txt').readline().rstrip()))

sum_p1, sum_p2 = 0,0
size = len(numbers)
step = size//2

for i,number in enumerate(numbers):
    if number == numbers[(i+1)%size]:
        sum_p1 += number
    if number == numbers[(i+step)%size]:
        sum_p2 += number

print(f'Part 1: {sum_p1}')
print(f'Part 2: {sum_p2}')

