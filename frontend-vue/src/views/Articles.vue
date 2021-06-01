<template>
<div>
    <v-container>
        <v-row>
            <v-col cols="12">
                <h2>Articles</h2>
                <v-btn class="subtitle-1 pa-3 my-3" color="success" dense @click="overlay=true">
                    <v-icon left>mdi-pencil</v-icon>
                    <span>Create article</span>
                </v-btn>
                <v-divider></v-divider>
                <v-simple-table fixed-header height="600px" >
                    <thead>
                        <tr>
                            <th class="text-left">#</th>
                            <th class="text-left">Image</th>
                            <th class="text-left">Title</th>
                            <th class="text-left">Tag</th>
                            <th class="text-left">Created at</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="a in articles" :key="a.id">
                            <td>{{a.id}}</td>
                            <td><v-img width="50px" :src="a.image_url"></v-img></td>
                            <td ><span class="text-truncate">{{a.title}}</span></td>
                            <td>{{a.tag}}</td>
                            <td>{{a.createdat}}</td>
                            <td>
                                <v-btn-toggle rounded dense>
                                    <v-btn color="warning" @click="editArticle(a)"><v-icon color="white">mdi-clipboard-edit</v-icon></v-btn>
                                    <v-btn color="red" @click="deleteArticle(a)"><v-icon color="white">mdi-delete</v-icon></v-btn>
                                </v-btn-toggle>
                            </td>
                        </tr>
                    </tbody>
                </v-simple-table> 
            </v-col>
        </v-row>
        <v-dialog v-model="overlay" hide-overlay>
            <v-card class="pa-3 deep-purple lighten-5" >
                <v-card-title class="d-flex justify-space-between">
                    <span>Enter new article</span>
                    <v-btn @click="overlay=false" icon fab small color="gray"><v-icon>mdi-close</v-icon></v-btn>
                </v-card-title>
                <v-divider></v-divider>
                <v-form>
                    <v-row class="mt-3">
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="article.title" label="Title" clearable dense required class="pb-2"></v-text-field>
                            <v-text-field v-model="article.image_url" label="Image url" clearable dense required class="pb-2"></v-text-field>
                            <v-select v-model="article.tag" label="Tag" :items="tags" clearable dense required class="pb-2"></v-select>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-textarea label="Text" v-model="article.text" clearable style="max-height=500px; overflow-y: auto;overflow-x: hidden;" class="sm-6"></v-textarea>
                        </v-col>
                    </v-row>
                </v-form>
                <v-btn-toggle dense >
                    <v-btn @click="saveArticleForm" color="success">Save</v-btn>
                    <v-btn @click="resetArticleForm" color="warning">Reset</v-btn>
                    <v-btn @click="openEditor" color="primary">WYSIWYG Editor</v-btn>
                </v-btn-toggle>
            </v-card>
        </v-dialog>
    </v-container>
</div>
  
</template>

<script>
export default {
    components: {
    },
    data(){
        return{
            overlay: false,
            article: {
                id: 1,
                title: "",
                image_url: "",
                tag: "",
                text: "",
                createdat: "12.04.2000 12:00"
            },
            tags:[
                "news",
                "analytics",
                "breaking",
                "history"
            ],
            articles: [
                {   
                    id: 1,
                    title: 'Ipsum culpa minim cillum laborum mollit commodo ullamco nisi aliqua ex dolore velit ullamco.',
                    text: 'Laborum labore cillum proident labore in veniam id id dolor sint id. Ut minim officia dolore aliquip ea. Aliquip ea consectetur nulla sint Lorem sunt ut voluptate anim exercitation ut. Consequat est nisi anim sunt incididunt nulla mollit qui. Commodo consequat anim qui laboris do ipsum anim excepteur deserunt. Dolor eu enim nostrud magna anim.',
                    createdat: "12.04.2000 12:00",
                    image_url: 'https://picsum.photos/1920/1080?random',
                    tag: 'news',
                },
                {   
                    id: 2,
                    title: 'Ipsum culpa minim cillum laborum mollit commodo ullamco nisi aliqua ex dolore velit ullamco.',
                    text: 'Laborum labore cillum proident labore in veniam id id dolor sint id. Ut minim officia dolore aliquip ea. Aliquip ea consectetur nulla sint Lorem sunt ut voluptate anim exercitation ut. Consequat est nisi anim sunt incididunt nulla mollit qui. Commodo consequat anim qui laboris do ipsum anim excepteur deserunt. Dolor eu enim nostrud magna anim.',
                    createdat: "12.04.2000 12:00",
                    image_url: 'https://picsum.photos/1920/1080?random',
                    tag: 'analytics',
                },
                {   
                    id: 3,
                    title: 'Ipsum culpa minim cillum laborum mollit commodo ullamco nisi aliqua ex dolore velit ullamco.',
                    text: 'Laborum labore cillum proident labore in veniam id id dolor sint id. Ut minim officia dolore aliquip ea. Aliquip ea consectetur nulla sint Lorem sunt ut voluptate anim exercitation ut. Consequat est nisi anim sunt incididunt nulla mollit qui. Commodo consequat anim qui laboris do ipsum anim excepteur deserunt. Dolor eu enim nostrud magna anim.',
                    createdat: "12.04.2000 12:00",
                    image_url: 'https://picsum.photos/1920/1080?random',
                    tag: 'news',
                },
            ]
        }
    },
    methods:{
        resetArticleForm(){
            this.article.title="";
            this.article.tag="";
            this.article.image_url="";
            this.article.text="";
        },
        saveArticleForm(){
            this.articles.push(this.article);
            this.overlay = false;
        },
        editArticle(elem){
            this.article = elem;
            this.overlay = true;
        },
        deleteArticle(elem){
            this.articles = this.articles.filter(item => item !== elem);
        },
        openEditor(){
            window.open("https://html-online.com/editor/", "_blank");
        }
    }

}
</script>

<style lang="scss">
textarea{
    width: 100%;
    max-height: 300px;
    overflow-y: auto;
}
</style>