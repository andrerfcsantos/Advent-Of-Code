# -*- coding: utf-8 -*-
import re
import itertools as it
from collections import deque

outputs = {}
bot_cargo = {}
bot_cmds = {}

for line in open(r'../inputs/day10.txt'):
    line =line.strip()
    r1 = 'bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)'
    r2 = 'value (\d+) goes to bot (\d+)'
    bot_action = re.match(r1,line)
    bot_input = re.match(r2,line)
    
    if bot_action:
        bot, r_low, r_high = map(int,re.group(1,3,5))
        rtype_low, rtype_high = re.group(2,4)
        
        if bot not in bot_cmds.keys():
            bot_cmds[bot] = deque([(rtype_low,r_low),(rtype_high,r_high)])
        else:
            bot_cmds[bot].extend([(rtype_low,r_low),(rtype_high,r_high)])
            
    elif bot_input:
        bot, chip = map(int,re.group(2,1))
        
        if bot not in bot_cargo.keys():
            bot_cargo[bot] = deque([chip])
        else:
            bot_cargo[bot].append(chip)
        
        if len(bot_cargo[bot]) >= 2:
            chip1=bot_cargo[bot].popleft()
            chip2=bot_cargo[bot].popleft()
            chips=sorted([chip1,chip2])
            (rtype_low,r_low) = bot_cmds[bot].popleft()
            (rtype_high,r_high) = bot_cmds[bot].popleft()
            
            if rtype_low == 'output':
                outputs[r_low].append(chips[0])
            else:
                
            
            if rtype_high == 'output'
                outputs[r_high].append(chips[1])
            


print('Part 1: ')
print('Part 2: ')



