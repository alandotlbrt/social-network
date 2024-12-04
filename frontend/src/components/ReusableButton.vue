<template>
  <div @click="onClick" class="reusableButton">
     {{ label }}
  </div>
</template>

<script>
   export default {
     props: {
       label: {
         type: String,
         required: true
       },
       onClick: {
         type: Function,
         required: true
       }
     }
   };
</script>
