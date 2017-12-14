# -*- coding: utf-8 -*-
import functools as ftools
import operator as op

DAY = 14
disk = []

def flood_fill(i,j,fill,target='#'):
    global disk

    nrows = len(disk)
    ncols = len(disk[i])
    if disk[i][j] == target:
        disk[i] = disk[i][:j] + fill + disk[i][(j+1):]
        if i > 0:       flood_fill(i-1,j,fill,target)
        if j > 0:       flood_fill(i,j-1,fill,target)
        if i < nrows-1: flood_fill(i+1,j,fill,target)
        if j < ncols-1: flood_fill(i,j+1,fill,target)

def chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i + n]

def knot_hash(string, rounds=1):
    res = list(range(0,256))
    lengths = [ord(x) for x in string] + [17, 31, 73, 47, 23]
    size = len(res)

    pos, skip_size = 0,0

    for rcount in range(0,rounds):
        for length in lengths:
            i_end = pos + length - 1
            if i_end >= size:
                until_end = size-pos
                from_start = i_end % size
                l_aux = list(reversed(res[pos:] + res[:(from_start+1)]))
                res[pos:] = l_aux[:until_end]
                res[:(from_start+1)] = l_aux[until_end:]
            elif i_end!=0:
                res[pos:(i_end+1)] = list(reversed(res[pos:i_end+1]))
            
            pos = (pos + length + skip_size)%size
            skip_size+=1
    
    res = [ftools.reduce(op.xor,chunk,0) for chunk in chunks(res, 16)]
    res = ''.join([hex(n)[2:].zfill(2) for n in res])
    
    return res        

base_string = open(f'../in/day{DAY:02}.txt').readline().strip()

total_fragments = 0

for i_row in range(0,128):
    row_string = f'{base_string}-{i_row}'
    row_khash = knot_hash(row_string, rounds=64)
    binary_str = ''.join([bin(int(x,16))[2:].zfill(4) for x in row_khash])
    disk.append(binary_str.translate(str.maketrans({'0':'.','1':'#'})))
    total_fragments += binary_str.count('1')

region = 0

for i in range(0,128):
    for j in range(0,128):
        if disk[i][j] == '#':
            region+=1
            flood_fill(i,j,target='#',fill='x')

print(f'Part 1: {total_fragments}')
print(f'Part 2: {region}')
