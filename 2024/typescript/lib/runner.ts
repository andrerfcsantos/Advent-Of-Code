import { info } from "@std/log";
import { parseArgs } from "@std/cli";
import { green, red } from "@std/fmt/colors";
import { getInput } from "./input.ts";
import { submitAnswer, SubmitResult } from "./submit.ts";

export async function runProblem<StateType>(
  day: number,
  year: number,
  parse: (lines: string) => StateType,
  part1: (state: StateType) => string,
  part2: (state: StateType) => string,
  session?: string
): Promise<void> {
  const parsedArgs = parseArgs(Deno.args, {
    string: ["submit"],
    boolean: ["force-download"],
  });

  performance.mark("start download");
  const { input, source } = await getInput(
    day,
    year,
    { forceDownload: parsedArgs["force-download"] },
    session
  );
  performance.mark("end download");

  performance.mark("start parse p1");
  const state = parse(input);
  performance.mark("end parse p1");

  performance.mark("start part1");
  const p1 = part1(state);
  performance.mark("end part1");

  performance.mark("start parse p2");
  const newState = parse(input);
  performance.mark("end parse p2");

  performance.mark("start part2");
  const p2 = part2(newState);
  performance.mark("end part2");

  const perf_download = performance.measure(
    "download",
    "start download",
    "end download"
  );
  const perf_parse_p1 = performance.measure(
    "parse",
    "start parse p1",
    "end parse p1"
  );
  const perf_parse_p2 = performance.measure(
    "parse",
    "start parse p2",
    "end parse p2"
  );
  const perf_part1 = performance.measure("part1", "start part1", "end part1");
  const perf_part2 = performance.measure("part2", "start part2", "end part2");

  const sourceMessage = source === "file" ? "Read from file" : "Download";

  info(`${sourceMessage} in ${perf_download.duration.toFixed(2)}ms`);
  info(
    `Part 1: ${p1} (${perf_part1.duration.toFixed(
      2
    )}ms) | Input parsing in ${perf_parse_p1.duration.toFixed(2)}ms`
  );
  info(
    `Part 2: ${p2} (${perf_part2.duration.toFixed(
      2
    )}ms) | Input parsing in ${perf_parse_p2.duration.toFixed(2)}ms`
  );

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
