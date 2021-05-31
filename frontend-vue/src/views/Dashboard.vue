<template>
<div>
  <v-container fluid>
  <v-row dense  justify-md="space-between">
    <v-col cols="12" md="6" class="pa-2">
      <v-card >
        <v-img height="150px" src="https://images.unsplash.com/photo-1586473219010-2ffc57b0d282">
          <v-card-title >
            <span class="d-block white--text">Tasks</span>
          </v-card-title>
          <v-card-subtitle class="mt-1">
            <v-text-field label="task" background-color="white" placeholder="Enter your task" color="black" outlined dense clearable v-model="task" @keypress.enter="addTask"/>
          </v-card-subtitle>
        </v-img>
        <v-divider></v-divider>
        <v-list v-if="tasks.length > 0" max-height="600px" class="overflow-y-auto">
          <v-container fluid v-for="task, index in tasks" :key="index">
            <v-list-item>
              <v-checkbox v-model="task.check" :label="task.text" />
              <v-spacer></v-spacer>
              <v-btn fab text icon small @click.prevent="deleteTask(task.text)">
                <v-icon color="red">mdi-delete</v-icon>
              </v-btn>
            </v-list-item>
            <v-divider></v-divider>
          </v-container>

        </v-list>
        <v-card-text v-else>
          Congrats, you have no tasks! Add new to continue work!!!
        </v-card-text>
      </v-card>
    </v-col>
    <v-col cols="12" md="6" class="pa-2">
      <v-card max-height="600px" class="overflow-y-auto" >
        <v-card-title>Logs</v-card-title>
        <v-spacer></v-spacer>
        <v-btn text @click.prevent="deleteLogs">
          <v-icon color="red" left>mdi-delete</v-icon>
          <span>Clear logs</span>
        </v-btn>
        <v-divider></v-divider>
        <v-card-text>
          {{ logs }}
        </v-card-text>
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
        tasks: [
          {text: "make articles", check: false},
          {text: "make subscribers", check: false}
        ],
        task: "",
        logs: "Labore est voluptate occaecat pariatur. Non culpa ea mollit velit. Incididunt labore incididunt ipsum in nulla exercitation in laboris eiusmod tempor sint tempor quis ullamco. Duis sit reprehenderit eu ullamco duis dolor sunt. Occaecat reprehenderit excepteur reprehenderit culpa reprehenderit reprehenderit ut ad non consequat id. Cupidatat consectetur incididunt eiusmod aliquip duis est in. Qui exercitation labore laborum elit minim."
      }
    },
    methods:{
      addTask(){
        //console.log(this.task)
        this.tasks.push({text: this.task, check: false});
        this.task = "";
      },
      deleteTask(text){
        //console.log(text)
        this.tasks = this.tasks.filter(item => item.text !== text);
      },
      deleteLogs(){
        this.logs = "Have no logs"
      },
    },
    

    components: {
    },
  }
</script>
