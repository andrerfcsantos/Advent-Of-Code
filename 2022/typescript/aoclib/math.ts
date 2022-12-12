export function lcm(...numbers: number[]): number {
  let lcm = 1;

  for (const num of numbers) {
    lcm = (lcm * num) / gcd(lcm, num);
  }

  return lcm;
}

export function gcd(a: number, b: number): number {
  if (a < b) [a, b] = [b, a];

  // Euclid's algorithm
  while (b > 0) {
    const temp = b;
    b = a % b;
    a = temp;
  }

  return a;
}
