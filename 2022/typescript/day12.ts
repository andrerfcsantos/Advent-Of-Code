import { runProblem } from "./aoclib/runner.ts";
import { nonEmptyLines } from "./aoclib/input.ts";
import { Point } from "./aoclib/math.ts";
import { PriorityQueue } from "./aoclib/priorityqueue.ts";

const DAY = 12;
const YEAR = 2022;

interface State {
  startingPosition: Point;
  goal: Point;
  heightMap: number[][];
}

export function parse(input: string): State {
  const lines = nonEmptyLines(input);
  let startingPosition = new Point(0, 0);
  let goal = new Point(0, 0);

  const heightMap = new Array<number[]>(lines.length);

  for (let i = 0; i < lines.length; i++) {
    heightMap[i] = new Array<number>(lines[i].length).fill(0);
  }

  const a = "a".charCodeAt(0);
  const z = "z".charCodeAt(0);

  for (let i = 0; i < lines.length; i++) {
    const chars = lines[i].split("");
    for (let j = 0; j < chars.length; j++) {
      if (chars[j] === "S") {
        startingPosition = new Point(j, i);
        heightMap[i][j] = a - a;
        continue;
      }

      if (chars[j] === "E") {
        goal = new Point(j, i);
        heightMap[i][j] = z - a;
        continue;
      }

      heightMap[i][j] = chars[j].charCodeAt(0) - a;
    }
  }
  return { startingPosition, heightMap, goal };
}

// A* finds a path from start to goal.
// h is the heuristic function.
// h(n) estimates the cost to reach goal from node n.
function aStar(
  heigthMap: number[][],
  start: Point,
  goal: Point,
  h: (p: Point) => number
) {
  const queue = new PriorityQueue<Point>();
  queue.insert(start, h(start));

  const visited = new Set<string>();
  visited.add(start.toString());

  const parents = new Map<string, Point>();

  // For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
  const gScore = new Map<string, number>();
  gScore.set(start.toString(), 0);

  // f(n)=g(n)+h(n)
  const fScore = new Map<string, number>();
  fScore.set(start.toString(), h(start));

  while (!queue.isEmpty()) {
    // This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
    const current = queue.pop();
    visited.add(current!.toString());

    if (current && current.equals(goal)) {
      return reconstructPath(parents, current);
    }

    const neighbors = current!
      .cardinalNeighbors()
      .filter(
        (p) =>
          p.x >= 0 &&
          p.x < heigthMap[0].length &&
          p.y >= 0 &&
          p.y < heigthMap.length &&
          heigthMap[p.y][p.x] - heigthMap[current!.y][current!.x] <= 1
      );

    for (const neighbor of neighbors) {
      const tentative_gScore =
        (gScore.get(current!.toString()) ?? Infinity) + 1;

      if (tentative_gScore < (gScore.get(neighbor.toString()) ?? Infinity)) {
        // This path to neighbor is better than any previous one. Record it!
        parents.set(neighbor.toString(), current!);
        gScore.set(neighbor.toString(), tentative_gScore);
        fScore.set(neighbor.toString(), tentative_gScore + h(neighbor));
        if (!queue.has(neighbor, (a, b) => a.equals(b))) {
          queue.insert(neighbor, tentative_gScore + h(neighbor));
        } else {
          queue.replace(neighbor, tentative_gScore + h(neighbor), (a, b) =>
            a.equals(b)
          );
        }
      }
    }
  }

  throw new Error("No path found - goal is never reached");
}

function reconstructPath(cameFrom: Map<string, Point>, current: Point) {
  const total_path = [current];
  while (cameFrom.has(current.toString())) {
    current = cameFrom.get(current.toString())!;
    total_path.unshift(current);
  }
  return total_path;
}

export function part1(parsed: State): string {
  const { startingPosition, heightMap, goal } = parsed;

  const path = aStar(heightMap, startingPosition, goal, (p) =>
    p.taxiCabDistance(goal)
  );

  return (path.length - 1).toString();
}

export function part2(parsed: State): string {
  // check all distances from points that have height 0 to the goal

  const pointsToCheck = new Map<string, Point>();

  for (let i = 0; i < parsed.heightMap.length; i++) {
    for (let j = 0; j < parsed.heightMap[i].length; j++) {
      if (parsed.heightMap[i][j] === 0) {
        const p = new Point(j, i);
        pointsToCheck.set(p.toString(), p);
      }
    }
  }

  const distances = new Map<string, number>();

  for (const [key, point] of pointsToCheck) {
    try {
      const path = aStar(parsed.heightMap, point, parsed.goal, (p) =>
        p.taxiCabDistance(parsed.goal)
      );
      distances.set(key, path.length - 1);
    } catch (e) {}
  }

  const min = Math.min(...distances.values());

  return min.toString();
}

await runProblem<State>(DAY, YEAR, parse, part1, part2);
