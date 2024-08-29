<script setup>
import axios from '../../../axiosConfig'
import { ref } from 'vue'
import {useToast} from 'vue-toastification';

const toast = useToast();
const form = ref(null)
const initialFormData = ref({
    Manufacturer: '',
    LeadModel: '',
    Name: '',
    Polarity: '',
    HasHazard: false,
    IsMri: true,
    comment: '',
})

const formData = ref({ ...initialFormData.value })

async function handleSubmit() {
    try {
        const response = await axios.post('api/leads', formData.value, {
        headers: {
            'Content-Type': 'application/json'
        }
        });
        console.log('Lead updated successfully', response.data);
        toast.success('Lead updated successfully',{
            timeout: 2000,
        });
        formData.value = { ...initialFormData };
    } catch (error) {
        console.error('Failed to update lead', error);
        toast.error('Failed to update lead',{
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
        <label for="Manufacturer" class="form-label">Manufacturer</label>
        <select class="form-select" id="Manufacturer" name="Manufacturer" v-model="formData.Manufacturer" @input="handleInput">
            <option disabled value="" >Please select a manufacturer</option>
            <option value="Abbott">Abbott</option>
            <option value="Biotronik">Biotronik</option>
            <option value="Boston Scientific">Boston Scientific</option>
            <option value="Medtronic">Medtronic</option>
            <option value="Sorin">Sorin</option>
        </select>
        <label for="LeadModel" class="form-label">Lead Model</label>
        <input type="text" class="form-control" id="LeadModel" name="LeadModel" v-model="formData.LeadModel" @input="handleInput">
        <label for="Name" class="form-label">Name</label>
        <input type="text" class="form-control" id="Name" name="Name" v-model="formData.Name" @input="handleInput">
        <label for="Polarity" class="form-label">Polarity</label>
        <select class="form-select" id="Polarity" name="Polarity" v-model="formData.Polarity" @input="handleInput">
            <option disabled value="">Please select a lead type</option>
            <option value="Unipolar">Unipolar</option>
            <option value="Bipolar">Bipolar</option>
            <option value="Multipolar">Multipolar</option>
        </select>
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
        <button type="submit" class="btn btn-primary mt-2">Submit</button>
    </form>
  </div>
</template>