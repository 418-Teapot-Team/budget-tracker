<template>
  <div
    class="bg-black shadow-sm text-white rounded-2xl p-2 flex flex-col justify-center items-center"
  >
    <div class="flex flex-col justify-center items-start gap-3">
      <span class="text-3xl font-bold">Total: </span>
      <span class="text-4xl font-thin">{{ amount }} UAH</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BalanceCard',
  props: {
    amount: Number,
  },
};
</script>
