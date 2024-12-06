export class Point  {
    constructor(public x: number, public y: number) {}

    addVector(vector: Vector): Point {
        return new Point(this.x + vector.dx, this.y + vector.dy);
    }

    toString(): string {
        return `(${this.x}, ${this.y})`;
    }
}

export class Vector {
    constructor(public dx: number, public dy: number) {}

    addVector(vector: Vector): Vector {
        return new Vector(this.dx + vector.dx, this.dy + vector.dy);
    }
}