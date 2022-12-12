import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";

const DAY = 6;
const YEAR = 2022;

interface State {
  signal: string;
}

export function parse(input: string): State {
  return { signal: nonEmptyLines(input)[0] };
}

export function part1(parsed: State): string {
  return seekMarker(parsed.signal, 4).toString();
}

export function part2(parsed: State): string {
  return seekMarker(parsed.signal, 14).toString();
}

export function seekMarker(signal: string, markerSize: number): number {
  for (let i = markerSize; i < signal.length; i += 1) {
    const lastX = signal.slice(i - markerSize, i);
    if (lastX.length === new Set(lastX).size) {
      return i;
    }
  }
  return NaN;
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
