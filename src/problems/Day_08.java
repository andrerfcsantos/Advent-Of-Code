package problems;

import utils.InputHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Created by Andre on 08-12-2015.
 */
public class Day_08 {

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 8 Problem 1: " + problem_01());
        System.out.println("Solution Day 8 Problem 2: " + problem_02());
    }


    public static int problem_01() throws IOException {
        int i;
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getNotEmptyLines("Inputfiles/day08_1.txt");
        int totaldif=0;
        int strLen=0,actualchars=0;

        for(String line:lines){
            strLen = line.length();
            actualchars=0;
            i=1;
            while(i < strLen-1){
                if(line.charAt(i)=='\\' && (line.charAt(i+1)=='\"' || line.charAt(i+1)=='\\')){
                    i+=2;
                }else if(line.charAt(i)=='\\' && line.charAt(i+1)=='x'){
                    i+=4;
                }else{
                    i++;
                }

                actualchars++;
            }
            totaldif += strLen - actualchars;
        }

        return totaldif;
    }

    public static int problem_02() throws IOException  {
        ArrayList<String> lines = (ArrayList<String>) InputHandler.getNotEmptyLines("Inputfiles/day08_1.txt");
        int totaldif=0;
        int strLen=0,encodedchars=0;

        for(String line:lines){
            strLen = line.length();
            encodedchars=2;

            for(int i=0;i < strLen;i++){
                if(line.charAt(i)=='\\' || line.charAt(i)=='\"'){
                    encodedchars+=2;
                }else {
                    encodedchars++;
                }
            }
            totaldif += encodedchars - strLen;
        }

        return totaldif;
    }

}
