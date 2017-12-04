# -*- coding: utf-8 -*-

import itertools as it

DAY = 4

inp = open(f'../in/day{DAY:02}.txt')

part_1, part_2 = 0,0

for line in inp:
    line = line.strip().split()
    valid_p1, valid_p2 = True, True

    for w1,w2 in it.permutations(line,r=2):
        if w1 == w2:
            valid_p1 = False
        if sorted(w1) == sorted(w2[::-1]):
            valid_p2 = False
        
    if valid_p1:
        part_1+=1
    if valid_p2:
        part_2+=1

inp.close()

print(f'Part 1: {part_1}')
print(f'Part 2: {part_2}')
