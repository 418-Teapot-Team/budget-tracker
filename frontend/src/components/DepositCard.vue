<template>
  <div
    class="w-full h-44 shadow-sm bg-white rounded-2xl flex flex-row justify-between items-center p-6 hover:bg-gray-50"
  >
    <div class="h-full w-28 mr-10">
      <deposit-card-icon />
    </div>
    <div class="w-full h-full flex flex-col justify-between items-center mr-5">
      <div class="w-full h-1/2 flex flex-row justify-between items-start">
        <span class="text-3xl"> {{ data?.name }} </span>
        <div class="flex flex-col justify-start items-end">
          <span class="text-black"
            >{{ data?.payed?.toFixed(2) }} / {{ data?.goalSum?.toFixed(2) }} UAH</span
          >
          <span class="text-black">Ставка {{ data?.percent }}% річних</span>
        </div>
      </div>
      <div class="w-full">
        <div class="w-full h-3 bg-black rounded-full bg-opacity-10 mb-3 overflow-hidden">
          <div
            class="h-3 bg-black rounded-full"
            :style="`width: ${(data?.payed / data?.goalSum) * 100}%;`"
          ></div>
        </div>
        <div class="w-full flex flex-row justify-between items-start">
          <span class="text-black">
            {{ data?.startAccountDate?.split('-').reverse().join('.') }}
          </span>
          <span class="text-black">
            {{ data?.finishAccountDate?.split('-').reverse().join('.') }}
          </span>
        </div>
      </div>
    </div>
    <div class="w-7 h-full flex flex-col justify-start items-end gap-3">
      <div class="w-7 h-7 cursor-pointer" @click="showConfirmPopup = true">
        <delete-icon />
      </div>
      <div class="w-7 h-7 cursor-pointer" @click="editDeposit">
        <edit-icon />
      </div>
    </div>
    <confirm-popup
      v-if="showConfirmPopup"
      @onClose="showConfirmPopup = false"
      @onConfirm="deleteDeposit"
    />
  </div>
</template>

<script>
import DepositCardIcon from '@/components/Icons/DepositCardIcon.vue';
import DeleteIcon from '@/components/Icons/DeleteIcon.vue';
import EditIcon from '@/components/Icons/EditIcon.vue';
import ConfirmPopup from '@/components/ConfirmPopup.vue';
export default {
  name: 'DepositCard',
  components: {
    DepositCardIcon,
    DeleteIcon,
    EditIcon,
    ConfirmPopup,
  },
  props: {
    data: Object,
  },
  data() {
    return {
      showConfirmPopup: false,
    };
  },
  methods: {
    editDeposit() {
      this.$emit('onEditDeposit', this.data.id);
    },
    deleteDeposit() {
      this.$emit('onDeleteDeposit', this.data.id);
    },
  },
};
</script>
