<template>
  <main class="flex flex-row justify-start gap-14">
    <app-menu />
    <section id="content" class="max-h-screen overflow-y-auto w-full pt-6">
      <router-view />
    </section>
  </main>
</template>

<script>
import AppMenu from '@/components/AppMenu.vue';
export default {
  name: 'MainLayout',
  components: {
    AppMenu,
  },
};
</script>
