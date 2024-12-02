import { DOMParser } from "jsr:@b-fuze/deno-dom";

export interface SubmitResult {
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
