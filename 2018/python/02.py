from pathlib import Path
import time
import datetime
import itertools


def part2():
    nwords = len(words)
    for i in range(0,nwords):
        for j in range(i+1,nwords):
            difs = 0
            difAt = None
            for k in range(0,len(words[i])):
                if words[i][k] != words[j][k]:
                    difs +=1
                    difAt = k
            
            if difs == 1:
                return words[i][:difAt] + words[i][difAt+1:]


INPUT_PATH = Path('../inputs').resolve()

input_file_path = INPUT_PATH.joinpath('2018_02.txt') 
words = []
before_parse = time.perf_counter()
with open(input_file_path, encoding='utf-8') as input_file:
    for line in input_file:
        if line != "" :
            words.append(line.strip())
parse_duration = time.perf_counter() - before_parse

print("Parse in {:.4f} ms".format(parse_duration*1000))

before_p1 = time.time()

count2, count3 = 0,0

for word in words:

    has2OfSame , has3OfSame = False, False
    letterSet = dict(zip(word, [0]*len(word)))

    for letter in word:
        letterSet[letter]+=1

    for letter,count in letterSet.items():
        if count==2:
            has2OfSame = True
        elif count==3:
            has3OfSame = True
        
    if has2OfSame:
        count2 += 1
    if has3OfSame:
        count3 += 1

after_p1 = time.time() - before_p1

print("Part 1: {} (in {:.4f} ms)".format(count2*count3,after_p1*1000))

before_p2 = time.time()


p2 = part2()

after_p2 = time.time() - before_p2

print("Part 2: {} (in {:.4f} ms)".format(p2,after_p2*1000))