<template>
  <pokedex />
</template>

<script>
import pokedex from './components/pokedex.vue'

export default {
  name: 'App',
  components: {
    pokedex,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
</style>
