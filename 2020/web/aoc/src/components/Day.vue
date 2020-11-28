<template>
  <div class="day">
    <v-row>
      <v-col
          cols="12"
          sm="6"
      >
        <v-textarea
            v-model="input"
            name="input-7-1"
            label="Problem Input"
            placeholder="Paste your problem input here"
            rows="2"
        ></v-textarea>

        <v-btn
            elevation="5"
            x-large
        >Solve!</v-btn>
      </v-col>
    </v-row>

    <v-row>
      <pre> {{ input }} </pre>
    </v-row>


  </div>
</template>

<script>
export default {
  name: "Day",

  data: () => ({
    input: ""
  })
}
</script>

<style scoped>

</style>
