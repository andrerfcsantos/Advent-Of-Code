# -*- coding: utf-8 -*-

DAY = 11

deltas = {
    'n' : (    0, 1   ),
    'nw': ( -0.5, 0.5 ),
    'ne': (  0.5, 0.5 ),
    's' : (    0, -1  ),
    'sw': ( -0.5, -0.5),
    'se': (  0.5, -0.5)
}

def shorthest_path_to_pos(x,y):
    x , y = abs(x),abs(y)
    dsteps = min(x/0.5,y/0.5)
    vsteps = max(x-(dsteps*0.5),y-(dsteps*0.5))
    return int(dsteps + vsteps)

def shorthest_furthest_path(directions):
    global deltas
    path_deltas = [deltas[x] for x in directions]
    
    pos = [0,0]
    furthest = 0

    for dx,dy in path_deltas:
        pos[0] += dx
        pos[1] += dy
        dist = shorthest_path_to_pos(pos[0],pos[1])
        if dist > furthest:
            furthest = dist

    return shorthest_path_to_pos(pos[0],pos[1]),furthest

directions = open(f'../in/day{DAY:02}.txt').readline().strip().split(',')
part_1, part_2 = shorthest_furthest_path(directions)

print(f'Part 1: {part_1}')
print(f'Part 2: {part_2}')
