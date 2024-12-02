import { join } from "@std/path/join";
import { dirname } from "@std/path";


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

interface GetInputOptions {
  forceDownload?: boolean;
}

export async function getInput(
  day: number,
  year: number,
  options: GetInputOptions = {},
  session?: string
): Promise<GetInputResult> {
  const filename = `${year}_${day.toString(10).padStart(2, "0")}.txt`;
  const filepath = join(INPUTS_DIR, filename);

  if (!options.forceDownload && (await fileExists(filepath))) {
    return { input: await Deno.readTextFile(filepath), source: "file" };
  }

  return {
    input: await downloadInput(day, year, filepath, session),
    source: "fetch",
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

const NEWLINE_REGEX = /\r?\n/;

export function nonEmptyLines(
  input: string,
  lineSeparator: string | RegExp = NEWLINE_REGEX
): string[] {
  return input.split(lineSeparator).filter((l) => l.length > 0);
}

export function groupedLines(
  input: string,
  groupSep: string | RegExp = /\r?\n\r?\n/
): string[][] {
  return input.split(groupSep).map((g) => nonEmptyLines(g));
}
