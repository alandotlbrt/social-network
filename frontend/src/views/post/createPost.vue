<template>
  <div class="container-post">
    <form id="form-post" @submit.prevent="submitFormPost">
      <div>
        <div v-if="imagePreview" @click="deleteImage" class="removePhoto-Post">
          <svg :class="{ darkmode: isDarkMode }" class="cross-createPost" fill="#000000" height="15px" width="15px" version="1.1" id="Capa_1"
            xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 490 490"
            xml:space="preserve">
            <g id="SVGRepo_iconCarrier">
              <polygon
                points="456.851,0 245,212.564 33.149,0 0.708,32.337 212.669,245.004 0.708,457.678 33.149,490 245,277.443 456.851,490 489.292,457.678 277.331,245.004 489.292,32.337 " />
            </g>
          </svg>
        </div>

        <div class="superposition-Post">
          <div v-if="imagePreview" class="overlay-disabled" @click="handleDisabledInputClick" @auxclick="updateImagePreview">
          </div>

          <input id="avatar-post" type="file" ref="fileInput" name="post-pictures-input" accept="image/png, image/jpeg"
            @change="handleFileChange" :disabled="imagePreview" />
          <label for="avatar-post" class="image-post">choose a picture: </label>
        </div>
      </div>

      <div class="Text-post">
        <div class="title-post">
          <div class="custom-select-container"
               v-on="eventHandler"
               @mouseleave="hideOptions">
            <select v-model="privacyValue" class="button-privacy" name="privacy" id="privacy">
              <option value="public">Public</option>
              <option value="private">Private</option>
              <option value="followers">Choose your followers</option>
            </select>
            <div class="custom-options" :class="{ active: isOptionsVisible }">
              <div class="custom-option" @click="selectOption('public')">
                Public
              </div>
              <div class="custom-option" @click="selectOption('private')">
                Private
              </div>
              <div class="custom-option" @click="selectOption('followers')">
                Choose your followers
              </div>
            </div>
          </div>
        </div>
        <textarea v-model="text" class="textarea-createPost"
          placeholder="Add a description or start a discuss about something ..."></textarea>
        <div class="error-Post">{{ this.errorPost }}</div>
      </div>
    </form>
  </div>
  <button class="button-post" @click="submitFormPost">Post! </button>
  <div class="Flou">
    <div class="imagePreview-Post">
      <div @click="closePreview" style="cursor: pointer;">
        <svg :class="{ darkmode: isDarkMode }" class="cross-createPost" fill="#000000" height="15px" width="15px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 490 490" xml:space="preserve">
          <polygon
            points="456.851,0 245,212.564 33.149,0 0.708,32.337 212.669,245.004 0.708,457.678 33.149,490 245,277.443 456.851,490 489.292,457.678 277.331,245.004 489.292,32.337 " />
        </svg>
      </div>
      <div class="Preview_image_post">
        <img />
      </div>
    </div>
  </div>
  <DarkMode />
</template>

<script>
import { formPostModels } from "../../services/postServices";
import { formMethodsRegister } from "../../services/authServices";

export default {
  data() {
    return {
      text: '',
      pictures: null,
      imagePreview: false,
      privacyValue: 'public',
      isOptionsVisible: false,
      errorPost: null,
      isMobile: window.innerWidth <= 650, 
    };
  },
  computed: {
    isDarkMode() {
      return document.body.classList.contains('darkmode');
    },
    eventHandler() {
      return this.isMobile
        ? { click: this.toggleOptions } 
        : { mouseenter: () => (this.isOptionsVisible = true) };
    },
  },
  methods: {
    async handleFileChange(event) {
      try {
        if (!this.imagePreview) {
          formMethodsRegister.handleFileChange(event, event.target.labels[0]);
          this.imagePreview = true;
        }
      } catch (error) {
        console.error('Error handling file change:', error);
      } finally {
        setTimeout(() => {
          this.updateImagePreview();
        }, 50);
      }
    },
    async submitFormPost() {
      try {
        await formPostModels(this);
      } catch (error) {
        console.error("Error submitting the form", error);
      }
    },
    deleteImage() {
      this.imagePreview = false;
      this.pictures = null;
      if (this.$refs.fileInput) {
        this.$refs.fileInput.value = '';
      }
      const label = document.querySelector('label[for="avatar-post"]');
      if (label) {
        label.style.backgroundImage = '';
        label.textContent = "choose a picture:";
      }
    },
    handleDisabledInputClick() {
      const blur = document.getElementsByClassName("Flou")
      blur[0].style.display = "block";
    },
    closePreview() {
      const blur = document.getElementsByClassName("Flou")
      blur[0].style.display = "none";
    },
    updateImagePreview() {
      const label = document.querySelector('label[for="avatar-post"]');
      const imageUrl = window.getComputedStyle(label).backgroundImage;
      if (imageUrl !== 'none') {
        const previewImage = document.querySelector('.Preview_image_post img');
        if (previewImage) {
          previewImage.src = imageUrl.slice(5, -2);
        }
      }
    },
    toggleOptions() {
      this.isOptionsVisible = !this.isOptionsVisible;
    },
    hideOptions() {
    
      this.isOptionsVisible = false;
    },
    selectOption(value) {
      this.privacyValue = value;
      this.hideOptions();
    },
    handleResize() {
      this.isMobile = window.innerWidth <= 650;
    },
  },
  mounted() {
    window.addEventListener("resize", this.handleResize);
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.handleResize);
  },
};
</script>
