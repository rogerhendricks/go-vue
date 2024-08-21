<script setup>
import axios from '../../axiosConfig'
import { ref } from 'vue'
import CreateForm from './partials/CreateForm.vue'
import EditForm from './partials/EditForm.vue'

const devices = ref([])
const loading = ref(false)
const currentComponent = ref(null) 
const selectedDevice = ref(null)
const formVisible = ref(false)

function toggleForm() {
  formVisible.value = !formVisible.value
  currentComponent.value = formVisible.value ? CreateForm : null
}


async function getDevices() {
  loading.value = true
  try {
    const response = await axios.get('/api/devices')
    devices.value = response.data.devices
  } catch (error) {
    console.error('Error fetching devices:', error)
  } finally {
    loading.value = false
  }
}
getDevices()

</script>

<template>
  <div class="row">
    <div class="col-md-3">
      <h1>Devices</h1>
      <!-- <button class="btn btn-success" @click="loadCreateForm">
        {{ formVisible ? 'Cancel' : 'New Device' }}
      </button> -->
      <CreateForm />
    </div>
    <div class="col-md-9">
      <button class="btn btn-primary" @click="getDevices">Refresh</button>
      <div v-if="loading">Loading...</div>
      <div v-else>
        <ul>
          <li v-for="device in devices" :key="device.ID">
            <span @click="selectedDevice = device">{{ device.name }}</span>
            <EditForm v-if="selectedDevice === device" :device="selectedDevice" @cancel="selectedDevice = null" />
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>