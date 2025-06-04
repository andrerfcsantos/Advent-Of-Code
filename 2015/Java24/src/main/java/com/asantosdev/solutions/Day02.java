package com.asantosdev.solutions;

import com.asantosdev.utils.ReadUtils;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Stream;

public class Day02 implements Solver {

  public record Dimensions(int length, int width, int height) {

    private static final Pattern pattern = Pattern.compile( "(?<l>\\d+)x(?<w>\\d+)x(?<h>\\d+)");

    public static Dimensions fromString(String dimensions) {
      String[] parts = dimensions.split("x");
      if (parts.length != 3) {
        throw new IllegalArgumentException("Invalid dimensions format: " + dimensions);
      }
      int length = Integer.parseInt(parts[0]);
      int width = Integer.parseInt(parts[1]);
      int height = Integer.parseInt(parts[2]);
      return new Dimensions(length, width, height);
    }

    public static Optional<Dimensions> fromStringRegex(String dimensions) {

      Matcher matcher = pattern.matcher(dimensions);

      if(matcher.matches()){
        int length = Integer.parseInt(matcher.group("l"));
        int width = Integer.parseInt(matcher.group("w"));
        int height = Integer.parseInt(matcher.group("h"));
        return Optional.of(new Dimensions(length, width, height));
      }

      return Optional.empty();
    }

    public int surfaceArea() {
      return 2*length*width + 2*width*height + 2*height*length;
    }

    public int wrappingPaperSlack() {
      return Math.min(Math.min(length*width, length*height), width*height);
    }

    public int wrappingPaper() {
      return surfaceArea() + wrappingPaperSlack();
    }

    public int ribbon() {
      return ribbonBox() + ribbonBow();
    }

    public int ribbonBox() {
      List<Integer> list = Stream.of(length, width, height).sorted().limit(2).toList();
      var min1 = list.get(0);
      var min2 = list.get(1);
      return min1*2 + min2*2;
    }

    public int ribbonBow() {
      return length*width*height;
    }
  }

  private final List<Dimensions> dimensionList = new ArrayList<>();

  public void processInput(String input) {
    List<String> lines = ReadUtils.nonEmptyLines(input);
    for (String line : lines) {
      Optional<Dimensions> d = Dimensions.fromStringRegex(line);
      d.ifPresent(this.dimensionList::add);
    }
  }

  public String part1() {
    int wrappingPaper = 0;
    for (Dimensions d : dimensionList) {
      wrappingPaper += d.wrappingPaper();
    }
    return String.valueOf(wrappingPaper);
  }

  public String part2() {
    int ribbon = 0;
    for (Dimensions d : dimensionList) {
      ribbon += d.ribbon();
    }
    return String.valueOf(ribbon);
  }
}
