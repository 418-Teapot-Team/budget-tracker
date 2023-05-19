<template>
  <div class="rounded-2xl shadow-sm pb-3 px-6 w-full bg-white transactions-wrapper">
    <div class="flex flex-col justify-start items-start border-b border-b-gray-100 bg-white pt-3">
      <span class="text-black text-xl"> Transactions </span>
      <span class="text-gray-500 text-sm font-thin"> Recent expenses </span>
    </div>
    <div class="flex flex-col w-full pt-2 pr-2 gap-6 overflow-y-auto transactions">
      <!-- item -->
      <div class="flex flex-row" v-for="item in data" :key="item.id">
        <div class="flex flex-col w-full">
          <span class="text-black text-lg">{{ item?.category?.category }}</span>
          <span class="text-gray-500 text-xs font-thin">
            {{
              new Date(item.createdAt?.Time).getDate() +
              '.' +
              (new Date(item.createdAt?.Time).getMonth() < 10
                ? '0' + new Date(item.createdAt?.Time).getMonth()
                : new Date(item.createdAt?.Time).getMonth()) +
              '.' +
              new Date(item.createdAt?.Time).getFullYear()
            }}</span
          >
        </div>
        <span
          class="text-lg font-thin self-center justify-end"
          :class="item.type === 'expenses' ? 'text-red-500' : 'text-green-500'"
          >{{ item.type === 'expenses' ? '-' : '+' }}{{ item.amount }}</span
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LastTransactions',
  props: {
    data: Array,
  },
};
</script>

<style>
.transactions-wrapper {
  max-height: calc(100vh - 130px);
  min-height: calc(100vh - 130px);
}

.transactions {
  max-height: calc(100vh - 200px);
}

@media screen and (min-width: 1280px) {
  .transactions-wrapper {
    max-height: calc(100vh - 114px);
    min-height: calc(100vh - 114px);
  }

  .transactions {
    max-height: calc(100vh - 187px);
  }
}
@media screen and (min-width: 1536px) {
  .transactions-wrapper {
    max-height: calc(100vh - 120px);
    min-height: calc(100vh - 120px);
  }

  .transactions {
    max-height: calc(100vh - 190px);
  }
}

.transactions::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}

.transactions::-webkit-scrollbar-track {
  border-radius: 100vh;
  background: transparent;
}

.transactions::-webkit-scrollbar-thumb {
  background: #1c1c21;
  border-radius: 100vh;
  border: 3px solid #f8f9fe;
}

.transactions::-webkit-scrollbar-thumb:hover {
  background: #1c1c21;
}
</style>
