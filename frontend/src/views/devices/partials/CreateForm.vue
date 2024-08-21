<script setup>
import axios from '../../../axiosConfig'
import { ref } from 'vue'


const form = ref(null)
const formData = ref({
    Manufacturer: '',
    DeviceModel: '',
    Name: '',
    Type: '',
    Header: '',
    HasHazard: false,
    IsMri: true,
    comment: '',

})

async function handleSubmit() {
    try {
        const response = await axios.post('api/devices', formData.value, {
        headers: {
            'Content-Type': 'application/json'
        }
        });
        console.log('Device updated successfully', response.data);
    } catch (error) {
        console.error('Failed to update device', error);
    }
}

function handleInput(event) {
  formData.value[event.target.name] = event.target.value
}
</script>

<template>
  <div>
    <form @submit.prevent="handleSubmit" ref="form">
        <label for="Manufacturer" class="form-label">Manufacturer</label>
        <select class="form-select" id="Manufacturer" name="Manufacturer" v-model="formData.Manufacturer" @input="handleInput">
            <option disabled value="" >Please select a manufacturer</option>
            <option value="Abbott">Abbott</option>
            <option value="Biotronik">Biotronik</option>
            <option value="Boston Scientific">Boston Scientific</option>
            <option value="Medtronic">Medtronic</option>
            <option value="Sorin">Sorin</option>
        </select>
        <label for="DeviceModel" class="form-label">DeviceModel</label>
        <input type="text" class="form-control" id="DeviceModel" name="DeviceModel" v-model="formData.DeviceModel" @input="handleInput">
        <label for="Name" class="form-label">Name</label>
        <input type="text" class="form-control" id="Name" name="Name" v-model="formData.Name" @input="handleInput">
        <label for="Type" class="form-label">Type</label>
        <select class="form-select" id="Type" name="Type" v-model="formData.Type" @input="handleInput">
            <option disabled value="">Please select a device type</option>
            <option value="Defibrillator">Defibrillator</option>
            <option value="Pacemaker">Pacemaker</option>
            <option value="Loop Recorder">Loop Recorder</option>
        </select>
        <label for="Header" class="form-label">Header</label>
        <input type="text" class="form-control" id="Header" name="Header" v-model="formData.Header" @input="handleInput">
        <div class="form-check">
            <input type="checkbox" class="form-check-input" id="HasHazard" name="HasHazard" v-model="formData.HasHazard" @input="handleInput">
            <label for="HasHazard" class="form-check-label">HasHazard</label>
        </div>
        <div class="form-check">
            <input type="checkbox" class="form-check-input" id="IsMri" name="IsMri" v-model="formData.IsMri" @input="handleInput">
            <label for="IsMri" class="form-check-label">IsMri</label>
        </div>
        <label for="comment" class="form-label">comment</label>
        <input type="text" class="form-control" id="comment" name="comment" v-model="formData.comment" @input="handleInput"> 
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>