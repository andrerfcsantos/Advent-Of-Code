import { nonEmptyLines, getInput } from "./lib/input.ts";
import { runProblem } from "./lib/runner.ts";

const DAY = 1;
const YEAR = 2024;

interface State {
  lines: string[];
  left_list: number[];
  right_list: number[];
}

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  const left_list: number[] = [];
  const right_list: number[] = [];

  const numbers = lines.map((line) => line
  .replace(/\s+/g, " ")
  .split(" ")
  .map((x) => parseInt(x, 10)));

  numbers.forEach(([left, right]) => {
    left_list.push(left);
    right_list.push(right);
  })

  return { lines, left_list, right_list };
}

export function part1(parsed: State): string {
  const right_sorted = parsed.right_list.toSorted();
  const left_sorted = parsed.left_list.toSorted();
  let sum = 0;
  for(let i = 0; i < right_sorted.length; i++) {
    sum += Math.abs(left_sorted[i] - right_sorted[i])
  }
  return sum.toString(10);
}

export function part2(parsed: State): string {

  const right_counts = new Map<number, number>();
  parsed.right_list.forEach((n) => {
    right_counts.set(n, 1+ (right_counts.get(n) || 0) )
  })

  const similarity = parsed.left_list.reduce((prev, current) => {
    return prev + (current * (right_counts.get(current) || 0))
  },0)

  return similarity.toString(10);
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
