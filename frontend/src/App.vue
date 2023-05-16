<template>
  <component :is="layout">
    <RouterView />
  </component>
</template>

<script>
import { RouterView } from 'vue-router';
import MainLayout from '@/layouts/MainLayout.vue';
import EmptyLayout from '@/layouts/EmptyLayout.vue';

// Chart.js config
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
);

export default {
  name: 'App',
  components: {
    RouterView,
    MainLayout,
    EmptyLayout,
  },
  computed: {
    layout() {
      return (this.$route.meta.layout || 'empty') + '-layout';
    },
  },
};
</script>

<style scoped></style>
