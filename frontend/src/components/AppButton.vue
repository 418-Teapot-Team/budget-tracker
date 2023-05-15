<template>
  <button
    class="border-2 border-black h-full text-center w-full rounded-xl flex items-center justify-center"
    :class="
      isOutline
        ? 'transition duration-300 hover:bg-black hover:text-white text-black'
        : 'bg-black text-white transition duration-300 hover:bg-slate-950 hover:border-slate-950'
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
