
import api from './apiconfig.js';

export async function followRequestJs(iduser, idtofollow, context, isInModal){  

    try {
        const response = await api.post("/api/follow", {iduser :iduser, idtofollow:idtofollow })
        if (response.data == true){
            addingOneFollower()
            if (isInModal==true){
                const follower = context.allfollowers.find(f => f.Id === idtofollow);
                if (follower) {
                    follower.HeIsFollowing = true; 
                }
                if (context.idUser == iduser){
                    addingOneFollowing()
                }
            }
            return true
        } else {

            return false
        }
    } catch(error){
        console.log(error)
        return false
    }
}


export async function removeTheUserInTheList(context, idtounfollow){

    const response = await api.post('/api/removefollower', {iduser:context.userState.id, idunfollow: idtounfollow})
    if (response.data == true){
        if (context.allfollowers){
            context.allfollowers = context.allfollowers.filter(f => f.Id !== idtounfollow);
        }
        return true
        
    } else {
        return false
    }
}



export async function unFollowRequestJs(iduser, idtounfollow,context, isInModal, typeoffollower){
    try {
        const response = await api.post("/api/unfollow", {iduser :iduser, idunfollow:idtounfollow })
        if (response.data == true){
            removeOneFollower()
            if (isInModal==true){
                const follower = context.allfollowers.find(f => f.Id === idtounfollow);
                if (follower) {
                    follower.HeIsFollowing = false;   
                }

                if (context.idUser == iduser){
                    if (typeoffollower == "following"){
                        context.allfollowers = context.allfollowers.filter(f => f.Id !== idtounfollow);
                    }
                    removeOneFollowing()
                }
            }
            return true
        } else {
            return false
        }
    } catch(error){
        console.log(error)
        return false
    }

}


export async function requestFollowJs(idWaiting, idRequest, context, isInModal){
    console.log(idWaiting, idRequest)
    try{
        const response = await api.post("/api/requestfollow", {iduser :idWaiting, idtofollow:idRequest })
        if (response.data == true){
            if (isInModal==true){
                const follower = context.allfollowers.find(f => f.Id === idRequest);
                if (follower) {
                    follower.HeIsWaiting = true; 
                }
            }
            return true
        } else {
            return false
        }
    }   
    catch(error){
        console.log(error)
        return false
    }
}




export async function unRequestFollowJs(idWaiting, idRequest, context,isInModal){
   
    try{
        const response = await api.post("/api/unrequestfollow", {iduser :idWaiting, idunfollow:idRequest })
        if (response.data == true){
            
            if (isInModal==true){
                const follower = context.allfollowers.find(f => f.Id === idRequest);
                if (follower) {
                    follower.HeIsWaiting = false; 
                }
            }
            return true
        } else {
            return false
        }
    }   
    catch(error){
        console.log(error)
        return false
    }
}   



export async function fetchWaitingFollowers(iduser){
    try{    
        const response = await api.post("/api/waitingFollowers")
        const divContainer = document.getElementById('modal-followers-waiting')
    

        response.data.forEach(element => {
            const divUser = document.createElement('div')
            divUser.className = "user-button-waiting"

            const img=document.createElement('img')
            img.src=element.Pictures

            const divButtonsName = document.createElement('div')
            divButtonsName.className = "div-all-in-divUser"

            const divButtonsRows = document.createElement('div')
            divButtonsRows.className= "div-buttons-in-divUser"

            const name=document.createElement('p')
            name.textContent = element.Username

            const buttonDecline = document.createElement('button')
            buttonDecline.classList = "button-user-waiting decline-user"
            buttonDecline.textContent = "decline"
            buttonDecline.addEventListener("click", async ()=>{
                try{
                    const responseDecline = await api.post("/api/declinefollower", {idtodecline:element.Id, iduser:iduser})
                    if (responseDecline.data == true){
                        divUser.remove()
                    }
                }catch(error){
                    console.log(error)
                }
            })

            const buttonAccept = document.createElement('button')
            buttonAccept.classList = "button-user-waiting accept-user"
            buttonAccept.textContent = "accept"
            buttonAccept.addEventListener("click", async ()=>{
                try{
                    const responseAccept = await api.post("/api/acceptfollower", {idtoaccept:element.Id, iduser:iduser})
                    if (responseAccept.data==true){
                        addingOneFollower()
                        divUser.remove()
                        
                    }
                } catch(error){
                    console.log(error)
                }
            })

            divButtonsRows.append(buttonDecline, buttonAccept)

            divButtonsName.append(name, divButtonsRows)

            divUser.append(img, divButtonsName)
            
            divContainer.appendChild(divUser)
        });

        return

    }catch{
        console.log(error)
        return
    }
}




function removeOneFollowing(){
    const followElement = document.getElementById("Following");
    const currentCount = parseInt(followElement.textContent, 10) || 0;
    followElement.textContent = currentCount -1;
}

function addingOneFollowing(){
    const followElement = document.getElementById("Following");
    const currentCount = parseInt(followElement.textContent, 10) || 0;
    followElement.textContent = currentCount + 1;
}

function removeOneFollower(){
    const followElement = document.getElementById("Followers");
    const currentCount = parseInt(followElement.textContent, 10) || 0;
    followElement.textContent = currentCount -1;
}

function addingOneFollower(){
    const followElement = document.getElementById("Followers");
    const currentCount = parseInt(followElement.textContent, 10) || 0;
    followElement.textContent = currentCount + 1;
}

export async function fetchFollowersFromTheProfile(iduser, userstateid, whatsWeLookingfor){
    try{
        const response = await api.post("/api/fetchfollowers", {iduser: iduser,iduserstate: userstateid, type:whatsWeLookingfor })
        if (response){  
            console.log(response.data)
            
            return response.data 
        }
    } catch(error){
        console.log(error)
        return
    }
}