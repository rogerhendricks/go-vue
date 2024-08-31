<script setup>
import axios from '../../../axiosConfig'
import { ref, reactive } from 'vue'
import {useToast} from 'vue-toastification';

const toast = useToast();

const initialDoctorState = reactive({
  name: '',
  email: '',
  phone: '',
  addresses: [
    {
      street: '',
      city: '',
      state: '',
      zip: '',
      country: ''
    }
  ]
})

const doctor = reactive({ ...initialDoctorState })

const addAddress = () => {
  doctor.addresses.push({
    street: '',
    city: '',
    state: '',
    zip: '',
    country: ''
  })
}

const removeAddress = (index) => {
  doctor.addresses.splice(index, 1)
}

const handleSubmit = async () => {
  try {
    const response = await axios.post('api/doctors', doctor)
    console.log('Doctor updated successfully', response.data);
    toast.success('Doctor updated successfully',{
            timeout: 2000,
        });
    Object.assign(doctor, initialDoctorState);
  } catch (error) {
    console.error('Failed to update doctor', error);
    toast.error('Failed to update doctor',{
            timeout: 2000,
        });
  }
}
</script>
<template>
    <div>
      <h1>Create Doctor</h1>
      <form @submit.prevent="handleSubmit">
        <div class="mb-3">
          <label for="name" class="form-label">Name</label>
          <input type="text" class="form-control" id="name" v-model="doctor.name" />
        </div>
        <div class="mb-3">
          <label for="email" class="form-label">Email</label>
          <input type="email" class="form-control" id="email" v-model="doctor.email" />
        </div>
        <div class="mb-3">
          <label for="phone" class="form-label">Phone</label>
          <input type="text" class="form-control" id="phone" v-model="doctor.phone" />
        </div>
  
        <div v-for="(address, index) in doctor.addresses" :key="index">
          <h3>Address {{ index + 1 }}</h3>
          <div class="mb-3">
            <label for="street" class="form-label">Street</label>
            <input type="text" class="form-control" id="street" v-model="address.street" />
          </div>
          <div class="mb-3">
            <label for="city" class="form-label">City</label>
            <input type="text" class="form-control" id="city" v-model="address.city" />
          </div>
          <div class="mb-3">
            <label for="state" class="form-label">State</label>
            <input type="text" class="form-control" id="state" v-model="address.state" />
          </div>
          <div class="mb-3">
            <label for="zip" class="form-label">Zip</label>
            <input type="text" class="form-control" id="zip" v-model="address.zip" />
          </div>
          <div class="mb-3">
            <label for="country" class="form-label">Country</label>
            <input type="text" class="form-control" id="country" v-model="address.country" />
          </div>
          <button type="button" class="btn btn-danger" @click="removeAddress(index)">Remove This Address</button>
        </div>
  
        <div class="btn-group mt-4">
          <button type="button" class="btn btn-secondary" @click="addAddress">Add Another Address</button>
          <button type="submit" class="btn btn-primary">Submit</button>
        </div>

      </form>
    </div>
  </template>