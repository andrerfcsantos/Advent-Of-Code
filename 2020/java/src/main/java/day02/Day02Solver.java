package day02;

import puzzle.Solver;

import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day02Solver implements Solver {

    private List<Password> passwords;

    public Day02Solver() {
        passwords = new ArrayList<>();
    }

    @Override
    public void processInput(List<String> lines) {
        Pattern passwordPattern = Pattern.compile("^(\\d+)-(\\d+) (\\w): (\\w+)$");

        for(String line: lines) {
            Matcher matcher = passwordPattern.matcher(line);
            if (matcher.matches()) {
                int lowest = Integer.parseInt(matcher.group(1));
                int highest = Integer.parseInt(matcher.group(2));
                char letter = matcher.group(3).charAt(0);
                String password = matcher.group(4);

                passwords.add(new Password(lowest, highest, letter, password));
            }
        }

    }

    @Override
    public String part1() {
        return Long.toString(passwords.stream().filter(p -> p.isValidPart1()).count());
    }

    @Override
    public String part2() {
        return Long.toString(passwords.stream().filter(p -> p.isValidPart2()).count());
    }
}
