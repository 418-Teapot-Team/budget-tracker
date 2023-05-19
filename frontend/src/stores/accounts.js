import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('accounts', {
  state: () => ({
    deposits: [],
    loans: [],
    depositTotal: {},
  }),
  actions: {
    async getAccounts({ type, orderBy = '', reverse = false }) {
      const { data } = await HttpClient.get(
        `api/accounts/${type}?orderBy=${orderBy}&reverse=${reverse}`
      );
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
    async getDepositTotal() {
      const { data } = await HttpClient.get('api/account-total');
      this.depositTotal = data;
    },
  },
});
