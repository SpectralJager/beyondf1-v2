<template>
<div>
    <v-app-bar app dense dark>
        <v-app-bar-nav-icon @click="drawer= !drawer" class="hidden-lg-and-up">
        </v-app-bar-nav-icon>
    </v-app-bar>
    <v-navigation-drawer app v-model="drawer" dark class="bg-mixer">
        <v-list-item>
            <v-list-item-content>
                <v-list-item-title class="display-2">Beyond<span class="red--text">F1</span></v-list-item-title>
                <v-list-item-subtitle>Formula 1 infomation resource</v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list dense nav>
            <v-list-item link v-for="cat in categories" :key="cat.name" :to="cat.path">
                <v-icon left>{{ cat.icon }}</v-icon>
                <v-list-item-content>
                    {{ cat.name }}
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</div>
</template>

<script>
export default {
    data() {
        return {
            drawer: true,
            categories: [
                {name: "Home", icon: "mdi-home", path: "/"},
                {name: "Articles", icon: "mdi-newspaper-variant", path: "/articles"},
                //{name: "Contacts", icon: "mdi-contacts", path: "/contacts"},
            ]
        }
    }
}
</script>

<style lang="scss">
.v-navigation-drawer{
    background-image: url("/navbar-back.jpg");
    
}
</style>