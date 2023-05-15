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
          <app-button text="Вхід" :isOutline="tab !== 'login'" @click="tab = 'login'" />
          <app-button text="Реєстрація" :isOutline="tab !== 'reg'" @click="tab = 'reg'" />
        </div>
        <auth-form v-if="tab === 'login'" :isLoading="isLoading" @onSubmit="loginSubmit" />
        <registration-form v-if="tab === 'reg'" :isLoading="isLoading" @onSubmit="regSubmit" />
      </div>
    </div>
  </section>
</template>

<script>
import TeapotIcon from '@/components/Icons/Teapot.vue';
import AuthForm from '@/components/AuthForm.vue';
import RegistrationForm from '@/components/RegistrationForm.vue';
import { mapState, mapActions } from 'pinia';
import useAuthStore from '@/stores/auth';

export default {
  name: 'Auth',
  components: {
    TeapotIcon,
    AuthForm,
    RegistrationForm,
  },
  data() {
    return {
      tab: 'login',
    };
  },
  watch: {
    isSuccess(val) {
      if (val === true) {
        this.$toast.success(this.message);
      }
      if (this.isAuthorized === true) {
        this.$router.push('/');
      }
    },
    isError(val) {
      if (val === true) {
        this.$toast.error(this.error?.message);
      }
    },
  },
  computed: {
    ...mapState(useAuthStore, [
      'user',
      'error',
      'message',
      'isAuthorized',
      'isLoading',
      'isError',
      'isSuccess',
    ]),
  },
  methods: {
    ...mapActions(useAuthStore, ['login', 'register']),
    regSubmit(values) {
      this.register({ fullName: values.name, email: values.email, password: values.password });
    },
    loginSubmit(values) {
      this.login({ email: values.email, password: values.password });
    },
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
