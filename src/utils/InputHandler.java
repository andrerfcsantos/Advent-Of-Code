package utils;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Class that hosts the code for reading the input files.
 */

public final class InputHandler {
    private InputHandler() {
    }

    public static List<String> getLines(String filePath) throws IOException {
        ArrayList<String> lines = new ArrayList<>();
        String line;
        File filein = new File(filePath);
        BufferedReader bin = new BufferedReader(new FileReader(filein));

        while (bin.ready()) {
            line = bin.readLine();
            if (!line.equalsIgnoreCase("")) {
                lines.add(line);
            }
        }

        bin.close();

        return lines;
    }

    /*
    Day 6 requires a special treatment for each line, that i decided to do here.
    This is a bodge, intended to "just work". In the future i'll consider other alternatives
     like changing the original function getLines to receive a function as argument (java 8 ftw!).
     */
    public static List<String> getLinesDay6(String filePath) throws IOException {
        ArrayList<String> lines = new ArrayList<>();
        String line;
        File filein = new File(filePath);
        BufferedReader bin = new BufferedReader(new FileReader(filein));

        while (bin.ready()) {
            line = bin.readLine();
            if (!line.equalsIgnoreCase("")) {
                lines.add(line.replace("turn on","turnon")
                                .replace("turn off","turnoff")
                                .replace(" through ", " "));
            }
        }

        bin.close();

        return lines;
    }

}
