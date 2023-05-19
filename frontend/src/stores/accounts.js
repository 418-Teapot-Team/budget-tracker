import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('accounts', {
  state: () => ({
    deposits: [],
    loans: [],
  }),
  actions: {
    async getAccounts({ type }) {
      const { data } = await HttpClient.get(`api/accounts/${type}`);
      if (type === 'deposit') {
        this.deposits = data.result;
      } else {
        this.loans = data.result;
      }
    },
    async createAccount(payload) {
      await HttpClient.post('api/accounts/create', {
        type: payload.type,
        name: payload.name,
        monthAmount: Number(payload.monthAmount),
        percent: Number(payload.percent),
        date: payload.date,
        sum: Number(payload.sum),
      });
    },
    async editAccount(payload) {
      await HttpClient.put('api/accounts/update', {
        id: payload.id,
        type: payload.type,
        name: payload.name,
        monthAmount: Number(payload.monthAmount),
        percent: Number(payload.percent),
        date: payload.date,
        sum: Number(payload.sum),
      });
    },
    async deleteAccount(payload) {
      await HttpClient.delete('api/accounts/delete', {
        data: { id: payload.id },
      });
    },
  },
});