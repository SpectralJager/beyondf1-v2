<template>
  <v-app>
    <Header />
    <v-content app>
      <v-container style="padding: 0;">
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Header from './components/Header.vue';
export default {
  components: {
    Header,
  }
  
}
</script>
