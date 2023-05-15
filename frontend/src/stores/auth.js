import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('auth', {
  state: () => ({
    user: {},
    error: {},
    isError: false,
    isLoading: false,
    message: '',
    isSuccess: false,
  }),
  getters: {
    isAuthorized() {
      const token = localStorage.getItem('tracker-auth-token');
      if (token && token.trim() !== '') {
        return true;
      } else {
        return false;
      }
    },
  },
  actions: {
    async login(payload) {
      try {
        this.isSuccess = false;
        this.isLoading = true;
        this.message = '';
        const { token } = await HttpClient.post('auth/sign-in', {
          email: payload.email,
          password: payload.password,
        });
        this.error = {};
        this.isError = false;
        this.message = 'Вітаємо у системі';
        localStorage.setItem('tracker-auth-token', token);
        this.isSuccess = true;
      } catch (e) {
        console.log(e);
        this.error = e?.request?.data || { message: e.message };
        this.isError = true;
      } finally {
        this.isLoading = false;
      }
    },
    async register(payload) {
      try {
        this.isSuccess = false;
        this.isLoading = true;
        this.message = '';
        await HttpClient.post('auth/sign-up', {
          fullName: payload.fullName,
          email: payload.email,
          password: payload.email,
        });
        this.error = {};
        this.isError = false;
        this.message = 'Акаунт успішно створено. Увійдіть в акаунт, щоб користуватись системою';
        this.isSuccess = true;
      } catch (e) {
        console.log(e);
        this.error = e?.request?.data || { message: e.message };
        this.isError = true;
      } finally {
        this.isLoading = false;
      }
    },
    // #TODO: get user info
    async whoAmI() {},
  },
});
