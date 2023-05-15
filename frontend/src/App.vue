<template>
  <component :is="layout">
    <RouterView />
  </component>
</template>

<script>
import { RouterView } from 'vue-router';
import MainLayout from '@/layouts/MainLayout.vue';
import EmptyLayout from '@/layouts/EmptyLayout.vue';
export default {
  name: 'App',
  components: {
    RouterView,
    MainLayout,
    EmptyLayout,
  },
  computed: {
    layout() {
      return (this.$route.meta.layout || 'empty') + '-layout';
    },
  },
};
</script>

<style scoped></style>
