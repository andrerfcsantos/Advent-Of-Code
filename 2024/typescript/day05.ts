import { runProblem } from "./lib/runner.ts";
import { groupedLines } from "./lib/input.ts";

const DAY = 5;
const YEAR = 2024;

interface State {
  rules: Map<number, number[]>;
  updates: number[][];
}

export function parse(input: string): State {
  const [ruleLines, updateLines] = groupedLines(input);

  const rules = new Map<number, number[]>();
  ruleLines.map((line) => {
    const [from, to] = line.split("|").map((rule) => parseInt(rule, 10));
    if(!rules.has(from)) {
      rules.set(from, []);
    }
    rules.get(from)?.push(to);
  });

  const updates = updateLines.map((line) => line.split(",").map((num) => parseInt(num, 10)));

  return { rules, updates };
}

export function part1(parsed: State): string {
  const middleSum = parsed.updates
  .filter((update) => isValidUpdate(update, parsed.rules))
  .map(middlePageNumber)
  .reduce((acc, num) => acc + num, 0);
  return middleSum.toString(10);
}

function middlePageNumber(update: number[]): number {
  return update[Math.floor(update.length / 2)];
}

function makeValid(update: number[], rules: Map<number, number[]>): number[] {

  let { valid, idxA, idxB } = isValidUpdateAux(update, rules);

  while(!valid) {
    const tmp = update[idxA];
    update[idxA] = update[idxB];
    update[idxB] = tmp;
    const validResult = isValidUpdateAux(update, rules);
    valid = validResult.valid;
    idxA = validResult.idxA;
    idxB = validResult.idxB;
  }

  return update;
}

interface ValidResult {
  valid: boolean;
  idxA: number;
  idxB: number;
}

function isValidUpdateAux(update: number[], rules: Map<number, number[]>): ValidResult {
  const nToIndex = new Map<number, number>();

  for(let i = 0; i < update.length; i++) {
    nToIndex.set(update[i], i);
  }

  for(let i = 0; i < update.length; i++) {
    const item = update[i];
    for(const candidate of rules.get(item) || []) {
      if(nToIndex.has(candidate) && ((nToIndex.get(candidate) || 0) < i)) {
        return { valid: false, idxA: i, idxB: nToIndex.get(candidate)! };
      }
    }
  }

  return { valid: true, idxA: -1, idxB: -1 };
}

function isValidUpdate(update: number[], rules: Map<number, number[]>): boolean {
  return isValidUpdateAux(update, rules).valid;
}

export function part2(parsed: State): string {
  const middleSum = parsed.updates
  .filter((update) => !isValidUpdate(update, parsed.rules))
  .map((update) => makeValid(update, parsed.rules))
  .map(middlePageNumber)
  .reduce((acc, num) => acc + num, 0);

  return middleSum.toString(10);
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
