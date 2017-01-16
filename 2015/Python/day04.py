# -*- coding: utf-8 -*-
import hashlib

key = open(r'..\in\day04.txt').readline().rstrip()

i=0

while True:
    tohash = key + str(i)
    m = hashlib.md5()
    m.update(tohash.encode())
    if m.hexdigest().startswith('0'*5): break
    i+=1

j=0
while True:
    tohash = key + str(j)
    m = hashlib.md5()
    m.update(tohash.encode())
    if m.hexdigest().startswith('0'*6): break
    j+=1

print('Part 1:',i)
print('Part 2:',j)
