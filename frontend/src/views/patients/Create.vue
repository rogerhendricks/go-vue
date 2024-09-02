<script setup>
import { ref, reactive, computed, watchEffect, onMounted } from 'vue'
import axios from '../../axiosConfig'
import { useToast } from 'vue-toastification'
import { useStore } from 'vuex'

const store = useStore()
const toast = useToast()

// Access the correct level of the store's state
const doctorsFromStore = computed(() => store.state.doctors.doctors || []) 
const devicesFromStore = computed(() => store.state.devices.devices || []) 
const leadsFromStore = computed(() => store.state.leads.leads || []) 

const initialPatientSate = reactive({
  name: '',
  email: '',
  phone: '',
  address: '',
  city: '',
  state: '',
  zip: '',
  country: '',
  sex: '',
  dob: '',
  active: true,
  doctors: []
})

const patient = reactive({ ...initialPatientSate })

const createPatient = async () => {
  try {
    const response = await axios.post('/api/patients', patient)
    console.log('Patient created:', response.data)
    router.push({ name: 'Patient', params: { id: response.data.patient.ID } })
  } catch (error) {
    console.error('Error creating patient:', error)
    toast.error('Failed to create patient')
  }
}
onMounted(() => {
    store.dispatch('fetchLeads')
    store.dispatch('fetchDevices')
  })
</script>

<template>
  <span>Create Form</span>
  <div class="row mt-2">
    <div class="col">
      <input class="form-control" v-model="patient.name" type="text" placeholder="Name" />
    </div>
    <div class="col">
      <input class="form-control" v-model="patient.email" type="email" placeholder="Email" />
    </div>
    <div class="col">
      <input class="form-control" v-model="patient.phone" type="text" placeholder="Phone" />
    </div>
  </div>
  <div class="row mt-2">
    <div class="col">
      <input class="form-control" v-model="patient.address" type="text" placeholder="Address" />
    </div>
    <div class="col">
      <input class="form-control" v-model="patient.city" type="text" placeholder="City" />
    </div>
  </div>
  <div class="row mt-2">
    <div class="col">
      <input class="form-control" v-model="patient.state" type="text" placeholder="State" />
    </div>
    <div class="col">
      <input class="form-control" v-model="patient.country" type="text" placeholder="Country" />
    </div>
    <div class="col">
      <input class="form-control" v-model="patient.zip" type="text" placeholder="Zip" />
    </div>
  </div>
  <div class="row mt-2">
    <div class="col">
      <select class="form-select" v-model="patient.doctors" multiple v-if="doctorsFromStore && doctorsFromStore.length">
        <option v-for="doctor in doctorsFromStore" :key="doctor.ID" :value="doctor.ID">
          {{ doctor.name }}
        </option>
      </select>
    </div>
    <!-- <div class="col">
      <select v-if="devicesFromStore && devicesFromStore.length">
        <option v-for="device in devicesFromStore" :key="device.ID" :value="device.ID">
          {{ device.name }}
        </option>
      </select>
    </div>
    <div class="col">
      <select v-if="leadsFromStore && leadsFromStore.length">
        <option v-for="lead in leadsFromStore" :key="lead.ID" :value="lead.ID">
          {{ lead.name }}
        </option>
      </select>
    </div> -->
  </div>


</template>
