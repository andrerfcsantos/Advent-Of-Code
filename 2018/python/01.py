from pathlib import Path
import time
import datetime
import itertools

INPUT_PATH = Path('../inputs').resolve()

input_file_path = INPUT_PATH.joinpath('2018_01.txt') 
input_frequencies = []

before_parse = time.perf_counter()
with open(input_file_path, encoding='utf-8') as input_file:
    for line in input_file:
        input_frequencies.append(int(line.strip()))
parse_duration = time.perf_counter() - before_parse

print("Parse in {:.4f} ms".format(parse_duration*1000))

before_p1 = time.time()
p1 = sum(input_frequencies)
after_p1 = time.time() - before_p1

print("Part 1: {} (in {:.4f} ms)".format(p1,after_p1*1000))

before_p2 = time.time()
frequenciesSeen = set()
frequency =0
frequenciesSeen.add(frequency)


for change in itertools.cycle(input_frequencies):

    frequency += change
    if frequency in frequenciesSeen:
        p2 = frequency
        break
    
    frequenciesSeen.add(frequency)

after_p2 = time.time() - before_p2

print("Part 2: {} (in {:.4f} ms)".format(p2,after_p2*1000))