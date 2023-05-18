<template>
  <div
    class="flex flex-row justify-start items-center w-full h-36 xl:h-40 2xl:h-44 rounded-2xl bg-black pl-6 gap-7 shadow-sm"
  >
    <div
      class="w-28 h-28 xl:w-32 xl:h-32 bg-white rounded-full p-6 flex flex-row justify-center items-center"
    >
      <StatsDownIcon v-if="amount < 0" />
      <StatsUpIcon v-else />
    </div>
    <div class="flex flex-col items-start justify-center gap-2 font-normal">
      <span class="text-gray-500 text-lg">This month you saved</span>
      <span class="text-white text-4xl font-bold tracking-widest"
        >{{ amount > 0 ? `+${amount}` : amount }}
        <span class="uppercase font-normal tracking-normal text-3xl">uah</span></span
      >
      <span class="text-gray-500 text-xs flex"
        >To get more details visit
        <router-link to="/statistic" class="mx-1 cursor-pointer hover:text-gray-200" exact>
          <span class="underline"> statistic </span>
        </router-link>
        page</span
      >
    </div>
  </div>
</template>

<script>
import StatsDownIcon from '@/components/Icons/StatsDownIcon.vue';
import StatsUpIcon from '@/components/Icons/StatsUpIcon.vue';
export default {
  name: 'CurrentMonthBanner',
  props: {
    amount: Number,
  },
  components: {
    StatsDownIcon,
    StatsUpIcon,
  },
};
</script>
