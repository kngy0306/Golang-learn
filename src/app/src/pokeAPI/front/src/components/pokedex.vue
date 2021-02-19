<template>
  <div class="pokedex">
    <h1>Pokedex</h1>
    <div v-for="pokemon in pokemons" :key="pokemon.id">
      <p>{{ pokemon.id }}</p>
      <p>{{ pokemon.name }}</p>
      <p>{{ pokemon.type1 }}</p>
      <img :src="pokemon.img" alt="pokemon" class="image" />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pokemons: [],
    };
  },
  mounted() {
    var req = new Request("/pokedex", {
      method: "GET",
      detaType: "json",
    });
    fetch(req)
      .then((res) => {
        return res.json();
      })
      .then((items) => {
        this.pokemons = items;
      });
  },
};
</script>

<style>
.image {
  width: 10%;
}
</style>