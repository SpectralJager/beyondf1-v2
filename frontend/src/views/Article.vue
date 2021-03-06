<template>
<v-container>
    <v-card dense flat>
        <v-row>
            <v-col cols="12" md="6">
                <v-card-title class="text-xl-h2 text-lg-h3 text-md-h4 text-sm-h2 text-h4">{{article.title}}</v-card-title>
                <v-card-subtitle class="my-1 text-lg-h5 text-subtitle-1 text-sm-h5 text-subtitle-1">{{moment(article.createdat).format('MMMM Do YYYY, h:mm:ss a')}}</v-card-subtitle>
                <v-chip class="pa-4 ml-3 white--text text-lg-h5 text-subtitle-1 text-sm-h5 text-subtitle-1" :color="article.tag">{{article.tag}}</v-chip>
                <v-btn dense text color="primary" :href="article.source">Source: {{hostname}}</v-btn>
            </v-col>
            <v-divider></v-divider>
            <v-col cols="12" md="6" class="pa-0">
                <v-img width="100%" max-height="500px" :src="article.image_url">
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
            <v-col cols="12" md="6">
                <vue-markdown :source="article.text"></vue-markdown>
            </v-col>
            <v-divider></v-divider>
            <v-col cols="12" md="6">
                <v-sheet class="articles--list">
                    <template v-for="a in articles" >
                        <v-hover v-slot="{ hover }">
                            <v-card height="750px" outlined :key="a.id" style="position:relative;" class="article-card" :dark=" hover ? true : false" :elevetion="hover ? 12 : 2">
                                <v-img height="50%" width="100%" :src="a.image_url">
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
                                <div>
                                    <v-card-title class="text-h6">{{ a.title }}</v-card-title>
                                    <v-card-subtitle class="my-1 ">
                                        <v-chip class="white--text text-subtitle-1" small dense flat :color="a.tag" >{{ a.tag }}</v-chip>
                                        <span class="ml-4 text-subtitle-1">{{ moment(a.createdat).format('MMMM Do YYYY, h:mm:ss a') }}</span>
                                    </v-card-subtitle>
                                    <v-card-text class="text-body-1">{{getSlice(a.text)}}**</v-card-text>
                                </div>
                                <v-btn color="red" class="white--text" absolute bottom left :to="'/article/' + a.id" @click.prevent="reload">Read more</v-btn>
                            </v-card>
                        </v-hover>
                    </template>
                </v-sheet>
            </v-col>
        </v-row>
    </v-card>
</v-container>
  
</template>

<script>
import VueMarkdown from '@adapttive/vue-markdown'

export default {
    data(){
        return{
            article: {},
            articles: [],
            hostname: '',
        }
    },
    components: {
        VueMarkdown
    },
    methods: {
        async fetchArticle(){
            var url = this.$store.state.backend_url + "/article/" + this.$route.params.id;
            await fetch(url, {
                method: "GET",
                mode: "cors",
                cache: "no-cache",
            })
            .then(response => response.json())
            .then((data) => {
                if(data.msg == "success"){
                    this.article = data.articles
                    this.domain() 
                }
                else {
                    this.$router.push('../articles')
                }
            });
        },
        async fetchArticles(){
            var url = this.$store.state.backend_url + "/articles/" + 1 + "/" + 2;
            await fetch(url, {
                method: "GET",
                mode: "cors",
                cache: "no-cache",
            })
            .then(response => response.json())
            .then((data) => {
                if(data.msg == "success"){
                this.articles = data.articles;
                }
            });
        },
        getSlice(text){
            return text.slice(0,150)+"...";
        },
        reload(){
            //vm.$forceUpdate();
            this.$store.commit('toggleLoad');
            this.fetchArticle();
            setTimeout(() => {
                this.$store.commit('toggleLoad');
            }, 1000)
        },
        domain() {
            var a = document.createElement('a');
            a.href = this.article.source;
            this.hostname = a.hostname;
        }
    },
    beforeCreate(){
        this.$store.commit('toggleLoad');
    },
    mounted(){
        this.fetchArticle();
        this.fetchArticles();
        setTimeout(() => {
            this.$store.commit('toggleLoad');
        }, 1000)
    }
}
</script>

<style>


</style>