# -*- coding: utf-8 -*-
import re

def decompress_v2(comp_str):
    decomp_size = 0
    start = 0
    
    while start < len(comp_str):
        marker = re.match(r'\((\d+)x(\d+)\)',comp_str[start:])
        data = re.match(r'([A-Za-z0-9\(\)]+?)\(?',comp_str[start:])
        
        if marker:
            match_size = marker.end(0)
            n_chars,reps = map(int,marker.group(1,2))
            match_end = start+match_size
            decomp_size += len(comp_str[match_end:match_end+n_chars])*reps
            start += match_size+n_chars
        elif data:
            match_size = len(data.group(1))
            decomp_size += match_size
            start += match_size
    
    return decomp_size

def decompress_v1(comp_str):
    marker = re.match(r'\((\d+)x(\d+)\)',comp_str)
    data = re.match(r'([A-Za-z0-9]+)',comp_str)
        
    if marker:
        msize = marker.end(0)
        n_chars,reps = map(int,marker.group(1,2))
        selection = msize+n_chars
        selection_size = reps*decompress_v1(comp_str[msize:selection])
        return selection_size + decompress_v1(comp_str[selection:])
    elif data:
        msize = data.end(0)
        return msize + decompress_v1(comp_str[msize:])
    else: 
        return 0

res_p1,res_p2=0,0

for line in open(r'../inputs/day09.txt'):
    line = line.strip()
    res_p1+=decompress_v2(line)
    res_p2+=decompress_v1(line)

print('Part 1: ',res_p1)
print('Part 2: ',res_p2)


