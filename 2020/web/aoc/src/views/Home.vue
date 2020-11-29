<template>
  <div class="pt-4 home">
    <h2>2020</h2>
    <span>Not started yet</span>
    <h2>2019</h2>
    <ul>
      <li>
        <router-link to="/2019/day/1">Day 01</router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {}
})
export default class Home extends Vue {}
</script>

<style scoped>
ul {
  color: var(--aoc-light-yellow-dark);
}
.v-application a {
  text-decoration: none;
}

.v-application a:link {
  color: var(--aoc-light-yellow-dark);
}

.v-application a:visited {
  color: var(--aoc-light-yellow-dark);
}

.v-application a:focus {
  color: var(--aoc-light-yellow);
}

.v-application a:hover {
  color: var(--aoc-light-yellow);
}

.v-application a:active {
  color: var(--aoc-light-yellow-dark);
}
</style>
