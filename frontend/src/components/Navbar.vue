<script setup>
  import {computed,ref,inject, watchEffect} from 'vue'
  import { useRouter } from 'vue-router'
  import axios from '../axiosConfig'
  const router = useRouter()
//   const user = computed(() => {
//       return JSON.parse(localStorage.getItem('user') || 'null')
//     })
    // const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))
    const user = inject('user');
    watchEffect(() => {
    user.value = JSON.parse(localStorage.getItem('user') || 'null')
  })
    const logout = async () => {
        try {
            await axios.post('/api/logout');
            localStorage.removeItem('user');
            user.value = null;
            router.push('/login');
        } catch (error) {
            console.error('Logout failed', error);
        }
        // await console.log('logout')
    }
</script>
<template>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link class="nav-link" to="/dashboard" v-if="user">Dashboard</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/doctors" v-if="user">Doctors</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/leads" v-if="user">Leads</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/devices" v-if="user">Devices</router-link>
          </li>
          <li class="nav-item">
            <button type="button" class="btn nav-link" @click="logout" v-if="user">Logout</button>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/login" v-if="!user">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/register" v-if="!user">Register</router-link>
          </li>
        </ul>
        <div class="d-flex">
         <span v-if="user">
          <router-link class="nav-link" to="/settings">
            {{ user.fullname }}
          </router-link>
        </span>
         
         <span v-else>
            <router-link class="nav-link" to="/login" v-if="!user">Login</router-link>
         </span>
        </div>
      </div>
    </div>
  </nav>
</template>


