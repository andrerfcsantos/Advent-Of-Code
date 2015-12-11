package problems;

import utils.Crono;

import javax.xml.bind.DatatypeConverter;
import java.io.IOException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;


/**
 * Problem 4
 */
public class Day_04 {
    private static Crono crono;

    public static String secretKey = "iwrupvqb";

    public static void main(String[] args) throws IOException, NoSuchAlgorithmException {
        crono = new Crono();

        crono.start();
        System.out.print("[Day 04] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 04] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }


    public static int problem_01() throws IOException, NoSuchAlgorithmException {
        int result = -1;
        byte[] digestResult;
        String digestStr;
        MessageDigest md =  MessageDigest.getInstance("MD5");

        for(int i=0;i<Integer.MAX_VALUE;i++){

            digestResult = md.digest((secretKey+i).getBytes());
            digestStr = DatatypeConverter.printHexBinary(digestResult);

            if(digestStr.startsWith("00000")){
                result = i;
                break;
            }

        }

        return result;
    }

    public static int problem_02() throws IOException, NoSuchAlgorithmException  {
        int result = -1;
        byte[] digestResult;
        String digestStr;
        MessageDigest md =  MessageDigest.getInstance("MD5");

        for(int i=0;i<Integer.MAX_VALUE;i++){

            digestResult = md.digest((secretKey+i).getBytes());
            digestStr = DatatypeConverter.printHexBinary(digestResult);

            if(digestStr.startsWith("000000")){
                result = i;
                break;
            }

        }

        return result;
    }

}
