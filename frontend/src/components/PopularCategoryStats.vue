<template>
  <div class="w-24 h-24 xl:w-40 xl:h-40 2xl:w-48 2xl:h-48 rounded-2xl p-6 shadow-sm cursor-pointer">
    <div class="flex flex-col justify-between items-start h-full">
      <div
        class="hidden xl:w-12 xl:h-12 2xl:w-16 2xl:h-16 rounded-full bg-white xl:flex flex-row justify-center items-center p-4"
      >
        <img :src="imageUrl" alt="img" class="bg-white w-full h-full" />
      </div>
      <div class="flex flex-col justify-center items-center xl:items-start gap-1">
        <span class="uppercase text-sm xl:text-2xl 2xl:text-3xl text-black">₴{{ amount }}</span>
        <span class="text-gray-400 font-thin text-xs xl:text-sm xl:pl-2">{{ title }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PopularCategoryStats',
  props: {
    imageUrl: String,
    amount: Number,
    title: String,
  },
};
</script>
