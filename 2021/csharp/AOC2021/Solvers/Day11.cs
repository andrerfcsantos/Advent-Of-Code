using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day11 : ISolver
    {

        public int[][] Energies { get; set; }

        public void ProcessInput(string input)
        {
            Energies = Input
                        .GetLines(input)
                        .Select(line => line.ToCharArray().Select(d => d - '0').ToArray())
                        .ToArray();
        }

        public int[][] EnergiesCopy()
        {
            int rows = Energies.Length;
            int cols = Energies[0].Length;

            int[][] copy = new int[rows][];
            for (int i = 0; i < rows; i++)
            {
                copy[i] = new int[cols];
                for (int j = 0; j < cols; j++)
                {
                    copy[i][j] = Energies[i][j];
                }
            }
            return copy;
        }

        private bool inBounds(int[][] arr, int i, int j)
        {
            if (i < 0 || j < 0 || i >= arr.Length || j >= arr[i].Length)
            {
                return false;
            }
            return true;
        }

        private void printMatrix(int[][] arr)
        {
            int rows = Energies.Length;

            int[][] copy = new int[rows][];
            for (int i = 0; i < rows; i++)
            {
                Console.WriteLine(String.Join("", arr[i]));
            }
        }

        private int getEnergyAt(int[][] arr, int n, int m, int defaultVal)
        {
            return inBounds(arr, n, m) ? arr[n][m] : defaultVal;
        }


        private void incEnergyAt(int[][] arr, int n, int m)
        {
            if (inBounds(arr, n,m))
            {
                arr[n][m] += 1;
            }
        }

        private void setEnergyAt(int[][] arr, int n, int m, int val)
        {
            if (inBounds(arr, n,m))
            {
                arr[n][m] = val;
            }
        }

        public string Part1()
        {
            int[][] energies = EnergiesCopy();
            int totalFlashed = 0;
            for (int i = 0; i < 100; i++)
            {

                HashSet<(int, int)> flashed = new();

                for (int j = 0; j < energies.Length; j++)
                {
                    for (int k = 0; k < energies[j].Length; k++)
                    {
                        energies[j][k] += 1;
                    }
                }

                int previouslyFlashed = -1;
                while (previouslyFlashed != flashed.Count)
                {
                    previouslyFlashed = flashed.Count;

                    for (int j = 0; j < Energies.Length; j++)
                    {
                        for (int k = 0; k < Energies[j].Length; k++)
                        {
                            int cell = energies[j][k];

                            if(cell > 9 && !flashed.Contains((j,k))) {
                                flashed.Add((j, k));
                                incEnergyAt(energies, j - 1, k);
                                incEnergyAt(energies, j + 1, k);
                                incEnergyAt(energies, j, k - 1);
                                incEnergyAt(energies, j, k + 1);
                                incEnergyAt(energies, j - 1, k - 1);
                                incEnergyAt(energies, j + 1, k - 1);
                                incEnergyAt(energies, j - 1, k + 1);
                                incEnergyAt(energies, j + 1, k + 1);

                            }
                        }
                    }

                }

                foreach (var (x, y) in flashed)
                {
                    setEnergyAt(energies, x, y, 0);
                }

                totalFlashed += flashed.Count;
            }
            return totalFlashed.ToString();
        }

        public string Part2()
        {

            int octopuses = Energies.Length * Energies[0].Length;
            int res = 0;
            for (int i = 0; true; i++)
            {
                int[][] previous = EnergiesCopy();

                HashSet<(int, int)> flashed = new();

                for (int j = 0; j < Energies.Length; j++)
                {
                    for (int k = 0; k < Energies[j].Length; k++)
                    {
                        previous[j][k] += 1;
                    }
                }

                int previouslyFlashed = -1;
                while (previouslyFlashed != flashed.Count)
                {
                    previouslyFlashed = flashed.Count;

                    for (int j = 0; j < Energies.Length; j++)
                    {
                        for (int k = 0; k < Energies[j].Length; k++)
                        {
                            int cell = previous[j][k];

                            if(cell > 9 && !flashed.Contains((j,k))) {
                                flashed.Add((j, k));
                                incEnergyAt(previous, j - 1, k);
                                incEnergyAt(previous, j + 1, k);
                                incEnergyAt(previous, j, k - 1);
                                incEnergyAt(previous, j, k + 1);
                                incEnergyAt(previous, j - 1, k - 1);
                                incEnergyAt(previous, j + 1, k - 1);
                                incEnergyAt(previous, j - 1, k + 1);
                                incEnergyAt(previous, j + 1, k + 1);

                            }
                        }
                    }

                }

                foreach (var (x, y) in flashed)
                {
                    setEnergyAt(previous, x, y, 0);
                }

                if (octopuses == flashed.Count)
                {
                    res = i+1;
                    break;
                }
                Energies = previous;
            }
            return (res).ToString();
        }

    }
}
