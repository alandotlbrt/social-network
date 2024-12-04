import { uploadImageSettings } from '../services/authServices.js';
import api from './apiconfig.js';
import router from '../router/index.js';

var Pictures;
export const Settings = {
     async HandleUpdate() {
          await uploadImageSettings()
               .then((base64String) => {
                    if (base64String == null) {
                         return Pictures;
                    }
                    return base64String;
               })
               .catch((error) => {
                    console.error("Erreur:", error);
               });

          try {
               var base64Photo = await uploadImageSettings()
                    .then((base64String) => {
                         return base64String;
                    })
                    .catch((error) => {
                         console.error("Erreur:", error);
                    });
               if (base64Photo !== null) {
                    Pictures = base64Photo;
               }

      const response = await api.post("/api/updatedata", {
        Id: document.getElementById("IdUser").value,
        Username: document.getElementById("Username").value,
        Firstname: document.getElementById("Firstname").value,
        Lastname: document.getElementById("LastName").value,
        email: document.getElementById("emailSettings").value,
        About: document.getElementById("About").value,
        Photo: Pictures,
        password: document.getElementById("Password").value,
        newPassword: document.getElementById("NewPassword").value,
        Privacy: document.getElementById("Text-Status").textContent,
      });

               const result = response.data;
               if (result.success) {
                    location.reload();
               } else {
                    alert(result.message);
               }
          } catch (error) {
               console.error("Error during Update:", error);
          }
     },
     async toOtherProfile(id) {
          router.push(`/${id}`);
     },
     async handleProfile(context) {
          try {
               const response = await api.post("/api/info", {
                    Id: context.idUser,
               });  
               if (response.data.success) {

                    if (response.data.Info.Profil_Pictures != "") {
                         let base64Image = response.data.Info.Profil_Pictures;
                         let PP = document.getElementsByClassName("circle-label-profile");
                         PP[0].style.backgroundImage = `url(${base64Image})`;
                    }
                    document.getElementById("Following").textContent += " " + response.data.Info.Follow;
                    document.getElementById("Followers").textContent += " " + response.data.Info.Follower;
                    document.getElementById("Username-profile").textContent += response.data.Info.Username;
                    document.getElementById("First-Name").textContent += response.data.Info.Firstname;
                    document.getElementById("Last-Name").textContent += response.data.Info.Lastname;
                    document.getElementById("Email").textContent += response.data.Info.Email;
                    document.getElementById("Birthday").textContent += response.data.Info.Birthday;
                    document.getElementById("About").textContent += response.data.Info.About;
                    setTimeout(()=>{
                         document.getElementById("number-follower-inbox").textContent = response.data.Info.FollowNumberInbox
                    }, 100)
                    return {isfollow:response.data.Info.IsFollowing, privacy:response.data.Info.Privacy, iswaiting:response.data.Info.IsWaitingFollow};
               }
          } catch (error) {
               console.log(error)
               router.push("/");
          }
     },
};

export function image(label, event) {
  if (event) {
    const reader = new FileReader();
    reader.onload = function (e) {
      label.style.backgroundImage = `url(${e.target.result})`;
      label.textContent = "";
    };
    reader.readAsDataURL(event);
  }
}

export function base64ToFile(base64String, filename) {
  // Décoder la chaîne Base64 en données binaires
  const arr = base64String.split(",");
  const mime = arr[0].match(/:(.*?);/)[1]; // Récupère le type MIME de la chaîne Base64
  const bstr = atob(arr[1]); // Décoder la partie Base64
  let n = bstr.length;
  const u8arr = new Uint8Array(n);
  // Convertir les caractères binaires en tableau d'octets
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  // Créer un Blob à partir du tableau d'octets
  const blob = new Blob([u8arr], { type: mime });

  // Créer un fichier à partir du Blob (nom et type définis)
  return new File([blob], filename, { type: mime });
}

export const Info = async (userState) => {
  try {
    const response = await api.post("/api/recupdata");

    const result = await response.data;

    if (result.success) {
      userState.id = result.Info.Id;
      userState.username = result.Info.Username;
      userState.profilePicture = result.Info.Profil_Pictures;
      userState.firstName = result.Info.Firstname;
      userState.lastName = result.Info.Lastname;
      userState.email = result.Info.Email;
      userState.about = result.Info.About;
      if (result.Info.Privacy == "Public") {
        userState.privacy = false;
      } else {
        userState.privacy = true;
      }
      return {
        username: result.Info.Username,
      };
    }
  } catch (error) {
    console.error("Error fetching user info:", error);
  }
  return null;
};


