package problems;

import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Problem 05
 */
public class Day_05 {

    private static ArrayList<String> lines;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String>) FileHandler.getNotEmptyLines("Inputfiles/day05_1.txt");
        System.out.println("[Day 05] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 05] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 05] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException{
        int niceStrings = 0, vowelCount =0;
        boolean doubleLetter = false, substrNotAllowed=false;
        char lastChar, currentChar;

        for(String line:lines){
            vowelCount =0; doubleLetter = false;
            substrNotAllowed=false;
            lastChar = '?';

            for(int i=0; i<line.length();i++){
                currentChar = line.charAt(i);

                if((currentChar=='b' && lastChar=='a') ||
                        (currentChar=='d' && lastChar=='c') ||
                        (currentChar=='q' && lastChar=='p') ||
                        (currentChar=='y' && lastChar=='x')){
                    substrNotAllowed=true;
                    break;
                }

                if(currentChar=='a' || currentChar=='e' ||
                        currentChar=='i' || currentChar=='o' || currentChar=='u'){
                    vowelCount++;
                }

                if(currentChar==lastChar) doubleLetter=true;
                lastChar = currentChar;
            }
            if(!substrNotAllowed && doubleLetter && vowelCount>=3) niceStrings++;
        }

        return niceStrings;
    }

    public static int problem_02() throws IOException {
        int niceStrings = 0, vowelCount =0;
        boolean repeatedLetter = false, repeatedGroup=false;
        char currentChar;

        for(String line:lines){
            repeatedLetter = false; repeatedGroup=false;

            for(int i=2; i<line.length() && !(repeatedGroup && repeatedLetter);i++){
                currentChar = line.charAt(i);

                if(currentChar == line.charAt(i-2)) repeatedLetter=true;
                if(line.substring(i).contains(""+line.charAt(i-2)+line.charAt(i-1)))
                    repeatedGroup=true;
            }
            if(repeatedGroup&&repeatedLetter) niceStrings++;
        }

        return niceStrings;
    }


}
