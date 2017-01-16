# -*- coding: utf-8 -*-

import numpy as np

p1,p2=0,0
matrix_p1 = np.zeros((1000,1000))
matrix_p2 = np.zeros((1000,1000))

for line in open(r'..\in\day06.txt'):
    parse = line.rstrip().split(' ')
    
    (xs,ys) = tuple(map(int,parse[-1].split(',')))
    
    if parse[0] == 'toggle':
        (ins,(xi,yi))=('toggle',tuple(map(int,parse[1].split(','))))
    elif parse[0] == 'turn':
        (ins,(xi,yi))=(parse[1],tuple(map(int,parse[2].split(','))))
        
    
    if ins == 'toggle':
        matrix_p1[yi:(ys+1),xi:(xs+1)] = (matrix_p1[yi:(ys+1),xi:(xs+1)] +1)%2
        matrix_p2[yi:(ys+1),xi:(xs+1)] += 2
    elif ins == 'on':
        matrix_p1[yi:(ys+1),xi:(xs+1)] = 1
        matrix_p2[yi:(ys+1),xi:(xs+1)] += 1
    else:
        matrix_p1[yi:(ys+1),xi:(xs+1)] = 0
        matrix_p2[yi:(ys+1),xi:(xs+1)] -= 1
        matrix_p2[(matrix_p2<0)] = 0

               
p1 = np.sum(matrix_p1)
p2 = np.sum(matrix_p2)

print('Part 1:',p1)
print('Part 2:',p2)