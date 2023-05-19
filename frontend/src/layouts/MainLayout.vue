<template>
  <main class="flex flex-row justify-start gap-14 pr-14">
    <app-menu />
    <div id="content" class="max-h-screen overflow-y-auto w-full pt-6">
      <div class="flex flex-row justify-between items-center mb-6">
        <ul class="flex flex-row gap-8" v-if="courses?.length">
          <li
            class="flex flex-col justify-start items-start border-r border-black border-opacity-30 w-fit pr-8 last:border-r-0"
            v-for="item in courses"
            :key="item?.code"
          >
            <div class="flex flex-row justify-start items-center gap-1">
              <div class="w-6 h-6">
                <USAFlagIcon v-if="item?.code === 'USD'" />
                <EUFlagIcon v-if="item?.code === 'EUR'" />
                <PLFlagIcon v-if="item?.code === 'PLN'" />
              </div>
              <span class="text-black">{{ item?.code }}</span>
            </div>
            <div class="flex flex-row gap-1">
              <span class="text-xs text-black opacity-80">Buy - {{ item?.buyRate }}</span>
              <span class="text-xs text-black opacity-80">Sell - {{ item?.saleRate }}</span>
            </div>
          </li>
        </ul>
        <div class="flex flex-row justify-end gap-2">
          <span
            class="text-black font-bold text-xl w-72 h-7 text-end text-ellipsis whitespace-nowrap overflow-hidden"
          >
            Welcome, {{ user?.fullName }}!
          </span>
          <logout-icon class="cursor-pointer" @click="logout" />
        </div>
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
import { mapState, mapActions } from 'pinia';
import useAuthStore from '@/stores/auth';
import useDashboardStore from '@/stores/dashboard';
import { useToast } from 'vue-toastification';
import LogoutIcon from '@/components/Icons/LogoutIcon.vue';

const toast = useToast();

export default {
  name: 'MainLayout',
  components: {
    AppMenu,
    USAFlagIcon,
    EUFlagIcon,
    PLFlagIcon,
    LogoutIcon,
  },
  computed: {
    ...mapState(useAuthStore, ['user']),
    ...mapState(useDashboardStore, ['courses']),
  },
  watch: {
    isAuthorized: {
      handler() {
        if (!!localStorage.getItem('tracker-auth-token') === false) {
          this.$router.push('/auth');
        }
      },
      immediate: true,
    },
    '$route.path'() {
      if (!!localStorage.getItem('tracker-auth-token') === false) {
        this.$router.push('/auth');
      }
    },
  },
  methods: {
    ...mapActions(useAuthStore, ['whoAmI']),
    ...mapActions(useDashboardStore, ['getCourses']),
    logout() {
      localStorage.removeItem('tracker-auth-token');
      this.$router.push('/auth');
    },
  },
  mounted() {
    if (!localStorage.getItem('tracker-auth-token')) {
      this.$router.push('/auth');
    }
    if (localStorage.getItem('tracker-auth-token')) {
      try {
        this.whoAmI();
        this.getCourses();
      } catch (e) {
        toast.error(e.message);
      }
    }
  },
};
</script>
