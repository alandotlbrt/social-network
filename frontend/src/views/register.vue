<template>
  <router-link to="/"><svg class="cross-register" fill="#000000" height="15px" width="15px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 490 490" xml:space="preserve"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <polygon points="456.851,0 245,212.564 33.149,0 0.708,32.337 212.669,245.004 0.708,457.678 33.149,490 245,277.443 456.851,490 489.292,457.678 277.331,245.004 489.292,32.337 "></polygon> </g></svg></router-link>
  <div id="center-register">
    <form @submit.prevent="submitForm">
      <div id="profil-register">
        <div class="circle-wrapper-register">
          <input type="file" id="file-input-register" accept="image/*" @change="handleFileChange"/>
          <label for="file-input-register" class="circle-label-register">Upload</label>
        </div>
        <div>
          <div class="registration">
            Registration
          </div>
          <div class="Usernamediv-register">
            <p class="p-register">Username</p>
            <input class="input-register" v-model="Username" @input="clearAllErrors" />
          </div>
        </div>
      </div>
      <div class="table-container-register">
        <table>
          <tbody>
            <tr>
              <td>
                <div class="input-container-register">
                  <p class="p-register">First Name <span style="color: red;">*</span></p>
                  <input class="input-register" type="text" v-model="Firstname" required />
                </div>
              </td>
              <td>
                <div class="input-container-register">
                  <p class="p-register">Last Name <span style="color: red;">*</span></p>
                  <input class="input-register" type="text" v-model="Lastname" required />
                </div>
              </td>
            </tr>
            <tr>
              <td>
                <div class="input-container-register">
                  <p class="p-register">Email <span style="color: red;">*</span></p>
                  <div class="emailandverif-register" style="display: flex; align-items: center;">
                    <input class="input-register" type="email" v-model="email" @input="handleInput2"  
                    :class="{ 'input-success-register': isEmailValid && email.length > 0,
                              'input-failed-register': !isEmailValid && email.length > 0
                            }"
                         
                    required />
                  </div>
                </div>
              </td>
              <td>
                <div class="input-container-register">
                  <p class="p-register"> date of birth <span style="color: red;">*</span></p>
                  <input class="input-register" type="date" v-model="Birthday" @input="clearAllErrors" required />
                </div>
              </td>
            </tr>
            <tr>
              <td>
                <div class="input-container-register">
                  <p class="p-register">Password <span style="color: red;">*</span></p>
                  <input class="input-register" type="password" v-model="password" @input="handleInput"
                    :class="{ 'input-success-register':  errors.length === 0 && password.length > 0}"
                    required />
                </div>
              </td>
              <td>
                <div class="input-container-register">
                  <p class="p-register">Confirmation Password <span style="color: red;">*</span></p>
                  <input class="input-register" type="password" v-model="confirmPassword" @input="clearAllErrors" required />
                </div>
              </td>
            </tr>
          </tbody>
        </table>
        <ul v-if="errors.length" class="errors-list-register">
          <li v-for="(error, index) in errors" :key="index">{{ error }}</li>
        </ul>
      </div>
      <div>
        <div class="about-div-register">
          <p class="p-register">About Me</p>
          <textarea v-model="About" class="about-input-register" @input="nomdemafonction"></textarea>
        </div>
      </div>
      <button v-show="Allerrors2.length === 0" class="Signup-register" type="submit">Sign Up</button>
      <div v-show="Allerrors2.length > 0" id="Allerrors-register"></div>
      <p class="p-register"><router-link to="/login" class="WhiteSignIn-register">Sign in</router-link></p>
    </form>
  </div>
  <div v-if="isLoading" class="load">
    <Loading  />
  </div>
  <DarkMode />
</template>
<script>
import { formMethodsRegister } from '../services/authServices.js';
export default {
  data() {
    return {
      Username: '',
      Firstname: '',
      Lastname: '',
      Birthday: '',
      email: '',
      password: '',
      confirmPassword: '',
      About: '',
      isEmailValid: false,
      errors: [],
      isLoading: false,
      Allerrors2: [],
    };
  },
  methods: {
    handleInput() {
      this.validatePassword();
      this.clearAllErrors();
    },
    handleInput2() {
      this.validateEmail();
      this.clearAllErrors();
    },
    clearAllErrors() {
      this.Allerrors2 = [];
    },
    handleFileChange(event) {
      formMethodsRegister.handleFileChange(event, event.target.labels[0]);
    },
    validatePassword() {
      this.errors = [];
      
      if (this.password.trim() !== '') {
        formMethodsRegister.validatePassword(this.password, this.errors);
        if (this.errors.length > 0) {
          this.errors = [this.errors[0]];
        }
      }
    },
    validateEmail() {
      const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      this.isEmailValid = emailPattern.test(this.email);
    },
    async submitForm() {
      this.isLoading = true;
      formMethodsRegister.Username = this.Username;
      formMethodsRegister.Firstname = this.Firstname;
      formMethodsRegister.Lastname = this.Lastname;
      formMethodsRegister.email = this.email;
      formMethodsRegister.password = this.password;
      formMethodsRegister.confirmPassword = this.confirmPassword
      formMethodsRegister.Birthday = this.Birthday;
      formMethodsRegister.About = this.About;
//coucou
      try {
        await formMethodsRegister.HandleRegister(this);
      } finally {
        this.isLoading = false;
      }
    },
  },
}
</script>
