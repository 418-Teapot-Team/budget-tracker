/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js}'],
  theme: {
    extend: {
      colors: {
        black: '#1c1c21',
        yellow: '#ffefe2',
        'blue-light': '#F0F5FF',
        pink: '#FEF8FC',
        grey: '#e8e8e8',
        'grey-light': '#f8f9fe',
      },
    },
  },
  plugins: [],
};
