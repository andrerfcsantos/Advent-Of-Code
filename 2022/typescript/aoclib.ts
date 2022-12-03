import { dirname, join } from "https://deno.land/std@0.166.0/path/posix.ts";
import { info } from "https://deno.land/std@0.166.0/log/mod.ts";
import { parse as argparse } from "https://deno.land/std@0.166.0/flags/mod.ts";
import { green, red } from "https://deno.land/std@0.167.0/fmt/colors.ts";
import { DOMParser } from "https://deno.land/x/deno_dom/deno-dom-wasm.ts";

const INPUTS_DIR = join("inputs");

export async function readLines(filename: string): Promise<string[]> {
  const file = await Deno.readTextFile(filename);
  return file.split(/\r?\n/);
}

export async function fetchInput(
  day: number,
  year: number,
  session?: string
): Promise<string> {
  if (!session) {
    session = Deno.env.get("AOC_SESSION");
  }

  const url = `https://adventofcode.com/${year}/day/${day}/input`;
  const response = await fetch(url, {
    headers: {
      Cookie: `session=${session}`,
    },
  });
  return response.text();
}

export async function fileExists(filename: string): Promise<boolean> {
  try {
    const fileInfo = await Deno.stat(filename);
    return fileInfo.isFile;
  } catch (e) {
    if (e instanceof Deno.errors.NotFound) {
      return false;
    }
  }

  return false;
}

interface GetInputResult {
  source: "file" | "fetch";
  input: string;
}

export async function getInput(
  day: number,
  year: number,
  session?: string
): Promise<GetInputResult> {
  const filename = `${year}_${day.toString(10).padStart(2, "0")}.txt`;
  const filepath = join(INPUTS_DIR, filename);

  if (await fileExists(filepath)) {
    return { input: await Deno.readTextFile(filepath), source: "file" };
  }

  return {
    input: await downloadInput(day, year, filepath, session),
    source: "fetch",
  };
}

interface SubmitResult {
  statusCode: number;
  message: string;
  fullText: string;
}

export async function submitAnswer(
  year: number,
  day: number,
  level: number,
  answer: string,
  session?: string
): Promise<SubmitResult> {
  if (!session) {
    session = Deno.env.get("AOC_SESSION");
  }

  const reply = await fetch(
    `https://adventofcode.com/${year}/day/${day}/answer`,
    {
      method: "POST",
      headers: {
        cookie: `session=${session}`,
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams([
        ["level", level.toString()],
        ["answer", answer],
      ]),
    }
  );
  const fullText = await reply.text();
  const doc = new DOMParser().parseFromString(fullText, "text/html");
  const p = doc?.querySelector("main > article > p");

  const message = p?.textContent
    .split("\n")
    .map((s) => s.trim())
    .filter((s) => s.length > 0)
    .join(" ");

  return {
    statusCode: reply.status,
    message: message ?? "Unknown error",
    fullText,
  };
}

export async function downloadInput(
  day: number,
  year: number,
  filepath: string,
  session?: string
): Promise<string> {
  const input = await fetchInput(day, year, session);

  const dirPath = dirname(filepath);
  await Deno.mkdir(dirPath, { recursive: true });
  await Deno.writeTextFile(filepath, input, { create: true });

  return input;
}

export async function runProblem<StateType>(
  day: number,
  year: number,
  parse: (lines: string) => StateType,
  part1: (state: StateType) => string,
  part2: (state: StateType) => string,
  session?: string
): Promise<void> {
  performance.mark("start download");
  const { input, source } = await getInput(day, year, session);
  performance.mark("end download");

  performance.mark("start parse");
  const state = parse(input);
  performance.mark("end parse");

  performance.mark("start part1");
  const p1 = part1(state);
  performance.mark("end part1");

  performance.mark("start part2");
  const p2 = part2(state);
  performance.mark("end part2");

  const perf_download = performance.measure(
    "download",
    "start download",
    "end download"
  );
  const perf_parse = performance.measure("parse", "start parse", "end parse");
  const perf_part1 = performance.measure("part1", "start part1", "end part1");
  const perf_part2 = performance.measure("part2", "start part2", "end part2");

  const sourceMessage = source === "file" ? "Read from file" : "Download";

  info(
    `${sourceMessage} in ${perf_download.duration.toFixed(
      2
    )}ms | Input parsing in ${perf_parse.duration.toFixed(2)}ms`
  );
  info(`Part 1: ${p1} (${perf_part1.duration.toFixed(2)}ms)`);
  info(`Part 2: ${p2} (${perf_part2.duration.toFixed(2)}ms)`);

  const parsedArgs = argparse(Deno.args, { string: ["submit"] });

  let submitResult: SubmitResult | undefined;

  switch (parsedArgs.submit) {
    case "p1":
      info(`Submitting part 1 answer: ${p1}`);
      submitResult = await submitAnswer(year, day, 1, p1, session);
      break;
    case "p2":
      info(`Submitting part 2 answer: ${p2}`);
      submitResult = await submitAnswer(year, day, 2, p2, session);
      break;
    default:
      break;
  }

  if (submitResult) {
    const success = submitResult.message.startsWith("That's the right answer!");

    if (success) {
      console.log(
        green(`(${submitResult.statusCode}) ${submitResult.message}`)
      );
    } else {
      console.log(red(`(${submitResult.statusCode}) ${submitResult.message}`));
      await Deno.writeTextFileSync("submit_error.html", submitResult.fullText);
    }
  }
}

export function sumNumbers(ns: number[]): number {
  return ns.reduce((a, b) => a + b, 0);
}

export function intersectSets<T>(...sets: Set<T>[]): Set<T> {
  const result = new Set<T>();
  for (const set of sets) {
    for (const item of set) {
      if (sets.every((s) => s.has(item))) {
        result.add(item);
      }
    }
  }
  return result;
}

export function chunkArray<T>(array: T[], chunkSize: number): T[][] {
  const chunks = [];
  for (let i = 0; i < array.length; i += chunkSize) {
    chunks.push(array.slice(i, i + chunkSize));
  }
  return chunks;
}

export function isUppercase(c: string): boolean {
  return c === c.toUpperCase();
}

const NEWLINE_REGEX = /\r?\n/;

export function nonEmptyLines(input: string): string[] {
  return input.split(NEWLINE_REGEX).filter((l) => l.length > 0);
}

export function groupedLines(
  input: string,
  groupSep: string | RegExp = /\r?\n\r?\n/
): string[][] {
  return input.split(groupSep).map(nonEmptyLines);
}
