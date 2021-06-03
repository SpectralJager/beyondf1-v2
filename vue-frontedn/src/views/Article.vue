<template>
<v-container>
    <v-card dense flat>
        <v-row>
            <v-col cols="12" md="6">
                <v-card-title class="text-xl-h2 text-md-h3 text-h4">{{article.title}}</v-card-title>
                <v-card-subtitle class="text-md-h4 text-title my-lg-4 my-1">{{moment(article.createdat).format('MMMM Do YYYY, h:mm:ss a')}}</v-card-subtitle>
                <v-chip class="pa-3 pa-lg-6 ml-4 white--text text-lg-h5" color="warning">{{article.tag}}</v-chip>
            </v-col>
            <v-divider></v-divider>
            <v-col cols="12" md="6" class="pa-0">
                <v-img width="100%" max-height="500px" :src="article.image_url"></v-img>
            </v-col>
            <v-col cols="12" md="6">
                <v-card-text v-html="article.text" class="text-lg-subtitle-1 text-body-1"></v-card-text>
            </v-col>
        </v-row>
    </v-card>
</v-container>
  
</template>

<script>
export default {
    data(){
        return{
            article: {},
        }
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
                console.log(data);
                if(data.msg == "success"){
                this.article = data.articles
                }
                else {
                    this.$router.push('Articles')
                }
            });
        },
    },
    mounted(){
        this.fetchArticle();
    }
}
</script>

<style>


</style>