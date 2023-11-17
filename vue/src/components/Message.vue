<template>
    <div class="hello">
      <div class="card">
        <Message :closable="false" style="width: 100%;" :severity="type" :icon="icon"> 
          <div class="pl-2 flex flex-column">
              <div class="flex text-lg">{{userName}} :</div>
              <div class="flex align-left">{{ msg }}</div>
            </div>
        </Message>
     </div>
    </div>
    
  </template>
  
  <script>

  import Message from 'primevue/message';
  export default {
    name: 'MessageComp',
    components:{
      Message,
    },
    props: {
      type: String,
      icon: String,
      msg: String,
      userName: String,
      
    }
  }
  
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  
  