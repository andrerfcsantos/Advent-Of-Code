# -*- coding: utf-8 -*-

import itertools as it
import random as rd
import fileinput  as fi
import numpy as np
import re
import collections as col
import math

DAY = 8

registers = {}

def getRegVal(register):
    if register not in registers:
        registers[register] = 0
    return registers[register]

def modifyIfExists(register, value, op):

    if register not in registers:
        registers[register] = 0
    
    if op == 'inc':
        registers[register] += value
    elif op == 'dec':
        registers[register] -= value


rs_dest    = '(?P<dest>\w+)'
rs_op      = '(?P<op>(inc)|(dec))'
rs_orig    = '(?P<orig>\w+)'
rs_comp    = '(?P<comp>(==|!=|<=|>=|>|<))'
rs_change  = '(?P<change>-?\d+)'
rs_rval    = '(?P<val>-?\d+)'

rline = re.compile(f'{rs_dest} {rs_op} {rs_change} if {rs_orig} {rs_comp} {rs_rval}')

inp = open(f'../in/day{DAY:02}.txt')

maxval = 0

for line in inp:
    line = line.strip()
    m = re.match(rline, line)
    dest   = m.group('dest')
    op     = m.group('op')
    orig   = m.group('orig')
    comp   = m.group('comp')
    change = int(m.group('change'))
    val    = int(m.group('val'))

    if eval(f'getRegVal(orig) {comp} {val}'):
        modifyIfExists(dest,change,op)
    
    if getRegVal(orig) > maxval:
        maxval = getRegVal(orig)

inp.close()

print(f'Part 1: {max(registers.values())}')
print(f'Part 2: {maxval}')
