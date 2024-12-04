<template>
    <DarkMode />
   
    <div class="containerOnlyImage-post">
        <div class="headerOnlyImage-post">
            <div class="profilePicture-post">
                <img id="profile" name="profile">
                <label id="name" for="profile">Name</label>
            </div>
            <p class="hours-post">Heure</p>
            <p>Date</p>
        </div>
        <div class="onlyImage-post">
            <img id="img">
        </div>
    </div>

</template>

<script>
import { loadingPost } from '../../services/postServices';
import { displayContentPost } from '../../services/postServices';

export default {
    name: 'imagePost',
    data() {
        return {
            responseData: null,
        };
    },
    methods: {
       
    },
    mounted(){
        loadingPost().then(response => {
            this.responseData = response
            const img = document.getElementById('img')
           
            img.src = this.responseData.data[1].Image
        }).catch(error => {
            console.log("Erreur lors de la récupération du post", error)
        });
        displayContentPost()
    }
}

</script>