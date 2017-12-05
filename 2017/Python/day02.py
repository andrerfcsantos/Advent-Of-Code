# -*- coding: utf-8 -*-
import itertools as it

DAY = 2
inp = open(f'../in/day{DAY:02}.txt')

sum_p1, sum_p2  = 0,0

for line in inp:
    numbers = list(map(int,line.rstrip().split()))
    sum_p1 += max(numbers)-min(numbers)
    sum_p2 += sum( (n1//n2 for (n1,n2) in it.permutations(numbers,r=2) if n1%n2==0) )

inp.close()

print(f'Part 1: {sum_p1}')
print(f'Part 2: {sum_p2}')

