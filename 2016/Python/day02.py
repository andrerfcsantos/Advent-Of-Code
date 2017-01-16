# -*- coding: utf-8 -*-

x1,y1 = 1,1
x2,y2=2,0
code_p1, code_p2= '',''

board_p1 = [[1,2,3],
            [4,5,6],
            [7,8,9]]

board_p2 = [[ '-',  '-', '1', '-', '-'],
            [ '-',  '2', '3', '4', '-'],
            [ '5',  '6', '7', '8', '9'],
            [ '-',  'A', 'B', 'C', '-'],
            [ '-',  '-', 'D', '-', '-']]

for line in open(r'../inputs/day02.txt'):
    for c in line:
        if c=='R':
            if x1<2: x1+=1
            if x2<4 and board_p2[y2][x2+1] != '-': x2+=1
        elif c=='D':
            if y1<2: y1+=1
            if y2<4 and board_p2[y2+1][x2] != '-': y2+=1
        elif  c=='U':
            if y1>0: y1-=1
            if y2>0 and board_p2[y2-1][x2] != '-': y2-=1 
        elif c=='L':
            if x1>0: x1-=1
            if x2>0 and board_p2[y2][x2-1] != '-': x2-=1
    
    code_p1 += str(board_p1[y1][x1])
    code_p2 += str(board_p2[y2][x2])

print('Part 1:',code_p1)
print('Part 2:',code_p2)


