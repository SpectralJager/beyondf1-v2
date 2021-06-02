<template>
<v-container>
    <v-card dense flat>
        <v-row>
            <v-col cols="12" md="6">
                <v-card-title class="text-xl-h1 text-md-h2 text-h4">{{article.title}}</v-card-title>
                <v-card-subtitle class="text-md-h4 text-title my-lg-4 my-1">{{article.createdat}}</v-card-subtitle>
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
            article: {
                id: "1", 
                title: "Sit non veniam cupidatat do qui voluptate deserunt.", 
                createdat:  "12.04.2000 12:00", 
                tag: "news", 
                text: `Ut esse dolor consectetur cupidatat eiusmod commodo aliquip ea do commodo dolor ut. Fugiat ipsum tempor est velit cupidatat laborum elit incididunt consectetur. Et consectetur laborum culpa dolor reprehenderit. Est ex mollit quis exercitation cillum sit nulla aliquip sit. Lorem culpa pariatur labore occaecat sunt occaecat nulla qui aliqua et nulla veniam.<br><br>

                Ad ut quis cillum proident exercitation ullamco ea. Nulla ullamco fugiat commodo deserunt. Elit excepteur id magna id cupidatat cupidatat amet nisi laborum veniam et ea incididunt. Tempor duis sit adipisicing dolor ipsum sit qui cupidatat elit. Labore in elit non exercitation fugiat Lorem et consectetur ut. Laboris officia nulla irure id minim pariatur aute.<br><br>

                Sint ut esse sunt in Lorem do ad aute est. Elit elit velit proident dolor adipisicing. Nisi tempor ex nostrud sit nulla. Id culpa est anim excepteur ad sint ex fugiat laboris adipisicing. Occaecat nisi sunt esse velit amet esse commodo cupidatat in veniam in irure consectetur deserunt. Ipsum adipisicing excepteur enim adipisicing aute pariatur proident. <br><br>

                Dolor dolore aute magna adipisicing dolor nostrud ullamco elit laborum ex ullamco eu. Sit do minim proident cupidatat Lorem Lorem cillum fugiat officia culpa aute. Dolore officia consequat adipisicing quis dolore proident do et do. Sit fugiat ipsum aute laborum pariatur cillum laborum ipsum quis nisi veniam esse amet ullamco. Sunt sit duis aute enim voluptate incididunt excepteur. <br><br>

                Aute cupidatat tempor sit exercitation ea eiusmod sit eiusmod amet quis. Dolore ut sunt tempor velit fugiat dolor aute id sit. Nisi ex irure adipisicing id et consectetur dolore dolore ipsum est aliqua adipisicing nulla voluptate. Enim deserunt dolor laboris mollit tempor exercitation mollit adipisicing sunt labore. Nulla cillum consectetur esse magna ea.`,
                image_url: "https://images.unsplash.com/photo-1552849397-3f599d225636"
            },
        }
    }
}
</script>

<style>

</style>