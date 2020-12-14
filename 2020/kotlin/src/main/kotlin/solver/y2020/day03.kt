package solver.y2020

import solver.Solver

class Year2020Day03 : Solver {

    var board: List<List<Boolean>> = emptyList()

    override fun processInput(content: String) {
        board = content.split("\n").map { it.trim() }.filter { it != "" }
            .map { l -> l.toCharArray().map { c -> c == '#' } }
    }

    override fun part1(): String = runSlope(1,3).toString()

    override fun part2(): String  {
        return listOf(
            runSlope(1,1),
            runSlope(1,3),
            runSlope(1,5),
            runSlope(1,7),
            runSlope(2,1),
        ).map { it.toLong() }.reduce { acc, i -> acc*i }.toString()
    }

    fun runSlope(down: Int, right: Int) : Int {
        var x = 0
        var y = 0
        var trees = 0

        val lines = board.size
        val cols = board[0].size

        while (y < lines) {
            if(board[y][x%cols]){
                trees++
            }

            x+=right
            y += down
        }
        return trees
    }
}