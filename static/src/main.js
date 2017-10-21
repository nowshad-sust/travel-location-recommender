import Vue from 'vue'
import App from './App.vue'


import BootstrapVue from 'bootstrap-vue'
import VueLazyload from 'vue-lazyload'
import VueResource from 'vue-resource'
import VueYouTubeEmbed from 'vue-youtube-embed'


/* ...there may be other imports here */
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'bootstrap/dist/css/bootstrap.css'

// or with options
Vue.use(VueLazyload, {
  preLoad: 1.3,
  error: 'dist/error.png',
  loading: 'dist/loading.gif',
  attempt: 1
})

Vue.use(BootstrapVue)
Vue.use(VueResource)
Vue.use(VueYouTubeEmbed)


var vm = new Vue({
  el: '#app',
  data: {
    choices: null
  },
  render: h => h(App)
})
