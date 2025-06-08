package com.asantosdev;

import com.asantosdev.solutions.Day01;
import com.asantosdev.solutions.Day02;
import com.asantosdev.solutions.Solver;
import com.asantosdev.utils.AOCUtils;
import com.asantosdev.utils.ClassUtils;
import com.google.common.base.Stopwatch;
import com.google.common.collect.ImmutableSet;
import com.google.common.reflect.ClassPath;
import java.lang.reflect.Constructor;
import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.Callable;
import picocli.CommandLine;
import picocli.CommandLine.Command;
import picocli.CommandLine.Parameters;

@Command(
    name = "solution",
    mixinStandardHelpOptions = true,
    version = "checksum 4.0",
    description = "Runs the Advent of Code solutions.")
public class AdventOfCode implements Callable<Integer> {

  @Parameters(
      index = "0",
      description = "The day of the Advent of Code challenge (e.g., 1 for Day 1).",
      defaultValue = "1")
  private int day;

  @Override
  public Integer call() throws Exception { // your business logic goes here...
    Map<Integer, Solver> solvers = Map.ofEntries(
            Map.entry(1, new Day01()),
            Map.entry(2, new Day02())
    );

    if (day < 1 || day > solvers.size()) {
      System.err.println("Invalid day. Please provide a day between 1 and " + solvers.size());
      return 1;
    }

    Solver solver = solvers.get(day);
    String input = AOCUtils.getStringInputForDay(day);

    Stopwatch stopwatch = Stopwatch.createStarted();
    solver.processInput(input);
    stopwatch.stop();
    System.out.println("Processed input in " + stopwatch);

    stopwatch.reset().start();
    System.out.println("Part 1: " + solver.part1() + " (in " + stopwatch + ")");
    stopwatch.reset().start();
    System.out.println("Part 2: " + solver.part2() + " (in " + stopwatch + ")");
    stopwatch.stop();

    return 0;
  }

  public static void main(String[] args) {
    int exitCode = new CommandLine(new AdventOfCode()).execute(args);
    System.exit(exitCode);
  }
}
