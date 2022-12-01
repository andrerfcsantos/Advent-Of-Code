import { dirname, join } from "https://deno.land/std@0.166.0/path/posix.ts";
import { info } from "https://deno.land/std@0.166.0/log/mod.ts";

const INPUTS_DIR = join("inputs");

export async function readLines(filename: string): Promise<string[]> {
    const file = await Deno.readTextFile(filename);
    return file.split(/\r?\n/);
}

export async function fetchInput(day: number, year: number, session?: string): Promise<string> {
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

export async function downloadInput(day: number, year: number, session?: string): Promise<string> {
    const filename = `${year}_${day.toString(10).padStart(2, '0')}.txt`;
    const filepath = join(INPUTS_DIR, filename);

    if (await fileExists(filepath)) {
        info(`Using cached input for ${year} day ${day}`);
        return await Deno.readTextFile(filepath);
    }
    
    const input = await fetchInput(day, year, session);

    const dirPath = dirname(filepath);
    await Deno.mkdir(dirPath, { recursive: true });
    await Deno.writeTextFile(filepath, input, { create: true });

    return input;
}
