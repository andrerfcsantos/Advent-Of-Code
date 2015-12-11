package problems;

import utils.Crono;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Created by Andre on 11-12-2015.
 */
public class Day_11 {

    private static ArrayList<String> lines;
    private static Crono crono;
    private static char[] password1 = "hxbxwxba".toCharArray();
    private static char[] password2 = "hxbxxyzz".toCharArray();

    public static void main(String[] args) throws IOException {

        //System.out.println(isValid("hijklmmn".toCharArray()));

        crono = new Crono();

        crono.start();
        System.out.print("[Day 11] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 11] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");


    }

    public static String problem_01() throws IOException {
        boolean found=false;

        while(!found){
            incrementaStr(password1);
            if(isValid(password1)) found=true;
        }

        return String.copyValueOf(password1);
    }

    public static String problem_02() throws IOException  {

        boolean found=false;

        while(!found){
            incrementaStr(password2);
            if(isValid(password2)) found=true;
        }

        return String.copyValueOf(password2);
    }


    public static boolean isValid(char []str){
        boolean hasForbiddenChars = false;
        boolean hasIncrementalStreak = false;
        int doubleLetterCooldown = 2;
        int incrementalStreak = 1;
        int nrDoubleLetters = 0;
        char currChar, nextChar;

        for(int i=0; i<str.length -1 && !hasForbiddenChars; i++){
            currChar = str[i];
            nextChar = str[i+1];

            if(currChar == 'i' || currChar == 'o' || currChar == 'l'){
                hasForbiddenChars = true;
                break;
            }

            if(currChar==nextChar && doubleLetterCooldown==0){
                nrDoubleLetters++;
                doubleLetterCooldown=2;
            }

            if(!hasIncrementalStreak && currChar == nextChar-1){
                incrementalStreak++;
                if(incrementalStreak>=3) hasIncrementalStreak=true;
            }else{
                incrementalStreak=1;
            }

            if(doubleLetterCooldown>0) doubleLetterCooldown--;

        }

        return nrDoubleLetters>=2 && hasIncrementalStreak && !hasForbiddenChars;
    }

    public static void incrementaStr(char []str){
        boolean incremented=false;
        int lim = str.length-1;

        while(lim > 0 && !incremented){
            if(str[lim] == 'z'){
                str[lim] = 'a';
            }else{
                str[lim]++;
                incremented= true;
            }
            lim--;
        }

    }

}
