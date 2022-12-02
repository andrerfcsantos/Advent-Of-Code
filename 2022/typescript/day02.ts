import { runProblem } from "./aoclib.ts";

const DAY = 2;
const YEAR = 2022;

enum Shape {
  Rock = "Rock",
  Paper = "Paper",
  Scissors = "Scissors",
}

interface Round {
  player1: Shape;
  player2: string;
}

enum RoundResult {
  Win = "Win",
  Draw = "Draw",
  Loss = "Loss",
}

const resultPoints = new Map<RoundResult, number>([
  [RoundResult.Win, 6],
  [RoundResult.Draw, 3],
  [RoundResult.Loss, 0],
]);

const letterToShape = new Map<string, Shape>([
  ["A", Shape.Rock],
  ["B", Shape.Paper],
  ["C", Shape.Scissors],
]);

// Interpretation of the second encrypted column for part 1,
// where the letter indicates the shape played by the second player
const secondColumnToShape = new Map<string, Shape>([
  ["X", Shape.Rock],
  ["Y", Shape.Paper],
  ["Z", Shape.Scissors],
]);

// Interpretation of the second encrypted column for part 2
// where the letter indicates the desired result of the round
const secondColumnToResult = new Map<string, RoundResult>([
  ["X", RoundResult.Loss],
  ["Y", RoundResult.Draw],
  ["Z", RoundResult.Win],
]);

const shapePoints = new Map<Shape, number>([
  [Shape.Rock, 1],
  [Shape.Paper, 2],
  [Shape.Scissors, 3],
]);

const xBeatsY = new Map<Shape, Shape>([
  [Shape.Rock, Shape.Scissors],
  [Shape.Paper, Shape.Rock],
  [Shape.Scissors, Shape.Paper],
]);

const xIsBeatenByY = new Map<Shape, Shape>([
  [Shape.Scissors, Shape.Rock],
  [Shape.Rock, Shape.Paper],
  [Shape.Paper, Shape.Scissors],
]);

interface State {
  rounds: Round[];
}

export function parse(lines: string[]): State {
  const rounds: Round[] = [];

  for (const line of lines) {
    if (line === "") {
      continue;
    }
    const [player1, player2] = line.split(" ");
    rounds.push({
      player1: letterToShape.get(player1)!,
      player2,
    });
  }

  return { rounds };
}

function getRoundResult(round: Round): RoundResult {
  const secondPlayerShape = secondColumnToShape.get(round.player2)!;

  if (round.player1 === secondPlayerShape) {
    return RoundResult.Draw;
  }

  if (xBeatsY.get(secondPlayerShape) === round.player1) {
    return RoundResult.Win;
  }

  return RoundResult.Loss;
}

export function part1(parsed: State): string {
  let ourScore = 0;

  for (const round of parsed.rounds) {
    const pointsForResult = resultPoints.get(getRoundResult(round))!;
    const pointsForShape = shapePoints.get(
      secondColumnToShape.get(round.player2)!
    )!;
    ourScore += pointsForResult + pointsForShape;
  }

  return ourScore.toString();
}

function shapeForDesiredResult(
  result: RoundResult,
  otherPlayerShape: Shape
): Shape {
  if (result === RoundResult.Draw) {
    return otherPlayerShape;
  }

  if (result === RoundResult.Win) {
    return xIsBeatenByY.get(otherPlayerShape)!;
  }

  return xBeatsY.get(otherPlayerShape)!;
}

export function part2(parsed: State): string {
  let ourScore = 0;

  for (const round of parsed.rounds) {
    const desiredResult = secondColumnToResult.get(round.player2)!;
    const shape = shapeForDesiredResult(desiredResult, round.player1);

    const pointsForResult = resultPoints.get(desiredResult)!;
    const pointsForShape = shapePoints.get(shape)!;
    ourScore += pointsForResult + pointsForShape;
  }

  return ourScore.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
