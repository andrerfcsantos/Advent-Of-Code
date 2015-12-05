package problems;

import javax.xml.bind.DatatypeConverter;
import java.io.IOException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;


/**
 * Problem 4
 */
public class Day_04 {

    public static void main(String[] args) throws IOException, NoSuchAlgorithmException {
        System.out.println("Solution Day 4 Problem 1: " + problem_01());
        System.out.println("Solution Day 4 Problem 2: " + problem_02());
    }


    public static int problem_01() throws IOException, NoSuchAlgorithmException {
        int result = -1;
        String secretKey = "iwrupvqb";
        byte[] digestResult;
        String digestStr;
        MessageDigest md =  MessageDigest.getInstance("MD5");
        long start = System.nanoTime();
        for(int i=0;i<Integer.MAX_VALUE;i++){

            digestResult = md.digest((secretKey+i).getBytes());
            digestStr = DatatypeConverter.printHexBinary(digestResult);

            if(digestStr.startsWith("00000")){
                result = i;
                break;
            }

        }
        long elapsedTime = (System.nanoTime() - start) / 1000000000;
        System.out.println("Result: " + result + " (in " + elapsedTime + " seconds)");
        return result;
    }

    public static int problem_02() throws IOException, NoSuchAlgorithmException  {
        int result = -1;
        String secretKey = "iwrupvqb";
        byte[] digestResult;
        String digestStr;
        MessageDigest md =  MessageDigest.getInstance("MD5");
        long start = System.nanoTime();
        for(int i=0;i<Integer.MAX_VALUE;i++){

            digestResult = md.digest((secretKey+i).getBytes());
            digestStr = DatatypeConverter.printHexBinary(digestResult);

            if(digestStr.startsWith("000000")){
                result = i;
                break;
            }

        }
        long elapsedTime = (System.nanoTime() - start) / 1000000000;
        System.out.println("Result: " + result + " (in " + elapsedTime + " seconds)");
        return result;
    }

}
