import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";

const DAY = 0;
const YEAR = 2022;

interface State {
  lines: string[];
}

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  return { lines };
}

export function part1(parsed: State): string {
  return "implement part1";
}

export function part2(parsed: State): string {
  return "implement part2";
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
