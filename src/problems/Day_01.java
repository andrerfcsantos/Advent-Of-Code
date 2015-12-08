package problems;

import utils.InputHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Problem 01
 */
public class Day_01 {

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 1 Problem 1: " + problem_01());
        System.out.println("Solution Day 1 Problem 2: " + problem_02());
    }

    public static int problem_01() throws IOException {
        int floor = 0;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getNotEmptyLines("Inputfiles/day01_1.txt");
        String line = lines.get(0);

        for (int i = 0; i < line.length(); i++) {
            if (line.charAt(i) == '(') floor++;
            if (line.charAt(i) == ')') floor--;
        }

        return floor;
    }

    public static int problem_02() throws IOException {
        int floor = 0;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getNotEmptyLines("Inputfiles/day01_1.txt");
        String line = lines.get(0);

        for (int i = 0; i < line.length(); i++) {

            if (line.charAt(i) == '(') floor++;
            if (line.charAt(i) == ')') floor--;

            if (floor == -1) {
                floor = i + 1;
                break;
            }

        }

        return floor;
    }

}
