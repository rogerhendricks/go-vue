<script setup>
import axios from './axiosConfig'
import {ref} from 'vue'
// import {useToast} from 'vue-toastification';
import {useRouter} from 'vue-router';

const router = useRouter();

const form =ref(null)
const formData = ref({
  Username: '',
  Fullname: '',
  Password: ''
});

async function handleSubmit() {
    try {
        await axios.put('/api/user', formData.value, {
            headers: {
            'Content-Type': 'application/json'
            }
    });
    console.log('User updated successfully', formData.value);
    localStorage.removeItem('user');
    router.push('/login')
    // console.log('User updated successfully:', response.data);
    } catch (error) {
    console.error('Error updating user:', error);
    }
}

</script>


<template>
    <div>
      <h1>User Settings</h1>
      <form @submit.prevent="handleSubmit" ref="form">
        <div>
          <label for="name" class="form-label">Name:</label>
          <input id="name" v-model="formData.Username" type="text"  class="form-control"/>
        </div>
        <div>
          <label for="fullname" class="form-label">Full Name:</label>
          <input id="fullname" v-model="formData.Fullname" type="text"  class="form-control"/>
        </div>
        <div>
          <label for="password" class="form-label">New Password:</label>
          <input id="password" v-model="formData.Password" type="password"  class="form-control"/>
        </div>
        <button type="submit" class="btn btn-primary mt-2">Update</button>
      </form>
    </div>
  </template>
  
