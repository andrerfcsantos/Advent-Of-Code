package com.asantosdev.utils;

import java.util.List;
import java.util.stream.Stream;

public class ReadUtils {

  public static List<String> nonEmptyLines(String input) {
    return input.lines().filter(line -> !line.isBlank()).toList();
  }

  public static List<List<String>> groupedLines(String input, String groupDelimiter) {
    String[] groups = input.split(groupDelimiter);

    return Stream.of(groups).map(group -> group.lines().toList()).toList();
  }
}
