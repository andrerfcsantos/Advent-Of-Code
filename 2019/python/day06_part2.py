import networkx as nx

g = nx.DiGraph()

with open("../inputs/day06.txt") as f:
    for line in f:
        src, dest = line.strip("\n\r\t").split(")")
        g.add_edge(src, dest)

lca = nx.lowest_common_ancestor(g, "YOU", "SAN")
to_you = len(nx.shortest_path(g, lca, "YOU"))-2
to_san = len(nx.shortest_path(g, lca, "SAN"))-2
print(to_you + to_san)
