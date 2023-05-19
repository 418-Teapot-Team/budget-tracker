<template>
  <div class="bg-black bg-opacity-70 fixed top-0 bottom-0 left-0 right-0 flex">
    <div class="bg-white shadow-sm rounded-2xl m-auto w-80 h-fit p-8 relative">
      <div class="absolute right-3 top-3 w-7 h-7 cursor-pointer" @click="closePopup">
        <close-icon />
      </div>
      <div class="flex flex-row justify-center items-center py-4">
        <span class="text-lg">Confirm deletion?</span>
      </div>
      <div class="h-10 flex flex-row gap-3">
        <app-button text="Cancel" :isOutline="true" @click="closePopup" />
        <app-button text="Delete" :isOutline="false" @click="confirm" />
      </div>
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
  methods: {
    closePopup() {
      this.$emit('onClose');
    },
    confirm() {
      this.$emit('onConfirm');
    },
  },
};
</script>
