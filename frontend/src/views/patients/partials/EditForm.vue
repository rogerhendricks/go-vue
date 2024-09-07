<script setup>
import { ref, onMounted, computed, reactive, defineProps } from 'vue'
import axios from '../../../axiosConfig'
// import { useRoute } from 'vue-router'
import { useToast } from 'vue-toastification'
import Multiselect from 'vue-multiselect'
import { useStore } from 'vuex'

const props = defineProps(['patient'])
const store = useStore()
// const route = useRoute()
const toast = useToast()
const doctorsFromStore = computed(() => store.state.doctors.doctors || [])
// const devicesFromStore = computed(() => store.state.devices.devices || [])
// const patient = computed(() => store.state.patient.patient)


const formData = reactive({
    name: props.patient.name,
    email: props.patient.email,
    phone: props.patient.phone,
    address: props.patient.address,
    city: props.patient.city,
    state: props.patient.state,
    zip: props.patient.zip,
    country: props.patient.country,
    sex: props.patient.sex,
    dob: props.patient.dob,
    active: props.patient.active,
    // doctors: props.patient.doctors || []
    doctors: props.patient.Doctors.map(doc => {
        return doctorsFromStore.value.find(d => d.ID === doc.ID) || { ID: doc.ID, name: doc.name }
    })
})
// const patient = reactive({ ...formData.value })
// const patient = computed(() => ({ ...formData.value }))

const handleSubmit = async () => {
    try {
        // Adjust the doctors mapping
        const updatedPatient = {
            ...formData,
            doctors: formData.doctors.map(doctor => ({ ID: doctor.ID }))
        }
        // formData.doctors = formData.doctors.map(doctor => ({ ID: doctor.ID }))
        const response = await axios.put(`/api/patients/${props.patient.ID}`, updatedPatient, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
        console.log('Patient updated:', response.data)
        toast.success('Patient updated successfully',{
            timeout: 2000,
        });
        store.commit('setPatient', response.data.patient)
        console.log('Updated patient:', store.state.patient);

    } catch (error) {
        console.error('Error updating patient:', error)
        toast.error('Failed to update patient',{
            timeout: 2000,
        });
    }
}
</script>
<template>
    <p>EditForm</p>
    <form @submit.prevent="handleSubmit">
        <div class="row mb-3">
            <div class="col">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" v-model="formData.name">
            </div>
            <div class="col">
                <label for="email" class="form-label">Email</label>
                <input type="email" class="form-control" id="email" v-model="formData.email">
            </div>
            <div class="col">
                <label for="phone" class="form-label">Phone</label>
                <input type="text" class="form-control" id="phone" v-model="formData.phone">
            </div>
        </div>
        <div class="row mb-3">
            <div class="col">
                <label for="address" class="form-label">Address</label>
                <input type="text" class="form-control" id="address" v-model="formData.address">
            </div>
            <div class="col">
                <label for="city" class="form-label">City</label>
                <input type="text" class="form-control" id="city" v-model="formData.city">
            </div>
            <div class="col">
                <label for="state" class="form-label">State</label>
                <input type="text" class="form-control" id="state" v-model="formData.state">
            </div>
        </div>
        <div class="row mb-3">
            <div class="col">
                <label for="zip" class="form-label">Zip</label>
                <input type="text" class="form-control" id="zip" v-model="formData.zip">
            </div>
            <div class="col">
                <label for="country" class="form-label">Country</label>
                <input type="text" class="form-control" id="country" v-model="formData.country">
            </div>
        </div>
        <div class="row mb-3">
            <div class="col">
                <label  class="form-label">Doctors</label>
                <multiselect 
                v-if="doctorsFromStore.length > 0"
                class="custom__tag multiselect" 
                v-model="formData.doctors" 
                
                
                :options="doctorsFromStore" 
                :multiple="true" 
                :close-on-select="false" 
                :clear-on-select="false"            
                :preserve-search="true" 
                placeholder="Pick some" 
                label="name" 
                track-by="ID" 
                :preselect-first="true">
                <template #selection="{ values, search, isOpen }">
                    <span class="multiselect__single"
                        v-if="values.length"
                        v-show="!isOpen">{{ values.length }} options selected</span>
                </template>
                <template #option="{ option, search, id }">
                    <span :id="id"
                        class="option"
                        :class="{ 'option--selected': formData.doctors.includes(option) }">
                    {{ option.name }}
                    </span>
                </template>
                </multiselect>
            </div>  
        </div>
        <div class="row mb-3">
            <div class="col">
                <button type="submit" class="btn btn-secondary" data-bs-dismiss="modal">Update</button>
            </div>
        </div>
    </form>
</template>