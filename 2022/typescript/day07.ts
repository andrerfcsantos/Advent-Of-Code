import { runProblem, nonEmptyLines } from "./aoclib.ts";

const DAY = 7;
const YEAR = 2022;

interface State {
  part1: number;
  part2: number;
}

type Dir = [name: string, size: number];

const FILESYSTEM_TOTAL_SPACE = 70000000;
const TARGET_UNUSED_SPACE = 30000000;

export function parse(input: string): State {
  const lines = nonEmptyLines(input).map((l) => l.split(" "));

  const stack = new Array<Dir>();
  stack.push(["/", 0]);

  let total_used = 0;
  const dirs = new Array<Dir>();

  let sumUnder100k = 0;
  for (const line of lines) {
    const fst = line[0];

    switch (fst) {
      // deno-lint-ignore no-case-declarations
      case "$":
        const [cmd, arg] = line.slice(1);
        if (cmd == "cd" && arg == "/") {
          // ignore root
          continue;
        } else if (cmd == "cd" && arg == "..") {
          // we know we are done with this directory
          const [name, size] = stack.pop()!;
          dirs.push([name, size]);

          // update parent size with the size of
          // the directory we just visited
          stack[stack.length - 1][1] += size;
          if (size <= 100000) {
            sumUnder100k += size;
          }
          continue;
        } else if (cmd == "cd") {
          stack.push([arg, 0]);
          continue;
        }
        break;
      case "dir":
        continue;
      // deno-lint-ignore no-case-declarations
      default:
        // line representing a file in the current directory
        const size = parseInt(line[0], 10);
        stack[stack.length - 1][1] += size;
        total_used += size;
    }
  }

  // deal with items still in the stack
  let rest = 0;
  for (const [name, size] of stack.reverse()) {
    dirs.push([name, size + rest]);
    rest += size;
  }

  const actual_unused = FILESYSTEM_TOTAL_SPACE - total_used;
  const space_to_free = TARGET_UNUSED_SPACE - actual_unused;

  const minDir = dirs
    .filter(([_, size]) => size >= space_to_free)
    .sort((a, b) => a[1] - b[1]);

  return { part1: sumUnder100k, part2: minDir[0][1] };
}

export function part1(parsed: State): string {
  return parsed.part1.toString();
}

export function part2(parsed: State): string {
  return parsed.part2.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
