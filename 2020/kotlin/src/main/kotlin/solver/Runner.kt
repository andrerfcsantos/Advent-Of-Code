package solver

import java.io.File

class Runner(val input: String = "", val solver: Solver) {

    fun run(): Solution {
        solver.processInput(input)
        val p1 = solver.part1()
        println("Part 1: $p1")
        val p2 = solver.part2()
        println("Part 2: $p2")
        return Solution(part1=p1,part2=p2)
    }
}