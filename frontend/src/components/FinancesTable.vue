<template>
  <div class="w-full overflow-y-auto table-wrapper">
    <table class="w-full border-separate border-spacing-y-2 mb-4 -mt-2">
      <thead class="h-14 sticky top-0 bg-grey-light">
        <th class="text-start pl-2 text-xl font-normal text-black w-1/12">Date</th>
        <th class="text-center text-xl font-normal text-black w-3/12">Category</th>
        <th class="text-start pl-2 text-xl font-normal text-black w-3/12">Amount</th>
        <th class="text-start pl-2 text-xl font-normal text-black w-3/12">Comment</th>
        <th class="text-end pr-4 text-xl font-normal text-black w-1/12">Actions</th>
      </thead>
      <tbody class="text-start mt-2">
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
        <tr class="h-16 bg-white shadow-sm rounded-lg p-2 hover:bg-gray-50">
          <td class="pl-2 rounded-l-lg text-black">16.05.23</td>
          <td class="text-center text-black">Bussiness</td>
          <td class="pl-2" :class="isIncome ? 'text-green-500' : 'text-red-500'">
            {{ isIncome ? '+' : '-' }}
            340 UAH
          </td>
          <td class="pl-2 text-black">Work income</td>
          <td class="pr-6 text-end rounded-r-lg text-black">
            <div class="w-full h-full flex justify-end">
              <div class="h-8 w-8 cursor-pointer mr-2" @click="editItem('id')">
                <edit-icon />
              </div>
              <div class="h-8 w-8 cursor-pointer" @click="deleteItem('id')">
                <delete-icon />
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="w-full flex h-10 flex-row justify-center items-center">
      <div class="w-36 h-full">
        <app-button text="Load more" @click="loadMore" />
      </div>
    </div>
  </div>
</template>

<script>
import DeleteIcon from '@/components/Icons/DeleteIcon.vue';
import EditIcon from '@/components/Icons/EditIcon.vue';
export default {
  name: 'FinancesTable',
  props: {
    data: Array,
    isIncome: Boolean,
  },
  components: {
    DeleteIcon,
    EditIcon,
  },
  methods: {
    deleteItem(id) {
      console.log(id);
      this.$emit('onDeleteItem', id);
    },
    loadMore() {
      this.$emit('onLoadMore');
    },
    editItem(id) {
      this.$emit('onEditItem', id);
    },
  },
};
</script>

<style scoped>
.table-wrapper {
  max-height: calc(100vh - 144px);
  padding-bottom: 20px;
}
.table-wrapper::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}

.table-wrapper::-webkit-scrollbar-track {
  border-radius: 100vh;
  background: #f8f9fe;
}

.table-wrapper::-webkit-scrollbar-thumb {
  background: #1c1c21;
  border-radius: 100vh;
  border: 3px solid #f8f9fe;
}

.table-wrapper::-webkit-scrollbar-thumb:hover {
  background: #1c1c21;
}
</style>
