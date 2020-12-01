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
    for (let j = i + 1; j < size; j++) {
      if (i != j && nums[i] + nums[j] == 2020) {
        return (nums[i] * nums[j]).toString();
      }
    }
  }

  throw "Could not find any 2 numbers that have sum 2020";
}

export function part2Solver(inp: string): string {
  const size: number = nums.length;

  for (let i = 0; i < size; i++) {
    for (let j = i + 1; j < size; j++) {
      for (let z = j + 1; z < size; z++) {
        if (i != j && j != z && z != i && nums[i] + nums[j] + nums[z] == 2020) {
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
