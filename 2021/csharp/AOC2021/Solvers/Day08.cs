using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day08 : ISolver
    {

        public class Display
        {
            public string[] Signals { get; set; } = new string[10];
            public string[] Output { get; set; } = new string[4];

            public Display(string[] signals, string[] output)
            {
                Signals = signals;
                Output = output;
            }

            public int DecodeOutput()
            {

                Dictionary<(bool, bool, bool, bool), char> fingerprints = new Dictionary<(bool, bool, bool, bool), char>{
                        { (false, false, true , true), 'a' },
                        { (false, true , false, true), 'b' },
                        { (true , true , true , true), 'c' },
                        { (false, false, false, false), 'd' },
                        { (false, false, false, false), 'e' },
                        { (false, false, false, false), 'f' },
                        { (false, false, false, false), 'g' },
                    };

                // Figure out the signal strings for 1, 4, 7 and 8
                // Create signal fingerprints from the signal strings

                // Match signal fingerprints to the current fingerprints
                // and create a translation map

                // 
            }

        }

        public List<Display> Displays { get; set; } = new();

        public void ProcessInput(string input)
        {
            Displays = Input.GetLines(input)
            .Select(line =>
            {
                string[][] parts = line.Split(" | ").Select(part => part.Split()).ToArray();
                return new Display(parts[0], parts[1]);
            })
            .ToList();
        }

        public string Part1()
        {
            return Displays.Select(
                d => d.Output
                    .Where(o => new int[] { 2, 3, 4, 7 }.Contains(o.Length))
                    .Count()
                ).Sum()
                .ToString();
        }

        public string Part2()
        {
            return "";
        }

    }
}
