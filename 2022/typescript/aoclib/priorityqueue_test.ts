import { assertEquals } from "https://deno.land/std@0.167.0/testing/asserts.ts";
import { PriorityQueue } from "./priorityqueue.ts";

Deno.test("Test queue with 3 elements", () => {
  const q = new PriorityQueue<number>();
  q.insert(2, 2);
  q.insert(1, 1);
  q.insert(3, 3);
  assertEquals(q.pop(), 1);
  assertEquals(q.pop(), 2);
  assertEquals(q.pop(), 3);
});
