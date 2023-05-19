<template>
  <div class="bg-black bg-opacity-70 fixed top-0 bottom-0 left-0 right-0 flex">
    <div class="bg-white shadow-sm rounded-2xl m-auto w-80 h-fit p-8 relative">
      <div class="absolute right-3 top-3 w-7 h-7 cursor-pointer" @click="closePopup">
        <close-icon />
      </div>
      <vee-form class="h-full">
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Title</label>
            <ErrorMessage name="title" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="title"
            type="text"
            placeholder="Title"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.title"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Month count</label>
            <ErrorMessage name="monthCount" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="monthCount"
            type="number"
            placeholder="Month count"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.monthCount"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Percentage</label>
            <ErrorMessage name="percentage" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="percentage"
            type="number"
            placeholder="Percentage"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.percentage"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Date</label>
            <ErrorMessage name="date" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="percentage"
            type="date"
            placeholder="Date"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.date"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Amount</label>
            <ErrorMessage name="percentage" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="amount"
            type="number"
            placeholder="Amount"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.amount"
          />
        </div>
        <div class="h-10">
          <app-button :text="isEdit ? 'Update' : 'Add'" :isOutline="false" type="submit" />
        </div>
      </vee-form>
    </div>
  </div>
</template>

<script>
import CloseIcon from '@/components/Icons/CloseIcon.vue';
export default {
  name: 'BankingPopup',
  components: {
    CloseIcon,
  },
  props: {
    isEdit: Boolean,
    defaltValues: Object,
  },
  watch: {
    defaltValues(values) {
      if (this.isEdit) {
        this.initialValues.title = values?.title ? values.title : '';
        this.initialValues.amount = values?.monthCount ? values.monthCount : 0;
        this.initialValues.comment = values?.percentage ? values.percentage : 0;
        this.initialValues.date = values?.date ? values.date : '2004-05-22';
        this.initialValues.amount = values?.amount ? values.amount : 0;
      }
    },
  },
  data() {
    return {
      initialValues: {
        title: '',
        monthCount: 0,
        percentage: 0,
        date: '2004-05-22',
        amount: 0,
      },
    };
  },

  methods: {
    closePopup() {
      this.$emit('onClose');
    },
    submitForm(values) {
      if (this.isEdit) {
        this.$emit('onEdit', values);
      } else {
        this.$emit('onAdd', values);
      }
    },
  },
};
</script>
