import { createApp } from 'vue'
import router from './router'
import './style.css'
import App from './App.vue'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import Toast from 'vue-toastification';
import "vue-toastification/dist/index.css";
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
