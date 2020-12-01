export function computeFuelForMass(mass: number): number {
  return mass / 3 - 2;
}

export function computeFuelForFuelMass(fuel: number): number {
  let res = 0;
  for (;;) {
    fuel = computeFuelForMass(fuel);
    if (fuel <= 0) {
      break;
    }
    res += Math.floor(fuel);
  }
  return res;
}

export function part1Solver(inp: string): string {
  let sum = 0;
  const lines = inp.split(/\r?\n/);

  for (let line of lines) {
    line = line.trim();
    const mass = parseInt(line);
    if (!isNaN(mass)) {
      sum += Math.floor(computeFuelForMass(mass));
    } else {
      throw `Could not parse '${line}' as integer. Please check your input.`;
    }
  }

  return sum.toFixed(0);
}

export function part2Solver(inp: string): string {
  let sum = 0;
  const lines = inp.split(/\r?\n/);

  for (let line of lines) {
    line = line.trim();
    const mass = parseInt(line);
    if (!isNaN(mass)) {
      const massFuel = Math.floor(computeFuelForMass(mass));
      const fuelForFuelMass = Math.floor(computeFuelForFuelMass(massFuel));
      sum += massFuel + fuelForFuelMass;
    } else {
      throw `Could not parse '${line}' as integer. Please check your input.`;
    }
  }

  return sum.toFixed(0);
}

export default {
  part1Solver,
  part2Solver
};
