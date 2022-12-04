import { runProblem, nonEmptyLines } from "./aoclib.ts";

const DAY = 4;
const YEAR = 2022;

interface Range {
  min: number;
  max: number;
}

interface RangePair {
  first: Range;
  second: Range;
}

interface State {
  rangePairs: RangePair[];
}

export function parse(input: string): State {
  return {
    rangePairs: nonEmptyLines(input).map((l) => {
      const [r1, r2] = l.split(",");
      const [min1, max1] = r1.split("-").map((x) => parseInt(x, 10));
      const [min2, max2] = r2.split("-").map((x) => parseInt(x, 10));
      return {
        first: { min: min1, max: max1 },
        second: { min: min2, max: max2 },
      };
    }),
  };
}

function isContainedBy(r1: Range, r2: Range): boolean {
  return r1.min >= r2.min && r1.max <= r2.max;
}

function doesOverlap(r1: Range, r2: Range): boolean {
  return r1.min <= r2.max && r1.max >= r2.min;
}

export function part1(parsed: State): string {
  const containedRanges = parsed.rangePairs.filter((rangePair) => {
    const { first, second } = rangePair;
    return isContainedBy(first, second) || isContainedBy(second, first);
  });

  return containedRanges.length.toString();
}

export function part2(parsed: State): string {
  const overlappedRanges = parsed.rangePairs.filter((rangePair) => {
    const { first, second } = rangePair;
    return doesOverlap(first, second) || doesOverlap(second, first);
  });

  return overlappedRanges.length.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
