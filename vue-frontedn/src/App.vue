<template>
  <v-app style="overflow-x: hidden;">
    <Header />
    <v-main app>
      <v-container fluid style="padding: 0;">
        <router-view></router-view>
      </v-container>
    <Footer />    
    </v-main>
  </v-app>
</template>

<script>
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';

export default {
  components: {
    Header,
    Footer,
  }
  
}
</script>

<style lang="scss">
.v-main{
  //background: linear-gradient(75deg, #ebf4f5, #b5c6e0);
}
.bg-mixer{
  background-position: center;
  background-size: cover;
  background-color: blue;
  background-blend-mode: multiply;
}
.v-card__text, .v-card__title {
  word-break: normal; /* maybe !important  */
}
img{
    max-width: 100%;
}
</style>
