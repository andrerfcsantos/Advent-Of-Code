# -*- coding: utf-8 -*-

import itertools as it
import random as rd
import fileinput  as fi
import numpy as np
import re
import collections as col

DAY = 1

inp = open(f'../in/day{DAY:02}.txt')

for line in inp:
    line = line.strip()

inp.close()

part1, part2 = None, None

print(f'Part 1: {part1}')
print(f'Part 2: {part2}')
