# -*- coding: utf-8 -*-

import itertools as it

DAY = 4

lines = list(map(lambda x: x.rstrip().split(),open(f'../in/day{DAY:02}.txt').readlines()))

part_1, part_2 = 0,0

for line in lines:
    valid_p1, valid_p2 = True, True

    for w1,w2 in it.combinations(line,r=2):
        if w1 == w2:
            valid_p1 = False
        if sorted(w1) == sorted(w2):
            valid_p2 = False
        
    if valid_p1:
        part_1+=1
    if valid_p2:
        part_2+=1

print(f'Part 1: {part_1}')
print(f'Part 2: {part_2}')
