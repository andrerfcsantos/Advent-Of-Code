# -*- coding: utf-8 -*-

DAY = 12

village = {}

def add_node(node_name):
    global village
    if node_name not in village:
        village[node_name] = set()

def add_edge(orig, dest):
    global village
    add_node(orig)
    add_node(dest)
    village[dest].add(orig)
    village[orig].add(dest)

def all_reachable(node):
    global village
    visited, to_visit = set(),set()

    visited.add(node)
    to_visit.update(village[node])

    while len(to_visit) > 0:
        next_node = to_visit.pop()
        to_visit.update(village[next_node]-visited)
        visited.add(next_node)

    return visited

def components():
    global village
    components = []
    nodes_to_check = set(village.keys())

    while len(nodes_to_check) > 0:
        node_to_check = nodes_to_check.pop()
        node_component = all_reachable(node_to_check)
        components.append(node_component)
        nodes_to_check -= node_component
    
    return components


inp = open(f'../in/day{DAY:02}.txt')

for line in inp:
    line = line.strip().split(' <-> ')
    left_hand_side,right_hand_side= line[0],line[1].split(', ')

    for program in right_hand_side:
        add_edge(left_hand_side,program)

print(f'Part 1: {len(all_reachable("0"))}')
print(f'Part 2: {len(components())}')
