<template>
  <button
    class="border-2 border-black p-2 text-center w-full rounded-xl"
    :class="
      isOutline
        ? 'transition duration-300 hover:bg-black hover:text-white text-black'
        : 'bg-black text-white'
    "
    :type="type ?? type"
    @click="onClick"
  >
    {{ text }}
  </button>
</template>

<script>
export default {
  name: 'AppButton',
  props: {
    text: String,
    isOutline: Boolean,
    type: String,
  },
  methods: {
    onClick() {
      this.$emit('click');
    },
  },
};
</script>
