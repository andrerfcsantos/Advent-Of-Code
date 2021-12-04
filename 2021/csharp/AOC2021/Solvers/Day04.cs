using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day04 : ISolver
    {


        public class Card
        {
            private int[][] numbers;
            private bool[][] marked;
            private int rows;
            private int cols;

            public Card(int[][] numbers)
            {
                this.rows = numbers.Length;
                this.cols = numbers[0].Length;

                this.numbers = new int[rows][];
                this.marked = new bool[rows][];

                for (int i = 0; i< rows; i++)
                {
                    this.marked[i] = new bool[cols];
                    this.numbers[i] = new int[cols];
                    for (int j = 0; j < cols; j++)
                    {
                        this.marked[i][j] = false;
                        this.numbers[i][j] = numbers[i][j];
                    }
                }

            }
               
            public void MarkNumber(int number)
            {
                for (int i = 0; i < rows; i++)
                {
                    for (int j = 0; j < cols; j++)
                    {
                        if(numbers[i][j] == number)
                        {
                            marked[i][j] = true;
                        }
                    }
                }
            }

            public int Score()
            {
                int score = 0;
                for (int i = 0; i < rows; i++)
                {
                    for (int j = 0; j < cols; j++)
                    {
                        if (!marked[i][j])
                        {
                            score += numbers[i][j];
                        }
                    }
                }
                return score;
            }
            public bool IsWinner()
            {
                return bingoInRow() || bingoInColumn();
            }

            private bool bingoInRow()
            {
                return marked.Any(row => row.All(b => b));
            }

            private bool bingoInColumn()
            {
                for(int icol = 0; icol < cols; icol++)
                {
                    if(marked.Select(row => row[icol]).All(b => b))
                    {
                        return true;
                    };
                }
                return false;
            }
        }


        public class Game
        {
            internal int currentDraw = 0;
            internal int[] draws;
            internal List<Card> cards = new();
            internal List<Card> winners = new();

            public void Draw()
            {
                int n = draws[currentDraw];


                List<Card> toRemove = new();
                foreach(Card card in cards)
                {
                    card.MarkNumber(n);
                    if(card.IsWinner())
                    {
                        toRemove.Add(card);
                        winners.Add(card);
                    }
                }


                foreach(Card card in toRemove)
                {
                    cards.Remove(card);
                }

                currentDraw++;
            }

            public int LastNumberCalled()
            {
                return draws[currentDraw-1];
            }

            public bool HasWinner()
            {
                return winners.Count() > 0;
            }

            public List<Card> Winners()
            {
                return winners;
            }

            public void RunUntilWinner() {

                while(!this.HasWinner() && currentDraw < draws.Length)
                {
                    this.Draw();
                }
            }

            public void RunUntilLastWinner()
            {

                while (this.cards.Count() > 0 && currentDraw < draws.Length)
                {
                    this.Draw();
                }
            }

        }

        internal Game game = new();

        public void ProcessInput(string input)
        {
            List<int[]> nums = new();
            var lineGroups = Input.GetGroupedLines(input);

            game.draws = lineGroups[0][0].Split(',').Select(n => int.Parse(n)).ToArray();

            var cards = lineGroups.GetRange(1, lineGroups.Count() - 1);

            foreach(List<string> card in cards)
            {
                int[][] cardNumbers = new int[card.Count()][];

                for (int i = 0; i < card.Count(); i++)
                {
                    var cardRow = card[i];
                    cardNumbers[i] = cardRow.Split(' ', StringSplitOptions.RemoveEmptyEntries).Select(n => int.Parse(n)).ToArray();
                }

                game.cards.Add(new Card(cardNumbers));
            }
        }

        public string Part1()
        {
            game.RunUntilWinner();
            List<Card> winners = game.Winners();
            Card firstWinner = winners[0];

            return (firstWinner.Score() * game.LastNumberCalled()).ToString();
        }




        public string Part2()
        {
            game.RunUntilLastWinner();
            List<Card> winners = game.Winners();
            Card lastWinner = winners[winners.Count()-1];

            return (lastWinner.Score() * game.LastNumberCalled()).ToString();
        }

    }
}
