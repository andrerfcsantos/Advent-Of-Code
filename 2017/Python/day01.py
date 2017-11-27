# -*- coding: utf-8 -*-

import itertools as it
import random as rd
import fileinput  as fi
import numpy as np
import re
import collections as col

DAY = 1

part1, part2 = None, None

inp = open(f'../in/{DAY:02}.txt')

for line in inp:
    line = line.strip()


print(f'Part 1: {part1}')
print(f'Part 2: {part2}')

inp.close()
