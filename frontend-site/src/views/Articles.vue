<template>
<v-container fluid>
    <v-sheet class="articles--list">
        <v-card height="600px" outlined v-for="article in articles" :key="article.id" style="position:relative;">
            <v-img height="50%" width="100%" :src="article.image_url"></v-img>
            <div>
                <v-card-title>{{ article.title }}</v-card-title>
                <v-card-subtitle class="my-1">
                    <v-chip class="white--text" small dense flat color="warning" >{{ article.tag }}</v-chip>
                    <span class="ml-4">{{ article.createdat }}</span>
                </v-card-subtitle>
                <v-card-text>{{ getSlice(article.text) + "..." }}</v-card-text>
            </div>
            <v-btn color="red" class="white--text" absolute bottom left :to="'/articles/' + article.id">Read more</v-btn>
        </v-card>
        
    </v-sheet>
    <v-row width="100%">
        <v-col class="d-flex justify-center my-5" width="100%">
            <v-btn-toggle rounded>
                <v-btn color="indigo" class="white--text" large>Prev</v-btn>
                <v-btn color="primary" class="white--text" large>Next</v-btn>
            </v-btn-toggle>
        </v-col>
    </v-row>
</v-container>
</template>

<script>
export default {
    data() {
        return {
            articles: [
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
                {id: "1", title: "Sit non veniam cupidatat do qui voluptate deserunt.", createdat:  "12.04.2000 12:00", tag: "news", text: "Irure duis laborum ea esse et amet. Commodo occaecat in labore commodo culpa amet amet. Dolore commodo proident exercitation tempor nulla eu excepteur aliquip cillum dolor esse deserunt quis. Id nulla sit in deserunt esse labore. Qui excepteur aute dolore fugiat id quis.", image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"},
            ],
        }
    },
    methods: {
        getSlice(text){
            return text.slice(100);
        }
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