<template>
  <main>
    <RouterView />
  </main>
</template>

<script>
import { RouterView } from 'vue-router';
export default {
  name: 'App',
  components: {
    RouterView,
  },
};
</script>

<style scoped></style>
