import { info } from "https://deno.land/std@0.167.0/log/mod.ts";
import { runProblem, nonEmptyLines, ocrScreenToText } from "./aoclib.ts";

const DAY = 10;
const YEAR = 2022;

interface Instruction {
  opcode: "noop" | "addx";
  operand: number;
}

interface State {
  instructions: Instruction[];
}

export function parse(input: string): State {
  return {
    instructions: nonEmptyLines(input).map((line) => {
      const [opCode, operand] = line.split(" ");
      return {
        opcode: opCode as "noop" | "addx",
        operand: opCode === "addx" ? parseInt(operand, 10) : 0,
      };
    }),
  };
}

export function part1(parsed: State): string {
  let signalStrength = 0;

  function doCycle(cycle: number, xRegister: number) {
    if ((cycle - 20) % 40 === 0) {
      signalStrength += xRegister * cycle;
    }
  }

  runInstructions(parsed.instructions, doCycle);

  return signalStrength.toString();
}

function pixel(cathodPos: number, pixelPos: number) {
  return pixelHitByCathode(cathodPos, pixelPos) ? "#" : ".";
}

function pixelHitByCathode(cathodPos: number, pixelPos: number) {
  return (
    pixelPos - 1 === cathodPos ||
    pixelPos + 1 === cathodPos ||
    pixelPos === cathodPos
  );
}

export function part2(parsed: State): string {
  const screen: string[] = [];
  let currentBuffer = "";

  function cycleFunc(cycle: number, xRegister: number) {
    currentBuffer += pixel(currentBuffer.length, xRegister);
    if (cycle % 40 === 0) {
      screen.push(currentBuffer);
      currentBuffer = "";
    }
  }

  runInstructions(parsed.instructions, cycleFunc);

  const text = ocrScreenToText(screen);
  info("\n" + screen.join("\n"));

  return text;
}

function runInstructions(
  instructions: Instruction[],
  cyclceFunc: (cycle: number, xRegister: number) => void
) {
  let xRegister = 1;
  let currentCycle = 0;

  function doCycle() {
    currentCycle++;
    cyclceFunc(currentCycle, xRegister);
  }

  for (const { opcode, operand } of instructions) {
    switch (opcode) {
      case "noop":
        doCycle();
        break;
      case "addx":
        doCycle();
        doCycle();
        xRegister += operand;
        break;
    }
  }
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
