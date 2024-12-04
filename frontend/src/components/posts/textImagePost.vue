<template>
    <Darkmode />

    <div class="containerOnlyImage-post">
        <div class="headerProfile-post">
            <div class="profilePicture-post">
                <img id="profile" name="profile">
                <label id="name" for="profile">Name</label>
            </div>
            <p class="hours-post">Heure</p>
            <p>Date</p>
        </div>
        <div class="imageAndText-post">
            <div class="text">
                <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Totam, temporibus?</p>
            </div>
            <div class="image">
                <img id="img1" class="Image-post">
            </div>
        </div>
    </div>

     
</template>

<script>
import { loadingPost } from '../../services/postServices';

export default {
    name: 'textImagePost',
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
            const img1 = document.getElementById('img1')
            img1.src = this.responseData.data[1].Image
        }).catch(error => {
            console.log("Erreur lors de la récupération du post", error)
        })
    }
}

</script>