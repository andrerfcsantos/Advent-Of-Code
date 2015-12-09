package utils;

import java.io.*;
import java.util.List;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.stream.Collectors;

/**
 * Class that hosts the code for reading the input files.
 */

public final class FileHandler {

    private FileHandler() {}

    public static Predicate<String> NO_BLANK_LINES = (e) -> !e.equalsIgnoreCase("");
    public static Function<String,String> DO_NOTHING = Function.identity();
    public static Predicate<String> NO_FILTER = (e) -> true;

    public static List<String> getNotEmptyLines(String filePath) throws IOException {
        return getAndTransformLines(filePath, NO_BLANK_LINES, DO_NOTHING);
    }

    /**
     * Generates a list with the contents of the file in filepath.
     *
     * @param filePath Path to the input file.
     * @param filter The condition for each line to be accepted and included in the returned list.
     * @param transformation The transformation function to apply to each line before adding it to the returned list.
     * @param <T> The type of each element on the returned list.
     * @return A list with elements from each line of the input file.
     * @throws IOException File could not be read.
     */
    public static <T> List<T> getAndTransformLines(String filePath,
                                                   Predicate<String> filter,
                                                   Function<String,T> transformation) throws IOException {
        BufferedReader bin = new BufferedReader(new FileReader(new File(filePath)));
        return bin.lines().filter(filter).map(transformation).collect(Collectors.toList());
    }

    public static void writeFile(String path, List<String> lines) throws IOException {
        BufferedWriter bwriter = new BufferedWriter(new PrintWriter(new File(path)));
        lines.stream().forEach((s)->{
            try {
                bwriter.write(s+"\n");
            } catch (IOException e1) {
                e1.printStackTrace();
            }
        });
        bwriter.flush();
        bwriter.close();
    }

    }



