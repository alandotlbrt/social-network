<template>
  <div>
  
    <label class="circle-labels-menu" :style="{ backgroundImage: `url(${userState.profilePicture})` }" @click="toggleDropdown"></label>

    <div
      class="drop-down-menu"
      :class="{ visible: isDropdownVisible }"
    >
    <div class="left-menu" @click="toggleDropdown" style="cursor: pointer;" ><svg fill="#fff" height="15px" width="15px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 490 490" xml:space="preserve"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <polygon points="456.851,0 245,212.564 33.149,0 0.708,32.337 212.669,245.004 0.708,457.678 33.149,490 245,277.443 456.851,490 489.292,457.678 277.331,245.004 489.292,32.337 "></polygon> </g></svg></div>
      <div class="circle-wrappers-menu second-circle">
        <label class="circle-labels-menu second-circle" id="image" v-if="userState && userState.profilePicture" :style="{ backgroundImage: `url(${userState.profilePicture})` }">
        </label>
      </div>
      <div class="redirectionButton-Menu">
        <reusableButton label="Profile" @click="handleClickProfile(userState.id)" />
        <reusableButton label="Settings" :onClick="handleClickSettings" />
        <reusableButton label="Groups" :onClick="handleClickGroups" />
      </div>
      <Deconnection />
    </div>
  </div>
</template>

<script>
import { inject } from 'vue';
import router from '../router';
import { Settings } from '../services/userServices';

export default {
  data() {
    return {
      isDropdownVisible: false,
      userState: null,
    };
  },

  created() {
    this.userState = inject("userState");
  },

  methods: {
    toggleDropdown() {
      this.isDropdownVisible = !this.isDropdownVisible;
    },

    handleClickProfile(id) {
      Settings.toOtherProfile(id)
    },

    handleClickSettings() {
      router.replace("/settings");
    },

    handleClickGroups() {
      router.replace("/invitations");
    },
  },
};
</script>
