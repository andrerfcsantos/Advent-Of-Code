# -*- coding: utf-8 -*-
from enum import Enum
import re

DAY = 16

class MType(Enum):
    S = 0
    X = 1
    P = 2

def parse_moves(moves):
    res = []
    s_move = '(?P<move_type>(s|x|p))'
    s_spin = '(?P<spin_size>\d+$)'
    s_exchange = '((?P<ex1>\d+)/(?P<ex2>\d+))'
    s_partner = '((?P<p1>\w+)/(?P<p2>\w+))'
    re_move = re.compile('{}({}|{}|{})'.format(s_move, s_spin,s_exchange, s_partner))

    for move in moves:
        m = re.match(re_move,move)
        move_type = m.group('move_type')

        if move_type == 's':
            res.append( (MType.S,int(m.group('spin_size'))) )
        elif move_type == 'x':
            res.append( (MType.X, int(m.group('ex1')), int(m.group('ex2'))) )
        elif move_type == 'p':
            p1, p2 = m.group('p1', 'p2')
            res.append( (MType.P, p1,p2) )
        else:
            res.append( None )
    
    return res

def do_moves(state, moves):
    res = state[:]
    for m in moves:
        move_type = m[0]

        if move_type == MType.S:
            spin_size = m[1]
            res[:-spin_size], res[-spin_size:] = res[-spin_size:], res[:-spin_size]
        elif move_type == MType.X:
            pos1, pos2 = m[1:]
            res[pos1], res[pos2] = res[pos2], res[pos1]
        elif move_type == MType.P:
            p1, p2 = m[1:]
            pos1, pos2 = res.index(p1), res.index(p2)
            res[pos1], res[pos2] = res[pos2], res[pos1]
        else:
            print('Move type unknown:', move_type)
    return res

def do_dance(dancers, moves, times=1):
    res = dancers[:]
    moves = moves[:]

    init_state = ''.join(dancers)
    last_seen = {init_state: 0}
    history = [init_state]

    i=1

    while i <= times:
        res = do_moves(res,moves)

        state = ''.join(res)
        if state in last_seen:
            cycle_size = i - last_seen[state]
            times, left = divmod(times-i,cycle_size)
            cycle = history[-cycle_size:]
            res = cycle[left-1]
            break
        else:
            last_seen[state] = i
        
        history.append(state)
        i+=1
    return res


moves = open('../in/day{}.txt'.format(DAY)).readline().strip().split(',')
moves = parse_moves(moves)

dancers = [chr(x) for x in range(ord('a'),ord('p')+1)]

part1 = do_dance(dancers=dancers,moves=moves, times=1)
part2 = do_dance(dancers=part1, moves=moves,times=1000000000)

print('Part 1:', "".join(part1))
print('Part 2:', "".join(part2))
