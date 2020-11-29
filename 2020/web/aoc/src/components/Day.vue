<template>
  <div class="day">
    <h2 class="day-title">Day {{ day }} - {{ title }}</h2>
    <a :href="'https://adventofcode.com/2019/day/' + day">
      [ Problem ]
    </a>
    <v-row>
      <v-col cols="8" xs="11" sm="8" md="8" lg="6" xl="6">
        <v-textarea
          v-model="input"
          name="input-7-1"
          label="Problem Input"
          placeholder="Paste your problem input here"
          rows="3"
          color="#ffff00"
          :rules="rules"
          :validate-on-blur="true"
          dark
          :error="textareaErr"
          :error-messages="textareaErrMessage"
        ></v-textarea>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="3">
        <v-btn
          :disabled="false"
          elevation="5"
          x-large
          @click="solve"
          class="solve-button align-self-start d-inline-block mr-5"
          >Solve!</v-btn
        >
        <div class="outputs">
          <span class="part-output-container" v-if="part1Out">
            <span class="part-label">
              Part 1:
            </span>
            <span class="part-output">
              {{ part1Out }}
            </span>
          </span>

          <span class="part-output-container" v-if="part2Out">
            <span class="part-label">
              Part 2:
            </span>
            <span class="part-output">
              {{ part2Out }}
            </span>
          </span>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "Day",
  props: {
    day: Number,
    title: String,
    part1Solver: Function,
    part2Solver: Function
  },
  data: () => ({
    input: "",
    part1Out: "",
    part2Out: "",
    rules: [
      function inputNotEmpty() {
        if (this.input == "") {
          return "please provide an input";
        }
        return true;
      }
    ],
    textareaErr: false,
    textareaErrMessage: ""
  }),
  methods: {
    solve: function(evt) {
      if (this.input == "") {
        this.textareaErrMessage = "Please provide an input";
        this.textareaErr = true;
        return;
      }
      this.textareaErrMessage = "";
      this.textareaErr = false;

      if (this.part1) {
        try {
          this.part1();
        } catch (e) {
          this.textareaErrMessage = e;
          this.textareaErr = true;
        }
      }

      if (this.part2) {
        try {
          this.part2();
        } catch (e) {
          this.textareaErrMessage = e;
          this.textareaErr = true;
        }
      }
    },
    part1: function() {
      this.part1Out = this.part1Solver(this.input);
    },
    part2: function() {
      this.part2Out = this.part2Solver(this.input);
    }
  }
};
</script>

<style scoped>
.outputs {
  display: inline-block;
}

.part-label {
  font-weight: bolder;
  color: var(--aoc-light-yellow-dark);
}

.part-output-container {
  display: block;
}
.solve-button {
  vertical-align: top;
}
</style>
