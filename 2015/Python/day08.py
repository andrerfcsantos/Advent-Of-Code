# -*- coding: utf-8 -*-

part1, part2 = 0,0

for line in open(r'..\in\day08.txt'):
    string = line.rstrip()
    part1 += len(string) - (len(eval(string)))
    part2 += string.count('\\') + string.count('\"') + 2

print('Part 1:',part1)
print('Part 2:',part2)
