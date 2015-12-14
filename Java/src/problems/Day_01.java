package problems;

import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Problem 01
 */
public class Day_01 {

    private static ArrayList<String> lines;
    private static String line;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String>) FileHandler.getNotEmptyLines("../inputfiles/day01.txt");
        line = lines.get(0);
        System.out.println("[Day 01] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 01] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 01] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException {
        int floor = 0;
        int nrInstructions = line.length();

        for (int i = 0; i < nrInstructions; i++) {
            if (line.charAt(i) == '(') floor++;
            if (line.charAt(i) == ')') floor--;
        }

        return floor;
    }

    public static int problem_02() throws IOException {
        int floor = 0;
        int nrInstructions = line.length();

        for (int i = 0; i < nrInstructions; i++) {
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
