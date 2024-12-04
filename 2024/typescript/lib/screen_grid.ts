export class ScreenPoint  {
    constructor(public x: number, public y: number) {}

    addVector(vector: ScreenVector): ScreenPoint {
        return new ScreenPoint(this.x + vector.dx, this.y + vector.dy);
    }
}

export class ScreenVector {
    constructor(public dx: number, public dy: number) {}

    addVector(vector: ScreenVector): ScreenVector {
        return new ScreenVector(this.dx + vector.dx, this.dy + vector.dy);
    }
}

export class ScreenGrid<T> {
    private _grid: T[][];
    private _width: number;
    private _height: number;

    constructor(grid: T[][]) {
        this._grid = grid;
        this._width = grid[0]?.length || 0;
        this._height = grid.length;
    }

    get width(): number {
        return this._width;
    }

    get height(): number {
        return this._height;
    }

    contentsAlongLine(start: ScreenPoint, direction: ScreenVector, times: number): T[] {
        const contents: T[] = [];
        let current = start;
        for (let i = 0; i < times; i++) {
            if (!this.inBounds(current)) {
                break;
            }
            contents.push(this.get(current));
            current = current.addVector(direction);
        }
        return contents;
    }

    get(point: ScreenPoint): T {
        return this._grid[point.y][point.x];
    }

    inBounds(point: ScreenPoint): boolean {
        return point.x >= 0 && point.x < this._width && point.y >= 0 && point.y < this._height;
    }


  }