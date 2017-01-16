package problems;

import com.google.common.collect.Collections2;
import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;

/**
 * Created by Andre on 09-12-2015.
 */
public class Day_09 {

    private static ArrayList<String[]> lines;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String[]>) FileHandler.getAndTransformLines("../inputfiles/day09.txt",
                FileHandler.NO_FILTER,
                ((String e) -> {return e.split(" ((to)|=) ");}));
        System.out.println("[Day 09] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 09] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 09] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException {
        int id=0, nVertices=0;
        int adjacencymatrix[][];
        int shorthestPath=Integer.MAX_VALUE, currentPath=0;

        HashMap<String, Integer> citiesMap = new HashMap<>();
        ArrayList<String> citiesList = new ArrayList<>();

        for(String[] line : lines){
            if(!citiesMap.containsKey(line[0])){
                citiesMap.put(line[0],id);
                id++;
            }
            if(!citiesMap.containsKey(line[1])){
                citiesMap.put(line[1],id);
                id++;
            }
        }
        nVertices=citiesMap.size();
        adjacencymatrix = new int[nVertices][nVertices];


        for(String[] line : lines){
            adjacencymatrix[citiesMap.get(line[0])][citiesMap.get(line[1])] = Integer.parseInt(line[2]);
            adjacencymatrix[citiesMap.get(line[1])][citiesMap.get(line[0])] = Integer.parseInt(line[2]);
        }

        for(String city:citiesMap.keySet()){
            citiesList.add(city);
        }

        Collection<List<String>> perms = Collections2.permutations(citiesList);

        for(List<String> caminho : perms){
            currentPath=0;
            for(int i=0; i<caminho.size()-1;i++){
                currentPath += adjacencymatrix[citiesMap.get(caminho.get(i))][citiesMap.get(caminho.get(i+1))];
            }
            if(currentPath<shorthestPath) shorthestPath=currentPath;
        }


        return shorthestPath;
    }

    public static int problem_02() throws IOException {
        int id=0, nVertices=0;
        int adjacencymatrix[][];
        int longestPath=Integer.MIN_VALUE, currentPath=0;

        HashMap<String, Integer> citiesMap = new HashMap<>();
        ArrayList<String> citiesList = new ArrayList<>();


        for(String[] line : lines){
            if(!citiesMap.containsKey(line[0])){
                citiesMap.put(line[0],id);
                id++;
            }
            if(!citiesMap.containsKey(line[1])){
                citiesMap.put(line[1],id);
                id++;
            }
        }
        nVertices=citiesMap.size();
        adjacencymatrix = new int[nVertices][nVertices];


        for(String[] line : lines){
            adjacencymatrix[citiesMap.get(line[0])][citiesMap.get(line[1])] = Integer.parseInt(line[2]);
            adjacencymatrix[citiesMap.get(line[1])][citiesMap.get(line[0])] = Integer.parseInt(line[2]);
        }

        for(String city:citiesMap.keySet()){
            citiesList.add(city);
        }

        Collection<List<String>> perms = Collections2.permutations(citiesList);

        for(List<String> caminho : perms){
            currentPath=0;
            for(int i=0; i<caminho.size()-1;i++){
                currentPath += adjacencymatrix[citiesMap.get(caminho.get(i))][citiesMap.get(caminho.get(i+1))];
            }
            if(currentPath>longestPath) longestPath=currentPath;
        }


        return longestPath;
    }

}
