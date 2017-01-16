# -*- coding: utf-8 -*-

import itertools as it
import math

shortest, longest = math.inf, -math.inf

cities = set()
distances = {}

for line in open(r'..\in\day09.txt'):
    string = line.rstrip().split(' ')
    c1,c2,d = string[0],string[2],int(string[4])
    cities.update([c1,c2])
    distances[(c1,c2)] = d
    distances[(c2,c1)] = d

for perm in it.permutations(cities):
    dist = 0
    
    for i in range(1,len(cities)):
        dist += distances[(perm[i-1],perm[i])]
    
    if dist < shortest: shortest = dist
    if dist > longest: longest = dist
        

print('Part 1:',shortest)
print('Part 2:',longest)

