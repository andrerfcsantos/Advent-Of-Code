# -*- coding: utf-8 -*-
import operator as op
import re
import itertools as it

alphabet = [chr(x) for x in range(ord('a'), ord('z') + 1)]

def comp(c): return c[1], chr(ord('z') - ord(c[0]))

sum_p1 = 0
sect_p2 = -1
cyphers, messages = [],[]

for line in open(r'../inputs/day04.txt'):
    sline = re.split(r'[-\[]', line.strip().strip(']'))
    sline[:-2] = [''.join(sline[:-2])]
    sline[-2] = int(sline[-2])
    cyphers.append(sline)
    
for (letters,sect_id,top5) in cyphers:
    occurrences = dict(zip(alphabet,it.repeat(0)))
    message = ''
    for letter in letters:
        message += alphabet[(ord(letter) - ord('a') + sect_id) % len(alphabet)]
        occurrences[letter] += 1
    messages.append(message)
    
    sorted_x = sorted(occurrences.items(), key=comp, reverse=True)
    top_letters = ''.join(map(op.itemgetter(0), sorted_x[:5]))
    
    if top_letters == top5:
        sum_p1 += sect_id
        
    if message == 'northpoleobjectstorage':
        sect_p2 = sect_id

print('Part 1:', sum_p1)
print('Part 2:', sect_p2)


