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

            MostCommonBit[] bitPopularity = mostCommonBits(numbers);

            int cols = numbers.First().Count();

            int gammaRate = 0;
            int epsilonRate = 0;

            for (int i = 0; i < cols; i++)
            {
                if (bitPopularity[i] == MostCommonBit.ZERO)
                {
                    gammaRate += 1 << (cols - 1 - i);
                }
                else
                {
                    epsilonRate += 1 << (cols - 1 - i);
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


        public enum BitCriteria
        {
            MOST_COMMON,
            LEAST_COMMON
        }


        private int bitsToNumber(int[] bits)
        {
            int n = 0;
            int size = bits.Count();

            for (int i = 0; i < size; i++)
            {
                if (bits[i] == 1)
                {
                    n += 1 << (size - 1 - i);
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


        private int[] filterByCriteria(int[][] bits, BitCriteria criteria)
        {
            bits = bits.ToArray();

            Func<int, bool> zeroMatch = (criteria == BitCriteria.MOST_COMMON) ?
                                        (i) => i == 0 :
                                        (i) => i == 1;

            for (int i = 0; i < bits.Length && bits.Count() > 1; i++)
            {
                MostCommonBit[] popularity = mostCommonBits(bits);
                var pop = popularity[i];

                switch (pop)
                {
                    case MostCommonBit.ZERO:
                        bits = bits.Where(ds => zeroMatch(ds[i])).ToArray();
                        break;
                    case MostCommonBit.ONE:
                        bits = bits.Where(ds => !zeroMatch(ds[i])).ToArray();
                        break;
                    case MostCommonBit.EQUAL:
                        bits = bits.Where(ds => !zeroMatch(ds[i])).ToArray();
                        break;
                }

            }

            return bits[0];
        }
        public string Part2()
        {
            int cols = numbers.First().Count();

            int[] oxygenBits= filterByCriteria(numbers.ToArray(), BitCriteria.MOST_COMMON);
            int oxygenRating = bitsToNumber(oxygenBits);

            int[] co2Bits= filterByCriteria(numbers.ToArray(), BitCriteria.LEAST_COMMON);
            int co2Rating = bitsToNumber(co2Bits);

            Console.WriteLine($"co2Rating={co2Rating} oxygenRating={oxygenRating}");
            return (co2Rating * oxygenRating).ToString();
        }

    }
}
