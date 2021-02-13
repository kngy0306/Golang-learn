<template>
  <div class="pokedex">
    <h1>Pokedex</h1>
    <div v-for="pokemon in pokemons" :key="pokemon.id">
      <p>{{pokemon.id}}</p>
      <p>{{pokemon.name}}</p>
      <p>{{pokemon.type1}}</p>
      <p>{{pokemon.img}}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return{
      pokemons: []
    }
  },
  mounted() {
    var req = new Request("/pokedex", {
        method: "GET",
        detaType: "json",
      });
      fetch(req).then(res => {
        return res.json();
      })
      .then(items => {
        this.pokemons = items;
      })
  }
}
</script>