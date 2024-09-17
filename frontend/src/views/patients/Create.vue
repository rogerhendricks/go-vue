<script setup>
import { ref, reactive, computed, watchEffect, onMounted } from 'vue'
import axios from '../../axiosConfig'
import { useToast } from 'vue-toastification'
import { useStore } from 'vuex'
import Multiselect from 'vue-multiselect'
import { useRouter } from 'vue-router'

const router = useRouter()
const store = useStore()
const toast = useToast()

// Access the correct level of the store's state
const doctorsFromStore = computed(() => store.state.doctors.doctors || []) 
// const devicesFromStore = computed(() => store.state.devices.devices || []) 
// const leadsFromStore = computed(() => store.state.leads.leads || []) 

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

const handleSubmit = async () => {
  try {
    const response = await axios.post('/api/patients', patient, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    console.log('Patient created:', response.data)
    // router.push({ name: 'Patient', params: { id: response.data.patient.ID } })
    toast.success('Doctor updated successfully',{
            timeout: 2000,
        });
    Object.assign(patient, initialPatientSate);
  } catch (error) {
    console.error('Error creating patient:', error)
    toast.error('Failed to create patient',{
            timeout: 2000,
        });
  }
}
onMounted(() => {
    store.dispatch('fetchLeads')
    store.dispatch('fetchDevices')
  })
</script>

<template>
  <span>Create Patient</span>
  <form @submit.prevent="handleSubmit">
    <div class="row mt-2">
      <div class="col">
        <input class="form-control" v-model="patient.name" type="text" placeholder="Name" />
      </div>
      <div class="col">
        <input class="form-control"  type="date" v-model="patient.dob" placeholder="Date of Birth" />
      </div>
    </div>
    <div class="row mt-2">      
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
    </div>
    <div class="row mt-2">
      <div class="col">
        <input class="form-control" v-model="patient.city" type="text" placeholder="City" />
      </div>
      <div class="col">
        <input class="form-control" v-model="patient.state" type="text" placeholder="State" />
      </div>
    </div>
    <div class="row mt-2">
      <div class="col">
        <input class="form-control" v-model="patient.country" type="text" placeholder="Country" />
      </div>
      <div class="col">
        <input class="form-control" v-model="patient.zip" type="text" placeholder="Zip" />
      </div>
    </div>
    <div class="row mt-2">
      <div class="col">
        <label  class="form-label">Doctors</label>
        <multiselect 
        class="custom__tag multiselect" 
        v-model="patient.doctors" 
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
                  :class="{ 'option--selected': patient.doctors.includes(option) }">
              {{ option.name }}
            </span>
          </template>
        </multiselect>
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
    <div class="mt-2">
      <button type="submit" class="btn btn-primary mt-2">Submit</button>
    </div>
  </form>

</template>
