<script setup>
import { computed, ref, defineProps, reactive, watch } from 'vue'
import { useStore } from 'vuex'
import axios from '../../../axiosConfig'
import { useRoute } from 'vue-router'
import { useToast } from 'vue-toastification'
import Multiselect from 'vue-multiselect'
import ImplantedDevices from './ImplantedDevices.vue'
import ImplantedLeads from './ImplantedLeads.vue'
// import EditForm from './EditForm.vue'
const props = defineProps({
  patient: {
    type: Object,
    required: true,
    default: () => ({})
  }
})

// const route = useRoute()
const toast = useToast()
const store = useStore()
const doctorsFromStore = computed(() => store.state.doctors.doctors || [])
const patientFromStore = computed(() => store.state.patient.patient || {})
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
watch(
  () => props.patient,
  (newPatient) => {
    if (newPatient) {
      console.log('Received updated patient:', newPatient); // Check if new patient data is received
      
      Object.assign(formData, {
        name: newPatient.name,
        email: newPatient.email,
        phone: newPatient.phone,
        address: newPatient.address,
        city: newPatient.city,
        state: newPatient.state,
        zip: newPatient.zip,
        country: newPatient.country,
        sex: newPatient.sex,
        dob: newPatient.dob,
        active: newPatient.active,
        doctors: newPatient.Doctors.map(doc => {
          return doctorsFromStore.value.find(d => d.ID === doc.ID) || { ID: doc.ID, name: doc.name };
        })
      });
    }
  },
  { immediate: true }
);


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
        store.commit('setPatient', response.data)
        store.dispatch('fetchPatients')
        toast.success('Patient updated successfully',{
            timeout: 2000,
        });
    } catch (error) {
        console.error('Error updating patient:', error)
        toast.error('Failed to update patient',{
            timeout: 2000,
        });
    }
}
</script>
<template>
    <nav>
        <div class="nav nav-tabs" id="nav-tab" role="tablist">
            <button class="nav-link active" id="nav-patient-tab" data-bs-toggle="tab" data-bs-target="#nav-patient"
                type="button" role="tab" aria-controls="nav-patient" aria-selected="true">Patient</button>
            <button class="nav-link" id="nav-devices-tab" data-bs-toggle="tab" data-bs-target="#nav-devices"
                type="button" role="tab" aria-controls="nav-devices" aria-selected="false">Devices</button>
            <button class="nav-link" id="nav-leads-tab" data-bs-toggle="tab" data-bs-target="#nav-leads" type="button"
                role="tab" aria-controls="nav-leads" aria-selected="false">Leads</button>
            <button type="button" class="btn" data-bs-toggle="modal" data-bs-target="#editPatientModal">Edit</button>
        </div>
    </nav>
    <div class="tab-content text-bg-light border border-top-0" id="nav-tabContent">
        <div class="tab-pane fade show active" id="nav-patient" role="tabpanel" aria-labelledby="nav-patient-tab"
            tabindex="0">
            <div class="p-2">
                <div class="row">
                    <div class="col">
                        <p>Name: {{ patient.name }}</p>
                    </div>
                    <div class="col">
                        <p>DOB: {{ patient.dob }}</p>
                    </div>
                    <div class="col">
                        <p>Phone: {{ patient.phone }}</p>
                    </div>
                    <div class="col">
                        <p>Email: {{ patient.email }}</p>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p>Address: {{ patient.address }}</p>
                    </div>
                    <div class="col">
                        <p>City: {{ patient.city }}</p>
                    </div>
                    <div class="col">
                        <p>State: {{ patient.state }}</p>
                    </div>
                    <div class="col">
                        <p>Zip: {{ patient.zip }}</p>
                    </div>
                    <div class="col">
                        <p>Country: {{ patient.country }}</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="tab-pane fade" id="nav-devices" role="tabpanel" aria-labelledby="nav-devices-tab" tabindex="0">
            <ImplantedDevices />
        </div>
        <div class="tab-pane fade" id="nav-leads" role="tabpanel" aria-labelledby="nav-leads-tab" tabindex="0">
            <ImplantedLeads />
        </div>
    </div>

    <!-- Modal -->
<div class="modal fade" id="editPatientModal" tabindex="-1" aria-labelledby="editPatientModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h1 class="modal-title fs-5" id="editPatientModalLabel">Update Patient</h1>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>

        <form @submit.prevent="handleSubmit">
            <div class="modal-body">
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
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                <button type="submit" class="btn btn-primary">Save changes</button>
            </div>
        </form>


    </div>
  </div>
</div>

</template>