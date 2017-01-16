# -*- coding: utf-8 -*-
from math import inf
import itertools as it


max_score = -inf
i_capacity,i_durability,i_flavor, i_texture, i_calories = range(0,5)

ingredients = {}
ing_list = []

for line in open(r'..\in\day15.txt'):
    split = line.strip('\n').split(' ')
    split = [x.strip(':,') for x in split]
    ing_name = split[0].strip(':')
    properties = (int(split[2]),int(split[4]),int(split[6]),int(split[8]),int(split[10]))
    ingredients[ing_name] = properties
    ing_list.append(ing_name)

nr_ingredients = len(ingredients)
                
for c in it.product(range(101), repeat=nr_ingredients):
    if sum(c)==100:
        capacity,durability,flavor,texture,calories = 0,0,0,0,0
        for i,ing in enumerate(ing_list):
            capacity += ingredients[ing][i_capacity]*c[i]
            durability += ingredients[ing][i_durability]*c[i]
            flavor += ingredients[ing][i_flavor]*c[i]
            texture += ingredients[ing][i_texture]*c[i]
        score = capacity*durability*flavor*texture
        if score > max_score and score > 0:
            print('Better Solution found: ',score,c)
            max_score=score

p1 = max_score
        
print('Part 1:',p1)
print('Part 2:')
        
