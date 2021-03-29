package day02;

import java.util.Arrays;

public class Password {

    private int lowestNum;
    private int highestNum;
    private char letter;
    private String password;

    private Password(){ }

    public Password(int lowestNum, int highestNum, char letter, String password) {
        this.lowestNum = lowestNum;
        this.highestNum = highestNum;
        this.letter = letter;
        this.password = password;
    }

    public boolean isValidPart1() {
        long count = password.chars().filter(c -> (char) c == letter).count();
        return count >= lowestNum && count <= highestNum;
    }

    public boolean isValidPart2() {
        return password.charAt(lowestNum-1) == letter ^ password.charAt(highestNum-1) == letter;
    }

}
