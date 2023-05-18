<template>
  <vee-form class="flex flxe-row justify-end gap-4 w-1/3 h-8" @submit="applyFilters">
    <vee-field
      v-if="withCategory"
      class="w-4/12 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white text-black"
      name="category"
      as="select"
      v-model="category"
    >
      <option value="default" selected>All</option>
      <option value="bsns">Bussiness</option>
      <option value="entrt">Enterteimant</option>
      <option value="oth">Others</option>
    </vee-field>
    <vee-field
      class="w-4/12 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white text-black"
      name="filter"
      as="select"
      v-model="filter"
    >
      <option value="date">Date</option>
      <option value="val">Value</option>
    </vee-field>
    <div class="w-4/12">
      <app-button text="Filter" type="submit" />
    </div>
  </vee-form>
</template>
<script>
export default {
  name: 'Filters',
  props: {
    withCategory: Boolean,
  },
  data() {
    return {
      category: 'default',
      filter: 'date',
    };
  },
  methods: {
    applyFilters() {
      if (this.category && this.filter) {
        this.$emit('onApplyFilters', { category: this.category, filter: this.filter });
      }
    },
  },
};
</script>
