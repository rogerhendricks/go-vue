<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import axios from '../../axiosConfig'
import { useRoute } from 'vue-router'
import PatientDetails from './partials/PatientDetails.vue';
import { useStore } from 'vuex'

const route = useRoute()
const store = useStore()
const doctorsFromStore = computed(() => store.state.doctors.doctors || []) 
const devicesFromStore = computed(() => store.state.devices.devices || [])
const leadsFromStore = computed(() => store.state.leads.leads || [])
const patient = computed(() => store.state.patient.patient || {})
const isLoading = computed(() => store.state.isLoading)

// watch(patient, (newPatient) => {
//   console.log('Updated patient data:', newPatient);
// });

onMounted(async () => {
  const patientId = route.params.id
  store.commit('setLoading', true)

  await store.dispatch('fetchPatient', patientId)
  store.commit('setLoading', false)
  if (doctorsFromStore.value.length === 0) {
    await store.dispatch('fetchDoctors')
  }
  if (devicesFromStore.value.length === 0) {
    await store.dispatch('fetchDevices')
  }
  if (leadsFromStore.value.length === 0) {
    await store.dispatch('fetchLeads')
  }
})
</script>
<template>
  <div v-if="!isLoading">
    <PatientDetails v-show="Object.keys(patient).length" :patient="patient"/>
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

  <!-- modal dialog for edit patient -->
</template>