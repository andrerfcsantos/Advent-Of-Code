# -*- coding: utf-8 -*-
import re

valid_triangles_p1 = 0
valid_triangles_p2 = 0
triangles = []

for line in open(r'../inputs/day03.txt'):
    sline = re.split(r'[ ]+',line.strip())
    triangles.append(list(map(int,sline)))

for triangle in triangles:
    striangle = sorted(triangle)
    
    if striangle[0] + striangle[1] > striangle[2]:
        valid_triangles_p1+=1

for i in range(0,3):
    for j in range(0,len(triangles),3):
        triangle = [triangles[j][i],triangles[j+1][i],triangles[j+2][i]]
        striangle = sorted(triangle)
        
        if striangle[0] + striangle[1] > striangle[2]:
            valid_triangles_p2+=1

print('Part 1:',valid_triangles_p1)
print('Part 2:',valid_triangles_p2)


