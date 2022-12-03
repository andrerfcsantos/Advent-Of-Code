import {
  runProblem,
  sumNumbers,
  intersectSets,
  chunkArray,
  isUppercase,
  nonEmptyLines,
} from "./aoclib.ts";

const DAY = 3;
const YEAR = 2022;

type Backpack = string;

interface State {
  backpacks: Backpack[];
}

export function parse(input: string): State {
  return { backpacks: nonEmptyLines(input) };
}

export function part1(parsed: State): string {
  const commonItems = parsed.backpacks.map(findCommonItemTypeInCompartiments);
  const priorities = commonItems.map(getItemPriority);
  return sumNumbers(priorities).toString();
}

export function part2(parsed: State): string {
  const groups = chunkArray(parsed.backpacks, 3);

  const badgesforGroups = groups.map((group) => {
    const backpacks = group.map((backpack) => new Set(backpack));
    return Array.from(intersectSets(...backpacks))[0];
  });

  const priorities = badgesforGroups.map(getItemPriority);

  return sumNumbers(priorities).toString();
}

function findCommonItemTypeInCompartiments(backpack: Backpack): string {
  const middle = backpack.length / 2;
  const compartiments = [
    new Set(backpack.slice(0, middle)),
    new Set(backpack.slice(middle, backpack.length)),
  ];
  return Array.from(intersectSets(...compartiments))[0];
}

const LOWER_A = "a".charCodeAt(0);
const UPPER_A = "A".charCodeAt(0);

function getItemPriority(item: string): number {
  const c = item.charCodeAt(0);
  const diff = isUppercase(item) ? c - UPPER_A + 26 : c - LOWER_A;
  return diff + 1;
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
