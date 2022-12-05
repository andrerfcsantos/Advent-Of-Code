import { runProblem, groupedLines } from "./aoclib.ts";

const DAY = 5;
const YEAR = 2022;

interface Instructions {
  amount: number;
  from: number;
  to: number;
}

interface State {
  stacks: string[][];
  instructions: Instructions[];
}

export function parse(input: string): State {
  const lines = groupedLines(input);
  const [initialStack, instructions] = lines;
  const colSize = initialStack[0].length;
  const nStacks = parseInt(initialStack.at(-1)![colSize - 2], 10);

  const stacks: string[][] = new Array(nStacks);
  for (let i = 0; i < nStacks; i++) {
    stacks[i] = [];
  }

  for (let colIdx = 1; colIdx < initialStack[0].length - 1; colIdx += 4) {
    for (let row = initialStack.length - 2; row >= 0; row--) {
      const val = initialStack[row][colIdx];
      if (val === " ") {
        continue;
      }

      const stackNo = Math.floor((colIdx - 1) / 4);
      stacks[stackNo].push(val);
    }
  }

  const parsedInstructions = instructions.map((line) => {
    const [_, amount, from, to] = line.match(/move (\d+) from (\d+) to (\d+)/)!;
    return {
      amount: parseInt(amount, 10),
      from: parseInt(from, 10),
      to: parseInt(to, 10),
    };
  });

  return { stacks, instructions: parsedInstructions };
}

export function part1(parsed: State): string {
  // copy parsed stacks contents to a new matrix
  // so we are not mutating state
  const stacks = parsed.stacks.map((stack) => [...stack]);

  for (const { amount, from, to } of parsed.instructions) {
    const fromStack = stacks[from - 1];
    const toStack = stacks[to - 1];
    const items = fromStack.splice(fromStack.length - amount, amount).reverse();
    toStack.push(...items);
  }

  // get item on top of each stack
  const topItems = stacks.map((stack) => stack[stack.length - 1]);

  return topItems.join("");
}

export function part2(parsed: State): string {
  // copy parsed stacks contents to a new matrix
  // so we are not mutating state
  const stacks = parsed.stacks.map((stack) => [...stack]);

  for (const { amount, from, to } of parsed.instructions) {
    const fromStack = stacks[from - 1];
    const toStack = stacks[to - 1];
    const items = fromStack.splice(fromStack.length - amount, amount);
    toStack.push(...items);
  }

  // get item on top of each stack
  const topItems = stacks.map((stack) => stack[stack.length - 1]);

  return topItems.join("");
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
