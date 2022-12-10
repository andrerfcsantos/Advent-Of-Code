import { runProblem, nonEmptyLines } from "./aoclib.ts";

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
  let xRegister = 1;
  let currentCycle = 0;

  for (const { opcode, operand } of parsed.instructions) {
    currentCycle++;
    if ((currentCycle - 20) % 40 === 0) {
      signalStrength += xRegister * currentCycle;
    }

    switch (opcode) {
      case "noop":
        break;
      case "addx":
        currentCycle++;
        if ((currentCycle - 20) % 40 === 0) {
          signalStrength += xRegister * currentCycle;
        }
        xRegister += operand;
        break;
    }
  }

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
  let xRegister = 1;
  let screen: string[] = [];
  let currentBuffer = "";
  let currentCycle = 0;

  for (const { opcode, operand } of parsed.instructions) {
    currentCycle++;
    currentBuffer += pixel(currentBuffer.length, xRegister);
    if (currentCycle % 40 === 0) {
      screen.push(currentBuffer);
      currentBuffer = "";
    }

    switch (opcode) {
      case "noop":
        break;
      case "addx":
        currentCycle++;
        currentBuffer += pixel(currentBuffer.length, xRegister);
        if (currentCycle % 40 === 0) {
          screen.push(currentBuffer);
          currentBuffer = "";
        }
        xRegister += operand;
        break;
    }
  }

  return "\n" + screen.join("\n");
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
