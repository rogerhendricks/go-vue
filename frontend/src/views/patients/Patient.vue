<script setup>
import { ref, onMounted } from 'vue'

onMounted(() => {
  console.log('Patient ID:', $route.params.id)
  $store.dispatch('fetchDoctors')
  console.log('Doctors:', $store.state.doctors)
})
</script>
<template>
  <div class="container text-bg-secondary p-3">
    <span>Patient details go here</span>
  </div>
  <div class="row">
    <div class="col">
      <span>Report List</span>
    </div>
    <div class="col">
      <span>Report Details</span>
    </div>
  </div>

  <div class="row text-bg-secondary p-3">
      <div class="col">
        <span>Device list</span>{{$store.state.devices}}    
      </div>
  </div>

  <!-- modal dialog for edit patient -->
</template>