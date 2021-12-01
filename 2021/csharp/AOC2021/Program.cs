using AOCLibrary;
using CommandLine;
using System.Reflection;

namespace AOC2021
{
    class Program
    {
        public class Options
        {
            [Option('d', "day", Required = true, HelpText = "Day of advent of code")]
            public int Day { get; set; }

            [Option('i', "input", Required = false, HelpText = "Input file")]
            public string? InputPath { get; set; }
        }

        static void Main(string[] args)
        {
			var result = Parser.Default.ParseArguments<Options>(args).MapResult(
				(opts) => RunOptionsAndReturnExitCode(opts), //in case parser sucess
				errs => HandleParseError(errs)); //in  case parser fail
		}

		static int RunOptionsAndReturnExitCode(Options o)
		{
			if (o.InputPath == null)
            {
				o.InputPath = $"Inputs\\Day{o.Day:D2}.txt";
			}

			Console.WriteLine($"day={o.Day:D2} input={o.InputPath}");

			string fileData = File.ReadAllText(o.InputPath);

			Type? t = Type.GetType($"AOC2021.Solvers.Day{o.Day:D2}");
			ConstructorInfo? c = t.GetConstructor(new Type[] { });
			var inst = c.Invoke(null);


			MethodInfo? pi = t.GetMethod("ProcessInput", new Type[] { typeof(string) });
            _ = pi.Invoke(inst, new object[] { fileData });

            MethodInfo? p1 = t.GetMethod("Part1", new Type[] {  });
			var p1Res = p1.Invoke(inst, new object[] { });

			MethodInfo? p2 = t.GetMethod("Part2", new Type[] { });
			var p2Res = p2.Invoke(inst, new object[] { });

			Console.WriteLine($"p1={p1Res} p2={p2Res}");

			return 0;
		}

		//in case of errors or --help or --version
		static int HandleParseError(IEnumerable<Error> errs)
		{
			var result = -2;
			Console.WriteLine("errors {0}", errs.Count());
			if (errs.Any(x => x is HelpRequestedError || x is VersionRequestedError))
				result = -1;
			Console.WriteLine("Exit code {0}", result);
			return result;
		}
	}
}