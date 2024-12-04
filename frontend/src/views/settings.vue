<template>
  <DarkMode />
  <div class="container-settings">
    <form @submit.prevent="Update" class="settingsForm">
      <input type="hidden" id="IdUser" v-model="userState.id" />
      <div class="circleAndstatut-settings">
        <div class="circle-wrapper-settings">
          <input type="file" id="file-input-settings" accept="image/*" @change="handleFileChange" />
          <label for="file-input-settings" class="circle-label-settings" id="image" 
          :style="{ backgroundImage: `url(${tempProfilePicture || userState.profilePicture})` }">
          </label>
        </div>
        <label class="switch-settings" @click="togglePrivacy">
          <input class="input-settings" type="checkbox" id="Private" v-model="userState.privacy" />
          <span class="slider-settings"></span>
          <p id="Text-Status">{{ userState.privacy ? "Private" : "Public" }}</p>
        </label>
      </div>
      <div class="usernameDiv-settings">
        <p class="pUsername-settings">Username</p>
        <input class="input-settings" type="text" id="Username" v-model="userState.username" />
      </div>
      <div class="table-container-settings">
        <table>
          <tbody>
            <tr>
              <td>
                <div class="firstNameDiv-settings">
                  <p class="pFirstName-settings">First Name</p>
                  <input class="input-settings" type="text" id="Firstname" v-model="userState.firstName" />
                </div>
              </td>
              <td id="td-right-settings">
                <div class="lastNameDiv-settings">
                  <p class="pLastName-settings">Last Name</p>
                  <input class="input-settings" type="text" id="LastName" v-model="userState.lastName" />
                </div>
              </td>
            </tr>
            <tr>
              <td>
                <div class="passwordDiv-settings">
                  <p class="pPassword-settings">Password</p>
                  <input class="input-settings" type="password" id="Password" />
                </div>
              </td>
              <td id="td-right-settings">
                <div class="newPasswordDiv-settings">
                  <p class="pNewPassword-settings">New password</p>
                  <input
                    class="input-settings"
                    type="password"
                    id="NewPassword"
                    v-model="password"
                    @input="validatePassword"
                    :class="{ 'input-success': passwordTouched && errors.length === 0 && password.length > 0 }"
                    @blur="passwordTouched = true"
                  />
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
        <div class="EmailDiv-settings">
          <p class="pEmail-settings">Email</p>
          <input class="input-settings" type="email" id="emailSettings" v-model="userState.email" />
        </div>
        <div class="aboutMeDiv-settings">
          <p class="pAboutMe-settings">About Me</p>
          <div class="AboutAndChar-settings">
            <textarea class="aboutInput-settings" id="About" maxlength="150" @input="characterCount">{{ userState.about }}</textarea>
            <div id="charCount-settings"> </div>
          </div>
        </div>
        <button class="submit-settings" type="submit">Save</button>    
    </form>
    <div id="Allerrors"></div>
  </div>
</template>
<script>
import { formMethodsRegister } from '../services/authServices.js';
import { Settings } from "../services/userServices.js";
import { inject, watch } from "vue";

export default {
  setup() {
    const userState = inject("userState");
    return { userState };
  },
  data() {
    return {
      errors: [],
      password: "",
      passwordTouched: false,
      tempProfilePicture: null,
    };
  },
  methods: {
    Update() {
      Settings.HandleUpdate();
    },
    validatePassword() {
      this.errors = [];
      if (this.password.trim() !== "") {
        formMethodsRegister.validatePassword(this.password, this.errors);
        if (this.errors.length > 0) {
          this.errors = [this.errors[0]];
        }
      }
    },
    characterCount() {
      const aboutInput = document.getElementById('About').value.length
      const charCount = document.getElementById('charCount-settings')

      charCount.textContent = aboutInput + '/150';
    }, 
    handleFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.tempProfilePicture = e.target.result; 
        };
        reader.readAsDataURL(file);
      }
    },
    togglePrivacy(event) {
      event.stopPropagation();
    }
  },
  mounted() {
    setTimeout(() => {
      this.characterCount();
    }, 100);
  },

};
</script>
