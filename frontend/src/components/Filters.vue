<template>
  <vee-form class="flex flxe-row justify-end gap-4 w-1/3 h-8" @submit="applyFilters">
    <vee-field
      v-if="withCategory"
      class="w-4/12 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white text-black text-sm"
      name="category"
      as="select"
      v-model="category"
    >
      <option value="all" selected>All</option>
      <option v-for="item in categories" :key="item.id" :value="item.id">
        {{ item.category }}
      </option>
    </vee-field>
    <vee-field
      class="w-4/12 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white text-black text-sm"
      name="filter"
      as="select"
      v-model="filter"
    >
      <option value="default" selected>Default</option>
      <option v-for="item in currentFilters" :key="item.value.id" :value="item.value.id">
        {{ item.title }}
      </option>
    </vee-field>
    <div class="w-4/12">
      <app-button text="Apply" type="submit" />
    </div>
  </vee-form>
</template>
<script>
import useFiltersStore from '@/stores/filters';
import { mapState } from 'pinia';
export default {
  name: 'Filters',
  props: {
    withCategory: Boolean,
    isStats: Boolean,
    categories: Array,
  },
  data() {
    return {
      category: 'all',
      filter: 'default',
    };
  },
  computed: {
    ...mapState(useFiltersStore, ['filters', 'filtersForStats']),
    currentFilters() {
      if (this.isStats) {
        return this.filtersForStats;
      } else {
        return this.filters;
      }
    },
  },
  methods: {
    applyFilters(values) {
      console.log(values, this.currentFilters);
      if (this.category && this.filter) {
        this.$emit('onApplyFilters', { category: this.category, filter: this.filter });
      }
    },
  },
  mounted() {
    const filter = this.$route.query?.filterId ? this.$route.query?.filterId : 'default';
    const category =
      this.$route.query?.categoryId && this.$route.query?.categoryId !== 'all'
        ? this.$route.query?.categoryId
        : 'all';
    this.category = category;
    this.filter = filter;
  },
};
</script>
