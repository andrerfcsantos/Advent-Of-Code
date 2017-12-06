# -*- coding: utf-8 -*-

import itertools as it
import random as rd
import fileinput  as fi
import numpy as np
import re
import collections as col


def redistribute(state):
    new_state = state[:]
    size,m = len(new_state),max(new_state)
    i = new_state.index(m)
    new_state[i]=0
    i+=1


    for blocks_left in range(m,0,-1):
        new_state[i%size]+=1
        i+=1


    return new_state

DAY = 6

start_state = list(map(int,open(f'../in/day{DAY:02}.txt').readline().rstrip().split()))

states_known = []

cycles_p1, cycles_p2=0,0
state = start_state[:]
states_known.append(state)
inf_loop_found = False

while not inf_loop_found:
    state = redistribute(state)
    cycles_p1 += 1
    if state in states_known:
        first_dup = state[:] 
        inf_loop_found=True
    else:
        states_known.append(state)

states_known=[]
states_known.append(first_dup)

inf_loop_found = False

while not inf_loop_found:
    state = redistribute(state)

    cycles_p2 += 1
    if state in states_known:
        inf_loop_found=True
    else:
        states_known.append(state)


    
print(f'Part 1: {cycles_p1}')
print(f'Part 2: {cycles_p2}')
