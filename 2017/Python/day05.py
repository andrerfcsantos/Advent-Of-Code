# -*- coding: utf-8 -*-

DAY = 5

instr = list(map(lambda x: int(x.strip()),open(f'../in/day{DAY:02}.txt').readlines()))

instr_p1, instr_p2 = instr[:],instr[:]
steps_p1, steps_p2 = 0,0
size = len(instr)

i=0

while i < size:
    step = instr_p1[i]
    instr_p1[i]+=1
    i+=step
    steps_p1+=1

i=0

while i < size:
    step = instr_p2[i]

    if step >= 3:
        instr_p2[i]-=1
    else:
        instr_p2[i]+=1
        
    i+=step
    steps_p2+=1


print(f'Part 1: {steps_p1}')
print(f'Part 2: {steps_p2}')
