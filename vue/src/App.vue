<template>
   <Dialog :visible="showDialog" v-model="showDialog" header="Entrez votre pseudo" :closable="false" :modal="true">
    <InputText v-model="tempUsername" placeholder="Pseudo"/>
    <Button label="Confirmer" @click="confirmUsername"/>
  </Dialog>
  <div>
    <h1>FUEGO CHAT APP</h1>
    <div class="p-3">
        <ScrollPanel ref="scrollPanel" id="scrollMsg" class="back overflow-auto main-chat-cont"  style="width: 100%; height: 70vh">
            <MessageComp
              v-for="(message, index) in messages"
              :class="message.username === this.username ? 'msgRcv' : 'msgEnv'"
              :key="index"
              :type="message.type"
              :msg="message.message"
              :userName="message.username"
              :icon="message.type === 'success' ? 'pi pi-user-plus' : message.type === 'error' ? 'pi pi-user-minus':'pi pi-send'"
            />
        </ScrollPanel>
     </div>
    
    <div class="container p-2">
      <Toolbar class="mt-3 w-full">
        <template #start>
          <p class="text-sm mb-0 sm:mt-0">Votre message:</p>
        </template>
        <template #center>
          <div class="w-full w-auto-md">
            <InputText type="text" v-model="value" class="w-full w-auto-md" ref="input" @keyup.enter="send" />
          </div>
        </template>
        <template #end>
          <Button icon="pi pi-send" label="envoyer" aria-label="Filter" @click="send"/>
        </template>
      </Toolbar>
    </div>
  </div>
</template>

<script>
import MessageComp from './components/Message.vue'
import ScrollPanel from 'primevue/scrollpanel'
import Dialog from 'primevue/dialog'
import Message from './model/Message'

export default {
  name: 'App',
  components: {
    MessageComp,
    ScrollPanel,
    Dialog,
  },
  data() {
    return {
      username: '',
      tempUsername: '', // Variable temporaire pour le pseudo
      showDialog: true, // Afficher le dialogue au démarrage
      value: '',
      socket: null,
      messages: [] // Tableau pour stocker les messages reçus
    }
  },
  methods: {
    receiveMessage(jsonString) {
      try {
        const jsonObject = JSON.parse(jsonString);

        if (jsonObject && jsonObject.username && jsonObject.message && jsonObject.type) {
          const newMessage = new Message(jsonObject.type, jsonObject.username, jsonObject.message);
          console.log(newMessage)
          this.messages.push(newMessage);
        } else {
          console.error("JSON invalide reçu");
        }
      } catch (error) {
        console.error("Erreur lors du parsing du JSON:", error);
      }
    },
    send() {
      if (this.value.trim() !== '') {
        const messageObject = {
        type: "info",  
        username: this.username,
        message: this.value
        };
        this.socket.send(JSON.stringify(messageObject));
        const newMessage = new Message("info",this.username, this.value);
        this.messages.push(newMessage);
        this.value = "";
        console.log(document.getElementById('scrollMsg'))
      }
    },
    confirmUsername() {
      if (this.tempUsername.trim() !== '') {
        this.username = this.tempUsername;
        this.showDialog = false;
        const messageObject = {
        type: "success",  
        username: this.username,
        message: "Utilisateur connécté"
        };
        this.socket.send(JSON.stringify(messageObject));
        const newMessage = new Message("success",this.username, "Utilisateur connécté");
        this.messages.push(newMessage);
      } else {
        alert("Veuillez entrer un pseudo."); // Vous pouvez utiliser un mécanisme plus sophistiqué pour la validation
      }
    }
    
  },
  mounted() {
    this.socket = new WebSocket("ws://rayanekaabeche.fr:8081/echo");

    this.socket.onopen = () => {
     
    };

    this.socket.onmessage = (e) => {
      this.receiveMessage(e.data);
    };
  }
}
</script>
<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;

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
.back{
  background-color: var(--surface-card);
  border: solid;
  border-color: var(--surface-border);
  border-radius: 10px;

}
.msgRcv{
  padding-right: 30px;
  padding-left: 30px;
  border-radius: 10px;
}
.msgEnv{
  padding-right: 30px;
  padding-left: 30px;
  border-radius: 10px;
  text-align: left !important;
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
    width: 90%;
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
