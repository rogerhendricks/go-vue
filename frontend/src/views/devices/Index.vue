<script setup>
import axios from '../../axiosConfig'
import { ref, computed } from 'vue'
import CreateForm from './partials/CreateForm.vue'
import EditForm from './partials/EditForm.vue'

const devices = ref([])
const loading = ref(false)
// const currentComponent = ref(null) 
const selectedDevice = ref(null)
// const formVisible = ref(false)
let searchTerm = ref('');
let currentPage = ref(1);
let pageSize = ref(10);

// function toggleForm() {
//   formVisible.value = !formVisible.value
//   currentComponent.value = formVisible.value ? CreateForm : null
// }

const filteredDevices = computed(() => {
  if (!searchTerm.value) {
    return devices.value;
  }
  return devices.value.filter(device =>
    device.name.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    device.manufacturer.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    device.deviceModel.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    device.type.toLowerCase().includes(searchTerm.value.toLowerCase())
  );
});

const paginatedDevices =computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    const end = currentPage.value * pageSize.value;
    return filteredDevices.value.slice(start,end);
});

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
    <!-- <div class="col-md-9">
      <button class="btn btn-primary" @click="getDevices">Refresh</button>
      <div v-if="loading">Loading...</div>
      <div v-else>
        <ul>
          <li v-for="device in devices" :key="device.ID">
            <button type="button" class="btn" @click="selectedDevice = device">{{ device.name }}</button>
            <EditForm v-if="selectedDevice === device" :device="selectedDevice" @cancel="selectedDevice = null" />
          </li>
        </ul>
      </div>
    </div> -->
    <div class="col-md-9">
        <div class="mt-2">
            <input v-model="searchTerm" placeholder="Search..." class="form-control">
        </div>
        <div class="my-2">
          <button class="btn btn-primary" @click="getDevices">Refresh</button>
        </div>
        <ul>
          <li v-for="device in paginatedDevices">
            <button type="button" class="btn" @click="selectedDevice = device">{{ device.name }}</button>
            <EditForm v-if="selectedDevice === device" :device="selectedDevice" @cancel="selectedDevice = null" />
          </li>
        </ul>
        <div class="d-flex justify-content-center">
          <div v-if="filteredDevices.length > 0" class="d-inline-flex mt-2 xs:mt-0">
            <div class="btn-group" role="group"> 
              <button @click="currentPage--" :disabled="currentPage === 1" class="btn btn-secondary">
                <!-- <svg class="w-3.5 h-3.5 me-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10"> -->
                  <!-- <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5H1m0 0 4 4M1 5l4-4"/> -->
                <!-- </svg> -->
                Prev
              </button>
              <span class="btn btn-secondary">Page {{ currentPage }}</span>
              <button @click="currentPage++" :disabled="currentPage === Math.ceil(filteredDevices.length / pageSize)" class="btn btn-secondary">
                Next
                <!-- <svg class="w-3.5 h-3.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10"> -->
                  <!-- <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/> -->
                <!-- </svg> -->
              </button>
            </div>
          </div>
        </div>
    </div>
  </div>

</template>