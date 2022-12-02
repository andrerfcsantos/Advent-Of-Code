import { runProblem } from "./aoclib.ts";

const DAY = 0;
const YEAR = 2022;

interface State {}

export function parse(lines: string[]): State {
  return {};
}

export function part1(parsed: State): string {
  return "implement part1";
}

export function part2(parsed: State): string {
  return "implement part2";
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
