import day01.Day01Solver;
import day02.Day02Solver;
import day03.Day03Solver;
import puzzle.Solver;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Main {

    private static Map<Integer, Solver> daySolvers = Map.of(
            1, new Day01Solver(2020),
            2, new Day02Solver(),
            3, new Day03Solver()
    );

    public static void main(String[] args) throws IOException, URISyntaxException {
        // Parse command line arguments to get the day to solve
        if (args.length < 1) {
            System.out.println("Pass the day you want to solve as an argument to the program");
            System.exit(1);
        }
        int day = Integer.parseInt(args[0]);

        // Find the solver for this day
        if (!daySolvers.containsKey(day)) {
            System.out.println("Could not find a solver for day " + day);
            System.exit(1);
        }
        Solver solver = daySolvers.get(day);

        // Try to fetch input file from resources for this day
        String filename = String.format("day%02d.txt", day);
        URI resourceURI = Main.class.getClassLoader().getResource(filename).toURI();
        String resourcePath = Paths.get(resourceURI).toString();
        List<String> fileLines = Files.readAllLines(Path.of(resourcePath));

        // Solve the problem
        solver.processInput(fileLines);
        String part1 = solver.part1();
        String part2 = solver.part2();


        // Print the solution
        System.out.println("====== Day " + day + " ======");
        System.out.println("Part 1: " + part1);
        System.out.println("Part 2: " + part2);
    }

}
