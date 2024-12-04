import { DOMParser, Element } from "jsr:@b-fuze/deno-dom";
import { parseArgs } from "@std/cli/parse-args";
import {
  red,
  gray,
  green,
  bold,
  rgb24,
} from "@std/fmt/colors";

const args = parseArgs(Deno.args, { string: ["year"] });
const year = parseInt(args.year || "0", 10) || new Date().getFullYear();
const reply = await fetch(`https://adventofcode.com/${year}/stats`);
const html = await reply.text();

const parser = new DOMParser().parseFromString(html, "text/html");

const anchors = parser!.querySelectorAll(
  "main > pre > a[href]"
) as unknown as Element[];

const data = Array.from(anchors).map((a) => {
  const href: string = a.getAttribute("href")!;
  const match = href.match(/\/(\d+)\/day\/(\d+)/);
  if (!match) {
    return null;
  }
  const [, year, day] = match;

  const both = parseInt(
    a.getElementsByClassName("stats-both")[0].textContent.trim(),
    10
  );

  const firstOnly = parseInt(
    a.getElementsByClassName("stats-firstonly")[0].textContent.trim(),
    10
  );

  const firstOnlyPercent = firstOnly / (both + firstOnly);

  return {
    year: parseInt(year),
    day: parseInt(day),
    both,
    firstOnly,
    firstOnlyPercent,
    total: both + firstOnly,
  };
});

data.sort((a, b) => {
  if (a.year < b.year) {
    return -1;
  }
  if (a.year > b.year) {
    return 1;
  }
  if (a.day < b.day) {
    return -1;
  }
  if (a.day > b.day) {
    return 1;
  }
  return 0;
});

const hydrated = data.map((d, i, arr) => {
  if (i === 0) {
    return { ...d, dropoffTotal: 1, dropoffFirst: 1, dropoffSecond: 1 };
  }
  const prevTotal = arr[i - 1]!.total;
  const diffTotal = prevTotal - d!.total;
  const dropoffTotal = -diffTotal / prevTotal;

  const prevFirst = arr[i - 1]!.firstOnly;
  const diffFirst = prevFirst - d!.firstOnly;
  const dropoffFirst = -diffFirst / prevFirst;

  const prevSecond = arr[i - 1]!.both;
  const diffSecond = prevSecond - d!.both;
  const dropoffSecond = -diffSecond / prevSecond;

  return { ...d, dropoffTotal, dropoffFirst, dropoffSecond };
});

hydrated.forEach((d) => {
  const totalVals = `${green(d.total!.toString())} ${formatPercentChange(
    d.dropoffTotal
  )}`;

  const bothValues = `${rgb24(
    d.both!.toString(),
    0xffff66
  )} ${formatPercentChange(d.dropoffSecond)}`;

  console.log(
    `${bold(d.day!.toString().padStart(2))} ${totalVals.padEnd(
      35
    )} ${bothValues.padEnd(49)} ${rgb24(
      d.firstOnly!.toString(),
      0x9999cc
    )} ${formatPercent(d.firstOnlyPercent!)}`
  );
});

function formatPercent(decimal: number): string {
  const percent = decimal * 100;

  return gray(`(${percent.toFixed(1)}%)`);
}

function formatPercentChange(decimal: number): string {
  const percent = decimal * 100;

  if (percent > 0) {
    return green(`+${percent.toFixed(1)}%`);
  }
  if (percent < 0) {
    return red(`${percent.toFixed(1)}%`);
  }
  return gray(`${percent.toFixed(1)}%`);
}
