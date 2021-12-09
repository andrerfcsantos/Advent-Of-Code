using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day09 : ISolver
    {

        public int[][] Heights { get; set; }

        public void ProcessInput(string input)
        {
            Heights = Input
                        .GetLines(input)
                        .Select(line => line.ToCharArray().Select(d => d - '0').ToArray())
                        .ToArray();
        }

        private int getHeight(int n, int m, int defaultVal)
        {
            return inBounds(n, m) ? Heights[n][m] : defaultVal;
        }

        private bool isLowPoint(int n, int m)
        {
            int h = getHeight(n, m, int.MaxValue);
            return h < getHeight(n + 1, m, int.MaxValue) &&
                    h < getHeight(n, m + 1, int.MaxValue) &&
                    h < getHeight(n - 1, m, int.MaxValue) &&
                    h < getHeight(n, m - 1, int.MaxValue);
        }
        public string Part1()
        {
            int riskLevel = 0;
            for (int i = 0; i < Heights.Length; i++)
            {
                for (int j = 0; j < Heights[i].Length; j++)
                {
                    int h = getHeight(i, j, 0);
                    if (isLowPoint(i, j))
                    {
                        riskLevel += 1 + h;
                    }
                }
            }
            return riskLevel.ToString();
        }

        private bool inBounds(int i, int j)
        {
            if (i < 0 || j < 0 || i >= Heights.Length || j >= Heights[i].Length)
            {
                return false;
            }
            return true;
        }

        private int basinSizeFrom(int i, int j, HashSet<(int x, int y)> visited)
        {
            if (!inBounds(i, j) || getHeight(i, j, 0) == 9 || visited.Contains((i, j)))
            {
                return 0;
            }

            visited.Add((i, j));

            int childBasinSizes = 0;

            childBasinSizes += basinSizeFrom(i + 1, j, visited);
            childBasinSizes += basinSizeFrom(i, j + 1, visited);
            childBasinSizes += basinSizeFrom(i - 1, j, visited);
            childBasinSizes += basinSizeFrom(i, j - 1, visited);

            return 1 + childBasinSizes;
        }
        public string Part2()
        {
            Dictionary<(int, int), int> basinSizes = new();

            for (int i = 0; i < Heights.Length; i++)
            {
                for (int j = 0; j < Heights[i].Length; j++)
                {
                    if (isLowPoint(i, j))
                    {
                        int size = basinSizeFrom(i, j, new());
                        basinSizes.Add((i, j), size);
                    }
                }
            }

            var res = (from entry in basinSizes
                       orderby entry.Value descending
                       select entry)
                        .Take(3)
                        .Select(e => e.Value)
                        .Aggregate((a, b) => a * b);
            return res.ToString();
        }

    }
}
