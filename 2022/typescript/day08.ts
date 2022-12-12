import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";

const DAY = 8;
const YEAR = 2022;

interface State {
  forest: number[][];
}

export function parse(input: string): State {
  return {
    forest: nonEmptyLines(input).map((l) =>
      l.split("").map((c) => parseInt(c, 10))
    ),
  };
}

export function part1(parsed: State): string {
  const nRows = parsed.forest.length;
  const nCols = parsed.forest[0].length;
  let visible = nRows * 2 + nCols * 2 - 4;

  for (let rowIdx = 1; rowIdx < nRows - 1; rowIdx++) {
    const row = parsed.forest[rowIdx];

    for (let colIdx = 1; colIdx < row.length - 1; colIdx++) {
      const col = parsed.forest.map((r) => r[colIdx]);

      const hiddenLeft = row.slice(0, colIdx).some((h) => h >= row[colIdx]);
      const hiddenRight = row.slice(colIdx + 1).some((h) => h >= row[colIdx]);
      const hiddenAbove = col.slice(0, rowIdx).some((h) => h >= row[colIdx]);
      const hiddenBelow = col.slice(rowIdx + 1).some((h) => h >= row[colIdx]);

      if (!hiddenLeft || !hiddenRight || !hiddenAbove || !hiddenBelow) {
        visible += 1;
      }
    }
  }

  return visible.toString();
}

export function part2(parsed: State): string {
  const nRows = parsed.forest.length;
  const nCols = parsed.forest[0].length;
  let maxScenicScore = -1;

  for (let rowIdx = 1; rowIdx < nRows - 1; rowIdx++) {
    const row = parsed.forest[rowIdx];

    for (let colIdx = 1; colIdx < row.length - 1; colIdx++) {
      const col = parsed.forest.map((r) => r[colIdx]);

      let scoreLeft = row
        .slice(0, colIdx)
        .reverse()
        .findIndex((h) => h >= row[colIdx]);

      if (scoreLeft == -1) {
        scoreLeft = colIdx;
      } else {
        scoreLeft += 1;
      }

      let scoreRight = row.slice(colIdx + 1).findIndex((h) => h >= row[colIdx]);
      if (scoreRight == -1) {
        scoreRight = nCols - colIdx - 1;
      } else {
        scoreRight += 1;
      }

      let scoreAbove = col
        .slice(0, rowIdx)
        .reverse()
        .findIndex((h) => h >= row[colIdx]);

      if (scoreAbove == -1) {
        scoreAbove = rowIdx;
      } else {
        scoreAbove += 1;
      }

      let scoreBelow = col.slice(rowIdx + 1).findIndex((h) => h >= row[colIdx]);
      if (scoreBelow == -1) {
        scoreBelow = nRows - rowIdx - 1;
      } else {
        scoreBelow += 1;
      }

      const score = scoreLeft * scoreRight * scoreAbove * scoreBelow;
      if (score > maxScenicScore) {
        maxScenicScore = score;
      }
    }
  }

  return maxScenicScore.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
