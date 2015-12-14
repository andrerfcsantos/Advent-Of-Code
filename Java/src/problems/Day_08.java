package problems;

import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;

/**
 * Created by Andre on 08-12-2015.
 */
public class Day_08 {

    private static ArrayList<String> lines;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String>) FileHandler.getNotEmptyLines("Inputfiles/day08_1.txt");
        System.out.println("[Day 08] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 08] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 08] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }


    public static int problem_01() throws IOException {
        int i;
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
