import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";

const DAY = 2;
const YEAR = 2022;

enum Shape {
  Rock = "Rock",
  Paper = "Paper",
  Scissors = "Scissors",
}

interface Round {
  otherPlayer: Shape;
  me: string;
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

export function parse(input: string): State {
  const lines = nonEmptyLines(input);

  return {
    rounds: lines.map((l) => {
      const [otherPlayer, me] = l.split(" ");
      return {
        otherPlayer: letterToShape.get(otherPlayer)!,
        me,
      };
    }),
  };
}

function getRoundResult(round: Round): RoundResult {
  const secondPlayerShape = secondColumnToShape.get(round.me)!;

  if (round.otherPlayer === secondPlayerShape) {
    return RoundResult.Draw;
  }

  if (xBeatsY.get(secondPlayerShape) === round.otherPlayer) {
    return RoundResult.Win;
  }

  return RoundResult.Loss;
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

export function part1(parsed: State): string {
  let myScore = 0;

  for (const round of parsed.rounds) {
    const pointsForResult = resultPoints.get(getRoundResult(round))!;
    const pointsForShape = shapePoints.get(secondColumnToShape.get(round.me)!)!;
    myScore += pointsForResult + pointsForShape;
  }

  return myScore.toString();
}

export function part2(parsed: State): string {
  let myScore = 0;

  for (const round of parsed.rounds) {
    const desiredResult = secondColumnToResult.get(round.me)!;
    const shape = shapeForDesiredResult(desiredResult, round.otherPlayer);

    const pointsForResult = resultPoints.get(desiredResult)!;
    const pointsForShape = shapePoints.get(shape)!;
    myScore += pointsForResult + pointsForShape;
  }

  return myScore.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
