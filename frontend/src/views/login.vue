<template>
  <router-link to="/">
    <svg class="cross-login" fill="#000000" height="15px" width="15px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 490 490" xml:space="preserve">
      <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
      <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
      <g id="SVGRepo_iconCarrier">
        <polygon points="456.851,0 245,212.564 33.149,0 0.708,32.337 212.669,245.004 0.708,457.678 33.149,490 245,277.443 456.851,490 489.292,457.678 277.331,245.004 489.292,32.337"></polygon>
      </g>
    </svg>
  </router-link>
  <div class="container-login">
    <div class="innerDiv-login">
      <h1 class="title-login">Welcome to <span class="titleSpan-login">Plume</span></h1>
      <form id="myForm" @submit.prevent="submitFormLogin" class="form-login">
        <div class="emailDiv-login">
          <label id="email-login" for="email">Email</label>
          <input class="emailInput-login" type="email" v-model="email" placeholder="Enter your email" required />
          <label id="errorEmail-login"></label>
        </div>
        <div class="passwordDiv-login">
          <label id="password-login" for="password">Password</label>
          <input id="inputPassword" class="passwordInput-login" type="password" v-model="password" placeholder="Enter your password" required />
          <label id="errorPassword-login"></label>
        </div>
        <p class="forgotPassword-login">Password forgot?</p>
        <button class="login" type="submit">Log in</button>
        <div class="separation-login">
          <div class="line-login"></div>
          <p class="or-login">OR</p>
          <div class="line-login"></div>
        </div>
        <div class="githubDiv">
          <img src="/src/assets/images/github.png" alt="Github" class="img-login">
          <p class="github">Continue with Github</p>
        </div>
        <p class="signUp"><router-link id="signUP-login" to="/register">Sign up</router-link></p>
      </form>
    </div>
  </div>
  <DarkMode />
</template>

<script>

import { loginBlock } from '../services/authServices';

export const socket = new WebSocket("ws://localhost:8080/ws"); 
export default {
  name: 'Login',
  data() {
    return {
      email: '',
      password: '',
    };
  },

  methods: {
    async submitFormLogin(){
      loginBlock.handleLogin(this)
    }
  },
};
</script>
