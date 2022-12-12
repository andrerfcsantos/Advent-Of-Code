import { chunkArray } from "./arrays.ts";

const ocrLetterFingerprintToLetters = new Map<string, string>([
  [".##..#..#.#..#.####.#..#.#..#.", "A"],
  ["###..#..#.###..#..#.#..#.###..", "B"],
  [".##..#..#.#....#....#..#..##..", "C"],
  ["####.#....###..#....#....####.", "E"],
  ["####.#....###..#....#....#....", "F"],
  [".##..#..#.#....#.##.#..#..###.", "G"],
  ["#..#.#..#.####.#..#.#..#.#..#.", "H"],
  [".###...#....#....#....#...###.", "I"],
  ["..##....#....#....#.#..#..##..", "J"],
  ["#..#.#.#..##...#.#..#.#..#..#.", "K"],
  ["#....#....#....#....#....####.", "L"],
  [".##..#..#.#..#.#..#.#..#..##..", "O"],
  ["###..#..#.#..#.###..#....#....", "P"],
  ["###..#..#.#..#.###..#.#..#..#.", "R"],
  [".###.#....#.....##.....#.###..", "S"],
  ["#..#.#..#.#..#.#..#.#..#..##..", "U"],
  ["#....#.....#.#...#....#....#..", "Y"],
  ["####....#...#...#...#....####.", "Z"],
]);

export function ocrScreenToText(screenLines: string[]): string {
  const letterWidth = 5;
  const letterHeight = 6;

  const chunkedlines = screenLines.map((line) =>
    chunkArray(line.split(""), letterWidth).map((x) => x.join(""))
  );

  // transpose chunked lines
  const letters = [];
  for (let i = 0; i < chunkedlines[0].length; i++) {
    const letterLines = [];
    for (let j = 0; j < letterHeight; j++) {
      letterLines.push(chunkedlines[j][i]);
    }
    const letterFingerprint = letterLines.join("");
    const letter = ocrLetterFingerprintToLetters.get(letterFingerprint);
    if (letter) {
      letters.push(letter);
    } else {
      letters.push("?");
    }
  }

  return letters.join("");
}
