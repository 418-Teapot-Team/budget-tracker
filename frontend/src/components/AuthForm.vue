<template>
  <vee-form class="w-full">
    <div class="flex flex-col justify-center gap-1 mb-3">
      <label class="text-black text-left">E-mail</label>
      <vee-field
        name="email"
        type="email"
        placeholder="Email"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
      <ErrorMessage />
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <label class="text-black text-left">Password</label>
      <vee-field
        name="password"
        type="password"
        placeholder="Password"
        class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
      >
      </vee-field>
      <ErrorMessage />
    </div>
    <div class="flex flex-col justify-center gap-1 mb-3">
      <button
        type="submit"
        class="bg-black border-2 border-black text-white p-2 text-center w-full rounded-xl mt-6"
      >
        Увійти
      </button>
    </div>
  </vee-form>
</template>

<script>
export default {
  name: 'AuthForm',
};
</script>
