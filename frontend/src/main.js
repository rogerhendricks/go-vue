import { createApp } from 'vue'
import router from './router'
import './style.css'
import App from './App.vue'
import "bootstrap/dist/css/bootstrap.min.css"
// import "bootstrap"
import * as bootstrap from "bootstrap"
// // Import Bootstrap CSS
// import 'bootstrap/dist/css/bootstrap.min.css';
import './style.css'
// // Import Bootstrap JS
// import 'bootstrap/dist/js/bootstrap.bundle.min.js'; 
// Make Bootstrap available globally
window.bootstrap = bootstrap;

import Toast from 'vue-toastification';
import "vue-toastification/dist/index.css";
import "vue-multiselect/dist/vue-multiselect.css"
import store from './store'


router.beforeEach((to, from, next) => {
    console.log(`Navigating to ${to.path} from ${from.path}`);
    next();
  });

  
createApp(App)
.use(router)
.use(store)
.use(Toast)
.mount('#app')
