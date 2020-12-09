from dataclasses import dataclass
from typing import List, Set
import fileinput
import networkx as nx

@dataclass(init=True, repr=True, eq=True)
class Instruction:
    op: str
    value: int

@dataclass(init=True, repr=True, eq=True)
class VM:
    instructions: List[Instruction]
    execed: Set[int]
    finished: bool = False
    acc: int = 0
    pc: int = 0
    lastPc: int = -1

graph = nx.DiGraph()

instructions: List[Instruction]

with fileinput.input(files=["../inputs/2020_08.txt"]) as f:
    for line in f:
        l = line.strip()
        if l != "":
            parts = l.split(" ")
            ins = Instruction(op=parts[0], value=int(parts[1]))
            instructions.append(ins)


