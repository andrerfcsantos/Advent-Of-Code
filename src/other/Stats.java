package other;

import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;
import java.util.function.Function;

/**
 * Created by Andre on 09-12-2015.
 */
public class Stats {

    public static void main(String[] args) throws IOException {
        Function<String, int[]> separator = (s) ->{
            int []res = new int[3];
            int i=0;
            StringTokenizer strTok = new StringTokenizer(s, " *\t\r\n");
            while(strTok.hasMoreTokens()){
                res[i++] = Integer.parseInt(strTok.nextToken());
            }
            return res;
        };
        ArrayList<String> linestoprint = new ArrayList<>();
        ArrayList<int[]> lines = (ArrayList<int[]>) FileHandler.getAndTransformLines("Inputfiles/stats.txt",
                                                                                        FileHandler.NO_BLANK_LINES,
                                                                                        separator);

        lines.sort((p1, p2)->p1[0]-p2[0]);
        lines.stream().forEach((e)-> {
            linestoprint.add(""+e[0]+"\t"+e[1]+"\t"+e[2]);
        });
        FileHandler.writeFile("others/outStats.txt",linestoprint);
    }
}
