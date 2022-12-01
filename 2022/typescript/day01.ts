import { downloadInput } from "./aoclib.ts"
import { info } from "https://deno.land/std@0.166.0/log/mod.ts";

const DAY = 1;
const YEAR = 2022;

interface State {
    lines: string[];
    elves: number[][];
    totalByElf?: number[];
}

export function parse(lines: string[]): State {

    const elves: number[][] = [];

    let currentElf: number[] = [];

    for (const line of lines) {
        if(line==="") {
            elves.push(currentElf);
            currentElf = [];
            continue;
        }
        currentElf.push(parseInt(line, 10));
    }

    return {lines, elves};
}

function sumNumbers(ns: number[]): number {
    return ns.reduce((a, b) => a + b, 0);
}
  
export function part1(parsed: State): string {
    const totalByElf = parsed.elves.map(sumNumbers);

    // save totalByElf for part 2
    parsed.totalByElf = totalByElf;
    
    return Math.max(...totalByElf).toString();
}

export function part2(parsed: State): string {
    const top3 = parsed.totalByElf!.sort((a, b) => b - a).slice(0, 3);
    return sumNumbers(top3).toString();
}

performance.mark("start download");
const lines = (await downloadInput(DAY, YEAR)).split(/\r?\n/);
performance.mark("end download");

performance.mark("start parse");
const state = parse(lines);
performance.mark("end parse");

performance.mark("start part1");
const p1 = part1(state);
performance.mark("end part1");

performance.mark("start part2");
const p2 = part2(state);
performance.mark("end part2");

const perf_download = performance.measure("download", "start download", "end download");
const perf_parse = performance.measure("parse", "start parse", "end parse");
const perf_part1 = performance.measure("part1", "start part1", "end part1");
const perf_part2 = performance.measure("part2", "start part2", "end part2");

info(`Download in ${perf_download.duration}ms | Input parsing in ${perf_parse.duration}ms`);
info(`Part 1: ${p1} (${perf_part1.duration}ms)`);
info(`Part 2: ${p2} (${perf_part2.duration}ms)`);
