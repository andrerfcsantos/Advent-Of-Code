using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day07 : ISolver
    {

        public int[] Positions { get; set; }

        public void ProcessInput(string input)
        {
            Positions = Input.GetLines(input).First().Split(",").Select(int.Parse).ToArray();
        }

        static public int computeMinFuelCost(int[] positions, Func<int, int> costFunc) {
            int minPos = positions.Min();
            int maxPos = positions.Max();

            int minFuel = int.MaxValue;
            for (int i = minPos; i <= maxPos; i++)
            {
                int fuel = 0;
                foreach(int pos in positions){
                    int displacement = Math.Abs(pos-i);
                    fuel += costFunc(displacement);
                }

                if (fuel < minFuel) {
                    minFuel = fuel;
                }

            }

            return minFuel;
        }

        public string Part1()
        {
            return computeMinFuelCost(Positions, n => n).ToString();
        }

        public string Part2()
        {
            return computeMinFuelCost(Positions, n => (n*n +n)/2).ToString();
        }

    }
}
