<template>
    <div class="hello">
      <div class="card">
        <Message :closable="false" style="width: 100%;" severity="info" icon="pi pi-send"> 
          <div class="pl-2 flex flex-column">
              <div class="flex text-lg">Username :</div>
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
      msg: String
    }
  }
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  
  