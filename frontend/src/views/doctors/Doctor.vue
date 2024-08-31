<script setup>
import axios from '../../axiosConfig'
import { ref, computed } from 'vue'
import EditForm from './partials/EditForm.vue'
import { useRoute } from 'vue-router'

const loading = ref(false)
const route = useRoute()
const doctor = ref(null)

async function getDoctor() {
    loading.value = true
    try {
        const response = await axios.get(`/api/doctors/${route.params.id}`)
        doctor.value = response.data.doctor
    } catch (error) {
        console.error('Error fetching doctor:', error)
    } finally {
    loading.value = false
  }
}

getDoctor()
</script>
<template>
        <div v-if="loading">Loading...</div>
        <div v-else>
            <EditForm :doctor="doctor" />
        </div>
</template>