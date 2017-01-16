# -*- coding: utf-8 -*-

file = open(r'..\in\day02.txt')

wrap = 0
ribbon = 0

for line in file:
    l,w,h = map(int,line.rstrip().split('x'))
    s1,s2,s3 = sorted((l*w , w*h , h*l))
    wrap += 2*s1+2*s2+2*s3 + s1 + s2
    s_sides = sorted((l,w,h))
    ribbon += 2*s_sides[0] + 2*s_sides[1] + l*w*h

print('Part 1:',wrap)
print('Part 2:',ribbon)