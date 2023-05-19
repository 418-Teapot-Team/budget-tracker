import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('dashboard', {
  state: () => ({
    expensesStats: [],
    expensesCategories: [],
    total: 0,
    incomeStats: [],
    incomeCategories: [],
  }),

  actions: {
    async getCourses() {
      const { data } = await HttpClient.get('api/get-courses');
      this.courses = data;
    },
    async getTopCaregories() {
      const { data } = await HttpClient.get(`api/lists/get-top-categories?type=expenses`);
      console.log(data);
      this.topCategories = data.result;
    },
    async getCurrentMonthSavings() {
      const { data } = await HttpClient.get('api/lists/current-mon-saved');
      this.currentMonthSaving = data.result;
    },
    async getSavingsStats() {
      const { data } = await HttpClient.get('api/lists/get-saving-stats');
      this.savingStats = data.result;
    },
    async getTransactions({ take, skip }) {
      const { data } = await HttpClient.get(`api/lists/all?take=${take}&skip=${skip}`);
      this.transactions = data.result;
    },
  },
});