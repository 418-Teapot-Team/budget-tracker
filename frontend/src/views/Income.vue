<template>
  <section>
    <!-- filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Income</span>
        <div class="h-8 w-8 cursor-pointer"><add-icon /></div>
      </div>
      <vee-form class="flex flxe-row gap-4 w-1/3 h-8">
        <vee-field
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
        <app-button class="w-4/12" text="Filter" type="submit" />
      </vee-form>
    </div>
    <!-- table -->
    <div class="pt-4 w-full"><finances-table /></div>
  </section>
</template>

<script>
import FinancesTable from '@/components/FinancesTable.vue';
import AddIcon from '@/components/Icons/AddIcon.vue';
export default {
  name: 'Income',
  components: {
    FinancesTable,
    AddIcon,
  },
  data() {
    return {
      category: 'default',
      filter: 'date',
    };
  },
};
</script>
