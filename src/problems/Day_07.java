package problems;

import utils.FileHandler;

import java.io.IOException;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.StringTokenizer;
import java.util.function.Function;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Created by Andre on 08-12-2015.
 */
public class Day_07 {

    private static ArrayList<ArrayList<String>> lines;
    private static HashMap<String, ArrayList<String>> instructions;
    private static HashMap<String, Integer> values;

    public static void main(String[] args) throws IOException {
        System.out.println("Solution Day 7 Problem 1: " + problem_01());
        System.out.println("Solution Day 7 Problem 2: " + problem_02());
    }

    public static int problem_01() throws IOException {
        Function<String,ArrayList<String>> spliter = (e) ->{
            ArrayList<String> res = new ArrayList<>();
            StringTokenizer strTok = new StringTokenizer(e, "-> \r\n");
            while(strTok.hasMoreTokens()) res.add(strTok.nextToken());
            return res;
        };

        lines = (ArrayList<ArrayList<String>>) FileHandler.getAndTransformLines("Inputfiles/day07_1.txt",
                FileHandler.NO_BLANK_LINES,
                spliter);

        instructions = new HashMap<>();
        values = new HashMap<>();

        for(ArrayList<String> line : lines){
            String lastElement = line.remove(line.size()-1);
            instructions.put(lastElement, line);
        }

        return calculate("a");
    }

    public static int calculate(String strVal){
        int res=-1;
        Pattern number = Pattern.compile("[0-9]*");
        Matcher m = number.matcher(strVal);

        if(values.containsKey(strVal)){
            res = values.get(strVal);
        }else if(m.matches()){
            res = Integer.parseInt(strVal);
        }else {
            switch (instructions.get(strVal).size()) {
                case 1:
                    res = calculate(instructions.get(strVal).get(0));
                    break;
                case 2:
                    res = 65535 - calculate(instructions.get(strVal).get(1));
                    break;
                case 3:
                    switch (instructions.get(strVal).get(1)) {
                        case "AND":
                            res = calculate(instructions.get(strVal).get(0)) & calculate(instructions.get(strVal).get(2));
                            break;
                        case "OR":
                            res = calculate(instructions.get(strVal).get(0)) | calculate(instructions.get(strVal).get(2));
                            break;
                        case "LSHIFT":
                            res = Integer.rotateLeft(calculate(instructions.get(strVal).get(0)), calculate(instructions.get(strVal).get(2)));
                            break;
                        case "RSHIFT":
                            res = Integer.rotateRight(calculate(instructions.get(strVal).get(0)), calculate(instructions.get(strVal).get(2)));
                            break;
                        default:
                            res = -1;
                    }
                    break;
                default:
                    res = -1;
            }
            values.put(strVal, res);
        }
        return res;
    }


    public static int problem_02() throws IOException {
        Function<String,ArrayList<String>> spliter = (e) ->{
            ArrayList<String> res = new ArrayList<>();
            StringTokenizer strTok = new StringTokenizer(e, "-> \r\n");
            while(strTok.hasMoreTokens()) res.add(strTok.nextToken());
            return res;
        };

        lines = (ArrayList<ArrayList<String>>) FileHandler.getAndTransformLines("Inputfiles/day07_1.txt",
                FileHandler.NO_BLANK_LINES,
                spliter);

        instructions = new HashMap<>();
        values = new HashMap<>();

        for(ArrayList<String> line : lines){
            String lastElement = line.remove(line.size()-1);
            instructions.put(lastElement, line);
        }
        values.put("b", 46065);
        return calculate("a");
    }

    public static void testBitwise() {
        int x = 123;
        int y = 456;
        BigInteger bx = new BigInteger(Integer.toString(x));
        BigInteger by = new BigInteger(Integer.toString(y));
        System.out.println("x AND y: " + (x & y));
        System.out.println("x OR y: " + (x | y));
        System.out.println("x LSHIFT Y: "+Integer.rotateLeft(x,2));
        System.out.println("x LSHIFT Y: "+Integer.rotateRight(y,2));
        System.out.println("NOT x: "+ (65535 - x));
        System.out.println("NOT y: "+ (65535- y));
    }

    public static void testRegex(){
        Pattern regIns = Pattern.compile("([a-z]|[0-9])* (AND|OR|LSHIFT|RSHIFT) ([a-z]|[0-9])* -> [a-z]*");
        Pattern notIns = Pattern.compile("NOT ([a-z]|[0-9])* -> [a-z]*");
        Pattern signalIns = Pattern.compile("[0-9]* -> [a-z]*");
        Matcher m1 = regIns.matcher("bo AND bu -> bw");
        Matcher m2 = regIns.matcher("NOT hn -> ho");
        System.out.println("Result: " + m1.matches());
        System.out.println("Result: " + m2.matches());

    }
}
