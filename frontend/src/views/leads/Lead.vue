<script setup>
import axios from '../../axiosConfig'
import { ref } from 'vue'
import Form from './partials/Form.vue'

const props = defineProps(['id'])
const lead = ref({})
const loading = ref(false)

async function getLead() {
  loading.value = true
  try {
    const response = await axios.get(`/api/leads/${props.id}`)        

    lead.value = response.data.lead
  } catch (error) {
    console.error('Error fetching lead:', error)
  } finally {
    loading.value = false
  }
}

getLead()
</script>
<template>
  <div>
    <h1>Lead</h1>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <Form :modelValue="lead" @update:modelValue="updateLead" />
    </div>
  </div>
</template> 