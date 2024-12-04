import { createElementBlock } from "vue";
import { socket } from "../views/login.vue";
import api from "./apiconfig";
import { base64ToFile } from "./userServices";
import { jsx } from "vue/jsx-runtime";

export const Websocks = {
  Message(value, image, id) {
    socket.send(JSON.stringify({ type: "Message", from: id, value: value, style: "Private", image: image }));
  },

  async AllMessage(id) {
    //rajouter l'id du mec priver ou du groupe
    try {
      const response = await api.post("/api/loadChat", {});

      const result = response.data;

      if (result != null) {
        result.forEach((element) => {
          if (id == element.From_Id) {
            myMessage(element);
          } else {
            otherMessage(element);
          }
        });
      }
    } catch (error) {
      console.error("Error during login:", error);
    }
  },

  async receiveMessaege(id) {
    socket.onmessage = (arg) => {
      let Message = JSON.parse(arg.data);
      if (Message.Socket.type === "Message") {
        if (Message.Socket.from != id) {
          otherMessage(Message);
        } else {
          myMessage(Message);
        }
      }
    };
  },

  async Userchat(id) {
    try {
      const response = await api.post("/api/loadUser", {});

      const result = response.data;

      if (result != null) {
        result.forEach((element) => {
        });
      }
    } catch (error) {
      console.error("Error during login:", error);
    }
  },
};

function myMessage(element) {
  let bigdiv = document.createElement("div");
  bigdiv.className = "boxMessage";

  let boxMessage = document.createElement("div");
  boxMessage.className = "sendSaid";

  if (element.Content != undefined) {
    let Message = document.createElement("p");
    Message.className = "Said";
    if (element.Content != "") {
      Message.textContent = element.Content;
      boxMessage.appendChild(Message);
    }
  } else {
    let Message = document.createElement("p");
    Message.className = "Said";
    if (element.Socket.value != "") {
      Message.textContent = element.Socket.value;
      boxMessage.appendChild(Message);
    }
  }

  if (element.Image == undefined) {
    if (element.Socket.image != "") {
      try {
        const file = base64ToFile(element.Socket.image, "example.png");
        let message_Image = document.createElement("img");
        message_Image.className = "image_chat";
        message_Image.src = URL.createObjectURL(file);
        boxMessage.appendChild(message_Image);
      } catch (error) {
        console.error("An error occurred:", error);
      }
    }
  } else {
    if (element.Image != "") {
      try {
        const file = base64ToFile(element.Image, "example.png");
        let message_Image = document.createElement("img");
        message_Image.className = "image_chat";
        message_Image.src = URL.createObjectURL(file);
        boxMessage.appendChild(message_Image);
      } catch (error) {
        console.error("An error occurred:", error);
      }
    }
  }

  bigdiv.appendChild(boxMessage);

  document.getElementById("chat").appendChild(bigdiv);
}

function otherMessage(element) {
  let bigdiv = document.createElement("div");
  bigdiv.className = "boxMessageReceive";

  let divMessage = document.createElement("div");
  divMessage.className = "allMessage";

  let user_Pictures = document.createElement("label");
  const previewImage = user_Pictures.querySelector(".imagePreview-Post img");
  user_Pictures.style.backgroundImage = `url(${element.Pictures})`;
  user_Pictures.className = "user_Picures";
  let boxMessage = document.createElement("div");
  boxMessage.className = "receiveSaid";

  if (element.Content != undefined) {
    if (element.Content != "") {
      let Message = document.createElement("p");
      Message.className = "Said";

      Message.textContent = element.Content;

      boxMessage.appendChild(Message);
    }
  } else {
    if (element.Socket.value != "") {
      let Message = document.createElement("p");
      Message.className = "Said";

      Message.textContent = element.Socket.value;

      boxMessage.appendChild(Message);
    }
  }

  if (element.Image == undefined) {
    if (element.Socket.image != "") {
      try {
        const file = base64ToFile(element.Socket.image, "example.png");
        let message_Image = document.createElement("img");
        message_Image.className = "image_chat_receive";
        message_Image.src = URL.createObjectURL(file);
        boxMessage.appendChild(message_Image);
      } catch (error) {
        console.error("An error occurred:", error);
      }
    }
  } else {
    if (element.Image != "") {
      try {
        const file = base64ToFile(element.Image, "example.png");
        let message_Image = document.createElement("img");
        message_Image.className = "image_chat_receive";
        message_Image.src = URL.createObjectURL(file);
        boxMessage.appendChild(message_Image);
      } catch (error) {
        console.error("An error occurred:", error);
      }
    }
  }

  divMessage.appendChild(user_Pictures);
  divMessage.appendChild(boxMessage);

  bigdiv.appendChild(divMessage);

  document.getElementById("chat").appendChild(bigdiv);
}
