# -*- coding: utf-8 -*-

x, y = 0, 0
hqx, hqy = 0, 0
instructions = []
visited = [(0, 0)]
hqfound = False

direction = 'north'
new_directions = {
    'north': ['west' , 'east' , 0, 1],
    'south': ['east' , 'west' , 0,-1],
    'west':  ['south', 'north',-1, 0],
    'east':  ['north', 'south', 1, 0]
}

for line in open(r'../inputs/day01.txt'):
    sline = line.strip().split(', ')
    for inst in sline:
        instructions.append((inst[0], int(inst[1:])))

for turn, steps in instructions:
    direction = new_directions[direction][turn == 'L']

    for i in range(0,steps):
        x,y = x+new_directions[direction][2],y+new_directions[direction][3]
        
        if (x,y) in visited and not hqfound:
            hqfound = True
            hqx, hqy = x,y
            
        visited.append((x,y))

dist_p1 = abs(x) + abs(y)
dist_p2 = abs(hqx) + abs(hqy)

print('Part 1:', dist_p1)
print('Part 2:', dist_p2)

