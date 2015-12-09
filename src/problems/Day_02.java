package problems;

import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;

/**
 * Problem 02
 */
public class Day_02 {

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 2 Problem 1: " + problem_01());
        System.out.println("Solution Day 2 Problem 2: " + problem_02());
    }

    public static int problem_01() throws IOException {
        int result = 0, smallestSide, l, w, h;
        ArrayList<String> lines = (ArrayList<String>) FileHandler.getNotEmptyLines("Inputfiles/day02_1.txt");

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, "x\n\r");

            l = Integer.parseInt(strTok.nextToken());
            w = Integer.parseInt(strTok.nextToken());
            h = Integer.parseInt(strTok.nextToken());

            smallestSide = l * w;
            if (w * h < smallestSide) smallestSide = w * h;
            if (h * l < smallestSide) smallestSide = h * l;

            result = result + 2 * l * w + 2 * w * h + 2 * h * l + smallestSide;
        }
        return result;
    }

    public static int problem_02() throws IOException {
        int result = 0, l, w, h;
        int shortestSides[] = new int[2];

        ArrayList<String> lines = (ArrayList<String>) FileHandler.getNotEmptyLines("Inputfiles/day02_1.txt");

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, "x\n\r");

            l = Integer.parseInt(strTok.nextToken());
            w = Integer.parseInt(strTok.nextToken());
            h = Integer.parseInt(strTok.nextToken());

            shortestSides[0] = l;
            shortestSides[1] = w;
            if (h < l && l >= w) shortestSides[0] = h;
            if (h < w && w > l) shortestSides[1] = h;

            result += 2 * (shortestSides[0] + shortestSides[1]) + l * w * h;

        }
        return result;

    }


}
