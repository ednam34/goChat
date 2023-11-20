<template>
   <Dialog :visible="showDialog" v-model="showDialog" header="Entrez votre pseudo" :closable="false" :modal="true">
    <InputText v-model="tempUsername" placeholder="Pseudo"/>
    <Button class="ml-1" label="Confirmer" @click="confirmUsername"/>
  </Dialog>
  <div>
    <h1>FUEGO CHAT APP</h1>
    <div class="p-3 grid">
        <div ref="scrollPanel" id="scrollMsg" class="field col-10 mr-2 back main-chat-cont"  style="height: 70vh">
            <MessageComp
              v-for="(message, index) in messages"
              :class="message.username === this.username ? 'msgRcv' : 'msgEnv'"
              :key="index"
              :type="message.type"
              :msg="message.message"
              :userName="message.username"
              :icon="message.type === 'success' ? 'pi pi-user-plus' : message.type === 'error' ? 'pi pi-user-minus':'pi pi-send'"
            />
        </div>
        <ScrollPanel ref="scrollPanel" id="scrollMsg2" class="field col back"  style="height: 70vh">
          <MessageComp
              v-for="(message, index) in users"
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
      messages: [], // Tableau pour stocker les messages reçus
      users: []
    }
  },
  methods: {
    receiveMessage(jsonString) {
      try {
        const jsonObject = JSON.parse(jsonString);
        if (jsonObject && jsonObject.username && jsonObject.message && jsonObject.type) {
          if(jsonObject.type==="error"){
            console.log("error")
            this.users = this.users.filter(message => message.username !== jsonObject.username);
            this.messages.push(new Message(jsonObject.type, jsonObject.username, jsonObject.message));
          }
          else if(jsonObject.type==="success"){
            this.addUser(jsonObject.username)
          }else{
            this.messages.push(new Message(jsonObject.type, jsonObject.username, jsonObject.message));
            var scrollingContainer = document.getElementById('scrollMsg');
            scrollingContainer.scrollTop = scrollingContainer.scrollHeight;
          }
        } else {
          console.error("JSON invalide reçu");
        }
      } catch (error) {
        console.error("Erreur lors du parsing du JSON:", error);
      }
    },
    createAndSendMessage(type, message) {
      const messageObject = { type, username: this.username, message };
      this.socket.send(JSON.stringify(messageObject));
      if(type !="success"){
        this.addNewMessage(type, message);
      }
      //this.addNewMessage(type, message);
    },
    addUser(userName){
      const newUser = new Message("success", userName, "Connecté");
      this.users.push(newUser);
      var scrollingContainer = document.getElementById('scrollMsg2');
      scrollingContainer.scrollTop = scrollingContainer.scrollHeight;

    },
    addNewMessage(type, message) {
      const newMessage = new Message(type, this.username, message);
      this.messages.push(newMessage);
      var scrollingContainer = document.getElementById('scrollMsg');
      scrollingContainer.scrollTop = scrollingContainer.scrollHeight;
    },
    send() {
      if (!this.value.trim()) return;
      this.createAndSendMessage("info", this.value);
      this.value = "";
    },
    confirmUsername() {
      if (!this.tempUsername.trim()) {
        alert("Veuillez entrer un pseudo.");
        return;
      }
      this.username = this.tempUsername;
      this.showDialog = false;
      this.createAndSendMessage("success", "Utilisateur connécté");
    }
  },
  mounted() {
    this.socket = new WebSocket('ws://rayanekaabeche.fr:8081/chat');

    this.socket.onopen = () => {
      console.log("Connexion avec le serveur réussi")
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
  overflow-y: scroll;
}

.main-chat-cont::-webkit-scrollbar {
  width: 0.5em; /* Largeur de la barre de défilement */
}

.main-chat-cont::-webkit-scrollbar-thumb {
  background-color: #888; /* Couleur de la poignée de la barre de défilement */
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
