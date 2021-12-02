using AOCLibrary;

namespace AOC2021.Solvers
{

    public enum Direction
    {
        UP,
        DOWN,
        FORWARD,
    }

    public class Command
    {
        internal Direction direction;
        internal int units = 0;
    }

    public class Position
    {
        internal int depth = 0;
        internal int horizontal = 0;
        internal int aim = 0;
    }

    public class Day02 : ISolver
    {
        List<Command> commands = new();

        public void ProcessInput(string input)
        {
            commands = new();

            foreach (string line in Input.GetLines(input))
            {
                string[] vs = line.Split(new char[] { ' ' });

                Direction dir;

                switch (vs[0]) {
                    case "up": dir = Direction.UP; break;
                    case "down": dir = Direction.DOWN; break;
                    case "forward": dir= Direction.FORWARD; break;
                    default: dir = Direction.UP; break;
                }

                int value = Convert.ToInt32(vs[1]);

                commands.Add(new Command
                {
                    direction = dir,
                    units = value
                });
            }
        }

        public string Part1()
        {
            Position pos = new Position
            {
                depth = 0,
                horizontal = 0,
            };

            foreach(Command c in commands)
            {
                switch (c.direction)
                {
                    case Direction.UP: pos.depth-=c.units; break;
                    case Direction.DOWN: pos.depth += c.units; break;
                    case Direction.FORWARD: pos.horizontal+=c.units; break;
                }
            }
            return (pos.depth*pos.horizontal).ToString();
        }

        public string Part2()
        {
            Position pos = new Position
            {
                depth = 0,
                horizontal = 0,
                aim = 0,
            };

            foreach (Command c in commands)
            {
                switch (c.direction)
                {
                    case Direction.UP: pos.aim -= c.units; break;
                    case Direction.DOWN: pos.aim += c.units; break;
                    case Direction.FORWARD:
                        pos.horizontal += c.units;
                        pos.depth += c.units*pos.aim;
                        break;
                }
            }
            return (pos.depth * pos.horizontal).ToString();
        }

    }
}
