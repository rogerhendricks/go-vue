<template>
    <div>
      <h2>Login</h2>
      <form @submit.prevent="login">
        <input class="form-control mb-2" v-model="username" type="text" placeholder="Username" required>
        <input type="password" class="form-control mb-2" v-model="password" placeholder="Password" required>
        <button type="submit" class="btn btn-primary">Login</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import { inject } from 'vue';
  
  export default {
    data() {
      return {
        username: '',
        password: '',
        user: null
      }
    },
    created() {
    this.user = inject('user');
    },
    methods: {
      async login() {
        try {
          const response = await axios.post('/api/login', {
            username: this.username,
            password: this.password
          });
          localStorage.setItem('user', JSON.stringify(response.data.user));
          this.user = response.data.user;
          this.$router.push('/dashboard');
        } catch (error) {
          alert(error.response.data.error);
        }
      }
    }
  }
  </script>