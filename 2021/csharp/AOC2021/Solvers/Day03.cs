using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day03 : ISolver
    {
        int[][] numbers;

        public void ProcessInput(string input)
        {
            List<int[]> nums = new();

            foreach (string line in Input.GetLines(input))
            {
                int[] digits = line.ToCharArray()
                    .Select(c => c - '0')
                    .ToArray();

                nums.Add(digits);
            }
            numbers = nums.ToArray();
        }

        public string Part1()
        {
            int cols = numbers.First().Count();

            int gammaRate = 0;
            int epsilonRate = 0;

            for (int i = 0; i < cols; i++)
            {
                var digits = numbers.Select(row => row[i]);

                int zeroes = digits.Count(d => d == 0);
                int ones = digits.Count(d => d == 1);

                gammaRate <<= 1;
                epsilonRate <<= 1;
                if (zeroes > ones)
                {
                    gammaRate |= 0b1;
                }
                else
                {
                    epsilonRate |= 0b1;
                }
            }
            return (gammaRate * epsilonRate).ToString();
        }


        public enum MostCommonBit
        {
            ZERO,
            ONE,
            EQUAL
        }


        private int bitsToNumber(int[] bits)
        {
            int n = 0;
            for (int i = 0; i < bits.Count(); i++)
            {
                n <<= 1;
                if (bits[i] == 1)
                {
                    n |= 0b1;
                }
            }

            return n;
        }

        private MostCommonBit[] mostCommonBits(int[][] digits)
        {
            int cols = digits.First().Count();

            MostCommonBit[] popularity = new MostCommonBit[cols];

            for (int i = 0; i < cols; i++)
            {
                var colDigits = digits.Select(row => row[i]);

                int zeroes = colDigits.Count(d => d == 0);
                int ones = colDigits.Count(d => d == 1);

                MostCommonBit comp;
                if (zeroes > ones) comp = MostCommonBit.ZERO;
                else if (zeroes < ones) comp = MostCommonBit.ONE;
                else comp = MostCommonBit.EQUAL;

                popularity[i] = comp;
            }


            return popularity;
        }
        public string Part2()
        {
            int cols = numbers.First().Count();

            int[][] oxygenNumbers = numbers.ToArray();

            for (int i = 0; i < cols && oxygenNumbers.Count() > 1; i++)
            {
                MostCommonBit[] popularity = mostCommonBits(oxygenNumbers);
                var pop = popularity[i];

                switch (pop)
                {
                    case MostCommonBit.ZERO:
                        oxygenNumbers = oxygenNumbers.Where(ds => ds[i] == 0).ToArray();
                        break;
                    case MostCommonBit.ONE:
                        oxygenNumbers = oxygenNumbers.Where(ds => ds[i] == 1).ToArray();
                        break;
                    case MostCommonBit.EQUAL:
                        oxygenNumbers = oxygenNumbers.Where(ds => ds[i] == 1).ToArray();
                        break;
                }

            }


            int oxygenRating = bitsToNumber(oxygenNumbers[0]);

            int[][] co2Numbers = numbers.ToArray();
            for (int i = 0; i < cols && co2Numbers.Count() > 1; i++)
            {
                MostCommonBit[] popularity = mostCommonBits(co2Numbers);
                var pop = popularity[i];

                switch (pop)
                {
                    case MostCommonBit.ZERO:
                        co2Numbers = co2Numbers.Where(ds => ds[i] == 1).ToArray();
                        break;
                    case MostCommonBit.ONE:
                        co2Numbers = co2Numbers.Where(ds => ds[i] == 0).ToArray();
                        break;
                    case MostCommonBit.EQUAL:
                        co2Numbers = co2Numbers.Where(ds => ds[i] == 0).ToArray();
                        break;
                }

            }

            int co2Rating = bitsToNumber(co2Numbers[0]);

            Console.WriteLine($"co2Rating={co2Rating} oxygenRating={oxygenRating}");
            return (co2Rating * oxygenRating).ToString();
        }

    }
}
