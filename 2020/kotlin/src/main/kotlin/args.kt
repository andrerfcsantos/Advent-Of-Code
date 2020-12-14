import picocli.CommandLine
import solver.Runner
import solver.y2020.*
import java.io.File
import java.nio.file.Paths
import java.time.*
import java.util.*
import kotlin.math.ceil

val solvers = mapOf(
    2020 to mapOf(
        1 to Year2020Day01(),
        2 to Year2020Day02(),
        3 to Year2020Day03(),
    ),
)

class AocFlagDefaultsProvider : CommandLine.IDefaultValueProvider {

    data class AocChallenge(val year: Int, val day: Int)

    companion object {
        var challenge: AocChallenge? = null
    }

    private fun getCurrentSession(): String = System.getenv("AOC_SESSION")

    private fun getCurrentAocChallenge(): AocChallenge {
        var year = Calendar.getInstance().get(Calendar.YEAR)

        val aocStart = ZonedDateTime.of(
            year,
            Month.DECEMBER.value,
            1, 5, 0, 0, 0,
            ZoneOffset.UTC
        )

        val daysSinceStart: Int = ceil(Duration.between(aocStart.toInstant(), Instant.now()).toHours() / 24.0).toInt()

        var day = 0

        when {
            daysSinceStart < 0 -> {
                day = 25
                year -= 1
            }
            daysSinceStart > 25 -> {
                day = 25
            }
            else -> {
                day = daysSinceStart
            }
        }

        return AocChallenge(day = day, year = year)
    }

    override fun defaultValue(p0: CommandLine.Model.ArgSpec?): String {
        if (challenge == null) {
            challenge = getCurrentAocChallenge()
        }

        return when (p0?.paramLabel()) {
            "<year>" -> challenge!!.year.toString()
            "<day>" -> challenge!!.day.toString()
            "<aocSession>" -> getCurrentSession()
            else -> p0?.defaultValue() ?: ""
        }
    }

}

class AocCommandLine : Runnable {
    @CommandLine.Option(
        names = ["-y", "--year"],
        description = ["advent of code year"],
    )
    var year: Int = 0

    @CommandLine.Option(
        names = ["-d", "--day"],
        defaultValue = "1",
        description = ["advent of code day"],
    )
    var day: Int = 0

    @CommandLine.Option(
        names = ["--download"],
        defaultValue = "false",
        description = ["download input file for the the day/year specified by the flags --day and --year"],
    )
    var download: Boolean = false

    @CommandLine.Option(
        names = ["--download-only"],
        defaultValue = "false",
        description = ["download input file for the the day/year specified by the flags --day and --year"],
    )
    var downloadOnly: Boolean = false

    @CommandLine.Option(
        names = ["-i", "--input-dir"],
        defaultValue = "inputs",
        description = ["download input file for the the day/year specified by the flags --day and --year"],
    )
    lateinit var inputPath: File

    @CommandLine.Option(
        names = ["-s", "--session"],
        description = ["download input file for the the day/year specified by the flags --day and --year"],
    )
    var aocSession: String = ""

    fun Int.format(digits: Int) = "%0${digits}d".format(this)

    override fun run() {
        if (downloadOnly) download = true

        var filePath = Paths.get(inputPath.toString(), "${year}_${day.format(2)}.txt").toString()

        if (inputPath.isDirectory) {
            filePath = Paths.get(inputPath.toString(), "${year}_${day.format(2)}.txt").toString()
        }

        var input: String = ""

        if (download) {
            inputPath.mkdirs()
            input = downloadInput(session = aocSession, file = File(filePath), year = year, day = day)
        }

        if (input == "") {
            input = loadInputFile(file = File(filePath))
        }

        solvers[year]?.get(day)?.let { Runner(input = input, solver = it) }?.run()
    }
}