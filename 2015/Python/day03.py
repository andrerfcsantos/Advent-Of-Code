# -*- coding: utf-8 -*-
import itertools as it

line = open(r'..\in\day03.txt').readline().rstrip()

x,y = 0,0
visited = set()
visited.add((0,0))

for (i,char) in enumerate(line):
    if char == '<': x-=1
    elif char == '^': y+=1
    elif char == '>': x+=1
    else: y-=1
    visited.add((x,y))

part1=len(visited)

s_x,s_y = 0,0
r_x,r_y = 0,0
visited = set()
visited.add((0,0))

for char in it.islice(line,0,len(line),2):
    if char == '<': s_x-=1
    elif char == '^': s_y+=1
    elif char == '>': s_x+=1
    else: s_y-=1
    visited.add((s_x,s_y))

for char in it.islice(line,1,len(line),2):
    if char == '<': r_x-=1
    elif char == '^': r_y+=1
    elif char == '>': r_x+=1
    else: r_y-=1    
    visited.add((r_x,r_y))

part2 = len(visited)
print('Part 1:',part1)
print('Part 2:',part2)
