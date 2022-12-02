import { runProblem } from "./aoclib.ts";

const DAY = 1;
const YEAR = 2022;

type Calories = number;
type ElfCalories = Calories[];

interface State {
  elvesCalories: ElfCalories[];
  totalCaloriesByElf?: Calories[];
}

export function parse(lines: string[]): State {
  const elvesCalories: ElfCalories[] = [];
  let currentElfCalories: ElfCalories = [];

  for (const line of lines) {
    if (line === "") {
      elvesCalories.push(currentElfCalories);
      currentElfCalories = [];
      continue;
    }
    currentElfCalories.push(parseInt(line, 10));
  }

  return { elvesCalories };
}

function sumNumbers(ns: number[]): number {
  return ns.reduce((a, b) => a + b, 0);
}

export function part1(parsed: State): string {
  const totalCaloriesByElf = parsed.elvesCalories.map(sumNumbers);

  // save totalCaloriesByElf for part 2
  parsed.totalCaloriesByElf = totalCaloriesByElf;

  return Math.max(...totalCaloriesByElf).toString();
}

export function part2(parsed: State): string {
  const top3 = parsed.totalCaloriesByElf!.sort((a, b) => b - a).slice(0, 3);
  return sumNumbers(top3).toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
