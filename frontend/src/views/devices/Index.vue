<script setup>
import axios from '../../axiosConfig'
import { ref, computed } from 'vue'
import CreateForm from './partials/CreateForm.vue'
import EditForm from './partials/EditForm.vue'

const devices = ref([])
const loading = ref(false)
const selectedDevice = ref(null)
let searchTerm = ref('');
let currentPage = ref(1);
let pageSize = ref(10);

const filteredDevices = computed(() => {
  if (!searchTerm.value) {
    return devices.value;
  }
  return devices.value.filter(device =>
    device.manufacturer.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    device.name.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    device.deviceModel.toLowerCase().includes(searchTerm.value.toLowerCase())
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
    <h1>Devices</h1>
  </div>

  <div class="row">
    <div class="col-md-3">
      <CreateForm />
    </div>
    <div class="col-md-9 mt-4">
        <div class="mt-2">
          <div class="input-group mb-3">
            <input v-model="searchTerm" placeholder="Search..." class="form-control">
            <button class="btn btn-outline-secondary" type="button" id="button-addon2">Search</button>
          </div>
        </div>
        <div class="my-2">
          <button class="btn btn-primary" @click="getDevices">Refresh</button>
        </div>
        <ol>
          <li v-for="device in paginatedDevices">
            <!-- <button type="button" class="btn" @click="selectedDevice = device">{{ device.name }}</button> -->
            <!-- Button trigger modal -->
            <button type="button" class="btn" @click="selectedDevice = device" data-bs-toggle="modal" data-bs-target="#staticBackdrop">
              {{device.manufacturer}} {{ device.name }}
            </button>
            <!-- <EditForm v-if="selectedDevice === device" :device="selectedDevice" @cancel="selectedDevice = null" /> -->
          </li>
        </ol>
        <div class="d-flex justify-content-center">
          <nav aria-label="Page navigation example">
                <ul class="pagination
                justify-content-center">
                    <li class="page-item" :class="{ disabled
                    : currentPage === 1 }">
                        <a class="page-link" href="#" @click="currentPage = 1">First</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === 1 }">
                        <a class="page-link" href="#" @click="currentPage--">Previous</a>
                    </li>
                    <li class="page-item" v-for="page in Math.ceil(filteredDevices.length / pageSize)" :key="page" :class="{ active
                    : currentPage === page }">
                        <a class="page-link" href="#" @click="currentPage = page">{{ page }}</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === Math.ceil(filteredDevices.length / pageSize) }">
                        <a class="page-link" href="#" @click="currentPage++">Next</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === Math.ceil(filteredDevices.length / pageSize) }">
                        <a class="page-link" href="#" @click="currentPage = Math.ceil(filteredDevices.length / pageSize)">Last</a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
  </div>




<!-- Modal -->
<div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h1 class="modal-title fs-5" id="staticBackdropLabel">Edit Device</h1>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <EditForm v-if="selectedDevice" :device="selectedDevice" :key="selectedDevice.ID" @cancel="selectedDevice = null" />
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="selectedDevices = null" >Close</button>
      </div>
    </div>
  </div>
</div>
</template>