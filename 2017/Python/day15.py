# -*- coding: utf-8 -*-

import re
from queue import Queue

DAY = 15

svalues = {}

inp = open(f'../in/day{DAY:02}.txt')

re_parse = re.compile('Generator (?P<gen>\w+) starts with (?P<sval>\d+)')

for line in inp:
    m = re.search(re_parse,line.strip())
    svalues[m.group('gen')] = int(m.group('sval'))

inp.close()

MASK = 0xffff

def judge(a,b):
    return (a & MASK) == (b & MASK) 


val_a, val_b =  svalues['A'],svalues['B']
mfact_a, mfact_b =  16807, 48271

judge_count_p1 = 0

for i in range(int(40e6)):
    val_a = (val_a * mfact_a) % 2147483647
    val_b = (val_b * mfact_b) % 2147483647
    judge_count_p1 += 1 if judge(val_a,val_b) else 0


val_a, val_b =  svalues['A'],svalues['B']
judge_count_p2 = 0
compared = 0
a_queue,b_queue = Queue(), Queue()

while compared < int(5e6):
    val_a = (val_a * mfact_a) % 2147483647
    val_b = (val_b * mfact_b) % 2147483647

    if val_a % 4 == 0:
        a_queue.put_nowait(val_a)
    if val_b % 8 == 0:
        b_queue.put_nowait(val_b)
    
    if not a_queue.empty() and not b_queue.empty():
        compared += 1
        a,b = a_queue.get_nowait(), b_queue.get_nowait()
        if judge(a,b): judge_count_p2+=1


print(f'Part 1: {judge_count_p1}')
print(f'Part 2: {judge_count_p2}')
