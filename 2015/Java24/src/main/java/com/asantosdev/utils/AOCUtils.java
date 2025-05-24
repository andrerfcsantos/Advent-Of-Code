package com.asantosdev.utils;

public class AOCUtils {

  public static String getStringInputForDay(int day) {
    String inputFileName = String.format("inputs/day%02d.txt", day);
    try {
      return ClassUtils.getResourceAsString(inputFileName);
    } catch (RuntimeException e) {
      String message = String.format("Failed to read input for day %d.", day);
      throw new RuntimeException(message, e);
    }
  }
}
