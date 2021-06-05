<template>
<div>
  <v-card dense dark >
    <v-row style="background-image: url('/home-bg.jpg'); height: 100vh;" class="bg-gray overflow-hidden">
        <v-container class=" d-flex align-center">
          <v-col cols="12" md="6" sm='8' offset-md="3" offset-sm='2'>
            <div class="bg-gray white--text">
              <v-card-title class="text-xl-h1 text-h3 text-weight-black">Welcome to the <br> BeyondF1</v-card-title>
              <v-card-subtitle class="text-xl-h6 my-3 text-subtitle-1">Your guide through of the Ocean of empty talks to the Island of meaning!</v-card-subtitle>
              <v-card-text class="text-xl-subtitle-1 text-body-1"></v-card-text>
            </div>
          </v-col>
        </v-container>
    </v-row>
    <v-divider></v-divider>
    <v-row class="hidden-md-and-down">
      <v-col cols="12" md="6" class="pa-0" v-for="event, index in events" :key="index">
        <v-hover>
          <template v-slot:default="{hover}">
            <v-card>
              <img :src="event.image_url" class="overflow-hidden">
                <template v-slot:placeholder>
                    <v-row
                    class="fill-height ma-0"
                    align="center"
                    justify="center"
                    >
                    <v-progress-circular
                        indeterminate
                        color="grey lighten-5"
                    ></v-progress-circular>
                    </v-row>
                </template>             
              </v-img>
              <v-fade-transition>
                <v-overlay v-if="hover" absolute class="pa-5">
                  <v-container>
                    <h3>{{ index == 0 ? 'Next': 'Previus' }} event: {{ event.name }}</h3>
                    <p class="subtitle-1">At: {{ moment(event.date_start).format('MMMM Do')+ ' - ' +moment(event.date_end).format('MMMM Do') }}</p>
                    <p class="body text-left"> {{ event.description}}</p>
                    <v-btn :href="event.link" color="red">Read more</v-btn>
                  </v-container>     
                </v-overlay>
              </v-fade-transition>
            </v-card>
          </template>
        </v-hover>
      </v-col>
    </v-row>
    <v-row class="hidden-lg-and-up">
      <v-col cols="12" md="6" class="pr-0">
        <v-img :src="events[0].image_url" width="100%" class="overflow-hidden">
          <template v-slot:placeholder>
              <v-row
              class="fill-height ma-0"
              align="center"
              justify="center"
              >
              <v-progress-circular
                  indeterminate
                  color="grey lighten-5"
              ></v-progress-circular>
              </v-row>
          </template>
        </v-img>
      </v-col>
      <v-col cols="12" md="6" class="pa-8 pr-sm-0">
        <v-container>
          <h3>Next Event: {{ events[0].name }}</h3>
          <p class="subtitle-1">At: {{ moment(events[0].date_start).format('MMMM Do') + ' - ' +moment(events[0].date_end ).format('MMMM Do') }}</p>
          <p class="body text-left"> {{ events[0].description }}</p>
          <v-btn :href="events[0].link" color="red">Read more</v-btn>
        </v-container>
      </v-col>
    </v-row>
    <v-divider></v-divider>
    <v-row class="hidden-lg-and-up">
      <v-col cols="12" class="hidden-md-and-up pl-0">
        <v-img :src="events[1].image_url" width="100%" class="overflow-hidden">
          <template v-slot:placeholder>
              <v-row
              class="fill-height ma-0"
              align="center"
              justify="center"
              >
              <v-progress-circular
                  indeterminate
                  color="grey lighten-5"
              ></v-progress-circular>
              </v-row>
          </template>
        </v-img>
      </v-col>
      <v-col cols="12" md="6" class="pa-8 pr-sm-0">
        <v-container>
          <h3>Next Event: {{ events[1].name }}</h3>
          <p class="subtitle-1">At: {{ moment(events[1].date_start).format('MMMM Do') + ' - ' +moment(events[1].date_end).format('MMMM Do') }}</p>
          <p class="body text-left"> {{ events[1].description }}</p>
          <v-btn :href="events[1].link" color="red">Read more</v-btn>
        </v-container>
      </v-col>
      <v-col cols="12" md="6" class="hidden-sm-and-down pl-0">
        <v-img :src="events[1].image_url" width="100%" class="overflow-hidden">
          <template v-slot:placeholder>
              <v-row
              class="fill-height ma-0"
              align="center"
              justify="center"
              >
              <v-progress-circular
                  indeterminate
                  color="grey lighten-5"
              ></v-progress-circular>
              </v-row>
          </template>
        </v-img>
      </v-col>
    </v-row>
  </v-card>
</div>

</template>

<script>

  export default {
    name: 'Home',
    data(){
      return{
        hover: false,
        events: [],  
      }   
    },
    methods: {
        async fetchEvents(){
            var url = this.$store.state.backend_url + "/events";
            await fetch(url, {
                method: "GET",
                mode: "cors",
                cache: "no-cache",
            })
            .then(response => response.json())
            .then((data) => {
                if(data.msg == "success"){
                this.events = data.events;
                }
            });
            
        },
    },

    components: {
    },
    mounted(){
      this.$store.commit('toggleLoad');
      this.fetchEvents();
      setTimeout(() => {
        this.$store.commit('toggleLoad');
      }, 1000)
    }
  }
</script>

<style lang="scss">
.bg-gray{
  background-position: center;
  background-size: cover;
  background-color: rgba(gray, 0.55);
  box-shadow: 0px 0px 80px gray;
  background-blend-mode: multiply;
}

</style>
