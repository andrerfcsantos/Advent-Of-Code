<template>
  <div class="day01">
    <Day
      :year="2020"
      :day="1"
      title="Report Repair"
      :processInput="inp"
      :part1Solver="p1"
      :part2Solver="p2"
    />
  </div>
</template>

<script>
import Day from "@/components/Day.vue";
import { processInput, part1Solver, part2Solver } from "@/solvers/2020/day01";

export default {
  name: "Day01_2020",
  components: {
    Day
  },
  data: () => ({
    inp: processInput,
    p1: part1Solver,
    p2: part2Solver
  })
};
</script>
