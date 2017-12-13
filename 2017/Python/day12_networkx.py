# -*- coding: utf-8 -*-

import networkx as nx

DAY = 12

village = nx.Graph()

inp = open(f'../in/day{DAY:02}.txt')

for line in inp:
    line = line.strip().split(' <-> ')
    left_hand_side,right_hand_side= line[0],line[1].split(', ')

    village.add_node(left_hand_side)

    for program in right_hand_side:
        village.add_edge(left_hand_side,program )

print(f'Part 1: {len(nx.node_connected_component(village, "0"))}')
print(f'Part 2: {nx.number_connected_components(village)}')


