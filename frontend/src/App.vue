<template>
  <div id="app">
    <DigimonHeader />
    <DigimonList />
    <DigimonFooter />
  </div>
</template>

<script>
import DigimonHeader from './components/DigimonHeader.vue'
import DigimonList from './components/DigimonList/DigimonListLayout.vue';
import DigimonFooter from './components/DigimonFooter.vue'

export default {
  name: 'App',
  components: {
    DigimonHeader,
    DigimonList,
    DigimonFooter
  }
};
</script>

<style>
body {
  margin: 0px;
  padding: 10px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  display: flex;
  flex-direction: column;
}
</style>
