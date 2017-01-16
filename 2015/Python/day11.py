# -*- coding: utf-8 -*-

def is_valid(s):
    inc_streak = False
    pairs, l_pairs= 0, []
    size = len(s)
    if 'i' in s or 'o' in s or 'l' in s:
        return False
    
    for i in range(0,size-2):
        if ord(s[i]) == ord(s[i+1])-1 and ord(s[i]) == ord(s[i+2])-2:
            inc_streak = True
            break
    
    if not inc_streak: return False
    
    for i in range(0,size-1):
        if s[i] == s[i+1] and s[i] not in l_pairs:
            pairs += 1
            l_pairs.append(s[i])
    
    if pairs < 2: return False
    
    return True
            

def inc_password(s):
    res = ''
    for i in range(len(s)-1,-1,-1):
        if s[i] == 'z':
            res = 'a' + res
        else:
            res = chr(ord(s[i])+1) + res
            break
    
    res = s[:i] + res
    return res
    
password = open(r'..\in\day11.txt').readline().rstrip()

found = 0
passwords = []


while found<2:
    password = inc_password(password)
    if is_valid(password):
        found +=1
        passwords.append(password)

print('Part 1:',passwords[0])
print('Part 2:',passwords[1])



        


