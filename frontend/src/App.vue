<template>
  <div id="app">
    <DigimonTitle />
    <DigimonList />
  </div>
</template>

<script>
import DigimonList from './components/DigimonList.vue';
import DigimonTitle from './components/DigimonTitle.vue'

export default {
  name: 'App',
  components: {
    DigimonTitle,
    DigimonList
  }
};
</script>

<style>
body {
  margin: 0px;
  padding: 10px;
  /* background-color: #3a3b3c; */
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  /* background-color: #3a3b3c; */
}

@font-face {
  font-family: 'Digivolve';
  src: url('@/assets/fonts/Digivolve.woff2') format('woff2'),
    url('@/assets/fonts/Digivolve.woff') format('woff');
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: 'Digivolve';
  src: url('@/assets/fonts/DigivolveItalic.woff2') format('woff2'),
    url('@/assets/fonts/DigivolveItalic.woff') format('woff');
  font-weight: normal;
  font-style: italic;
}
</style>
