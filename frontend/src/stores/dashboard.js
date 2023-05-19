import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('dashboard', {
  state: () => ({
    courses: [],
    categories: [],
    currentMonthSaving: 0,
  }),

  actions: {
    async getCourses() {
      const { data } = await HttpClient.get('api/get-courses');
      this.courses = data;
    },
    async getTopCaregories() {
      const { data } = await HttpClient.get(`api/lists/get-top-categories?type=expenses`);
      console.log(data);
    },
    async getCurrentMonthSavings() {
      const { data } = await HttpClient.get('api/lists/current-mon-saved');
      this.currentMonthSaving = data.result;
    },
  },
});
