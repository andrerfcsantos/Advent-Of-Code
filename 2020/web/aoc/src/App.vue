<template>
  <v-app>
    <v-app-bar dense app class="navbar" color="#0f0f23">
      <router-link to="/">
        <div class="d-flex align-center">
          <h2>Advent of Code Solver / Visualizer</h2>
        </div>
      </router-link>
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

  components: {},

  data: () => ({
    //
  })
});
</script>

<style>
/*
  GLOBAL STYLES
 */
:root {
  --aoc-dark-blue: #0f0f23;
  --aoc-darklight-blue: #2c2e48;
  --aoc-yellow: #ffffcc;
  --aoc-green: #009900;
  --aoc-light-green: #99ff99;
  --aoc-white: #cccccc;
  --aoc-light-yellow: #fae284;
  --aoc-light-yellow-dark: #e0bc1c;
  --aoc-black: #000000;
}

body,
div {
  background-color: var(--aoc-darklight-blue);
}

h1,
h2,
h3,
h4 {
  color: var(--aoc-light-yellow-dark);
}

span,
p {
  color: var(--aoc-white);
}

.v-application a {
  text-decoration: none;
}

.v-application a:link {
  color: var(--aoc-green);
}

.v-application a:visited {
  color: var(--aoc-green);
}

.v-application a:focus {
  color: var(--aoc-light-green);
}

.v-application a:hover {
  color: var(--aoc-light-green);
}

.v-application a:active {
  color: var(--aoc-green);
}
/*
  NAVBAR
 */
.navbar {
  background-color: var(--aoc-dark-blue);
}

.navbar div {
  background-color: var(--aoc-dark-blue);
}

.navbar a {
  text-decoration: none;
}

.navbar h1,
.navbar h2,
.navbar h3,
.navbar h4 {
  color: var(--aoc-green);
}
</style>
