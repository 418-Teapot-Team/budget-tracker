import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('dashboard', {
  state: () => ({
    courses: [],
    categories: [],
  }),

  actions: {
    async getCourses() {
      const { data } = await HttpClient.get('api/get-courses');
      this.courses = data;
    },
    async getTopCaregories() {
      const { data } = await HttpClient.get('api/lists/get-top-expenses');
      console.log(data);
    },
    async getCurrentMonthSavings() {
      const { data } = await HttpClient.get('api/lists/current-mon-saved');
      console.log(data);
    },
  },
});
