<template>
  <vee-form class="w-full" :validation-schema="schema" @submit="onSubmit">
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
    <div class="flex flex-col justify-center gap-1 mb-3 h-10">
      <app-button v-if="!isLoading" text="Sign In" :isOutline="false" type="submit" />
      <div class="w-full flex justify-center items-center" v-else>
        <div class="w-10 h-10">
          <Preloader />
        </div>
      </div>
    </div>
  </vee-form>
</template>

<script>
export default {
  name: 'AuthForm',
  props: {
    isLoading: Boolean,
  },
  data() {
    return {
      schema: {
        email: 'required|email|max:255',
        password: 'required|min:6|max:100',
      },
    };
  },
  methods: {
    onSubmit(values, { resetForm }) {
      this.$emit('onSubmit', { values, resetForm });
    },
  },
};
</script>
