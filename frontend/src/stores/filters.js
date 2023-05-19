import { defineStore } from 'pinia';

export default defineStore('filters', {
  state: () => ({
    filters: [
      {
        title: 'Date DESC',
        value: {
          id: 1,
          orderBy: 'created_at',
          reverse: false,
        },
      },
      {
        title: 'Date ASC',
        value: {
          id: 2,
          orderBy: 'created_at',
          reverse: true,
        },
      },
      {
        title: 'Amount DESC',
        value: {
          id: 3,
          orderBy: 'amount',
          reverse: false,
        },
      },
      {
        title: 'Amount ASC',
        value: {
          id: 4,
          orderBy: 'amount',
          reverse: true,
        },
      },
    ],
    filtersForStats: [
      {
        title: '1 Month',
        value: {
          id: 5,
          months: 1,
        },
      },
      {
        title: '3 Month',
        value: {
          id: 6,
          months: 3,
        },
      },
      {
        title: '6 Month',
        value: {
          id: 7,
          months: 6,
        },
      },
      {
        title: '12 Month',
        value: {
          id: 8,
          months: 12,
        },
      },
    ],
  }),
});
