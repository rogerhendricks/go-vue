<script setup>
import axios from '../../../axiosConfig'
import { ref, defineProps } from 'vue'
import {useToast} from 'vue-toastification';

const toast = useToast();
const props = defineProps(['doctor'])

const form = ref(null)
const formData = ref({
  Name: props.doctor.name,
  Email: props.doctor.email,
  Phone: props.doctor.phone,
  Addresses: props.doctor.Addresses ? [...props.doctor.Addresses] : [],
//   addresses: props.doctor.addresses
//   addresses: [...props.doctor.addresses]
// addresses: JSON.parse(JSON.stringify(props.doctor.addresses))
})

// console.log('formData', formData.value)
const addAddress = () => {
  formData.value.Addresses.push({
    street: '',
    city: '',
    state: '',
    zip: '',
    country: ''
  })
}

const removeAddress = (index) => {
  formData.value.Addresses.splice(index, 1)
}

const handleSubmit = async () => {
    try {
        const url = `/api/doctors/${props.doctor.ID}`;
        const response = await axios.put(url, formData.value,{
            headers: {
                'Content-Type': 'application/json'
            }
        });
        console.log('Doctor updated successfully', response.data);
        toast.success('Doctor updated successfully',{
        timeout: 2000,
        });
    } catch (error) {
        console.error('Failed to update doctor', error);
        toast.error('Failed to update doctor',{
        timeout: 2000,
        });
    }
}
function handleInput(event) {
  formData.value[event.target.name] = event.target.value
}
</script>
<template>
    <div>
        <form @submit.prevent="handleSubmit" ref="form">

            <div class="row">
                <div class="col">
                    <label for="name" class="form-label">Name</label>
                    <input type="text" class="form-control" id="name" v-model="formData.Name" @input="handleInput"/>
                </div>
                <div class="col">
                    <label for="email" class="form-label">Email</label>
                    <input type="email" class="form-control" id="email" v-model="formData.Email" @input="handleInput"/>
                </div>
                <div class="col">
                    <label for="phone" class="form-label">Phone</label>
                    <input type="text" class="form-control" id="phone" v-model="formData.Phone" @input="handleInput"/>
                </div>  
            </div>
  

        <div v-for="(address, index) in formData.Addresses" :key="index" class="mt-4">
            <h4>Address {{ index + 1 }}</h4>
            <div class="row">
                <div class="col">
                    <label for="street" class="form-label">Street</label>
                    <input type="text" class="form-control" id="street" v-model="address.street" />
                </div>
                <div class="col">
                    <label for="city" class="form-label">City</label>
                    <input type="text" class="form-control" id="city" v-model="address.city" />
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <label for="state" class="form-label">State</label>
                    <input type="text" class="form-control" id="state" v-model="address.state" />
                </div>
                <div class="col">
                    <label for="country" class="form-label">Country</label>
                    <input type="text" class="form-control" id="country" v-model="address.country" />
                </div>
                <div class="col">
                    <label for="zip" class="form-label">Zip</label>
                    <input type="text" class="form-control" id="zip" v-model="address.zip" />
                </div>
            </div>
          <button type="button" class="btn btn-danger mt-2" @click="removeAddress(index)">Remove This Address</button>
        </div>
  
        <div class="btn-group mt-4">
          <button type="button" class="btn btn-secondary" @click="addAddress">Add Another Address</button>
          <button type="submit" class="btn btn-primary">Submit</button>
        </div>

      </form>
    </div>
    
</template>