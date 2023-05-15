<template>
  <svg
    width="100%"
    height="100%"
    viewBox="0 0 24 24"
    xmlns="http://www.w3.org/2000/svg"
    id="price-tag"
    class="icon glyph"
  >
    <path
      d="M20,2H14a1,1,0,0,0-.7.29L3,12.55a2,2,0,0,0,0,2.83L8.62,21a2,2,0,0,0,1.41.59h0A2,2,0,0,0,11.45,21L21.71,10.7A1,1,0,0,0,22,10V4A2,2,0,0,0,20,2ZM16.5,9A1.5,1.5,0,1,1,18,7.5,1.5,1.5,0,0,1,16.5,9Z"
      :style="'fill:' + (isActive ? '#1c1c21' : '#fff')"
    ></path>
  </svg>
</template>

<script>
export default {
  name: 'ExpensesIcon',
  props: {
    isActive: Boolean,
  },
};
</script>
