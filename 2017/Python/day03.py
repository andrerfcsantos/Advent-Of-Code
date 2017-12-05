# -*- coding: utf-8 -*-

import math as m
import numpy as np

def getAdjSum(mt, px, py):
    return mt[px-1, py] + mt[px-1, py-1] + mt[px-1, py+1]+mt[px+1, py] + mt[px+1, py-1] + mt[px+1, py+1]+mt[px, py+1] + mt[px, py-1]

DAY = 3

pos = int(open(f'../in/day{DAY:02}.txt').readline().rstrip())

r=1
result_p1 = None

while result_p1 == None:
    # Number at right bottom and difference between corners
    rb, cdiff = r**2, r-1
    # Numbers at left bottom, left top and right top
    lb, lt, rt = rb-cdiff*1,rb-cdiff*2,rb-cdiff*3
    # Corner to center distance 
    ctc = cdiff//2;

    if pos <= rb and pos >= lb: 
        result_p1 = ctc + abs(pos-(rb-ctc))
    elif pos < lb and pos >= lt:
        result_p1 = ctc + abs(pos-(lb-ctc))
    elif pos < lt and pos >= rt:
        result_p1 = ctc + abs(pos-(lt-ctc))
    elif pos < rt and pos > rb:
        result_p1 = ctc + abs(pos-(rt-ctc))

    r+=2


result_p2=None

dim = 21
matrix = np.zeros( (dim,dim) )

cx,cy = (dim-1)//2,(dim-1)//2

matrix[cx,cy] = 1
cx += 1
r=3

while result_p2==None:
    cdiff = r-1

    for aux in range(0,cdiff-1):
        matrix[cx,cy] = getAdjSum(matrix, cx, cy)
        result_p2 = matrix[cx,cy] if result_p2==None and matrix[cx,cy] > pos else result_p2
        cy-=1
    
    cy+=1
    
    for aux in range(0,cdiff):
        matrix[cx,cy] = getAdjSum(matrix, cx, cy)
        result_p2 = matrix[cx,cy] if result_p2==None and matrix[cx,cy] > pos else result_p2
        cx-=1
    
    cx+=1

    for aux in range(0,cdiff):
        matrix[cx,cy] = getAdjSum(matrix, cx, cy)
        result_p2 = matrix[cx,cy] if result_p2==None and matrix[cx,cy] > pos else result_p2
        cy+=1
    
    cy-=1

    for aux in range(0,cdiff):
        matrix[cx,cy] = getAdjSum(matrix, cx, cy)
        result_p2 = matrix[cx,cy] if result_p2==None and matrix[cx,cy] > pos else result_p2
        cx+=1
    
    r+=2

result_p2 = int(result_p2)

print(f'Part 1: {result_p1}')
print(f'Part 2: {result_p2}')
