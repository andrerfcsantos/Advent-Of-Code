# -*- coding: utf-8 -*-
import re

def getBABs(string):
    res=[]
    for i in range(3,len(string)+1,1):
        seq= string[i-3:i]
        if seq == seq[::-1] and seq[0] != seq[1]:
            res.append(seq[1]+seq[0]+seq[1])
    return res

def hasABBA(string):
    found = False
    for i in range(4,len(string)+1,1):
        seq= string[i-4:i]
        if seq == seq[::-1] and seq[0] != seq[1]:
            found = True
            break
    return found

count_p1 =0
count_p2 =0

for line in open(r'../inputs/day07.txt'):
    babs = []
    has_bab = False
    tls_out, tls_in = False, False
    sline = re.split(r'[\[\]]',line.strip())
    
    for seq in sline[::2]:
        babs.extend(getBABs(seq))
        if hasABBA(seq): tls_out = True
    
    for seq in sline[1::2]:
        for bab in babs:
            if bab in seq: has_bab = True
        
        if hasABBA(seq): tls_in = True
    
    if tls_out and not tls_in: count_p1+=1
    if has_bab: count_p2+=1    


print('Part 1:',count_p1)
print('Part 2:',count_p2)
