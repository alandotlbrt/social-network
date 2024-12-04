<template>
   <DarkMode />
   <div class="containerPost-home">
      <imagePost />
      <textImagePost />
      <textPost />
   </div>

</template>

<script>
import { loadingPost } from "../services/postServices"

export default{
   name:"loadingPost",
   mounted() {
      loadingPost();
   }
}

</script>


