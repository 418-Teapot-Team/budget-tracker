<template>
  <section class="flex flex-row justify-center xl:justify-start gap-0 h-screen">
    <div
      class="hidden w-1/3 bg-black xl:flex xl:flex-col justify-center items-center gap-5 auth-bg rounded-r-3xl"
    >
      <TeapotIcon fill="#000" />
      <div class="flex flex-col justify-start gap-1">
        <span class="text-3xl text-white text-center"> Welcome </span>
        <span class="text-xl text-white text-center">to Budget Tracker</span>
      </div>
      <span class="text-gray-500 text-sm justify-self-end">All rights reserved © 418 Teapot</span>
    </div>
    <div class="w-full flex justify-center items-center">
      <div class="w-80 lg:w-96">
        <div class="flex flex-row justify-center gap-2 mb-6 h-10">
          <app-button text="Sign In" :isOutline="tab !== 'login'" @click="tab = 'login'" />
          <app-button text="Sign Up" :isOutline="tab !== 'reg'" @click="tab = 'reg'" />
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
import { useToast } from 'vue-toastification';

const toast = useToast();

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
      isLoading: false,
    };
  },
  computed: {
    ...mapState(useAuthStore, ['user', 'isAuthorized']),
  },
  methods: {
    ...mapActions(useAuthStore, ['login', 'register']),
    async regSubmit({ values, resetForm }) {
      try {
        this.isLoading = true;
        await this.register({
          fullName: values.name,
          email: values.email,
          password: values.password,
        });
        toast.success('Successfully registered. Please log in');
        this.tab = 'login';
        resetForm();
      } catch (e) {
        toast.error(e.message);
      } finally {
        this.isLoading = false;
      }
    },
    async loginSubmit({ values, resetForm }) {
      try {
        this.isLoading = true;
        await this.login({ email: values.email, password: values.password });
        resetForm();
        toast.success('Welcome!');
        this.$router.push('/');
      } catch (e) {
        toast.error(e.message);
      } finally {
        this.isLoading = false;
      }
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
