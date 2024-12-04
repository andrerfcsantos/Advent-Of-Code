import { runProblem } from "./lib/runner.ts";
import { nonEmptyLines } from "./lib/input.ts";
import { ScreenGrid, ScreenPoint, ScreenVector } from "./lib/screen_grid.ts";

const DAY = 4;
const YEAR = 2024;

interface State {
  grid: ScreenGrid<string>;
}

const cardinalDirections: ScreenVector[] = [
  // left
  new ScreenVector(-1, 0),
  // right
  new ScreenVector(1, 0),
  // up
  new ScreenVector(0, -1),
  // down
  new ScreenVector(0, 1),
];

const diagonalDirections: ScreenVector[] = [
  // up-left
  new ScreenVector(-1, -1),
  // up-right
  new ScreenVector(1, -1),
  // down-left
  new ScreenVector(-1, 1),
  // down-right
  new ScreenVector(1, 1),
];

const directions: ScreenVector[] = [
  ...cardinalDirections,
  ...diagonalDirections,
];

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  const grid = new ScreenGrid(lines.map((line) => line.split("")));
  return { grid };
}

export function part1(parsed: State): string {
  let xmasCount = 0;
  const xmas = "XMAS";

  for (let y = 0; y < parsed.grid.height; y++) {
    for (let x = 0; x < parsed.grid.width; x++) {
      const start = new ScreenPoint(x, y);
      for (const direction of directions) {
        const contents = parsed.grid.contentsAlongLine(
          start,
          direction,
          xmas.length,
        );
        if (contents.join("") === xmas) {
          xmasCount++;
        }
      }
    }
  }

  return xmasCount.toString();
}

export function part2(parsed: State): string {
  let masCount = 0;
  const mas = "MAS";

  for (let y = 0; y < parsed.grid.height; y++) {
    for (let x = 0; x < parsed.grid.width; x++) {
      const topLeftStart = new ScreenPoint(x, y);
      const toCheck = [
        {
          start: topLeftStart,
          direction: new ScreenVector(1, 1),
        },
        {
          start: topLeftStart.addVector(new ScreenVector(2, 0)),
          direction: new ScreenVector(-1, 1),
        },
        {
          start: topLeftStart.addVector(new ScreenVector(0, 2)),
          direction: new ScreenVector(1, -1),
        },
        {
          start: topLeftStart.addVector(new ScreenVector(2, 2)),
          direction: new ScreenVector(-1, -1),
        },
      ];

      let localMasCount = 0;

      for (const { start, direction } of toCheck) {
        const contents = parsed.grid.contentsAlongLine(
          start,
          direction,
          mas.length,
        );
        if (contents.join("") === mas) {
          localMasCount++;
        }
      }

      if (localMasCount == 2) {
        masCount += 1
    }
  }
}

  return masCount.toString(10);
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
