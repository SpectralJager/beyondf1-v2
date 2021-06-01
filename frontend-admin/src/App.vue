<template>
  <v-app >
    <Header v-if="$store.state.token"/>
    <v-main>
      <v-container fluid>
        <router-view/>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Header from './components/Header.vue'

export default {
  name: 'App',
  components:{
    Header,
  },
  data() {
    return {
    }
  },
  created(){
    if(this.$store.state.token == ""){
      this.$router.push("login");
    }
    else{
      this.$router.push("dashboard");
    }
  }
  

};
</script>
