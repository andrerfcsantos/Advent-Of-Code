# -*- coding: utf-8 -*-
import hashlib

pass_p1, pass_p2 = '','_'*8
size_p2=0

key = open(r'..\inputs\day05.txt').readline().strip()

i=0
while len(pass_p1)<8 or size_p2<8:
    tohash = key + str(i)
    m = hashlib.md5()
    m.update(tohash.encode())
    dg = m.hexdigest()
    
    if dg.startswith('0'*5):
        if len(pass_p1)<8:pass_p1 += dg[5]
        
        if dg[5].isdigit() and int(dg[5])<8 and pass_p2[int(dg[5])]=='_' and size_p2 <8:
            pos = int(dg[5])
            pass_p2 = pass_p2[:pos] + dg[6] + pass_p2[pos+1:]
            size_p2+=1
            
    i+=1

print('Part 1:',pass_p1)
print('Part 2:',pass_p2)