import { Vector } from "./point.ts";

export const CartesianDirections = {
  UP: {
    vector: new Vector(0, -1),
    right: "RIGHT",
  },
  DOWN: {
    vector: new Vector(0, 1),
    right: "LEFT",
  },
  LEFT: {
    vector: new Vector(-1, 0),
    right: "UP",
  },
  RIGHT: {
    vector: new Vector(1, 0),
    right: "DOWN",
  },
};
