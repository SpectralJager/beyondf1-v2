<template>
<v-row class="justify-center">
    <v-snackbar v-model="message.show" :color="message.code" timeout="5000">
        {{message.msg}}
    </v-snackbar>
    <v-col cols="12" sm="8" md="6" lg="3" >
        <v-card color="cyan lighten-4" dark>
            <v-img  src="https://picsum.photos/1920/1080?random" class="pa-5">
                <h2 class="mb-4">Login</h2>
                <v-divider></v-divider>
                <v-text-field v-model="user" label="Username" outlined background-color="indigo darken-4"></v-text-field>
                <v-text-field v-model="password" label="Password" outlined background-color="indigo darken-4" type="password"></v-text-field>
                <v-divider></v-divider>
                <v-btn color="primary" class="mt-2" @click="LogIn">Submit</v-btn>
            </v-img>
        </v-card>
    </v-col>
</v-row>
</template>

<script>
export default {
    data(){
        return {
            user: "",
            password: "",
            message: {
                msg: "",
                code: "",
                show: false,
            }
        }
    },
    methods: {
        LogIn(){
            if(this.user == "admin" && this.password == "1234"){
                this.$store.state.token = "12354";
                this.$router.push('dashboard');
            }
            else{
                this.message = {
                    msg: "Incorect username or password!",
                    code: "warning",
                    show: true,
                }
            }
        }
    }

}
</script>