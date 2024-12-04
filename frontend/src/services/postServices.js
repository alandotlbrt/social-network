import api from './apiconfig.js';
export async function formPostModels(context) {
     // Téléchargement de l'image
     const base64Photo = await uploadImage().catch((error) => {
          console.error("Error uploading image:", error);
          return null;
     });

     try {
          const response = await api.post('/api/createpost', { message: context.text, pictures: base64Photo, privacy: context.privacyValue });
     
          return response.data
     }
     catch (error) {
          console.error("Error during login:", error);
     }
}


export function uploadImage() {
     const fileInput = document.getElementById("avatar-post");
     const file = fileInput.files[0];
     if (file) {
          return new Promise((resolve, reject) => {
               const reader = new FileReader();
               reader.readAsDataURL(file);
               reader.onload = () => {
                    const base64String = reader.result;
                    resolve(base64String);
               };
               reader.onerror = (error) => {
                    reject(error);
               };
          });
     } else {
          return Promise.resolve(null);
     }
}
export async function loadingPost() {
     try {
          const response = await api.post('/api/loadpost');
          
          if (response) {
               return response
          }
     } catch (error) {
          console.log(error)
     }
}

export function displayContentPost() {
     const onlyImagePost = document.querySelector('.onlyImage-post')
     const headerOnlyImagePost = document.querySelector('.headerOnlyImage-post')
     
     onlyImagePost.addEventListener('mouseover', function() {
          headerOnlyImagePost.style.display = "flex;"

     })

     onlyImagePost.addEventListener('mouseout', function() {
          headerOnlyImagePost.style.display = 'none';

     })
}

document.addEventListener("DOMContentLoaded", function() {
     const textContent = document.getElementById("text-content")
     const readMoreLink = document.getElementById("read-more")
     const modal = document.getElementById("modalText")
     const modalText = document.getElementById("modal-text")
     const closeModal = document.getElementsByClassName("close")[0]

     const originalText = textContent.innerText

     function checkOverFlow() {
          if (textContent.scrollHeight > textContent.clientHeight) {
               const truncatedText = originalText.substring(0, 100) + "..."
               textContent.innerText = truncatedText;
               readMoreLink.style.display = "inline"
          }
     }
})
