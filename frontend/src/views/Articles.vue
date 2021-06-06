<template>
<v-container fluid>
    <v-sheet class="articles--list">
        <template v-for="article in articles" >
            <v-hover v-slot="{ hover }">
                <v-card height="750px" outlined :key="article.id" style="position:relative;" class="article-card" :dark=" hover ? true : false" :elevetion="hover ? 12 : 2">
                    <v-img height="50%" width="100%" :src="article.image_url">
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
                        <v-card-title class="text-h6">{{ article.title }}</v-card-title>
                        <v-card-subtitle class="my-1 ">
                            <v-chip class="white--text text-subtitle-1" small dense flat :color="article.tag" >{{ article.tag }}</v-chip>
                            <span class="ml-4 text-subtitle-1">{{ moment(article.createdat).format('MMMM Do YYYY, h:mm:ss a') }}</span>
                        </v-card-subtitle>
                        <v-card-text v-html="getSlice(article.text)" class="text-body-1"></v-card-text>
                    </div>
                    <v-btn color="red" class="white--text" absolute bottom left :to="'/article/' + article.id">Read more</v-btn>
                </v-card>
            </v-hover>
        </template>
    </v-sheet>
    <v-row width="100%">
        <v-col class="d-flex justify-center my-5" width="100%">
            <v-btn-toggle rounded>
                <v-btn color="indigo" class="white--text" large @click="decPage" :disabled="isPrev">Prev</v-btn>
                <v-btn color="primary" class="white--text" large @click="incPage" :disabled="isNext">Next</v-btn>
            </v-btn-toggle>
        </v-col>
    </v-row>
</v-container>

</template>

<script>
export default {
    data() {
        return {
            articles: [],
            next: 0,
            prev: 0,
            page: 1,
            num: 8,
            isPrev: true,
            isNext: true,
        }
    },
    methods: {
        getSlice(text){
            return text.slice(0,150)+"...";
        },
        async fetchArticles(){
            var url = this.$store.state.backend_url + "/articles/" + this.page + "/" + this.num;
            await fetch(url, {
                method: "GET",
                mode: "cors",
                cache: "no-cache",
            })
            .then(response => response.json())
            .then((data) => {
                if(data.msg == "success"){
                this.articles = data.articles;
                this.next = data.next;
                this.prev = data.prev;
                }
            });
            
        },
        incPage(){
            this.$store.commit('toggleLoad');
            if(this.next > 0){
                this.page += 1
                this.fetchArticles()
            }
            setTimeout(() => {
                    this.$store.commit('toggleLoad');
            }, 1000)
        },
        decPage(){
            this.$store.commit('toggleLoad');
            if(this.prev > 0){
                this.page -= 1
                this.fetchArticles()
            }
            setTimeout(() => {
                    this.$store.commit('toggleLoad');
            }, 1000)
        },
        cancelAutoUpdate () {
            clearInterval(this.timer);
        },

    },
    watch: {
        prev: function (n, o) {
            if(n == 0){
                this.isPrev = true
            }
            else {
                this.isPrev = false
            }
        },
        next: function (n, o) {
            if(n == 0){
                this.isNext = true
            }
            else{
                this.isNext = false
            }
        }
    },
    beforeCreate(){
        this.$store.commit('toggleLoad');
    },
    mounted(){
        this.fetchArticles();
        setTimeout(() => {
                this.$store.commit('toggleLoad');
        }, 1000)
        this.timer = setInterval(this.fetchArticles, 600000);
        
    },
    beforeDestroy () {
      this.cancelAutoUpdate();
    }

}
</script>

<style lang="scss">
.articles--list{
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    width: 100%;
}
.card-img{
    height: 50%;
}
</style>