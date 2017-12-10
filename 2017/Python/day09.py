# -*- coding: utf-8 -*-

DAY = 9

line = open(f'../in/day{DAY:02}.txt').readline().strip()

def score(stream):
    size = len(stream)

    current_score, total_score = 0,0
    in_group, in_garbage  = False, False
    chars_in_garbage = 0
    i = 0

    while i < size:

        if stream[i] == '{' and not in_garbage:
            in_group= True
            current_score += 1
        elif stream[i] == '}' and in_group and not in_garbage:
            total_score += current_score
            current_score -= 1
            if current_score ==0:
                in_group = False
        elif stream[i] == '!' and in_garbage:
            i+=1
        elif stream[i] == '<' and not in_garbage:
            in_garbage = True
        elif stream[i] == '>' and in_garbage:
            in_garbage=False
        elif in_garbage:
            chars_in_garbage +=1
        
        i+=1
    
    return total_score,chars_in_garbage

print(f'Part 1: {score(line)[0]}')
print(f'Part 2: {score(line)[1]}')
