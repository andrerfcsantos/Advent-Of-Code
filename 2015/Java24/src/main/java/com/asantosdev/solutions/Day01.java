package com.asantosdev.solutions;

import com.asantosdev.utils.ReadUtils;
import java.util.ArrayList;
import java.util.List;

public class Day01 implements Solver {

  private final List<Character> parenthesis = new ArrayList<>();

  public void processInput(String input) {
    List<String> lines = ReadUtils.nonEmptyLines(input);
    for (String line : lines) {
      for (char c : line.toCharArray()) {
        parenthesis.add(c);
      }
    }
  }

  public String part1() {
    int floor = 0;
    for (char c : parenthesis) {
      switch (c) {
        case '(' -> floor++;
        case ')' -> floor--;
        default -> throw new IllegalArgumentException("Unexpected character: " + c);
      }
    }
    return String.valueOf(floor);
  }

  public String part2() {
    int floor = 0;
    for (int i = 0; i < parenthesis.size(); i++) {
      char c = parenthesis.get(i);

      switch (c) {
        case '(' -> floor++;
        case ')' -> floor--;
        default -> throw new IllegalArgumentException("Unexpected character: " + c);
      }

      if (floor < 0) {
        return String.valueOf(i + 1);
      }
    }
    return "Not found";
  }
}
