# -*- coding: utf-8 -*-
import functools as ftools
import operator as op

DAY = 10

def chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i + n]

def knot_hash(numbers, lengths, rounds=1):
    res = numbers[:]
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
    
    return res

def dense_hash(numbers):
    return [ftools.reduce(op.xor,chunk,0) for chunk in chunks(numbers, 16)]
        

str_list = open(f'../in/day{DAY:02}.txt').readline().strip()

lengths = [int(x) for x in str_list.split(',')]
numbers = list(range(0,256))
knot_hash_p1 = knot_hash(numbers, lengths,1)
part_1 = knot_hash_p1[0]*knot_hash_p1[1]

lengths = [ord(x) for x in str_list] + [17, 31, 73, 47, 23]
knot_hash_p2 = knot_hash(numbers, lengths,64)
dense_hash_p2 = dense_hash(knot_hash_p2)
hexa_string = ''.join([f'{n:x}' for n in dense_hash_p2])


print(f'Part 1: {part_1}')
print(f'Part 2: {hexa_string}')
