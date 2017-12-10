#!/usr/bin/python
# -*- coding: <encoding name> -*-
import re
import requests as req

r_day = re.compile('<a *href="/2017/day/\d+"> *(?P<day>\d+) *<span class="stats-both"> *(?P<two_stars>\d+) *</span> *<span class="stats-firstonly"> *(?P<one_star>\d+) *</span>')

reply  = req.get('https://adventofcode.com/2017/stats')

day_stats = {}

for m in re.finditer(r_day,reply.text):
    day, two_stars, one_star = list(map(int, m.group('day','two_stars','one_star')))
    day_stats[day] = (two_stars, one_star)

print(f"{'Day':>3} {'2*':>6} {'1*':>5} {'Total':>6} {'2*(%)':>8}")
for day in sorted(day_stats.keys()):
    two, one = day_stats[day]
    total = two + one
    print(f'{day:3} {two:6} {one:5} {total:6} {(two*100)/total:6.1f}%')