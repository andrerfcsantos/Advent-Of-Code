# -*- coding: utf-8 -*-
import re 
import numpy as np

screen = np.array(['.']*(50*6)).reshape(6,50)

for line in open(r'../inputs/day08.txt'):
    sline = line.strip()
    re_rect = re.search(r'rect (\d+)x(\d+)', sline)
    re_rotate = re.search(r'rotate (column|row) (x|y)=(\d+) by (\d+)', sline)
    
    if re_rect:
        width, height = map(int,re_rect.group(1,2))
        screen[:height,:width] = '#'
    elif re_rotate:
        what, [index, shift] = re_rotate.group(1),map(int,re_rotate.group(3,4))
        if what == 'column':
            screen[:,index] = np.roll(screen[:,index],shift)
        elif what == 'row':
            screen[index] = np.roll(screen[index],shift)
    
print('Part 1:',screen[(screen == '#')].size)
print('Part 2:')

for screen_line in screen.tolist():
    print(''.join(screen_line))

