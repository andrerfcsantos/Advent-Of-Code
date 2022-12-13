import { runProblem } from "./aoclib/runner.ts";
import { groupedLines, nonEmptyLines } from "./aoclib/input.ts";
import { sumNumbers } from "./aoclib/arrays.ts";

const DAY = 13;
const YEAR = 2022;

type PacketContent = number[] | number | PacketContent[];

interface PacketPair {
  left: PacketContent;
  right: PacketContent;
}

interface State {
  packetPairs: PacketPair[];
  allPackets: PacketContent[];
}

enum PacketOrdering {
  EQUAL = "Equal",
  ORDERED = "Ordered",
  NOT_ORDERED = "Not Ordered",
}

export function parse(input: string): State {
  return {
    packetPairs: groupedLines(input)
      .map((group) => group.map((l) => JSON.parse(l)))
      .map((group) => ({ left: group[0], right: group[1] } as PacketPair)),
    allPackets: nonEmptyLines(input).map((l) => JSON.parse(l)),
  };
}

function comparePackets(
  left: PacketContent,
  right: PacketContent
): PacketOrdering {
  if (typeof left === "number" && typeof right === "number") {
    if (left === right) {
      return PacketOrdering.EQUAL;
    } else if (left < right) {
      return PacketOrdering.ORDERED;
    }
    return PacketOrdering.NOT_ORDERED;
  }

  if (typeof left === "number" && Array.isArray(right)) {
    return comparePackets([left], right);
  }
  if (typeof right === "number" && Array.isArray(left)) {
    return comparePackets(left, [right]);
  }

  if (Array.isArray(left) && Array.isArray(right)) {
    const leftLength = left.length;
    const rightLength = right.length;

    if (leftLength === 0 && rightLength != 0) {
      return PacketOrdering.ORDERED;
    }

    if (leftLength != 0 && rightLength === 0) {
      return PacketOrdering.NOT_ORDERED;
    }

    if (leftLength === 0 && rightLength === 0) {
      return PacketOrdering.EQUAL;
    }

    const leftOrdering = comparePackets(left[0], right[0]);
    if (leftOrdering === PacketOrdering.EQUAL) {
      return comparePackets(left.slice(1), right.slice(1));
    }

    return leftOrdering;
  }

  return PacketOrdering.EQUAL;
}

export function part1(parsed: State): string {
  const orderedRight = parsed.packetPairs.map((pair, idx) => {
    const comparison = comparePackets(pair.left, pair.right);
    return comparison === PacketOrdering.ORDERED ? idx + 1 : 0;
  });

  const sumIndexes = sumNumbers(orderedRight);
  return sumIndexes.toString();
}

export function part2(parsed: State): string {
  parsed.allPackets.push([[2]], [[6]]);

  const sortedPackets = parsed.allPackets.toSorted((left, right) => {
    const comparison = comparePackets(left, right);
    if (comparison === PacketOrdering.ORDERED) {
      return -1;
    } else if (comparison === PacketOrdering.NOT_ORDERED) {
      return 1;
    }
    return 0;
  });

  const twoPacket = sortedPackets.findIndex((p) => {
    return (
      Array.isArray(p) &&
      p.length === 1 &&
      Array.isArray(p[0]) &&
      p[0].length === 1 &&
      typeof p[0][0] === "number" &&
      p[0][0] === 2
    );
  });

  const sixPacket = sortedPackets.findIndex((p) => {
    return (
      Array.isArray(p) &&
      p.length === 1 &&
      Array.isArray(p[0]) &&
      p[0].length === 1 &&
      typeof p[0][0] === "number" &&
      p[0][0] === 6
    );
  });

  return ((twoPacket + 1) * (sixPacket + 1)).toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
