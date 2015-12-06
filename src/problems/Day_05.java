package problems;

import utils.Crono;
import utils.InputHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Problem 05
 */
public class Day_05 {

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 5 Problem 1: " + problem_01());
        System.out.println("Solution Day 5 Problem 2: " + problem_02());
    }

    public static int problem_01() throws IOException{
        int niceStrings = 0, vowelCount =0;
        boolean doubleLetter = false, substrNotAllowed=false;
        char lastChar, currentChar;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getLines("Inputfiles/day05_1.txt");

        Crono time = new Crono();
        time.start();

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
        time.stop();

        System.out.println("Time (ms): " + time.getElapsedTime().toMillis());
        return niceStrings;
    }

    public static int problem_02() throws IOException {
        int niceStrings = 0, vowelCount =0;
        boolean repeatedLetter = false, repeatedGroup=false;
        char currentChar;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getLines("Inputfiles/day05_1.txt");
        Crono time = new Crono();
        time.start();
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
        time.stop();

        System.out.println("Time (ms): " + time.getElapsedTime().toMillis());
        return niceStrings;
    }


}
