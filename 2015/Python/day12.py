# -*- coding: utf-8 -*-

import json

def obj_sum_p1(obj):
    res=0
    
    if type(obj)==int:
        return obj
    elif type(obj)==list:
        for elem in obj:
            res+=obj_sum_p1(elem)
        return res
    elif type(obj)==dict:
        for key in obj:
            res+=obj_sum_p1(obj[key])
        return res
    else:
        return 0

def obj_sum_p2(obj):
    res=0
    if type(obj)==int:
        return obj
    elif type(obj)==list:
        for elem in obj:
            res+=obj_sum_p2(elem)
        return res
    elif type(obj)==dict:
        if 'red' in obj.values():
            print('red is value in',obj)
            return 0
        for key in obj:
            res+=obj_sum_p2(obj[key])
        return res
    else:
        return 0
    

json_obj = json.load(open(r'..\in\day12.txt'))

p1 = obj_sum_p1(json_obj)
p2 = obj_sum_p2(json_obj)

print('Part 1:',p1)
print('Part 2:',p2)
