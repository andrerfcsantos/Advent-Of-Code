package problems;

import com.google.common.collect.Collections2;
import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.*;

/**
 * Created by Andre on 13-12-2015.
 */
public class Day_12 {


    private static ArrayList<String> lines;
    private static String line;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String>) FileHandler.getNotEmptyLines("Inputfiles/day12_1.txt");
        line = lines.get(0);

        System.out.println("[Day 012] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 012] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 012] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException {
        int soma =0;
        StringTokenizer strTok = new StringTokenizer(line, "\"{}[],:\n\r");

        while (strTok.hasMoreTokens()){
            try {
                soma += Integer.parseInt(strTok.nextToken());
            } catch (NumberFormatException e){

            }
        }

        return soma;
    }

    public static int problem_02() throws IOException {
        return -1;
    }
}
