import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
    themes: {
        light: {
            breaking: '#DC2F02',
            analytics: '#0F8997',
            news: '#F48C06',
        },
    },
  },
});
