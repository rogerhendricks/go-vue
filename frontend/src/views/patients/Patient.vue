<script setup>
import { ref, onMounted, computed } from 'vue'
import axios from '../../axiosConfig'
import { useRoute } from 'vue-router'
import PatientDetails from './partials/PatientDetails.vue';
import { useStore } from 'vuex'

const route = useRoute()
const store = useStore()
const doctorsFromStore = computed(() => store.state.doctors.doctors || []) 
const devicesFromStore = computed(() => store.state.devices.devices || [])
const patient = computed(() => store.state.patient.patient || {})
const isLoading = computed(() => store.state.isLoading)

onMounted(async () => {
  const patientId = route.params.id
  store.commit('setLoading', true)
  await store.dispatch('fetchPatient', patientId)
  store.commit('setLoading', false)
  console.log('patient from patient.vue', store.state.patient)
  if (doctorsFromStore.value.length === 0) {
    await store.dispatch('fetchDoctors')
  }
  if (devicesFromStore.value.length === 0) {
    await store.dispatch('fetchDevices')
  }
})
</script>
<template>
  <div v-if="!isLoading">
    <PatientDetails v-if="Object.keys(patient).length" :patient="patient"/>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
  <div class="row">
    <div class="col-md-2">
      <span>Report List</span>
    </div>
    <div class="col-md-10">
      <span>Report Details</span>
    </div>
  </div>

  <div class="row text-bg-secondary p-3">
      <div class="col">
        <span>Device list</span> {{$store.state.devices}}    
      </div>
      <div class="col">
        <span>Doctors List:</span> 
        <div v-for="doctor in doctorsFromStore">
          {{ doctor.name }}
        </div>
      </div>
  </div>

  <!-- modal dialog for edit patient -->
</template>