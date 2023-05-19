<template>
  <div class="bg-black bg-opacity-70 fixed top-0 bottom-0 left-0 right-0 flex">
    <div class="bg-white shadow-sm rounded-2xl m-auto w-80 h-fit p-8 relative">
      <div class="absolute right-3 top-3 w-7 h-7 cursor-pointer" @click="closePopup">
        <close-icon />
      </div>
      <vee-form class="h-full" :validation-schema="schema" @submit="submitForm">
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Title</label>
            <ErrorMessage name="name" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="name"
            type="text"
            placeholder="Title"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.name"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Month count</label>
            <ErrorMessage name="monthAmount" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="monthAmount"
            type="number"
            placeholder="Month count"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.monthAmount"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Percentage</label>
            <ErrorMessage name="percent" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="percent"
            type="number"
            placeholder="Percentage"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.percent"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Date</label>
            <ErrorMessage name="date" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="date"
            type="date"
            placeholder="Date"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.date"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Amount</label>
            <ErrorMessage name="sum" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="sum"
            type="number"
            placeholder="Amount"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.sum"
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
    defaltValues: {
      handler(values) {
        if (this.isEdit) {
          this.initialValues.name = values?.name ? values.name : '';
          this.initialValues.monthAmount = values?.monthAmount ? values.monthAmount : 0;
          this.initialValues.percent = values?.percent ? values.percent : 0;
          this.initialValues.date = values?.date ? values.date : '2004-05-22';
          this.initialValues.sum = values?.sum ? values.sum : 0;
        }
      },
      immediate: true,
    },
  },
  data() {
    return {
      initialValues: {
        name: '',
        monthAmount: 0,
        percent: 0,
        date: '2023-01-12',
        sum: 0,
      },
      schema: {
        name: 'required|min:3|max:100',
        monthAmount: 'required|min_value:1',
        percent: 'required|min_value:0|max_value:100',
        date: 'required',
        sum: 'required|min_value:0',
      },
    };
  },

  methods: {
    closePopup() {
      this.$emit('onClose');
    },
    submitForm(values) {
      if (this.isEdit) {
        this.$emit('onEdit', { ...values, id: this.defaltValues.id });
      } else {
        this.$emit('onAdd', values);
      }
    },
  },
};
</script>
