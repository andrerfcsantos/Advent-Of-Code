# -*- coding: utf-8 -*-
import operator

SECONDS = 2503
RESTING,FLYING=0,1
reindeers = set()
stats = {}
distances_p1 ={}
distances_p2, status,cooldowns,points = {},{},{},{}

for line in open(r'day14.txt'):
    split = line.strip('\n.').split(' ')
    reindeers.update([split[0]])
    stats[split[0]] = (int(split[3]),int(split[6]),int(split[-2]))

for r in reindeers:
    speed,fly_time,rest_time = stats[r]
    cycle_time = fly_time + rest_time
    
    cycles,remainder = divmod(SECONDS,cycle_time)
    
    distance = speed*(cycles*fly_time + min([remainder,fly_time]))
    distances_p1[r] = distance

p1 = sorted(distances_p1.values())[-1]

for r in reindeers:
    distances_p2[r] = 0
    status[r] = FLYING
    cooldowns[r] = 0
    points[r] = 0

for i in range(0,SECONDS):
    for r in reindeers:
        speed,fly_time,rest_time = stats[r]
        if status[r] == FLYING:
            distances_p2[r] += speed
            cooldowns[r] = (cooldowns[r] + 1)%fly_time
        else:
            cooldowns[r] = (cooldowns[r] + 1)%rest_time
        
        if cooldowns[r]==0:
            status[r] = (status[r] +1)%2
    leader = sorted(distances_p2.items(), key=operator.itemgetter(1))[-1][0]
    points[leader] +=1

p2 = sorted(points.values())[-1]
print('Part 1:',p1)
print('Part 2:',p2)



