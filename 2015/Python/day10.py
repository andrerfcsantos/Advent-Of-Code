# -*- coding: utf-8 -*-

def look_and_say(s):
    res = ''
    size_s = len(s)
    c=1
    for i in range(0,size_s):
        if i+1 == size_s or s[i] != s[i+1]:
            res += str(c) + s[i]
            c=1
        else: 
            c+=1
    return res


its = []
seed = open(r'..\in\day10.txt').readline().rstrip()

for i in range(0,50):
    seed = look_and_say(seed)
    its.append(seed)

p1 = len(its[40-1])
p2 = len(its[50-1])

print('Part 1:',p1)
print('Part 2:',p2)

