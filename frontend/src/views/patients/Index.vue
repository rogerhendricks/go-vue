<script setup>
  import { onMounted,ref, computed } from 'vue'
  import axios from '../../axiosConfig'
  // import { useRouter } from 'vue-router'
  import { useStore } from 'vuex'
  import Create from './Create.vue'

  // const router = useRouter()

  const store = useStore()
  const searchTerm = ref('');
  const patients = ref([])
  onMounted(() => {
    store.dispatch('fetchDoctors')
  })
  
  async function getPatientsFromSearch() {
    try {
      const response = await axios.get('/api/patients', {
        params: {
          search: searchTerm.value
        }
      }).then((response) => {
        patients.value = response.data
        console.log('success', response.data)
      })
    } catch (error) {
      console.error('Error fetching patients:', error)
    }
  }
</script>
<template>
  <div>
    <h1>Patients</h1>
  </div>
  <div class="row">
    <div class="col-4">
      <Create />
    </div>
    <div class="col-8">
      <div class="input-group mb-3">
        <input type="text" class="form-control" placeholder="Search" aria-label="Search" aria-describedby="search-button">
        <button class="btn btn-outline-secondary" type="button" id="search-button" @click="getPatientsFromSearch">Search</button>
      </div>
      <div>
        <div >
        <table class="table">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Name</th>
              <th scope="col">Phone</th>
              <!-- <th scope="col">Actions</th> -->
            </tr>
          </thead>
          <tbody>
            <tr v-for="patient in patients.patients" :key="patient.ID">
              <th scope="row">{{ patient.ID }}</th>
              <td>{{ patient.name }}
                <router-link :to="'/patients/' + patient.ID" class="btn btn-primary">View</router-link>
              </td>
              <td>{{ patient.phone }}</td>
              <!-- <td>
                <router-link :to="'/patients/' + patient.ID" class="btn btn-primary">View</router-link>
                <router-link :to="'/patients/' + patient.ID + '/edit'" class="btn btn-warning">Edit</router-link>
                <button class="btn btn-danger" @click="deletePatient(patient.id)">Delete</button>
              </td> -->
            </tr>
          </tbody>
        </table>
        </div>
        <!-- <div v-else>
          <p>No matching patients found</p>
        </div> -->
      </div>
    </div>
  </div>
  <!-- <div class="text-bg-secondary p-3">
    <span>Patient list with search goes here</span>
  </div>
  <div class="row text-bg-secondary p-3">
      <div class="col">
        <button type="button" @click="$store.dispatch('fetchDevices')">Click Me!</button>
      </div>
      <div class="col">
        <span>Device list</span>{{$store.state.devices}}    
      </div>
      <div class="col">
        <span>Doctor list</span>{{$store.state.doctors}}    
      </div>
  </div> -->
  <!-- <router-link to="/patients/1">To patient</router-link>

  <router-link to="/patients/create">Create Patient</router-link> -->
</template>