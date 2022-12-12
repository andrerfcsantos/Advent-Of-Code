import { runProblem } from "./aoclib/runner.ts";
import { sumNumbers } from "./aoclib/arrays.ts";
import { groupedLines } from "./aoclib/input.ts";

const DAY = 1;
const YEAR = 2022;

type Calories = number;
type ElfCalories = Calories[];

interface State {
  elvesCalories: ElfCalories[];
}

export function parse(input: string): State {
  return {
    elvesCalories: groupedLines(input).map((g) => {
      return g.map((l) => parseInt(l, 10));
    }),
  };
}

export function part1(parsed: State): string {
  const totalCaloriesByElf = parsed.elvesCalories.map(sumNumbers);

  return Math.max(...totalCaloriesByElf).toString();
}

export function part2(parsed: State): string {
  const totalCaloriesByElf = parsed.elvesCalories.map(sumNumbers);
  const top3 = totalCaloriesByElf.sort((a, b) => b - a).slice(0, 3);

  return sumNumbers(top3).toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
