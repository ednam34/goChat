<template>
  <div>
    <HelloWorld msg="Welcome to The Chat App"/>
    <div class="p-3">
    <ScrollPanel ref="scrollPanel" class="overflow-auto main-chat-cont" scrollYRatio={0.99} style="width: 100%; height: 600px">
      <MessageComp v-for="(message, index) in messages" :key="index" :msg="message" />
    </ScrollPanel>
  </div>
    
    <div class="container">
      <Toolbar class="mt-3 w-full">
        <template #start>
          <p class="text-sm text-white mb-0 sm:mt-0">Votre message:</p>
        </template>
        <template #center>
          <div class="w-full w-auto-md">
            <InputText type="text" v-model="value" class="w-full w-auto-md" ref="input"/>
          </div>
        </template>
        <template #end>
          <Button icon="pi pi-send" aria-label="Filter" @click="send"/>
        </template>
      </Toolbar>
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import MessageComp from './components/Message.vue'
import ScrollPanel from 'primevue/scrollpanel'

export default {
  name: 'App',
  components: {
    HelloWorld,
    MessageComp,
    ScrollPanel
  },
  data() {
    return {
      value: '',
      socket: null,
      messages: [] // Tableau pour stocker les messages reçus
    }
  },
  methods: {
    send() {
      this.socket.send(this.value);
      this.value = "";
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:8080/echo");

    this.socket.onopen = () => {
      this.messages.push("Status : connected");
    };

    this.socket.onmessage = (e) => {
      // Ajoutez le message au tableau des messages
      this.messages.push(e.data);
    };
  }
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

.main-chat-cont {
  height: 83vh;
}

.p-message-text {
  word-break: break-all;
}

.p-toolbar-group-center {
  width: 85%;
}

.dialog-width-responsive {
  width: 50vw;
}

@media (max-width: 768px) {
  .w-auto-md {
    width: auto;
  }

  .main-chat-cont {
    max-height: 69.5vh;
  }

  .p-toolbar-group-start {
    width: 100%;
    font-size: 1.3em !important;
  }

  .p-toolbar-group-center {
    width: 80%;
  }

  .p-toolbar-group-end {
    width: auto;
  }

  .dialog-width-responsive {
    width: 90vw;
  }
}

@media (max-width: 500px) {
  .main-chat-cont {
    max-height: 66vh;
  }
}
</style>
