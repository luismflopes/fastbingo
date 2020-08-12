<template>
  <div id="app">
    <div id="nav">
      Logo
      <router-link to="/">Home</router-link> |
      <router-link to="/games">Jogos</router-link> |
      <router-link to="/play">Play</router-link>

      PT
    </div>
    <router-view />
    <footer>footer here</footer>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: gray;
}

#nav {
  padding: 30px;
  background-color: black;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

footer {
  padding: 30px;
  background-color: black;
  color: gray;
}
</style>
