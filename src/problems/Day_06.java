package problems;

import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;
import java.util.function.Function;

/**
 * Problem 06
 */
public class Day_06 {
    public static boolean lights[][] = new boolean[1000][1000];
    public static byte lightsPart2[][] = new byte[1000][1000];

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 6 Problem 1: " + problem_01());
        System.out.println("Solution Day 6 Problem 2: " + problem_02());
    }


    public static int problem_01() throws IOException {
        int lightsOn = 0;
        String instruction, lowerCoords, higherCoords;
        int x1,y1,x2,y2, deltaX, deltaY;
        Function<String, String> simplifyInstr = (e) -> e.replace("turn on","turnon")
                                                            .replace("turn off","turnoff")
                                                            .replace(" through ", " ");

        ArrayList<String> lines = (ArrayList<String>)
                        FileHandler.getAndTransformLines("Inputfiles/day06_1.txt",
                        FileHandler.NO_BLANK_LINES,
                                simplifyInstr);

        Crono time = new Crono();
        time.start();

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, " ,\n\r");

            instruction = strTok.nextToken();
            x1 = Integer.parseInt(strTok.nextToken());
            y1 = Integer.parseInt(strTok.nextToken());
            x2 = Integer.parseInt(strTok.nextToken());
            y2 = Integer.parseInt(strTok.nextToken());

            deltaX = x2 - x1;
            deltaY = y2 - y1;

            switch (instruction) {
                case "turnon":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            lights[i][j] = true;
                    break;
                case "turnoff":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            lights[i][j] = false;
                    break;
                case "toggle":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            lights[i][j] = !lights[i][j];
                    break;
                default:
                    break;
            }
        }

        for (int i = 0; i < lights[0].length; i++)
            for(int j = 0; j < lights.length;j++)
                if(lights[i][j]) lightsOn++;


        time.stop();
        System.out.println("Time (ms): " + time.getElapsedTime().toMillis());
        return lightsOn;

    }

    public static int problem_02() throws IOException {
        int brightness = 0;
        String instruction, lowerCoords, higherCoords;
        int x1,y1,x2,y2, deltaX, deltaY;
        Function<String, String> simplifyInstr = (e) -> e.replace("turn on","turnon")
                .replace("turn off","turnoff")
                .replace(" through ", " ");
        ArrayList<String> lines = (ArrayList<String>)
                FileHandler.getAndTransformLines("Inputfiles/day06_1.txt",
                        FileHandler.NO_BLANK_LINES,
                        simplifyInstr);

        Crono time = new Crono();
        time.start();

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, " ,\n\r");

            instruction = strTok.nextToken();
            x1 = Integer.parseInt(strTok.nextToken());
            y1 = Integer.parseInt(strTok.nextToken());
            x2 = Integer.parseInt(strTok.nextToken());
            y2 = Integer.parseInt(strTok.nextToken());

            deltaX = x2 - x1;
            deltaY = y2 - y1;

            switch (instruction) {
                case "turnon":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            lightsPart2[i][j]++;
                    break;
                case "turnoff":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            if(lightsPart2[i][j]!=0) lightsPart2[i][j]--;
                    break;
                case "toggle":
                    for (int i = x1; i <= x2; i++)
                        for (int j = y1; j <= y2; j++)
                            lightsPart2[i][j] += 2;
                    break;
                default:
                    break;
            }
        }

        for (int i = 0; i < lights[0].length; i++)
            for(int j = 0; j < lights.length;j++)
                brightness += lightsPart2[i][j];


        time.stop();
        System.out.println("Time (ms): " + time.getElapsedTime().toMillis());
        return brightness;
    }
}
