<script setup>
import { ref, onMounted, computed, watch, defineProps } from 'vue'
import axios from '../../axiosConfig'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import {useToast} from 'vue-toastification';
import UpdatForm from './partials/UpdateForm.vue'

const toast = useToast();
const route = useRoute()
const store = useStore()

const props = defineProps(['patient', 'report_id'])
const isLoading = computed(() => store.state.isLoading)
const patientReport = ref({})
const reportId = computed(() => props.report_id);

// Function to fetch the report
const fetchReport = async (reportId) => {
  try {
    const response = await axios.get(`/api/reports/${reportId}/`)
    console.log('Patient report:', response.data.report)
    patientReport.value = response.data.report
  } catch (error) {
    console.error('Error fetching report:', error)
    toast.error('Error fetching report',{
        timeout: 2000
    })
  }
}

// Fetch report on mount
onMounted(async () => {
    console.log('Mounted: report_id is', props.report_id); 
    await fetchReport(props.report_id)
})
watch(() => reportId, async (newReportId) => {
  if (newReportId) {
    console.log('Report ID changed:', newReportId)
    await fetchReport(newReportId)
  }
})

</script>
<template>
    <div class="row">
        <div class="col">
            <span>Report Side</span>
            <div v-if="patientReport">
                <p>Report ID: {{ patientReport.ID }}</p>
                <p>Report Date: {{ patientReport.report_date }}</p>
            </div>
            <div v-else>
                <p>No report data available.</p>
            </div>
        </div>
    </div>
</template>