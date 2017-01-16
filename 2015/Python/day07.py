# -*- coding: utf-8 -*-
def is_int(s):
    try:
        x=int(s)
        return (True,x)
    except ValueError:
        return (False,None)


def calc_signal(signal):
    int_t = is_int(signal)
    if int_t[0]==True:
        return int_t[1]
    
    if signal in values:
        return values[signal]

    op = ops[signal]

    if op[0] == 'VALUE':
        res = calc_signal(op[1])        
    elif op[0] == 'NOT':
        res = 65535 - calc_signal(op[1])
    elif op[0] == 'AND':
        res = calc_signal(op[1]) & calc_signal(op[2])
    elif op[0] == 'OR':
        res = calc_signal(op[1]) | calc_signal(op[2])
    elif op[0] == 'LSHIFT':
        res = calc_signal(op[1]) << calc_signal(op[2])
    elif op[0] == 'RSHIFT':
        res = calc_signal(op[1]) >> calc_signal(op[2])
        
    values[signal] = res
    return res


values = {}
ops = {}


for line in open(r'..\in\day07.txt'):
    inst = line.rstrip().split(' ')
    size = len(inst)
    
    if size==3:
        ops[inst[-1]] = ('VALUE',inst[0])
    elif size==4:
        ops[inst[-1]] = ('NOT',inst[1])
    else:
        ops[inst[-1]] = (inst[1],inst[0],inst[2])

p1 = calc_signal('a')
values = {}
values['b'] = 46065
p2= calc_signal('a')

print('Part 1:',p1)
print('Part 2:',p2)
    
    
    