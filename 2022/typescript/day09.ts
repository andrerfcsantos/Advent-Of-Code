import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";
import { Vector, Point } from "./aoclib/math.ts";

const DAY = 9;
const YEAR = 2022;

interface State {
  instructions: Instruction[];
}

enum Direction {
  Up = "Up",
  Down = "Down",
  Left = "Left",
  Right = "Right",
}

const directionMap: Record<string, Direction> = {
  U: Direction.Up,
  D: Direction.Down,
  L: Direction.Left,
  R: Direction.Right,
};

interface Instruction {
  direction: Direction;
  distance: number;
}

const directionVectors: Record<Direction, Vector> = {
  [Direction.Up]: new Vector(0, 1),
  [Direction.Down]: new Vector(0, -1),
  [Direction.Left]: new Vector(-1, 0),
  [Direction.Right]: new Vector(1, 0),
};

function runSimulation(instructions: Instruction[], numberOfKnots: number) {
  const visited = new Set<string>();

  const knots = new Array<Point>(numberOfKnots);
  for (let i = 0; i < numberOfKnots; i++) {
    knots[i] = new Point(0, 0);
  }
  const [head, tail] = [knots.at(0)!, knots.at(-1)!];

  visited.add(tail.toString());

  for (const { direction, distance } of instructions) {
    const directionVector = directionVectors[direction];

    for (let i = 0; i < distance; i++) {
      head.move(directionVector);

      for (let j = 0; j < numberOfKnots - 1; j++) {
        const [lead, follow] = [knots.at(j)!, knots.at(j + 1)!];

        if (!lead.isTouching(follow)) {
          const deltaVector = Vector.fromPoints(
            follow,
            lead
          ).toDirectionVector();
          follow.move(deltaVector);
        }
      }

      visited.add(tail.toString());
    }
  }
  return visited;
}

export function parse(input: string): State {
  return {
    instructions: nonEmptyLines(input).map((l) => {
      const [direction, distance] = l.split(" ");
      return {
        direction: directionMap[direction],
        distance: parseInt(distance, 10),
      };
    }),
  };
}

export function part1(parsed: State): string {
  return runSimulation(parsed.instructions, 2).size.toString();
}

export function part2(parsed: State): string {
  return runSimulation(parsed.instructions, 10).size.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
