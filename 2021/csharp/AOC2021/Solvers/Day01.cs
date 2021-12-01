using AOCLibrary;

namespace AOC2021.Solvers
{
    public class Day01 : ISolver
    {
        List<int> depths = new();

        public void ProcessInput(string input)
        {
            depths = new List<int>();
            foreach (string line in Input.GetLines(input))
            {
                depths.Add(Convert.ToInt32(line));
            }
        }

        public string Part1()
        {
            int count = 0;
            for(int i = 1; i < depths.Count; i++)
            {
                if (depths[i]>depths[i-1])
                {
                    count++;
                }
            }
            return count.ToString();
        }

        public string Part2()
        {
            int count = 0;
            for (int i = 0; i < depths.Count-3; i++)
            {
                if (depths[i+3] > depths[i])
                {
                    count++;
                }
            }
            return count.ToString();
        }

    }
}
