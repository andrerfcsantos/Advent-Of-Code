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

export class Vector {
  dx: number;
  dy: number;

  constructor(dx: number, dy: number) {
    this.dx = dx;
    this.dy = dy;
  }

  static fromPoints(source: Point, dest: Point): Vector {
    return new Vector(dest.x - source.x, dest.y - source.y);
  }

  toDirectionVector(): Vector {
    return new Vector(
      this.dx != 0 ? this.dx / Math.abs(this.dx) : 0,
      this.dy != 0 ? this.dy / Math.abs(this.dy) : 0
    );
  }

  toString(): string {
    return `(${this.dx},${this.dy})`;
  }
}

export class Point {
  x: number;
  y: number;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  add(v: Vector): Point {
    return new Point(this.x + v.dx, this.y + v.dy);
  }

  move(v: Vector): Point {
    this.x += v.dx;
    this.y += v.dy;
    return this;
  }

  isTouching(other: Point): boolean {
    const dx = Math.abs(this.x - other.x);
    const dy = Math.abs(this.y - other.y);
    return Math.abs(dx) <= 1 && Math.abs(dy) <= 1;
  }

  cardinalNeighbors(): Point[] {
    return [
      new Point(this.x, this.y + 1),
      new Point(this.x, this.y - 1),
      new Point(this.x + 1, this.y),
      new Point(this.x - 1, this.y),
    ];
  }

  taxiCabDistance(other: Point): number {
    return Math.abs(this.x - other.x) + Math.abs(this.y - other.y);
  }

  equals(other: Point): boolean {
    return this.x === other.x && this.y === other.y;
  }

  toString(): string {
    return `(${this.x},${this.y})`;
  }
}
