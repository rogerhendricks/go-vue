<script setup>
import axios from '../../../axiosConfig'
import { ref} from 'vue'
import {useToast} from 'vue-toastification';

const toast = useToast();
const props = defineProps(['device']);
const emit = defineEmits('cancel');

const form = ref(null)
const formData = ref({
    Manufacturer: props.device.manufacturer,
    DeviceModel: props.device.deviceModel,
    Name: props.device.name,
    Type: props.device.type,
    Header: props.device.header,
    HasHazard: props.device.hasHazard,
    IsMri: props.device.isMri,
    Comment: props.device.comment
})



async function handleSubmit() {
    try {
        const response = await axios.put(`api/devices/${props.device.ID}`, formData.value, {
            headers: {
                'Content-Type': 'application/json'
            }
        });
        console.log('Device updated successfully', response.data);
        toast.success('Device updated successfully',{
            timeout: 2000,
        });
    } catch (error) {
        console.error('Failed to update device', error);
        toast.error('Failed to update device',{
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
                <label for="DeviceModel" class="form-label">DeviceModel</label>
                <input type="text" class="form-control" id="DeviceModel" name="DeviceModel" v-model="formData.DeviceModel" @input="handleInput">
            </div>
            <div class="col">
                <label for="Name" class="form-label">Name</label>
                <input type="text" class="form-control" id="Name" name="Name" v-model="formData.Name" @input="handleInput">
            </div>
        </div>
        <div class="row mt-2">
            <div class="col">
                <label for="Type" class="form-label">Type</label>
                <select class="form-select" id="Type" name="Type" v-model="formData.Type" @input="handleInput">
                    <option disabled value="">Please select a device type</option>
                    <option value="Defibrillator">Defibrillator</option>
                    <option value="Pacemaker">Pacemaker</option>
                    <option value="Loop Recorder">Loop Recorder</option>
                </select>
            </div>
            <div class="col">
                <label for="Header" class="form-label">Header</label>
                <input type="text" class="form-control" id="Header" name="Header" v-model="formData.Header" @input="handleInput">
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
                    <!-- <div class="btn-group" id="actions" role="group" >
                        <button type="button" class="btn btn-secondary" @click="cancel">Cancel</button>
                    </div> -->
            </div>
        </div>
    </form>
  </div>
</template>