using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day06 : ISolver
    {

        public int[] numbers { get; set; }

        public void ProcessInput(string input)
        {
            List<string> lines = Input.GetLines(input);
            numbers = lines.First().Split(",").Select(int.Parse).ToArray();
        }


        public Int64 RunSimulation(int days) {
            Int64[] counts = new Int64[9];


            foreach(int num in numbers) {
                counts[num]++;
            }

            for(int d = 0; d < days ; d++) {
                Int64[] previous = counts.ToArray();

                for(int i = 0; i < previous.Count()-1; i++)
                {    
                    counts[i] += previous[i+1];
                }

                for(int i = 0; i < previous.Count(); i++)
                {    
                    counts[i] -= previous[i];
                }

                counts[6]+=previous[0];
                counts[8]+=previous[0];
            }

            return counts.Sum();
        }

        public string Part1()
        {
            return RunSimulation(80).ToString();
        }

        public string Part2()
        {
            return RunSimulation(256).ToString();
        }

    }
}
