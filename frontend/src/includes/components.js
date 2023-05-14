import AppButton from '@/components/AppButton.vue';
import Prealoder from '@/components/Preloader.vue';
export default {
  install(app) {
    app.component('AppButton', AppButton);
    app.component('Prealoader', Prealoder);
  },
};
