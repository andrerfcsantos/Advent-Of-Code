# -*- coding: utf-8 -*-

vowels = 'aeiou'
forbidden = ['ab', 'cd', 'pq', 'xy']

def is_nice_str_p1(s):
    n_vowels = 0
    double_l = False
    
    for char in s:
        if char in vowels: n_vowels+=1
        
    if n_vowels < 3: return False
    
    for i in range(1,len(s)):
        if s[i] == s[i-1]: 
            double_l = True
            break
    
    if not double_l: return False
    
    for sf in forbidden:
        if sf in s: return False
    
    return True

def is_nice_str_p2(s):
    found_pair = False
    found_rep = False
    
    for i in range(1,len(s)):
        if s[i-1]+s[i] in s[(i+1):]:
            found_pair = True
            break
    
    if not found_pair: return False
    
    for i in range(2,len(s)):
        if s[i] == s[i-2]:
            found_rep=True
            break
    
    if not found_rep: return False
    return True


file = open(r'..\in\day05.txt')
count_p1, count_p2 = 0,0

for line in file:
    line = line.rstrip()
    if is_nice_str_p1(line): count_p1+=1
    if is_nice_str_p2(line): count_p2+=1

print('Part 1:',count_p1)
print('Part 2:',count_p2)

        
    
    
    