import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('auth', {
  state: () => ({
    user: {},
  }),
  getters: {
    isAuthorized() {
      return !!localStorage.getItem('tracker-auth-token');
    },
  },
  actions: {
    async login(payload) {
      const { data } = await HttpClient.post('auth/sign-in', {
        email: payload.email,
        password: payload.password,
      });
      localStorage.setItem('tracker-auth-token', data?.token);
    },
    async register(payload) {
      await HttpClient.post('auth/sign-up', {
        fullName: payload.fullName,
        email: payload.email,
        password: payload.password,
      });
    },

    async whoAmI() {
      const { data } = await HttpClient.get('api/who-am-i');
      this.user = data;
    },
  },
});
