<template>
<div>
    <v-app-bar flat dark app color="#6A76AB" prominent src="https://picsum.photos/1920/1080?random">
        <v-app-bar-nav-icon @click="drawer = !drawer" class="hidden-lg-and-up"></v-app-bar-nav-icon>
        <v-app-bar-title>BeyondF1 - Admin panel</v-app-bar-title>
        <v-spacer></v-spacer>
        <v-btn color="red" @click="$store.state.token=''" href="#/admin/login">
            <span>Logout</span>
            <v-icon class="pl-1">mdi-logout</v-icon>
        </v-btn>
    </v-app-bar>
    <v-navigation-drawer app v-model="drawer" dark class="deep-purple">
        <v-list-item>
            <v-list-item-content>
                <v-list-item-title>Admin panel</v-list-item-title>
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
                {name: "Dashboard", icon: "mdi-view-dashboard", path: "/admin/dashboard"},
                {name: "Articles", icon: "mdi-text-box", path: "/admin/articles"},
                {name: "Subscribers", icon: "mdi-account-multiple-check", path: "/admin/subscribers"},
                {name: "Admin settings", icon: "mdi-account-box", path: "/admin/settings"},
            ]
        }
    }
}
</script>