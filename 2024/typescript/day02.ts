import { runProblem } from "./lib/runner.ts";
import { nonEmptyLines } from "./lib/input.ts";

const DAY = 2;
const YEAR = 2024;

interface State {
  reports: number[][];
}

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  const reports = lines.map((line) => line.split(" ").map((x) => parseInt(x, 10)))
  return { reports };
}

export function part1(parsed: State): string {
  const count = parsed.reports.filter(validReportWithoutDampener).length
  return count.toString(10);
}

export function part2(parsed: State): string {
  const count = parsed.reports.filter(validReportWithDampener).length
  return count.toString(10);
}

function validReportWithoutDampener(report: number[]): boolean {
  
  const isIncreasing = report[0] < report[1];

  for(let i = 1; i<report.length; i++) {
    const diff = report[i] - report[i-1];
    if(isIncreasing) {
      if(diff > 3 || diff <= 0){
        return false;
      }
    } else {
      if(diff < -3 || diff >= 0){
        return false;
      }
  }
  }
  return true;
} 

function validReportWithDampener(report: number[]): boolean {

  const reports = [report]
  for(let i = 0; i < report.length; i++) {
    reports.push(report.toSpliced(i, 1))
  }

  for(const report of reports) {
    if(validReportWithoutDampener(report)) {
      return true
    }
  }

  return false;
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);