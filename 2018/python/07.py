import re
import networkx as nx

PATTERN = re.compile(r"Step (?P<source>\w) must be finished before step (?P<dest>\w) can begin.")

def getInputGraph():
    G = nx.DiGraph()
    with open("../inputs/2018_07.txt") as in_file:
        for line in in_file:
            line = line.strip()
            m = re.match(PATTERN, line)
            if m:
                source, dest = m.groups(["source", "dest"])
                G.add_edge(source, dest)
    return G

def getRoots(graph):
    nodes = set(graph.nodes)

    for source, dest in graph.edges:
        if dest in nodes:
            nodes.remove(dest)

    return nodes

def canStart(graph, s, visited):
    predecessors = list(graph.predecessors(s))

    for predecessor in predecessors:
        if predecessor not in visited:
            return False
    
    return True

def get_task_order(graph):

    n_nodes = len(list(graph.nodes))
    order, available = [], []
    roots = sorted(list(getRoots(graph)),reverse=True)
    available.extend(roots)


    while len(order) != n_nodes:
        available = list(sorted(available,reverse=True))
        current_node = available.pop()
        order.append(current_node)
        successors = list(graph.successors(current_node))
        for s in successors:
            if s not in set(order) and canStart(graph, s, set(order)):
                available.append(s)

    return ''.join(order)

def time_completion(graph):
    time_tick = 0
    total_workers = 5
    n_nodes = len(list(graph.nodes))
    done, available = set(), []
    available_workers, time_to_complete = total_workers, {}

    available = getRoots(graph)

    while len(done) != n_nodes:
        available = list(sorted(available,reverse=True))

        while available_workers > 0 and len(available) >0:
            task = available.pop()
            available_workers -= 1
            time_to_complete[task] = 60 + (ord(task)-ord('A')+1)

        tasks_in_completion = list(time_to_complete.keys())
        for task_in_completion in tasks_in_completion:
            time_to_complete[task_in_completion]-=1

            if time_to_complete[task_in_completion]==0:
                print(f"Task {task_in_completion} done")
                done.add(task_in_completion)
                available_workers+=1
                del time_to_complete[task_in_completion]
                next_tasks = list(graph.successors(task_in_completion))

                for next_task in next_tasks:
                    if next_task not in done and canStart(graph, next_task, done):
                        available.append(next_task)
        
        time_tick+=1

    return time_tick

def main():
    graph = getInputGraph()
    task_order = get_task_order(graph)
    print("Part 1:", task_order)
    print("Part 2:", time_completion(graph))

if __name__ == "__main__":
    main()
    
