<template>
  <div v-if="!isLoading" class="visibleApp">
    <header v-if="showHeader && !isConnected">
      <headernc />
    </header>
    <header v-if="showHeader && isConnected">
      <headerc />
    </header>
      <router-view></router-view>
  </div>
  <div v-else>
    <Loading />
  </div>
</template>

<script>
import { checkConnection } from '../services/authServices.js';
import { inject } from "vue"; // Import inject to access userState
import { Info } from '../services/userServices.js';
import { socket } from './login.vue';



export default {
  name: "App",

  data() {
    return {
      isConnected: false,
      isLoading: false,
    };
  },

  computed: {
    showHeader() {
      return this.$route.name !== "Login" && this.$route.name !== "Register";
    },
  },

  watch: {
    async $route(to, from) {
      this.isLoading = true;
      try {
        await checkConnection(this, to);
      } finally {
        this.isLoading = false;
      }
      const restrictedPaths = ["/", "/login", "/register"];
      if (!restrictedPaths.includes(to.path) && !this.isConnected) {
        this.$router.push("/"); // Redirect if user tries to access restricted paths
      }
    },

    isConnected(newValue) {
      setTimeout(() => {
        socket.send(JSON.stringify({ type: "delete", value: this.userState.id }));
        socket.send(JSON.stringify({ type: "add", value: this.userState.id }));
      }, 40); 
      if (newValue) {
        this.setupUserState();
      }
    },
  },

  setup() {
    const userState = inject("userState"); // Accéder à l'état utilisateur
    return { userState };
  },

  methods: {
    async setupUserState() {
      try {
        await Info(this.userState);
      } catch (error) {
        console.error("Error setting up user state:", error);
      }
    },
  },
};

     let docTitle = document.title
     window.addEventListener("blur", () => {
          document.title = "Come back 👉👈";
     })
     window.addEventListener("focus", () => {
          document.title = docTitle;
     })

</script>