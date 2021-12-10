using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day10 : ISolver
    {
        public enum BracketType
        {
            ROUND,
            SQUARE,
            CURLY,
            ANGLE
        }

        public class BracketToken
        {
            public BracketType BracketType;
            public int Position = 0;

            public BracketToken(BracketType type, int pos)
            {
                BracketType = type;
                Position = pos;
            }
        }

        public class BracketParser
        {
            private Dictionary<BracketType, Stack<int>> stacks { get; set; } =
                new()
                {
                    { BracketType.ROUND, new Stack<int>() },
                    { BracketType.CURLY, new Stack<int>() },
                    { BracketType.SQUARE, new Stack<int>() },
                    { BracketType.ANGLE, new Stack<int>() },
                };

            public bool InError { get; set; } = false;

            public bool OpeningsAfterPos(int pos)
            {
                return stacks[BracketType.ROUND].Any(p => p > pos) ||
                stacks[BracketType.CURLY].Any(p => p > pos) ||
                stacks[BracketType.SQUARE].Any(p => p > pos) ||
                stacks[BracketType.ANGLE].Any(p => p > pos);
            }

            public void Consume(char c, int pos)
            {
                int lastOpening;
                switch (c)
                {
                    case '(':
                        stacks[BracketType.ROUND].Push(pos);
                        break;
                    case ')':
                        if (stacks[BracketType.ROUND].Count() == 0)
                        {
                            InError = true;
                            break;
                        }
                        lastOpening = stacks[BracketType.ROUND].Pop();
                        if (OpeningsAfterPos(lastOpening))
                        {
                            InError = true;
                        }
                        break;
                    case '[':
                        stacks[BracketType.SQUARE].Push(pos);
                        break;
                    case ']':
                        if (stacks[BracketType.SQUARE].Count() == 0)
                        {
                            InError = true;
                            break;
                        }
                        lastOpening = stacks[BracketType.SQUARE].Pop();
                        if (OpeningsAfterPos(lastOpening))
                        {
                            InError = true;
                        }
                        break;
                    case '{':
                        stacks[BracketType.CURLY].Push(pos);

                        break;
                    case '}':
                        if (stacks[BracketType.CURLY].Count() == 0)
                        {
                            InError = true;
                            break;
                        }
                        lastOpening = stacks[BracketType.CURLY].Pop();
                        if (OpeningsAfterPos(lastOpening))
                        {
                            InError = true;
                        }
                        break;
                    case '<':
                        stacks[BracketType.ANGLE].Push(pos);
                        break;
                    case '>':
                        if (stacks[BracketType.ANGLE].Count() == 0)
                        {
                            InError = true;
                            break;
                        }
                        lastOpening = stacks[BracketType.ANGLE].Pop();
                        if (OpeningsAfterPos(lastOpening))
                        {
                            InError = true;
                        }
                        break;
                }

            }

            public List<char> Complete()
            {

                List<BracketToken> tokens = new();
                List<int> list;

                list = stacks[BracketType.ROUND].ToList();
                for (int i = 0; i < list.Count(); i++)
                {
                    tokens.Add(new BracketToken(BracketType.ROUND, list[i]));
                }


                list = stacks[BracketType.ANGLE].ToList();
                for (int i = 0; i < list.Count(); i++)
                {
                    tokens.Add(new BracketToken(BracketType.ANGLE, list[i]));
                }


                list = stacks[BracketType.CURLY].ToList();
                for (int i = 0; i < list.Count(); i++)
                {
                    tokens.Add(new BracketToken(BracketType.CURLY, list[i]));
                }

                list = stacks[BracketType.SQUARE].ToList();
                for (int i = 0; i < list.Count(); i++)
                {
                    tokens.Add(new BracketToken(BracketType.SQUARE, list[i]));
                }

                tokens.Sort(delegate (BracketToken a, BracketToken b)
                {
                    var diff = b.Position - a.Position;
                    return diff >= 1 ? 1 : (diff <= -1) ? -1 : 0;
                });

                return tokens.Select(t =>
                {
                    switch (t.BracketType)
                    {
                        case BracketType.ROUND:
                            return ')';
                        case BracketType.ANGLE:
                            return '>';
                        case BracketType.CURLY:
                            return '}';
                        default:
                            return ']';
                    }
                }).ToList();
            }


        }


        public char[][] Lines { get; set; }

        public void ProcessInput(string input)
        {
            Lines = Input
                        .GetLines(input)
                        .Select(line => line.ToCharArray())
                        .ToArray();
        }

        public string Part1()
        {

            Dictionary<char, int> scoreTable = new Dictionary<char, int>() {
                    {')', 3},
                    {']', 57},
                    {'}', 1197},
                    {'>', 25137}
            };


            int totalScore = 0;

            foreach (var line in Lines)
            {
                BracketParser parser = new();
                var lineScore = 0;
                for (int i = 0; i < line.Count(); i++)
                {
                    char c = line[i];
                    parser.Consume(c, i);
                    if (parser.InError)
                    {
                        // Console.WriteLine($"Error at pos {i} in char {c} for string {new String(line)}");
                        lineScore += scoreTable[c];
                        break;
                    }
                }
                totalScore += lineScore;
            }

            return totalScore.ToString();
        }

        public string Part2()
        {

            List<BracketParser> incompletes = new();
            foreach (var line in Lines)
            {
                BracketParser parser = new();
                for (int i = 0; i < line.Count(); i++)
                {
                    parser.Consume(line[i], i);
                    if (parser.InError)
                    {
                        break;
                    }
                }

                if (!parser.InError)
                {
                    incompletes.Add(parser);
                }
            }

            Dictionary<char, int> scoreTable = new Dictionary<char, int>() {
                    {')', 1},
                    {']', 2},
                    {'}', 3},
                    {'>', 4}
                };

            List<Int64> scores = new();
            foreach (BracketParser incomplete in incompletes)
            {
                List<char> completion = incomplete.Complete();

                Int64 score = 0;
                foreach (char c in completion)
                {
                    score = (score * 5) + scoreTable[c];
                }

                scores.Add(score);

            }

            scores.Sort();

            return scores[scores.Count() / 2].ToString();
        }

    }
}
