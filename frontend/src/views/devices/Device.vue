<script setup>
import axios from '../../axiosConfig'
import { ref } from 'vue'
import Form from './partials/Form.vue'

const props = defineProps(['id'])
const device = ref({})
const loading = ref(false)

async function getDevice() {
  loading.value = true
  try {
    const response = await axios.get(`/api/devices/${props.id}`)        

    device.value = response.data.device
  } catch (error) {
    console.error('Error fetching device:', error)
  } finally {
    loading.value = false
  }
}

getDevice()
</script>
<template>
  <div>
    <h1>Device</h1>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <Form :modelValue="device" @update:modelValue="updateDevice" />
    </div>
  </div>
</template> 

