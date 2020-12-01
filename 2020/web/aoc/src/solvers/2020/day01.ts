// processed input
const nums: number[] = [];

export function processInput(inp: string) {
  const lines = inp.split(/\r?\n/);

  for (let line of lines) {
    line = line.trim();
    if (line == "") {
      continue;
    }
    const n = parseInt(line);
    if (!isNaN(n)) {
      nums.push(n);
    } else {
      throw `Could not parse '${line}' as integer. Please check your input.`;
    }
  }
}

export function part1Solver(inp: string): string {
  const size: number = nums.length;

  for (let i = 0; i < size; i++) {
    for (let j = i + 1; j < size - 1; j++) {
      if (nums[i] + nums[j] == 2020) {
        return (nums[i] * nums[j]).toString();
      }
    }
  }

  throw "Could not find any 2 numbers that have sum 2020";
}

export function part2Solver(inp: string): string {
  const size: number = nums.length;

  for (let i = 0; i < size; i++) {
    for (let j = i + 1; j < size - 1; j++) {
      for (let z = j + 1; z < size - 2; z++) {
        if (nums[i] + nums[j] + nums[z] == 2020) {
          return (nums[i] * nums[j] * nums[z]).toString();
        }
      }
    }
  }

  throw "Could not find any 3 numbers that have sum 2020";
}

export default {
  processInput,
  part1Solver,
  part2Solver
};
