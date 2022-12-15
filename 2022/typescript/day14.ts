import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";
import { Point, Vector } from "./aoclib/math.ts";

const DAY = 14;
const YEAR = 2022;

interface BoundingBox {
  minX: number;
  maxX: number;
  minY: number;
  maxY: number;
}

interface State {
  rocks: Map<string, Point>;
  boundingBox: BoundingBox;
}

export function parse(input: string): State {
  const points = nonEmptyLines(input).map((line) => {
    return line.split(" -> ").map((s) => {
      const [x, y] = s.split(",");
      return new Point(parseInt(x, 10), parseInt(y, 10));
    });
  });

  const rocks = new Map<string, Point>();

  for (const pointLine of points) {
    for (let i = 1; i < pointLine.length; i++) {
      const minX = Math.min(pointLine[i].x, pointLine[i - 1].x);
      const maxX = Math.max(pointLine[i].x, pointLine[i - 1].x);

      const minY = Math.min(pointLine[i].y, pointLine[i - 1].y);
      const maxY = Math.max(pointLine[i].y, pointLine[i - 1].y);

      for (let x = minX; x <= maxX; x++) {
        for (let y = minY; y <= maxY; y++) {
          const point = new Point(x, y);
          rocks.set(point.toString(), point);
        }
      }
    }
  }

  return {
    rocks,
    boundingBox: {
      minX: Math.min(...Array.from(rocks.values()).map((p) => p.x)),
      maxX: Math.max(...Array.from(rocks.values()).map((p) => p.x)),
      minY: Math.min(...Array.from(rocks.values()).map((p) => p.y)),
      maxY: Math.max(...Array.from(rocks.values()).map((p) => p.y)),
    },
  };
}

export function part1(parsed: State): string {
  const stillSand = new Map<string, Point>();
  const [DOWN, DOWN_LEFT, DOWN_RIGHT] = [
    new Vector(0, 1),
    new Vector(-1, 1),
    new Vector(1, 1),
  ];

  let simulationFinished = false;

  while (!simulationFinished) {
    const currentPosition = new Point(500, 0);
    let standingStill = false;

    while (!standingStill) {
      if (currentPosition.y > parsed.boundingBox.maxY) {
        simulationFinished = true;
        break;
      }

      const d = currentPosition.add(DOWN);

      const isDownStillSand = stillSand.has(d.toString());
      const isDownRock = parsed.rocks.has(d.toString());

      if (!isDownStillSand && !isDownRock) {
        currentPosition.move(DOWN);
        continue;
      }

      const dl = currentPosition.add(DOWN_LEFT);

      const isDownLeftStillSand = stillSand.has(dl.toString());
      const isDownLeftRock = parsed.rocks.has(dl.toString());

      if (!isDownLeftStillSand && !isDownLeftRock) {
        currentPosition.move(DOWN_LEFT);
        continue;
      }

      const dr = currentPosition.add(DOWN_RIGHT);

      const isDownRightStillSand = stillSand.has(dr.toString());
      const isDownRightRock = parsed.rocks.has(dr.toString());

      if (!isDownRightStillSand && !isDownRightRock) {
        currentPosition.move(DOWN_RIGHT);
        continue;
      }

      standingStill = true;
      stillSand.set(currentPosition.toString(), currentPosition);
    }
  }

  return stillSand.size.toString();
}

export function part2(parsed: State): string {
  const stillSand = new Map<string, Point>();
  const [DOWN, DOWN_LEFT, DOWN_RIGHT] = [
    new Vector(0, 1),
    new Vector(-1, 1),
    new Vector(1, 1),
  ];

  let simulationFinished = false;
  const starting_point = new Point(500, 0);
  while (!simulationFinished) {
    const currentPosition = new Point(500, 0);
    let standingStill = false;

    while (!standingStill) {
      if (stillSand.has(starting_point.toString())) {
        simulationFinished = true;
        break;
      }

      const d = currentPosition.add(DOWN);

      const isDownStillSand = stillSand.has(d.toString());
      const isDownRock =
        parsed.rocks.has(d.toString()) || d.y >= parsed.boundingBox.maxY + 2;

      if (!isDownStillSand && !isDownRock) {
        currentPosition.move(DOWN);
        continue;
      }

      const dl = currentPosition.add(DOWN_LEFT);

      const isDownLeftStillSand = stillSand.has(dl.toString());
      const isDownLeftRock =
        parsed.rocks.has(dl.toString()) || dl.y >= parsed.boundingBox.maxY + 2;

      if (!isDownLeftStillSand && !isDownLeftRock) {
        currentPosition.move(DOWN_LEFT);
        continue;
      }

      const dr = currentPosition.add(DOWN_RIGHT);

      const isDownRightStillSand = stillSand.has(dr.toString());
      const isDownRightRock =
        parsed.rocks.has(dr.toString()) || dr.y >= parsed.boundingBox.maxY + 2;

      if (!isDownRightStillSand && !isDownRightRock) {
        currentPosition.move(DOWN_RIGHT);
        continue;
      }

      standingStill = true;
      stillSand.set(currentPosition.toString(), currentPosition);
    }
  }

  return stillSand.size.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
