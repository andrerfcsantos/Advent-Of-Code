# -*- coding: utf-8 -*-

DAY = 13

inp = open(f'../in/day{DAY:02}.txt')

class PassResults:
    def __init__(self, caught=False, severity=0, caught_at=None):
        self.caught = caught
        self.severity = severity
        self.caught_at = [] if caught_at == None else caught_at

    def __str__(self):
        if not self.caught:
            return 'PassResult {Not caught}'
        else:
            return ('Caught: ' + str(self.caught) + ' | ' +
                    'Severity: ' + str(self.severity) + ' | ' +
                    'Caugth at layers: ' + str(self.caught_at) )
 

firewall = {}

def pass_with_delay(delay):
    global firewall
    res = PassResults()

    run_lenght = max(firewall.keys())
    total_severity=0
    for layer_depth in range(0,run_lenght+1):
        layer_range = firewall.get(layer_depth, None)
        picoseconds = layer_depth + delay

        if layer_range != None and (picoseconds % ((layer_range-1)*2)) == 0:
            total_severity += layer_depth*layer_range
            res.caught_at.append(layer_depth)
    
    res.severity = total_severity
    res.caught = False if len(res.caught_at)==0 else True
    return res

for line in inp:
    layer, depth = [int(x) for x in line.strip().split(': ')]
    firewall[layer] = depth

part1 = pass_with_delay(0)
part2 = None

valid_run_found = False
delay = 0

while not valid_run_found:
    run_res = pass_with_delay(delay)
    if not run_res.caught:
        valid_run_found = True
        part2 = delay
    delay += 1


print(f'Part 1: {part1.severity}')
print(f'Part 2: {part2}')
