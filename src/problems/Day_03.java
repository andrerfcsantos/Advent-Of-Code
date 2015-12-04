package problems;

import utils.InputHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashSet;

/**
 * Problem 03
 */
public class Day_03 {

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 2 Problem 1: " + problem_01());
        System.out.println("Solution Day 2 Problem 2: " + problem_02());
    }

    public static int problem_01() throws IOException {
        int result = 0;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getLines("Inputfiles/day03_1.txt");
        String line = lines.get(0);
        Coordinates currentPos = new Coordinates();
        HashSet<Coordinates> placesVisited = new HashSet<>(10000);
        char c;

        for (int i = 0; i < line.length(); i++) {

            if (!placesVisited.contains(currentPos)) {
                placesVisited.add(currentPos.clone());
                result++;
            }

            c = line.charAt(i);
            switch (c) {
                case '^':
                    currentPos.goUp();
                    break;
                case 'v':
                    currentPos.goDown();
                    break;
                case '<':
                    currentPos.goLeft();
                    break;
                case '>':
                    currentPos.goRight();
                    break;
                default:
            }
        }

        return result;
    }

    public static int problem_02() throws IOException {
        int result = 0;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getLines("Inputfiles/day03_1.txt");
        String line = lines.get(0);
        Coordinates posAux;
        Coordinates santaPos = new Coordinates();
        Coordinates robotSantaPos = new Coordinates();
        HashSet<Coordinates> placesVisited = new HashSet<>(10000);
        char c;

        for (int i = 0; i < line.length(); i++) {

            if (i % 2 == 0) posAux = santaPos;
            else posAux = robotSantaPos;

            if (!placesVisited.contains(posAux)) {
                placesVisited.add(posAux.clone());
                result++;
            }

            c = line.charAt(i);
            switch (c) {
                case '^':
                    posAux.goUp();
                    break;
                case 'v':
                    posAux.goDown();
                    break;
                case '<':
                    posAux.goLeft();
                    break;
                case '>':
                    posAux.goRight();
                    break;
                default:
            }
        }

        return result;
    }

    static class Coordinates {
        private int x, y;

        public Coordinates() {
            this(0, 0);
        }

        public Coordinates(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public Coordinates(Coordinates c) {
            this.x = c.x;
            this.y = c.y;
        }

        public int getX() {
            return x;
        }

        public int getY() {
            return y;
        }

        public void goUp() {
            this.y++;
        }

        public void goDown() {
            this.y--;
        }

        public void goRight() {
            this.x++;
        }

        public void goLeft() {
            this.x--;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Coordinates that = (Coordinates) o;
            return x != that.x && y == that.y;
        }

        @Override
        public int hashCode() {
            int result = x;
            result = 31 * result + y;
            return result;
        }

        public Coordinates clone() {
            return new Coordinates(this);
        }

        @Override
        public String toString() {
            return "(" + x + "," + y + ")";
        }
    }
}
