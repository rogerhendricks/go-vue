<script setup>
  import { onMounted,ref, computed } from 'vue'
  import { useRouter } from 'vue-router'
  import { useStore } from 'vuex'
  
  const router = useRouter()

  const store = useStore()

  onMounted(() => {
    store.dispatch('fetchDoctors')
  })
  
</script>
<template>
  <div>
    <h1>Patients</h1>
  </div>
  <div class="text-bg-secondary p-3">
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
  </div>
  <router-link to="/patients/1">To patient</router-link>

  <router-link to="/patients/create">Create Patient</router-link>
</template>