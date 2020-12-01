<template>
  <div class="day01">
    <Day
      :day="1"
      title="The Tyranny of the Rocket Equation"
      :part1Solver="p1"
      :part2Solver="p2"
    />
  </div>
</template>

<script>
import Day from "@/components/Day.vue";
import { part1Solver, part2Solver } from "@/solvers/2019/day01_ts";

export default {
  name: "Day01",
  components: {
    Day
  },
  data: () => ({
    p1: part1Solver,
    p2: part2Solver
  })
};
</script>
