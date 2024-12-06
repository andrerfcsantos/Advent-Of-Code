import { runProblem } from "./lib/runner.ts";
import { nonEmptyLines } from "./lib/input.ts";
import { Point, Vector } from "./lib/point.ts";
import { CartesianDirections } from "./lib/directions.ts";

const DAY = 6;
const YEAR = 2024;

interface Bounds {
  maxX: number;
  maxY: number;
}

interface State {
  obstacles: Set<string>;
  gridBounds: Bounds;
  startingPosition: Point;
  direction: string;
}

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  const obstacles = new Set<string>();
  let startingPosition: Point = new Point(0, 0);

  for(let y = 0; y < lines.length; y++) {
    const line = lines[y];
    for(let x = 0; x < line.length; x++) {
      if(line[x] === "#") {
        obstacles.add(new Point(x, y).toString());
      } else if(line[x] === "^") {
        startingPosition = new Point(x, y);
      }
    }
  }

  return { 
    obstacles,
    startingPosition,
    direction: "UP",
    gridBounds: { maxX: lines[0].length, maxY: lines.length }
   };
}

function InBounds(point: Point, bounds: Bounds): boolean {
  return point.x >= 0 && point.x < bounds.maxX && point.y >= 0 && point.y < bounds.maxY;
}

export function part1(parsed: State): string {

  let currentPosition = parsed.startingPosition;
  let currentDirection: keyof typeof CartesianDirections = parsed.direction as keyof typeof CartesianDirections;

  const visited = new Set<string>();

  while(InBounds(currentPosition, parsed.gridBounds)) {
    visited.add(currentPosition.toString());
    const inFront = currentPosition.addVector(CartesianDirections[currentDirection].vector);
    if(parsed.obstacles.has(inFront.toString())) {
      const directionRight: keyof typeof CartesianDirections = CartesianDirections[currentDirection].right as keyof typeof CartesianDirections;
      const newPosition = currentPosition.addVector(CartesianDirections[directionRight].vector);
      currentPosition = newPosition;
      currentDirection = directionRight;
    } else {
      currentPosition = inFront;
    }

  }

  return visited.size.toString(10);
}

export function part2(parsed: State): string {
  return "implement part2";
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
