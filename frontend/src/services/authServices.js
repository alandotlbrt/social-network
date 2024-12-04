import api from './apiconfig.js';
import router from '../router/index.js';

export const formMethodsRegister = {
     handleFileChange(event, label) {
          const file = event.target.files[0];
          if (file) {
               const reader = new FileReader();
               reader.onload = function (e) {
                    label.style.backgroundImage = `url(${e.target.result})`;
                    label.textContent = "";
               };
               reader.readAsDataURL(file);
          }
     },

     

     async HandleRegister(component) {
          // Assuming uploadImage is a function that uploads the user's image
          const base64Photo = await uploadImage().catch((error) => {
               console.error("Error uploading image:", error);
               return null; // Handle the case where uploadImage fails
          });

          try {

               const response = await api.post('/register', {
                    Username: component.Username,
                    Firstname: component.Firstname,
                    Lastname: component.Lastname,
                    email: component.email,
                    password: component.password,
                    confirmPassword: component.confirmPassword,
                    Birthday: component.Birthday,
                    About: component.About,
                    Photo: base64Photo,
                    Privacy: 'Public',
                    Follower: 0,
               });

               const result = response.data;
               component.Allerrors2 = [];

               if (result.success) {
                    router.replace("/login");
               } else {
                    var errorContainer = document.getElementById("Allerrors");
                    errorContainer.innerHTML = "";
                    if (result.message != "") {
                         component.Allerrors2.push(result.message);
                         var errorMessage = document.createElement("p");
                         errorMessage.textContent = result.message;
                         errorContainer.appendChild(errorMessage);
                    }

                    if (result.errors && result.errors.length > 0) {
                         result.errors.forEach((error) => {
                              component.Allerrors2.push(error);
                              var errorMessage = document.createElement("p");
                              errorMessage.textContent = error;
                              errorContainer.appendChild(errorMessage);
                         });
                    }
               }
          } catch (error) {
               console.error("Error during registration:", error);
          }
     },

     checkEmail(email) {
          const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
          const isEmailValid = emailPattern.test(email);
          return { isEmailValid };
     },

     validateDate(birthday) {
          const today = new Date();
          const birthDate = new Date(birthday);

          if (birthday === "" || isNaN(birthDate.getTime())) {
               return false;
          }

          let age = today.getFullYear() - birthDate.getFullYear();
          const monthDiff = today.getMonth() - birthDate.getMonth();

          if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
               age--;
          }
          

          return age >= 15;
     },

     validatePassword(password, errors) {
          if (!/[A-Z]/.test(password)) {
               errors.push("At least one capital letter");
          }

          if (!/[a-z]/.test(password)) {
               errors.push("At least one lowercase letter");
          }

          if (!/\d/.test(password)) {
               errors.push("At least one digit");
          }

          if (password.length < 8) {
               errors.push("At least 8 characters");
          }

          if (!/[@$!%*?&._-]/.test(password)) {
               errors.push("at least one of these @, $, !, %, *, ?, &, .,-,_");
          }
     },
};

export function uploadImage() {
     const fileInput = document.getElementById("file-input-register");
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

export function uploadImageSettings() {
     const fileInput = document.getElementById("file-input-settings");
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


export async function checkConnection(context, to) {
     const token = getCookiename();

     if (!token) {
          context.isConnected = false;
          return false;
     }

     try {
          const response = await api.post('/api/check', { session_token: token });

          const data = response.data;
          context.isConnected = data.valid;

          if (context.isConnected === true) {
               const restrictedPaths = ['/', '/login', '/register'];
               if (restrictedPaths.includes(to.path)) {
                    router.push('/');
               }
          } else {
               deleteCookie('session_token');
          }
     } catch (error) {
          console.error('Erreur lors de la vérification de connexion :', error);
          context.isConnected = false;
          deleteCookie('session_token');
     }

     return context.isConnected;
}

export function getCookiename() {
     const value = `; ${document.cookie}`;
     const parts = value.split(`; session_token=`);
     if (parts.length === 2) {
          return parts.pop().split(';').shift();
     }
     return null;
}

// Fonction pour supprimer le cookie spécifié
export function deleteCookie(name) {
     document.cookie = `${name}=; Max-Age=0; path=/;`;
}

///////////////////////////////
///////////////////////////////
///////////////////////////////

//LOGINNNNN

///////////////////////////////
///////////////////////////////
///////////////////////////////


export const loginBlock = {
     async handleLogin(context) {
          const errorEmail = document.getElementById("errorEmail-login");
          const errorPassword = document.getElementById("errorPassword-login");

          errorEmail.textContent = "";
          errorPassword.textContent = "";

          try {

               const response = await api.post('/login', { email: context.email, password: context.password, });

               const result = response.data;

               if (result.success) {
                    router.replace("/");
               } else {
                    if (result.message.includes("Email")) {
                         errorEmail.textContent = result.message;
                         context.password = "";
                    } else if (result.message.includes("password")) {
                         errorPassword.textContent = result.message;
                         context.password = "";
                    }
               }
          } catch (error) {
               console.error("Error during login:", error);
          }
     }
}
