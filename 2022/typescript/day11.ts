import { runProblem, nonEmptyLines, lcm } from "./aoclib.ts";

const DAY = 11;
const YEAR = 2022;

interface Monkey {
  number: number;
  items: number[];
  divisibleTtest: number;
  ifTrueThrow: number;
  ifFalseThrow: number;
  operation: Op;
}

interface State {
  monkeys: Monkey[];
}

type Op = (old: number) => number;

const monkeyRegex =
  /Monkey (?<monkeyNo>\d+):\r?\n\s*Starting items: (?<numbers>.*)\r?\n\s*Operation: new = old (?<op>.) (?<arg2>.*)\r?\n\s*Test: divisible by (?<divisbleBy>\d+)\r?\n\s*If true: throw to monkey (?<throwTrue>\d+)\r?\n\s*If false: throw to monkey (?<throwFalse>\d+)/m;

function getOperationFuncForOld(op: string): Op {
  switch (op) {
    case "+":
      return (old: number) => old + old;
    case "-":
      return (old: number) => old - old;
    case "*":
      return (old: number) => old * old;
    case "/":
      return (old: number) => old / old;
    default:
      return (old: number) => old + old;
  }
}

function getOperationFuncForArg(op: string, arg: number): Op {
  switch (op) {
    case "+":
      return (old: number) => old + arg;
    case "-":
      return (old: number) => old - arg;
    case "*":
      return (old: number) => old * arg;
    case "/":
      return (old: number) => old / arg;
    default:
      return (old: number) => old + arg;
  }
}

export function parse(input: string): State {
  const newlines = /\r?\n\r?\n/;

  const monkeys = nonEmptyLines(input, newlines).map((monkey) => {
    const { groups } = monkeyRegex.exec(monkey)!;
    const items = groups!.numbers
      .split(", ")
      .map((n: string) => parseInt(n, 10));

    const operation =
      groups!.arg2 == "old"
        ? getOperationFuncForOld(groups!.op)
        : getOperationFuncForArg(groups!.op, parseInt(groups!.arg2, 10));

    return {
      number: parseInt(groups!.monkeyNo, 10),
      items,
      divisibleTtest: parseInt(groups!.divisbleBy, 10),
      ifTrueThrow: parseInt(groups!.throwTrue, 10),
      ifFalseThrow: parseInt(groups!.throwFalse, 10),
      operation,
    };
  });

  return {
    monkeys,
  };
}

function computeMonkeyBusiness(
  monkeys: Monkey[],
  rounds: number,
  relaxFunction: (worryLevel: number) => number
): number {
  const counters = new Array(monkeys.length).fill(0);

  for (let i = 0; i < rounds; i++) {
    for (let j = 0; j < monkeys.length; j++) {
      const monkey = monkeys[j];
      counters[j] += monkey.items.length;
      while (monkey.items.length != 0) {
        const item = monkey.items.shift();
        const newWorryLevel = monkey.operation(item!)!;
        const relaxedWorryLevel = relaxFunction(newWorryLevel);

        const monkeyToThrow =
          relaxedWorryLevel % monkey.divisibleTtest == 0
            ? monkey.ifTrueThrow
            : monkey.ifFalseThrow;

        monkeys[monkeyToThrow].items.push(relaxedWorryLevel);
      }
    }
  }

  counters.sort((a, b) => b - a);

  return counters[0] * counters[1];
}

export function part1(parsed: State): string {
  return computeMonkeyBusiness(parsed.monkeys, 20, (worryLevel) =>
    Math.floor(worryLevel / 3)
  ).toString();
}

export function part2(parsed: State): string {
  const relaxDivisor = lcm(...parsed.monkeys.map((m) => m.divisibleTtest));
  return computeMonkeyBusiness(
    parsed.monkeys,
    10_000,
    (worryLevel) => worryLevel % relaxDivisor
  ).toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
