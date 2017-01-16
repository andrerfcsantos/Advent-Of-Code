# -*- coding: utf-8 -*-

floor = 0
basement_pos = -1
line = open(r'..\in\day01.txt').readline().rstrip()

for (i,char) in enumerate(line):
    if char == '(': floor+=1
    else: floor-=1
    
    if floor == -1 and basement_pos == -1:
        basement_pos = i+1

print('Part 1:',floor)
print('Part 2:',basement_pos)
