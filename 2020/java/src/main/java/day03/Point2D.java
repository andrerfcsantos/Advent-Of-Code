package day03;

public class Point2D {
    public int x;
    public int y;

    public Point2D() {
        this(0, 0);
    }

    public Point2D(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public Point2D add(Point2D other) {
        this.x += other.x;
        this.y += other.y;
        return this;
    }

}
