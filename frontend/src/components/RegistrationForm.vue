<template>
  <vee-form class="w-full" :validation-schema="schema" @submit="onSubmit">
    <div class="flex flex-col justify-center gap-1 mb-3">
      <div class="flex justify-between gap-2">
        <label class="text-black text-left">Name</label>
        <ErrorMessage name="name" class="text-red-600" />
      </div>
      <vee-field
        name="name"
        type="text"
        placeholder="Name"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <div class="flex justify-between gap-2">
        <label class="text-black text-left">E-mail</label>
        <ErrorMessage name="email" class="text-red-600" />
      </div>
      <vee-field
        name="email"
        type="email"
        placeholder="Email"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <div class="flex justify-between gap-2">
        <label class="text-black text-left">Password</label>
        <ErrorMessage name="password" class="text-red-600" />
      </div>
      <vee-field
        name="password"
        type="password"
        placeholder="Password"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <div class="flex justify-between gap-2">
        <label class="text-black text-left">Confirm password</label>
        <ErrorMessage name="confirm_password" class="text-red-600" />
      </div>
      <vee-field
        name="confirm_password"
        type="password"
        placeholder="Confirm password"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
      <ErrorMessage />
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <app-button text="Зареєструватись" v-if="!isLoading" :isOutline="false" type="submit" />
      <div class="w-10 h-10 flex justify-center items-center" v-else>
        <preloader />
      </div>
    </div>
  </vee-form>
</template>

<script>
export default {
  name: 'RegistrationForm',
  props: {
    isLoading: Boolean,
  },
  data() {
    return {
      schema: {
        name: 'required|alpha_spaces|min:2|max:255',
        email: 'required|email|max:255',
        password: 'required|min:6|max:100',
        confirm_password: 'passwords_mismatch:@password',
      },
    };
  },
  methods: {
    onSubmit(values) {
      this.$emit('onSubmit', values);
    },
  },
};
</script>
