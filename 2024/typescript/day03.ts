import { runProblem } from "./lib/runner.ts";
import { nonEmptyLines } from "./lib/input.ts";

const DAY = 3;
const YEAR = 2024;

interface State {
  operations: Operation[];
}

interface Operation {
  op: string;
  args: number[];
}

const OP_REGEX = /(?<op>do)\(\)|(?<op>don't)\(\)|(?<op>mul)\((?<a>\d{1,3}),(?<b>\d{1,3})\)/g

export function parse(input: string): State {
  const jumbledInstructions = nonEmptyLines(input).join("");
  const operations: Operation[] = [];

  for(const match of jumbledInstructions.matchAll(OP_REGEX)){
    const op = match?.groups?.["op"];

    switch(op) {
      case "mul": {
        const a = parseInt(match?.groups?.["a"] || "0", 10);
        const b = parseInt(match?.groups?.["b"] || "0", 10);
        operations.push({op, args: [a,b]});
        break;
      } 
      case "do":
      case "don't":
        operations.push({op, args: []});
        break
      default:
        operations.push({op: `unknown operation: ${op}`, args: []})

    }

  }
  return { operations };
}

export function part1(parsed: State): string {
  const result = parsed.operations.reduce((prev, current) => {
    if(current.op !== "mul") {
      return prev
    } 
    return prev + current.args[0]*current.args[1];
  }, 0)

  return result.toString(10);
}

export function part2(parsed: State): string {
  const result = parsed.operations.reduce((prev, current) => {
    switch(current.op) {
      case "mul": {
        if(prev.mulEnabled) {
          prev.sum += current.args[0]*current.args[1]
        }
        break;
      }
      case "do":
        prev.mulEnabled = true;
        break;
      case "don't":
        prev.mulEnabled = false;
        break;
      default:
        console.log(`unrecognized operation: ${current.op}`)
    }
    return prev
  }, {sum: 0, mulEnabled: true})

  return result.sum.toString(10);
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);