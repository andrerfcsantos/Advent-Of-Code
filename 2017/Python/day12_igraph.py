# -*- coding: utf-8 -*-

from igraph import Graph

DAY = 12

# Programs and pipes as sets since igraph
# doesn't check for duplicate vertices and
# edges and will happily repeat vertices with
# the same name and edges with the same endpoints
# Since that's not a desired behavior, set's 
# were used to handle duplicates before inserting
# them on the graph.
programs, pipes = set(),set()

village = Graph()

# Function renaming for role-playing puposes :)
village.add_pipes = village.add_edges
village.add_programs = village.add_vertices

inp = open(f'../in/day{DAY:02}.txt')

for line in inp:
    line = line.strip().split(' <-> ')
    left_hand_side,right_hand_side= line[0],line[1].split(', ')

    programs.add(left_hand_side)

    for program in right_hand_side:
        pipes.add( (left_hand_side,program) )

village.add_programs(list(programs))
village.add_pipes(list(pipes))

print(f'Part 1: {len(village.subcomponent("0"))}')
print(f'Part 2: {len(village.clusters())}')


