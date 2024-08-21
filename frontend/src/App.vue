<script setup>
  import {computed, provide, ref} from 'vue'
  import { useRouter } from 'vue-router'
  // import axios from './axiosConfig'
  import Default from './views/layouts/Default.vue';
  import Navbar from './components/Navbar.vue';
  // const router = useRouter()
  // const user = computed(() => {
  //     return JSON.parse(localStorage.getItem('user') || 'null')
  //   })
  
    const user = ref(JSON.parse(localStorage.getItem('user') || 'null'));
    provide('user', user);
  // watchEffect(() => {
  //   // This will run whenever `user` changes
  //   Navbar.props = { user: user.value }
  // })
  // const logout = async () => {
  //   try {
  //       await axios.post('/api/logout');
  //       localStorage.removeItem('user');
  //       router.push('/login');
  //   } catch (error) {
  //       console.error('Logout failed', error);
  //   }
  // }
</script>

<template>
    <div id="app">
    <Navbar />
      <Default>
        <router-view/>
      </Default>

  </div>
</template>