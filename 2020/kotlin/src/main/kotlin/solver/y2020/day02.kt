package solver.y2020

import solver.Solver

data class PasswordPolicy(val num1: Int, val num2: Int, val char: Char, val password: String) {
    fun isValidByCharacterFrequency(): Boolean = password.count { it == char } in num1..num2
    fun isValidByCharacterPositions(): Boolean = (password[num1 - 1] == char).xor(password[num2 - 1] == char)
}

class Year2020Day02 : Solver {

    var passwords: MutableList<PasswordPolicy> = mutableListOf()

    override fun processInput(content: String) {
        val parts = content.split("\n").map { it.trim() }.filter { it != "" }.map { it.split(" ") }

        for (part in parts) {
            val nums = part[0].split("-").map { it.toInt() }
            val char = part[1].trim(':').first()
            passwords.add(PasswordPolicy(nums[0], nums[1], char, part[2]))
        }
    }

    override fun part1(): String = passwords.count { it.isValidByCharacterFrequency() }.toString()

    override fun part2(): String = passwords.count { it.isValidByCharacterPositions() }.toString()

}