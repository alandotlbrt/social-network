<template>
  <DarkMode />
  <div class="Menu">
    <div class="buttonchat">
      <button class="Friend">Friend</button>
      <button class="Group">Group</button>
    </div>
    <p>pseudo des coupains</p>
  </div>
  <div class="chatbox">
    <div id="chat"></div>
    <div class="inputChat">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" class="icon">
        <path
          d="M3.25 23.25C2.5625 23.25 1.96875 23.0104 1.46875 22.5313C0.989584 22.0313 0.75 21.4375 0.75 20.75V3.25C0.75 2.5625 0.989584 1.97917 1.46875 1.5C1.96875 1 2.5625 0.75 3.25 0.75H20.75C21.4375 0.75 22.0208 1 22.5 1.5C23 1.97917 23.25 2.5625 23.25 3.25V20.75C23.25 21.4375 23 22.0313 22.5 22.5313C22.0208 23.0104 21.4375 23.25 20.75 23.25H3.25ZM3.25 20.75H20.75V3.25H3.25V20.75ZM4.5 18.25H19.5L14.8125 12L11.0625 17L8.25 13.25L4.5 18.25ZM3.25 20.75V3.25V20.75Z"
          fill="#1D1B20"
        />
      </svg>
      <input class="inputTextChat" type="text" v-model="Message" />
      <button class="Send-Chat" @click="sendTchat(Message)">send</button>
    </div>
  </div>
</template>

<script>
import { Websocks } from "../services/Websocket";
import { inject } from "vue";

const regex = new RegExp(/[a-zA-Z]/);

export default {
  setup() {
    const userState = inject("userState"); // Accéder à l'état utilisateur
    return { userState };
  },
  data() {
    return {
      Message: "", // Initialisez `Message` avec une chaîne vide ou toute autre valeur par défaut
    };
  },
  methods: {
    sendTchat(message) {
      if (regex.test(message)) {
        Websocks.Message(message, this.userState.id);
        this.Message = "";
      }
    },
  },
};
</script>
