<script setup>

import { ref, onMounted, computed } from 'vue'
import { useStore } from 'vuex'
import axios from '../../../axiosConfig'
import { useRoute } from 'vue-router'
import {useToast} from 'vue-toastification';

const toast = useToast();

const route = useRoute()
const store = useStore()
const doctorsFromStore = computed(() => store.state.doctors.doctors || [])
const devicesFromStore = computed(() => store.state.devices.devices || [])
const patient = computed(() => store.state.patient.patient || {})
const loading = ref(false)
const newDevice = ref({ 
  implant_date: '',
  explant_date: '',
  device_id: '',
  patient_id: route.params.id,
  doctor_id: '', 
})

const editedDevice = ref({
  ID: '',
  implant_date: '',
  explant_date: '',
  device_id: '',
  doctor_id: ''
});

const implantedDevices = ref([])
onMounted(async () => {
    getImplantedDevices()
    if (!devicesFromStore.value) {
        await store.dispatch('fetchDevices')
    }
})

async function getImplantedDevices() {
        loading.value = true
        const patientId = route.params.id
        try {
         const response = await axios.get(`/api/${patientId}/implantedDevices`)
         implantedDevices.value = response.data.implantedDevices
            console.log('Implanted Devices:', implantedDevices.value)
        } catch (error) {
            console.error('Error fetching implanted devices:', error)
        } finally {
            loading.value = false
        }
}

async function createDevice() {
    const patientId = Number(route.params.id)
    try {
      const response = await axios.post(`/api/${patientId}/implantedDevices`, {
            implant_date: newDevice.value.implant_date,
            explant_date: newDevice.value.explant_date || null, // Send null for optional fields
            device_id: newDevice.value.device_id,
            patient_id: patientId, // Send patient_id as a number
            doctor_id: newDevice.value.doctor_id
        });        
        
        implantedDevices.value.push(response.data.device)
        newDevice.value = { 
          implant_date: '',
          explant_date: '',
          device_id: '',
          patient_id: patientId,
          doctor_id: '',
         }
        hideModal('createDeviceModal')
        toast.success('Device created successfully',{
            timeout: 2000,
        });
    } catch (error) {
        console.error('Error creating device:', error)
        toast.error('Failed to create device',{
            timeout: 2000,
        });
    }
}

function openEditModal(device) {
    editedDevice.value = { ...device }
    console.log('Edited Device:', editedDevice.value)
    showModal('editDeviceModal')
}

async function updateDevice() {
    try {
        const response = await axios.put(`/api/implantedDevices/${editedDevice.value.ID}`, editedDevice.value)
        const editedDeviceId = Number(editedDevice.value.ID);
        const index = implantedDevices.value.findIndex(device => device.ID === editedDeviceId);
        if (index !== -1) {
          implantedDevices.value[index] = response.data.implantedDevice; // Assuming the response contains the updated device
        } else {
          console.warn('API response does not contain a valid device object');
        }
        hideModal('editDeviceModal')
        toast.success('Device updated successfully',{
            timeout: 2000,
        });
    } catch (error) {
      toast.error('Failed to update device',{
            timeout: 2000,
        });
        console.error('Error updating device:', error)
    }
}

  function confirmDelete(deviceId) {
      if (window.confirm('Do you want to delete this device?')) {
        deleteDevice(deviceId)
      }
    }

  async function deleteDevice(deviceId) {
      try {
          await axios.delete(`/api/implantedDevices/${deviceId}`)
          implantedDevices.value = implantedDevices.value.filter(d => d.ID !== deviceId)
          toast.success('Device deleted successfully',{
            timeout: 2000,
        });
        } catch (error) {
          toast.error('Error deleting device',{
            timeout: 2000,
        });
          console.error('Error deleting device:', error)
      }
  }

  function showModal(modalId) {
      const modal = new bootstrap.Modal(document.getElementById(modalId))
      modal.show()
  }

function hideModal(modalId) {
    const modalEl = document.getElementById(modalId)
    const modal = bootstrap.Modal.getInstance(modalEl)
    modal.hide()
}

</script>
<template>
  <div class="p-2">
    <div class="row">
      <div class="col-md-2">
        <div class="d-grid d-md-flex justify-content-md-start"> 
          <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#createDeviceModal">Add Device</button>
        </div>
      </div>
      <div v-if="loading" class="col">Loading...</div>
      <div class="col px-4" v-else>
        <div v-for="device in implantedDevices" :key="device.id">
          <div class="row d-flex flex-wrap text-bg-secondary p-3 rounded-2 mb-3"> 
                <div class="col">
                  <b>Implant Date:</b> {{ device.implant_date }}
                </div>
                <div class="col">
                  <b>Device:</b> {{ device.Device.manufacturer }} {{ device.Device.name }}
                </div>
                <div class="col">
                  <b>Implanter:</b> {{ device.Doctor.name }}
                </div>
                <div class="col">
                  <b>Explant Date:</b> {{ device.explant_date }}
                </div>
                <div class="col">
                  <div class="btn-group">
                    <button @click="openEditModal(device)" class="btn btn-sm btn-primary">Edit</button>
                    <button @click="confirmDelete(device.ID)" class="btn btn-sm btn-danger">Delete</button>
                  </div>
                </div>
          </div>
        </div>
      </div>
    </div>
  </div>
    <!-- Create Device Modal -->
    <div class="modal fade" id="createDeviceModal" tabindex="-1" aria-labelledby="createDeviceModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="createDeviceModalLabel">Create Device</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="createDevice">
            <div class="row mb-3">
              <div class="col">
                  <label for="implantDate" class="form-label">Implant Date</label>
                  <input type="date" v-model="newDevice.implant_date" class="form-control" id="implantDate" required>
              </div>
              <div class="col">
                <label for="explantDate" class="form-label">Explant Date</label>
                <input type="date" v-model="newDevice.explant_date" class="form-control border border-danger" id="explantDate">
              </div>
              <div class="col">
                <label for="device" class="form-label">Device</label>
                <select v-model="newDevice.device_id" class="form-select" id="device" required>
                  <option v-for="device in devicesFromStore" :key="device.ID" :value="device.ID">{{ device.name }}</option>
                </select>
              </div>
              <div class="col">
                <label for="doctor" class="form-label">Doctor</label>
                <select v-model="newDevice.doctor_id" class="form-select" id="doctor">
                  <option v-for="doctor in doctorsFromStore" :key="doctor.ID" :value="doctor.ID">{{ doctor.name }}</option>
                </select>
              </div>
            </div>
            <button type="submit" class="btn btn-primary">Save</button>
          </form>
        </div>
      </div>
    </div>
  </div>
    <!-- Edit Device Modal -->
    <div class="modal fade" id="editDeviceModal" tabindex="-1" aria-labelledby="editDeviceModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="editDeviceModalLabel">Edit Device</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateDevice">
              <div class="row mb-3">
                  <div class="col">
                    <label for="editDeviceImplantDate" class="form-label">Implant Date</label>
                    <input type="date" v-model="editedDevice.implant_date" class="form-control" id="editDeviceImplantDate" required>
                  </div>
                  <div class="col">
                    <label for="editDeviceExplantDate" class="form-label">Explant Date</label>
                    <input type="date" v-model="editedDevice.explant_date" class="form-control border border-danger" id="editDeviceExplantDate">
                  </div>
                  <div class="col">
                    <label for="editDeviceDevice" class="form-label">Device</label>
                    <select v-model="editedDevice.device_id" class="form-select" id="editDeviceDevice" required>
                      <option v-for="device in devicesFromStore" :key="device.ID" :value="device.ID">{{ device.name }}</option>
                    </select>
                  </div>
                  <div class="col">
                    <label for="editDeviceDoctor" class="form-label">Doctor</label>
                    <select v-model="editedDevice.doctor_id" class="form-select" id="editDeviceDoctor">
                      <option v-for="doctor in doctorsFromStore" :key="doctor.ID" :value="doctor.ID">{{ doctor.name }}</option>
                    </select>
                  </div>
              </div>
              <button type="submit" class="btn btn-primary">Update</button>
            </form>
          </div>
        </div>
      </div>
    </div>
</template>