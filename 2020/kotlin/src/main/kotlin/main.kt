import picocli.CommandLine
import kotlin.system.exitProcess

fun main(args: Array<String>) {
        val cmd = CommandLine(AocCommandLine())
        cmd.defaultValueProvider = AocFlagDefaultsProvider()
        exitProcess(cmd.execute(*args))
}