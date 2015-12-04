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

    public static void printLines(String filePath) throws IOException {
        ArrayList<String> lines = (ArrayList<String>) getLines(filePath);

        for (int i = 0; i < lines.size(); i++) {
            System.out.println("Line " + i + ": " + lines.get(i));
        }

    }
}
