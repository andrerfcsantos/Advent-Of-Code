using CommandLine;
using System.Diagnostics;
using System.Reflection;
using AOCLibrary;

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
				(opts) => RunWithReflection(opts), //in case parser sucess
				errs => HandleParseError(errs)); //in  case parser fail
		}

		static int RunWithReflection(Options o)
		{
			if (o.InputPath == null)
            {
				o.InputPath = $"..\\..\\inputs\\2021_{o.Day:D2}.txt";
			}

			Console.WriteLine($"day={o.Day:D2} input={o.InputPath}");

			Stopwatch stopWatch = new Stopwatch();

			stopWatch.Start();
			string fileData = File.ReadAllText(o.InputPath);
			TimeSpan readTime = stopWatch.Elapsed;
			stopWatch.Reset();

			Type? t = Type.GetType($"AOC2021.Solvers.Day{o.Day:D2}");
			ConstructorInfo? c = t.GetConstructor(new Type[] { });
			var inst = c.Invoke(null);


			stopWatch.Start();
			MethodInfo? pi = t.GetMethod("ProcessInput", new Type[] { typeof(string) });
            _ = pi.Invoke(inst, new object[] { fileData });
			TimeSpan processTime = stopWatch.Elapsed;
			stopWatch.Reset();

			stopWatch.Start();
			MethodInfo? p1 = t.GetMethod("Part1", new Type[] {  });
			var p1Res = p1.Invoke(inst, new object[] { });
			TimeSpan p1Time = stopWatch.Elapsed;
			stopWatch.Reset();

			stopWatch.Start();
			MethodInfo? p2 = t.GetMethod("Part2", new Type[] { });
			var p2Res = p2.Invoke(inst, new object[] { });
			TimeSpan p2Time = stopWatch.Elapsed;

			stopWatch.Reset();

			Console.WriteLine($"p1={p1Res} p2={p2Res}");
			Console.WriteLine($"Timings: read={readTime.TotalMilliseconds}ms process={processTime.TotalMilliseconds}ms p1={p1Time.TotalMilliseconds} p2={p2Time.TotalMilliseconds}ms");

			return 0;
		}

		static int Run(Options o)
		{
			if (o.InputPath == null)
			{
				o.InputPath = $"..\\..\\inputs\\2021_{o.Day:D2}.txt";
			}

			Console.WriteLine($"day={o.Day:D2} input={o.InputPath}");

			Stopwatch stopWatch = new Stopwatch();

			stopWatch.Start();
			string fileData = File.ReadAllText(o.InputPath);
			TimeSpan readTime = stopWatch.Elapsed;
			stopWatch.Reset();

			ISolver solver = new Solvers.Day01();

			stopWatch.Start();
			solver.ProcessInput(fileData);
			TimeSpan processTime = stopWatch.Elapsed;
			stopWatch.Reset();

			stopWatch.Start();
			string p1Res = solver.Part1();
			TimeSpan p1Time = stopWatch.Elapsed;
			stopWatch.Reset();

			stopWatch.Start();
			string p2Res = solver.Part2();
			TimeSpan p2Time = stopWatch.Elapsed;
			stopWatch.Reset();

			Console.WriteLine($"p1={p1Res} p2={p2Res}");
			Console.WriteLine($"Timings: read={readTime.TotalMilliseconds}ms process={processTime.TotalMilliseconds}ms p1={p1Time.TotalMilliseconds} p2={p2Time.TotalMilliseconds}ms");

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