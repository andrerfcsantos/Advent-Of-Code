<template>
  <v-app>
    <v-app-bar class="navbar" dense app color="#0f0f23">
      <div class="d-flex align-center">
        <h2>Advent of Code solver + visualizer</h2>
      </div>
    </v-app-bar>

    <v-main class="pl-6">
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "App",

  components: {
  },

  data: () => ({
    //
  })
});
</script>

<style>

:root {
  --aoc-dark-blue: #0f0f23;
  --aoc-yellow: #ffffcc;
  --aoc-green: #009900;
  --aoc-white: #cccccc;
  --aoc-light-yellow: #fae284;
}

body, div {
  background-color: var(--aoc-light-yellow);

}

h1, h2, h3, h4 {
  color: var(--aoc-green);
}

span, p {
  color: var(--aoc-white);
}

v-main {
  background-color: var(--aoc-light-yellow);
}

.navbar {
  background-color: var(--aoc-dark-blue);
}

</style>
