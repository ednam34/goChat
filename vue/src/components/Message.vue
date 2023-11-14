<template>
    <div class="hello">
        <Card>
            <template #title> UserName</template>
            <template #content>
                <p class="m-0">
                    {{ msg }}
                </p>
            </template>
        </Card>
    </div>
    
  </template>
  
  <script>
  export default {
    name: 'MessageComp',
    props: {
      msg: String
    }
  }
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  
  