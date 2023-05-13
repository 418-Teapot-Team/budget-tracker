<template>
  <section class="flex flex-row justify-start gap-0 h-screen">
    <div
      class="w-1/3 bg-black flex flex-col justify-center items-center gap-5 auth-bg rounded-r-3xl"
    >
      <TeapotIcon />
      <div class="flex flex-col justify-start gap-1">
        <span class="text-3xl text-white text-center"> Welcome </span>
        <span class="text-xl text-white text-center">to Budget Tracker</span>
      </div>
      <span class="text-gray-500 text-sm justify-self-end">All rights reserved C 418 Teapot</span>
    </div>
    <div class="w-full flex justify-center items-center">
      <div class="w-80 lg:w-96">
        <div class="flex flex-row justify-center gap-2 mb-10">
          <button
            class="bg-black border-2 border-black text-white p-2 text-center w-full rounded-xl"
          >
            Вхід
          </button>
          <button
            class="border-2 border-black text-black p-2 text-center w-full rounded-xl transition duration-300 hover:bg-black hover:text-white"
          >
            Реєстрація
          </button>
        </div>
        <auth-form />
      </div>
    </div>
  </section>
</template>

<script>
import TeapotIcon from '@/components/Icons/Teapot.vue';
import AuthForm from '@/components/AuthForm.vue';
export default {
  name: 'Auth',
  components: {
    TeapotIcon,
    AuthForm,
  },
};
</script>

<style scoped>
.auth-bg {
  background-image: url(/assets/images/auth-bg.jpg);
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}
</style>
