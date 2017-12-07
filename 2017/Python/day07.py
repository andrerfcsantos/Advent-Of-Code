# -*- coding: utf-8 -*-
import re
import random as rd
import collections as col
from mem import memoize_me_please

DAY = 7

lines = list(map(lambda x: x.rstrip(),open(f'../in/day{DAY:02}.txt').readlines()))
re_line = re.compile('(?P<tower>\w+) \((?P<weight>\d+)\)(?: -> (?P<related_towers>((\w+)+(?:, )?)*))?')

towers_below = {}
tree = {}

for line in lines:
    m = re.match(re_line,line)
    tower,weight,related = m.group('tower', 'weight', 'related_towers')

    if related:
        childs = related.split(', ')
        for related in childs:
            towers_below[related] = tower
        tree[tower] = (int(weight),childs[:])
    else:
        tree[tower] = (int(weight),[])
            

def find_bottom_tower(start):
    global towers_below
    
    if start in towers_below:
        return find_bottom_tower(towers_below[start])
    else:
        return start

@memoize_me_please
def sub_tree_weight(root):
    global tree
    own_weight = tree[root][0]
    childs_weight = sum( (sub_tree_weight(child) for child in tree[root][1])  )
    return own_weight+childs_weight

def find_unbalanced(start, expected_weight):
    global tree
    own_weight = tree[start][0]
    subtrees = [(child,sub_tree_weight(child)) for child in tree[start][1]]
    weights = list(map(lambda x:x[1],subtrees))
    right_weight = col.Counter(weights).most_common(1)[0][0]
    
    for subtree in subtrees:
        if subtree[1] != right_weight:
            return find_unbalanced(subtree[0],right_weight)
    
    return (start,expected_weight-sum(weights))


starting_tower = rd.choice(list(towers_below.keys()))
bottom_tower = find_bottom_tower(starting_tower)
unbalanced = find_unbalanced(bottom_tower,tree[bottom_tower][0])

print(f'Part 1: {bottom_tower}')
print(f'Part 2: {unbalanced}')
