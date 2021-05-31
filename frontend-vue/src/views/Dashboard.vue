<template>
<div>
  <v-container fluid>
  <v-row no-gutters wrap justify="space-between">
    <v-col md="6" sm="12" class="pa-2">
      <v-card >
        <v-card-title>
          <span class="d-inline-block mr-2">Tasks</span>
          <v-btn text color="success">Add Task</v-btn>
        </v-card-title>
        <v-divider></v-divider>
        <v-list>

        </v-list>
      </v-card>
    </v-col>
    <v-col md="6" sm="12" class="pa-2">
      <v-card >
        <v-card-title>Logs</v-card-title>
        <v-divider></v-divider>
        <v-list>

        </v-list>
      </v-card>
    </v-col>
  </v-row>

  </v-container>
</div>
</template>

<script>

  export default {
    name: 'Home',
    data(){
      return{
        tab
      }
    },

    components: {
    },
  }
</script>
