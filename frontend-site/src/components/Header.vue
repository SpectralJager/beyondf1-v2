<template>
<div>
    <v-fab-transition >
        <v-btn @click="drawer = !drawer" absolute top right color="primary" fab small class="hidden-lg-and-up mt-9" style="z-index: 4;"><v-icon>mdi-menu</v-icon></v-btn>
    </v-fab-transition>
    <v-navigation-drawer app v-model="drawer" dark>
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
                {name: "Contacts", icon: "mdi-contacts", path: "/contacts"},
            ]
        }
    }
}
</script>

<style lang="scss">
.v-navigation-drawer{
    background-image: url("/navbar-back.jpg");
    background-position: center;
    background-size: cover;
    background-color: blue;
    background-blend-mode: multiply;
}
</style>