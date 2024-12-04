// src/router/index.js
import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/home.vue";
import Register from "../views/register.vue";
import Login from "../views/login.vue";
import settings from "../views/settings.vue";
import Profile from "../views/profile.vue";
import createPost from "../views/post/createPost.vue";
import chat from "../views/chat.vue";
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/settings",
    name: "settings",
    component: settings,
  },
  {
    path: "/create-post",
    name: "createPost",
    component: createPost,
  },
  {
    path: "/:idUser",
    name: "Profile",
    component: Profile,
    props: true,
  },
  {
    path: "/chat",
    name: "chat",
    component: chat,
  },
  { path: "/:pathMatch(.*)*", name: "NotFound", redirect: "/" },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
