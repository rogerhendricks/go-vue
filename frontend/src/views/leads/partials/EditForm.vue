<script setup>
import axios from '../../../axiosConfig'
import { ref} from 'vue'
import {useToast} from 'vue-toastification';

const toast = useToast();
const props = defineProps(['lead']);
const emit = defineEmits('cancel');
console.log('These are the lead props',props.lead);

const form = ref(null)
const formData = ref({
    Manufacturer: props.lead.manufacturer,
    LeadModel: props.lead.leadModel,
    Name: props.lead.name,
    Polarity: props.lead.polarity,
    HasHazard: props.lead.hasHazard,
    IsMri: props.lead.isMri,
    Comment: props.lead.comment
})

console.log('This is the form data',formData.value);

async function handleSubmit() {
    try {
        const response = await axios.put(`api/leads/${props.lead.ID}`, formData.value, {
        headers: {
            'Content-Type': 'application/json'
        }
        });
        console.log('Lead updated successfully', response.data);
        toast.success('Lead updated successfully',{
            timeout: 2000,
        });
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
function cancel() {
      emit('cancel');
    }
</script>

<template>
  <div class="border border-success p-2 mb-2 border-opacity-50 rounded shadow-sm bg-light text-dark">
    <form @submit.prevent="handleSubmit" ref="form">
        <div class="row">
            <div class="col">
                <label for="Manufacturer" class="form-label">Manufacturer</label>
                <select class="form-select" id="Manufacturer" name="Manufacturer" v-model="formData.Manufacturer" @input="handleInput">
                    <option disabled value="" >Please select a manufacturer</option>
                    <option value="Abbott">Abbott</option>
                    <option value="Biotronik">Biotronik</option>
                    <option value="Boston Scientific">Boston Scientific</option>
                    <option value="Medtronic">Medtronic</option>
                    <option value="Sorin">Sorin</option>
                </select>
            </div>
            <div class="col">
                <label for="LeadModel" class="form-label">Lead Model</label>
                <input type="text" class="form-control" id="LeadModel" name="LeadModel" v-model="formData.LeadModel" @input="handleInput">
            </div>
            <div class="col">
                <label for="Name" class="form-label">Name</label>
                <input type="text" class="form-control" id="Name" name="Name" v-model="formData.Name" @input="handleInput">
            </div>
        </div>
        <div class="row mt-2">
            <div class="col">
                <label for="Polarity" class="form-label">Polarity</label>
                <select class="form-select" id="Type" name="Polarity" v-model="formData.Polarity" @input="handleInput">
                    <option disabled value="">Please select a lead type</option>
                    <option value="Unipolar">Unipolar</option>
                    <option value="Bipolar">Bipolar</option>
                    <option value="Multipolar">Multipolar</option>
                </select>
            </div>
            <div class="col">
                <label for="HasHazard" class="form-label">Hazard</label>
                <div class="form-check">
                    <input type="checkbox" class="form-check-input" id="HasHazard" name="HasHazard" v-model="formData.HasHazard" @input="handleInput">
                    <label for="HasHazard" class="form-check-label">Has Hazard</label>
                </div>
            </div>
            <div class="col">
                <label for="IsMri" class="form-label">MRI</label>
                <div class="form-check">
                    <input type="checkbox" class="form-check-input" id="IsMri" name="IsMri" v-model="formData.IsMri" @input="handleInput">
                    <label for="IsMri" class="form-check-label">Is Mri</label>
                </div>
            </div>
        </div>
        <div class="row mt-2">
            <div class="col">
                <label for="comment" class="form-label">Comment</label>
                <input type="text" class="form-control" id="comment" name="comment" v-model="formData.Comment" @input="handleInput"> 
            </div>
        </div>
        <div class="row mt-2">
            <div class="col">
                <button type="submit" class="btn btn-primary ">Update</button>
                    <div class="btn-group" id="actions" role="group" >
                        <button type="button" class="btn btn-secondary" @click="cancel">Cancel</button>
                    </div>
            </div>
        </div>
    </form>
  </div>
</template>