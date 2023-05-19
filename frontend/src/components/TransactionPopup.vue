<template>
  <div class="bg-black bg-opacity-70 fixed top-0 bottom-0 left-0 right-0 flex">
    <div class="bg-white shadow-sm rounded-2xl m-auto w-80 h-fit p-8 relative">
      <div class="absolute right-3 top-3 w-7 h-7 cursor-pointer" @click="closePopup">
        <close-icon />
      </div>
      <vee-form class="h-full" :validation-schema="schema" @submit="submitForm">
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Category</label>
            <ErrorMessage name="category" class="text-red-600 text-xs" />
          </div>
          <vee-field
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            name="category"
            as="select"
            v-model="initialValues.category"
          >
            <option v-for="item in categories" :key="item.id" :value="item.id">
              {{ item.category }}
            </option>
          </vee-field>
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Amount</label>
            <ErrorMessage name="amount" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="amount"
            type="number"
            placeholder="Amount"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.amount"
          />
        </div>
        <div class="flex flex-col justify-center gap-1 mb-3">
          <div class="flex justify-between gap-2">
            <label class="text-black text-left">Comment</label>
            <ErrorMessage name="comment" class="text-red-600 text-xs" />
          </div>
          <vee-field
            name="comment"
            type="text"
            placeholder="Comment"
            class="block w-full h-10 py-1.5 px-3 border-b-2 border-black bg-grey-light transition duration-300 focus:outline-none focus:bg-white"
            v-model="initialValues.comment"
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
  name: 'TransactionPopup',
  components: {
    CloseIcon,
  },
  props: {
    isEdit: Boolean,
    defaltValues: Object,
    categories: Array,
  },
  data() {
    return {
      initialValues: {
        category: null,
        amount: '',
        comment: '',
      },
      schema: {
        category: 'required',
        amount: 'required|min_value:1',
        comment: 'max:255',
      },
    };
  },
  watch: {
    defaltValues: {
      handler(values) {
        if (this.isEdit) {
          this.initialValues.category = values?.categoryId
            ? values.categoryId
            : this.categories[0].id;
          this.initialValues.amount = values?.amount ? values.amount : '';
          this.initialValues.comment = values?.comment ? values.comment : '';
        }
      },
      immediate: true,
    },
    categories: {
      handler(values) {
        if (!this.initialValues.category) {
          this.initialValues.category = values[0].id;
        }
      },
      immediate: true,
    },
  },

  methods: {
    closePopup() {
      this.$emit('onClose');
      this.initialValues = {
        category: null,
        amount: '',
        comment: '',
      };
    },
    submitForm(values) {
      if (this.isEdit) {
        this.$emit('onEdit', {
          ...values,
          id: this.defaltValues.id,
          createdAt: this.defaltValues?.createdAt?.Time,
        });
      } else {
        this.$emit('onAdd', values);
      }
    },
  },
};
</script>
