import re
import networkx as nx
import matplotlib.pyplot as plt
PATTERN = re.compile(r"Step (?P<source>\w) must be finished before step (?P<dest>\w) can begin.")

# LAPFCRGHZTKWENBXIMVOSUDJQY
global_graph = None
def getInputGraph():
    G = nx.DiGraph()
    with open("../inputs/2018_07_test2.txt") as in_file:
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



def topologicalOrder(graph):
    order, stack, visited = [], [], set()
    
    # Check roots of the graph, i.e nodes with no predeccessors
    # and add them to the stack
    roots = sorted(list(getRoots(graph)),reverse=True)
    stack.extend(roots)

    while len(stack) > 0:
        # Get node from the stack, mark it visited and add it
        # to the topological order.
        current_node = stack.pop()
        
        print("Current node", current_node)
        order.append(current_node)
        visited.add(current_node)

        # Get successors of current node and check if they can
        # be addded to the stack
        successors = sorted(list(graph.successors(current_node)),reverse=True)

        for s in successors:
            if s not in visited and s not in set(stack) and canStart(graph, s, visited):
                # A successor of the node can be added to the stack
                # if it was not already visited, it's not currently on the stack
                # and the tasks it depended on are completed
                stack.append(s)

        print("Stack",stack, "visited", visited)
    return ''.join(order)


def main():
    global global_graph
    graph = getInputGraph()
    global_graph = graph
    to = topologicalOrder(graph)
    print(to)
    print(set(graph.nodes)==set(to))

    


if __name__ == "__main__":
    main()
    
