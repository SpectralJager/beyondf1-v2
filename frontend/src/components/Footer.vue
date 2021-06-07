<template>
    <v-footer
    dark
    padless
    width="100%"
  >
    <v-card
      flat
      tile
      class="white--text text-center"
      width="100%"
    >
      <v-card-text>
        <v-btn
          v-for="link in links"
          :key="link.icon"
          class="mx-4 white--text"
          :href="link.href"
          icon
        >
          <v-icon size="24px">
            {{ link.icon }}
          </v-icon>
        </v-btn>
      </v-card-text>

      <v-card-text class="white--text pt-0">
      </v-card-text>

      <v-divider></v-divider>
      <v-card-text class="white--text">
        {{ new Date().getFullYear() }} — <strong>SpectralJager</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  data(){
    return{
      links: [
        {icon: "mdi-twitter", href: "https://twitter.com/SpectrumUnkn"},
        {icon: "mdi-linkedin", href: "https://www.linkedin.com/in/spectraljager/"},
        {icon: "mdi-email", href: "mailto:unkn.spectrum@gmail.com"}
      ]
    }
  }
}
</script>