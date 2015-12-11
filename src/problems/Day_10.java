package problems;

import utils.Crono;

import java.io.IOException;

/**
 * Created by Andre on 10-12-2015.
 */
public class Day_10 {
    private static String originalString= "1321131112";
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();

        System.out.print("[Day 10] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 10] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException {
        int totalL=0, startStrL=0;
        int qtdChar=0;
        char currentChar='z', lastChar='z';
        StringBuilder startingString = new StringBuilder(originalString);
        StringBuilder newstr;

        for(int times=0;times<40;times++) {
            lastChar=startingString.charAt(0);
            newstr = new StringBuilder();
            startStrL= startingString.length();

            for (int i = 0; i < startStrL; i++) {

                currentChar = startingString.charAt(i);
                if (currentChar == lastChar) {
                    qtdChar++;
                }
                else {
                    newstr.append(qtdChar).append(lastChar);
                    qtdChar = 1;
                }
                lastChar = currentChar;
            }
            newstr.append(qtdChar).append(lastChar);
            qtdChar = 0;

            startingString = new StringBuilder((newstr.toString()));
        }

        return startingString.length();
    }

    public static int problem_02() throws IOException {
        int totalL=0, startStrL=0;
        int qtdChar=0;
        char currentChar='z', lastChar='z';
        StringBuilder startingString = new StringBuilder(originalString);
        StringBuilder newstr;

        for(int times=0;times<50;times++) {
            lastChar=startingString.charAt(0);
            newstr = new StringBuilder();
            startStrL= startingString.length();

            for (int i = 0; i < startStrL; i++) {

                currentChar = startingString.charAt(i);
                if (currentChar == lastChar) {
                    qtdChar++;
                }
                else {
                    newstr.append(qtdChar).append(lastChar);
                    qtdChar = 1;
                }
                lastChar = currentChar;
            }
            newstr.append(qtdChar).append(lastChar);
            qtdChar = 0;

            startingString = new StringBuilder((newstr.toString()));
        }

        return startingString.length();
    }
}
