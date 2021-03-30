package day03;

import puzzle.Solver;

import java.util.ArrayList;
import java.util.List;

public class Day03Solver implements Solver {

    int width;
    int height;
    char[][] mountain;

    @Override
    public void processInput(List<String> lines) {
        height = lines.size();
        width = lines.get(0).length();

        mountain = new char[height][width];

        for (int h = 0; h < height; h++) {
            mountain[h] = lines.get(h).toCharArray();
        }

    }

    @Override
    public String part1() {
        long res = countTrees(new Point2D(3, 1));
        return Long.toString(res);
    }

    @Override
    public String part2() {
        long res = countTrees(
                new Point2D(1, 1),
                new Point2D(3, 1),
                new Point2D(5, 1),
                new Point2D(7, 1),
                new Point2D(1, 2)
        );

        return Long.toString(res);
    }

    public long countTrees(Point2D... slopes) {
        List<Long> counts = new ArrayList<>();

        for (Point2D slope : slopes) {
            long treeCount = 0;
            Point2D currentTile = new Point2D(0, 0);

            while (currentTile.y < height) {
                if (mountain[currentTile.y][(currentTile.x) % width] == '#') {
                    treeCount++;
                }
                currentTile.add(slope);
            }

            counts.add(treeCount);
        }

        return counts.stream().reduce((a, b) -> a * b).orElse(0L);
    }
}
