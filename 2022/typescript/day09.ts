import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";

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

class Vector {
  dx: number;
  dy: number;

  constructor(dx: number, dy: number) {
    this.dx = dx;
    this.dy = dy;
  }

  static fromPoints(source: Point, dest: Point): Vector {
    return new Vector(dest.x - source.x, dest.y - source.y);
  }

  toDirectionVector(): Vector {
    return new Vector(
      this.dx != 0 ? this.dx / Math.abs(this.dx) : 0,
      this.dy != 0 ? this.dy / Math.abs(this.dy) : 0
    );
  }

  toString(): string {
    return `(${this.dx},${this.dy})`;
  }
}

class Point {
  x: number;
  y: number;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  move(v: Vector): Point {
    this.x += v.dx;
    this.y += v.dy;
    return this;
  }

  isTouching(other: Point): boolean {
    const dx = Math.abs(this.x - other.x);
    const dy = Math.abs(this.y - other.y);
    return Math.abs(dx) <= 1 && Math.abs(dy) <= 1;
  }

  toString(): string {
    return `(${this.x},${this.y})`;
  }
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
