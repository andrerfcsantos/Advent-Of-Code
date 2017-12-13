# -*- coding: utf-8 -*-

# A more efficient version of day13.py

DAY = 13


inp = open(f'../in/day{DAY:02}.txt')


firewall = {}

def mindelay():
    global firewall

    run_lenght = max(firewall.keys())
    delay = 0
    mindelay_found = False

    while not mindelay_found:
        mindelay_found = True
        for layer_depth in range(0,run_lenght+1):
            layer_range = firewall.get(layer_depth, None)

            if layer_range != None and ((layer_depth + delay) % ((layer_range-1)*2)) == 0:
                mindelay_found = False
                break
        
        delay+=1
        
    return delay-1



def severity(delay):
    global firewall

    run_lenght = max(firewall.keys())
    total_severity=0
    for layer_depth in range(0,run_lenght+1):
        layer_range = firewall.get(layer_depth, None)
        picoseconds = layer_depth + delay

        if layer_range != None and (picoseconds % ((layer_range-1)*2)) == 0:
            total_severity += layer_depth*layer_range
    
    return total_severity

for line in inp:
    layer, depth = [int(x) for x in line.strip().split(': ')]
    firewall[layer] = depth

part1 = severity(0)
part2 = mindelay()

print(f'Part 1: {part1}')
print(f'Part 2: {part2}')
