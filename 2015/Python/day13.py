# -*- coding: utf-8 -*-

import itertools as it
import math

max_happiness = -math.inf
guests = set()
relationships = {}

for line in open(r'..\in\day13.txt'):
    string = line.strip('\n.').split(' ')
    p1,p2 = string[0], string[-1]
    
    if string[2] == 'gain':
        r = int(string[3])
    else:
        r = -int(string[3])
        
    guests.update((p1,p2))
    relationships[(p1,p2)]= r

n_guests=len(guests)
for perm in it.permutations(guests):
    happiness=0
    for i in range(0,len(perm)):
        happiness += relationships[(perm[i],perm[(i+1)%n_guests])]
        happiness += relationships[(perm[(i+1)%n_guests],perm[i])]
        if happiness > max_happiness:
            max_happiness = happiness

p1 = max_happiness
max_happiness = -math.inf

for guest in guests:
    relationships[(guest,'me')] = 0
    relationships[('me',guest)] = 0
guests.update(('me',))

n_guests=len(guests)
for perm in it.permutations(guests):
    happiness=0
    for i in range(0,len(perm)):
        happiness += relationships[(perm[i],perm[(i+1)%n_guests])]
        happiness += relationships[(perm[(i+1)%n_guests],perm[i])]
        if happiness > max_happiness:
            max_happiness = happiness

p2 = max_happiness

print('Part 1:',p1)
print('Part 2:',p2)
