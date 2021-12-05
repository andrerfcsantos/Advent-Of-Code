using AOCLibrary;

namespace AOC2021.Solvers
{

    public class Day05 : ISolver
    {

        public class Point : IEquatable<Point>, ICloneable
        {
            public int X { get; set; }
            public int Y { get; set; }


            public Point(int X, int Y)
            {
                this.X = X;
                this.Y = Y;
            }

            public bool Equals(Point other)
            {
                if (other is null)
                    return false;

                return X == other.X && Y == other.Y;
            }

            public override bool Equals(object obj) => Equals(obj as Point);

            public override int GetHashCode() => (X, Y).GetHashCode();

            public override string ToString() => $"({X},{Y})";

            public object Clone()
            {
                return (Point)MemberwiseClone();
            }
        }


        public class Line : ICloneable
        {
            public Point From { set; get; } = new(0, 0);
            public Point To { set; get; } = new(0, 0);

            public Line(Point from, Point to)
            {
                From = from;
                To = to;
            }
            public bool isHorizontal()
            {
                return From.X == To.X;
            }

            public bool isVertical()
            {
                return From.Y == To.Y;
            }

            public bool isOneDimensionalLine()
            {
                return this.isHorizontal() || this.isVertical();
            }

            public List<Point> Points()
            {
                int dx = To.X - From.X;
                int dy = To.Y - From.Y;

                (int x, int y) dv = (
                    x: dx != 0 ? dx / Math.Abs(dx) : 0,
                    y: dy != 0 ? dy / Math.Abs(dy) : 0
                );
                (int x, int y) current = (From.X, From.Y);

                List<Point> res = new();
                while (current.x != To.X || current.y != To.Y)
                {
                    res.Add(new Point(current.x, current.y));
                    current = (current.x + dv.x, current.y + dv.y);
                }
                res.Add(new Point(To.X, To.Y));

                return res;

            }

            public object Clone()
            {
                var clone = (Line)MemberwiseClone();
                clone.From = (Point)From.Clone();
                clone.To = (Point)To.Clone();
                return clone;
            }
        }

        public class Board
        {
            public List<Line> Lines { get; set; } = new();

            public void AddLine(Line l)
            {
                Lines.Add(l);
            }

            public Dictionary<Point, int> OneDimensionalOverlaps()
            {
                List<Line> oneDimLines = Lines.Where(p => p.isOneDimensionalLine()).ToList();
                Dictionary<Point, int> res = new();

                foreach (Line l in oneDimLines)
                {

                    List<Point> linePoints = l.Points();

                    foreach (Point pt in linePoints)
                    {
                        res[pt] = res.GetValueOrDefault(pt) + 1;
                    }

                }


                return res;
            }

            public Dictionary<Point, int> Overlaps()
            {
                Dictionary<Point, int> res = new();

                foreach (Line l in Lines)
                {

                    List<Point> linePoints = l.Points();

                    foreach (Point pt in linePoints)
                    {
                        res[pt] = res.GetValueOrDefault(pt) + 1;
                    }

                }


                return res;
            }


        }

        public Board board { get; set; } = new();

        public void ProcessInput(string input)
        {
            foreach (string line in Input.GetLines(input))
            {
                var parts = line.Split(" -> ");
                var from = parts[0].Split(',').Select(n => int.Parse(n)).ToArray();
                var to = parts[1].Split(',').Select(n => int.Parse(n)).ToArray();


                var fromPt = new Point(from[0], from[1]);
                var toPt = new Point(to[0], to[1]);

                board.AddLine(new Line(fromPt, toPt));
            }

        }

        public string Part1()
        {
            return board
                .OneDimensionalOverlaps()
                .Where(entry => entry.Value > 1)
                .Count()
                .ToString();
        }




        public string Part2()
        {
            return board
                .Overlaps()
                .Where(entry => entry.Value > 1)
                .Count()
                .ToString(); ;
        }

    }
}
