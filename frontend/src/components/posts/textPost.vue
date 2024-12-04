<template>
    <Darkmode />
    <div class="containerOnlyText-post">
        <div class="headerProfile-post">
            <div class="profilePicture-post">
                <img id="profile" name="profile">
                <label id="name" for="profile">Name</label>
            </div>
            <p class="hours-post">Heure</p>
            <p>Date</p>
        </div>
        <div class="onlyText-post">
            <p>Lorem, ipsum Quaerat, blanditiis maiores. Rem laudantium enim aspernatur officia laborum. Commodi t, blanditiis maiores. Rem laudantium enim aspernatur officia laborum. Commodi laboriosam corrupti u t, blanditiis maiores. Rem laudantium enim aspernatur officia laborum. Commodi laboriosam corrupti u laboriosam corrupti ut vitae quibusdam sint iste impedit sequi deserunt? Laudantium iste quia enim ab, aut reprehenderit nulla! Iusto vel id quo pariatur eligendi. Totam?</p>
        </div>
    </div>

    <div id="modalText" class="modal">
        <div class="modal-content">
            <span class="close">&times</span>
            <p id="modal-text"></p>
        </div>
    </div>

</template>

<script>
import { loadingPost } from '../../services/postServices';

export default {
    name: 'textPost',
}
</script>