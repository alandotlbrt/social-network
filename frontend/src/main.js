// src/main.js
import { reactive, createApp } from "vue";
import App from "./views/App.vue";
import router from "./router";
import LoadingComponent from "./components/shared/loading.vue";
import DarkMode from "./components/shared/darkMode.vue";
import headernotconnected from "./components/shared/notconnectedheader.vue";
import headerconnected from "./components/shared/headerconnected.vue";
import Info from "./views/info.vue";
import Decobutton from "./components/deconnection.vue";
import Menu from "./components/menu.vue"
import ReusableButton from './components/reusableButton.vue';
import './assets/styles/main.css';
import ImagePost from "./components/posts/imagePost.vue";
import textPost from './components/posts/textPost.vue';
import textImagePost from './components/posts/textImagePost.vue';

const userState = reactive({
  id: "",
  username: "",
  profilePicture: "",
  Firstname: "",
  Lastname: "",
  email: "",
  About: "",
});

const app = createApp(App);
app.component("headerc", headerconnected);
app.component("reusableButton", ReusableButton);
app.component("ProfilMenu", Menu);
app.component("Deconnection", Decobutton);
app.component("Loading", LoadingComponent);
app.component("DarkMode", DarkMode);
app.component("Profile", Info);
app.component("headernc", headernotconnected);
app.component("imagePost", ImagePost);
app.component("textPost", textPost);
app.component("textImagePost", textImagePost);

app.provide("userState", userState);

app.use(router);
app.mount("#app");