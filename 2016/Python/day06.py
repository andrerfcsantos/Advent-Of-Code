# -*- coding: utf-8 -*-
import collections as col
import operator as op

code_p1,code_p2='',''
alphabet = [chr(x) for x in range(ord('a'), ord('z') + 1)]
corrupted_codes = []

for line in open(r'../inputs/day06.txt'):
    corrupted_codes.append(line.strip())

for i in range(0,len(corrupted_codes[0])):
    occurrences = col.Counter([x[i] for x in corrupted_codes])    
    order = sorted(occurrences.items(), key=op.itemgetter(1), reverse=True)
    code_p1 += order[0][0]
    code_p2 += order[-1][0]

print('Part 1:',code_p1)
print('Part 2:',code_p2)



