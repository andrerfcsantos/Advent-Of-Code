package day01;

import puzzle.Solver;

import java.util.ArrayList;
import java.util.List;

public class Day01Solver implements Solver {

    private int sumTarget;
    private List<Integer> intList;

    public Day01Solver(int target) {
        sumTarget = target;
        intList = new ArrayList<>();
    }

    @Override
    public void processInput(List<String> lines) {
        for (String line : lines) {
            this.intList.add(Integer.parseInt(line));
        }
    }

    @Override
    public String part1() {

        int n = intList.size();

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int n1 = intList.get(i);
                int n2 = intList.get(j);
                if (n1 + n2 == sumTarget) {
                    return Integer.toString(n1*n2);
                }
            }
        }

        return "Could not find two numbers that add to " + sumTarget;
    }

    @Override
    public String part2() {
        int n = intList.size();

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                for (int k = j + 1; k < n; k++) {
                    int n1 = intList.get(i);
                    int n2 = intList.get(j);
                    int n3 = intList.get(k);
                    if (n1 + n2 +n3 == sumTarget) {
                        return Integer.toString(n1 * n2 * n3);
                    }
                }
            }
        }

        return "Could not find two numbers that add to " + sumTarget;
    }
}
