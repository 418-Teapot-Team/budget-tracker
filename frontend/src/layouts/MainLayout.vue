<template>
  <main class="flex flex-row justify-start gap-14 pr-14">
    <app-menu />
    <div id="content" class="max-h-screen overflow-y-auto w-full pt-6">
      <div class="flex flex-row justify-between items-center mb-6">
        <ul class="flex flex-row gap-8">
          <li
            class="flex flex-col justify-start items-start border-r border-black border-opacity-30 w-fit pr-8"
          >
            <div class="flex flex-row justify-start items-center gap-1">
              <div class="w-6 h-6">
                <USAFlagIcon />
              </div>
              <span class="text-black">USD</span>
            </div>
            <div class="flex flex-row gap-1">
              <span class="text-xs text-black opacity-80">Official- 36.5686</span>
              <span class="text-xs text-black opacity-80">Marketplace - 37.77</span>
            </div>
          </li>
          <li
            class="flex flex-col justify-start items-start border-r border-black border-opacity-30 w-fit pr-8"
          >
            <div class="flex flex-row justify-start items-center gap-1">
              <div class="w-6 h-6">
                <EUFlagIcon />
              </div>
              <span class="text-black">EUR</span>
            </div>
            <div class="flex flex-row gap-1">
              <span class="text-xs text-black opacity-80">Official- 39.8287</span>
              <span class="text-xs text-black opacity-80">Marketplace - 41.35</span>
            </div>
          </li>
          <li class="flex flex-col justify-start items-start">
            <div class="flex flex-row justify-start items-center gap-1">
              <div class="w-6 h-6">
                <PLFlagIcon />
              </div>
              <span class="text-black">PLN</span>
            </div>
            <div class="flex flex-row gap-1">
              <span class="text-xs text-black opacity-80">Official- 8.785</span>
              <span class="text-xs text-black opacity-80">Marketplace - 9.0</span>
            </div>
          </li>
        </ul>
        <span
          class="text-black font-bold text-xl w-72 h-7 text-end text-ellipsis whitespace-nowrap overflow-hidden"
        >
          Welcome, Username!
        </span>
      </div>
      <router-view />
    </div>
  </main>
</template>

<script>
import AppMenu from '@/components/AppMenu.vue';
import USAFlagIcon from '@/components/Icons/USAFlagIcon.vue';
import EUFlagIcon from '@/components/Icons/EUFlagIcon.vue';
import PLFlagIcon from '@/components/Icons/PLFlagIcon.vue';
export default {
  name: 'MainLayout',
  components: {
    AppMenu,
    USAFlagIcon,
    EUFlagIcon,
    PLFlagIcon,
  },
};
</script>
