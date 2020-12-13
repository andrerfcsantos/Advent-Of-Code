package solver.y2020

import solver.Solver

class Year2020Day01 : Solver {

    var nums: List<Int> = emptyList()

    override fun processInput(content: String) {
        nums = content.split("\n").map { it.trim() }.filter { it != "" }.map { it.toInt() }
    }

    override fun part1(): String {
        val size = nums.size
        for(i in 0 until size) {
            for(j in i+1 until size - 1) {
                if ((nums[i] + nums[j]) == 2020) {
                    return (nums[i] * nums[j]).toString()
                }
            }
        }

        throw Exception("no valid numbers found")
    }

    override fun part2(): String {
        val size = nums.size
        for(i in 0 until size) {
            for(j in i+1 until size - 1) {
                for(k in j+1 until size - 2) {
                    if ((nums[i] + nums[j] + nums[k]) == 2020) {
                        return (nums[i] * nums[j] * nums[k]).toString()
                    }
                }
            }
        }

        throw Exception("no valid numbers found")
    }

}