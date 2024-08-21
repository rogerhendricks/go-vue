<template>
    <div>
      <h2>Register</h2>
      <form @submit.prevent="register">
        <input class="form-control mb-2" v-model="username" type="text" placeholder="Username" required>
        <input class="form-control mb-2" v-model="fullname" type="text" placeholder="Full Name" required>
        <input class="form-control mb-2" v-model="password" type="password" placeholder="Password" required>
        <button type="submit" class="btn btn-primary">Register</button>
      </form>
      </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',
        fullname: '',
        password: '',
      }
    },
    methods: {
      async register() {
        try {
          await axios.post('/api/register', {
            username: this.username,
            fullname: this.fullname,
            password: this.password
          });
          this.$router.push('/login');
        } catch (error) {
          console.error('Registration error:', error);
        }
      }
    }


  }
  </script>

<style scoped>
.error {
  color: red;
}
</style>