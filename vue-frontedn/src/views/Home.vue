<template>
<div>
  <v-card dense dark >
    <v-row style="background-image: url('/home-bg.jpg'); height: 100vh;" class="bg-gray overflow-hidden">
        <v-container class=" d-flex align-center">
          <v-col cols="12" md="6" offset-md="3">
            <div class="bg-gray white--text">
              <v-card-title class="text-xl-h1 text-h4 text-weight-black">Welcome to the <br> BeyondF1</v-card-title>
              <v-card-subtitle class="text-xl-h6 my-3 text-sm-subtitle-2">Your guide through of the Ocean of empty talks to the Island of meaning!</v-card-subtitle>
              <v-card-text class="text-xl-subtitle-1 text-sm-body-1">Qui reprehenderit non qui incididunt occaecat laborum. Dolor excepteur amet occaecat velit exercitation. Sunt tempor incididunt aliquip ut nostrud laboris nulla. Laborum velit nostrud sunt in tempor eiusmod amet est fugiat consectetur commodo ullamco.</v-card-text>
            </div>
          </v-col>
        </v-container>
    </v-row>
    <v-divider></v-divider>
    <v-row class="hidden-md-and-down">
      <v-col cols="12" md="6" class="pr-0">
        <v-hover>
          <template v-slot:default="{hover}">
            <v-card>
              <img :src="event.next.img_url" class="overflow-hidden"></v-img>
              <v-fade-transition>
                <v-overlay v-if="hover" absolute class="pa-5">
                  <v-container>
                    <h3>Next Event: {{ event.next.title }}</h3>
                    <p class="subtitle-1">At: {{ event.next.date }}</p>
                    <p class="body text-left"> {{ event.next.text }}</p>
                    <v-btn :href="event.next.href" color="red">Read more</v-btn>
                  </v-container>     
                </v-overlay>
              </v-fade-transition>
            </v-card>
          </template>
        </v-hover>
      </v-col>
      <v-col cols="12" md="6" class="pl-0">
        <v-hover>
          <template v-slot:default="{hover}">
            <v-card>
              <v-img :src="event.prev.img_url" width="100%" class="overflow-hidden"></v-img>
              <v-fade-transition>
                <v-overlay v-if="hover" absolute class="pa-5">
                  <v-container>
                    <h3>Prev Event: {{ event.prev.title }}</h3>
                    <p class="subtitle-1">At: {{ event.prev.date }}</p>
                    <p class="body text-left"> {{ event.prev.text }}</p>
                    <v-btn :href="event.prev.href" color="red">Read more</v-btn>
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
        <v-img :src="event.next.img_url" width="100%" class="overflow-hidden"></v-img>
      </v-col>
      <v-col cols="12" md="6" class="pa-8 pr-sm-0">
        <v-container>
          <h3>Next Event: {{ event.next.title }}</h3>
          <p class="subtitle-1">At: {{ event.next.date }}</p>
          <p class="body text-left"> {{ event.next.text }}</p>
          <v-btn :href="event.next.href" color="red">Read more</v-btn>
        </v-container>
      </v-col>
    </v-row>
    <v-divider></v-divider>
    <v-row class="hidden-lg-and-up">
      <v-col cols="12" class="hidden-md-and-up pl-0">
        <v-img :src="event.prev.img_url" width="100%" class="overflow-hidden"></v-img>
      </v-col>
      <v-col cols="12" md="6" class="pa-8 pr-sm-0">
        <v-container>
          <h3>Prev Event: {{ event.prev.title }}</h3>
          <p class="subtitle-1">At: {{ event.prev.date }}</p>
          <p class="body text-left"> {{ event.prev.text }}</p>
          <v-btn :href="event.prev.href" color="red">Read more</v-btn>
        </v-container>
      </v-col>
      <v-col cols="12" md="6" class="hidden-sm-and-down pl-0">
        <v-img :src="event.prev.img_url" width="100%" class="overflow-hidden"></v-img>
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
        event: {
          next: {
            title: "Monaco GP",
            date: "12.04.2000",
            text: "Esse pariatur proident pariatur aute ad ipsum excepteur labore fugiat nisi reprehenderit. Nostrud non tempor do nulla. Consequat exercitation aliquip qui ad. Ullamco ut laborum id ut et dolor quis excepteur eu. Sit quis sint enim aliqua sit non dolor duis non nisi aute sint. Voluptate culpa occaecat laboris ad dolor cillum ea proident.",
            href: "https://www.formula1.com/en/racing/2021/Monaco.html",
            img_url: "https://www.formula1.com/content/dam/fom-website/2018-redesign-assets/Racehub%20header%20images%2016x9/Monaco.jpg.transform/fullbleed/image.jpg",
          },
          prev: {
            title: "Spanish GP",
            date: "10.04.2000",
            text: "Esse pariatur proident pariatur aute ad ipsum excepteur labore fugiat nisi reprehenderit. Nostrud non tempor do nulla. Consequat exercitation aliquip qui ad. Ullamco ut laborum id ut et dolor quis excepteur eu. Sit quis sint enim aliqua sit non dolor duis non nisi aute sint. Voluptate culpa occaecat laboris ad dolor cillum ea proident.",
            href: "https://www.formula1.com/en/racing/2021/Spain.html",
            img_url: "https://www.formula1.com/content/dam/fom-website/2018-redesign-assets/Racehub%20header%20images%2016x9/Spain.jpg.transform/fullbleed/image.jpg",
          }
        }
      }
    },

    components: {
    },
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
