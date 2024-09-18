<script setup>
import { ref, onMounted, computed, watch, defineProps } from 'vue'
import axios from '../../axiosConfig'
// import { useRoute } from 'vue-router'
// import { useStore } from 'vuex'
import {useToast} from 'vue-toastification';
// import UpdatForm from './partials/UpdateForm.vue'

const toast = useToast();
const props = defineProps(['patient', 'report_id'])
const patientReport = ref({})
const reportId = computed(() => props.report_id);

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

onMounted(async () => {
    console.log('Mounted: report_id is', props.report_id); 
    await fetchReport(props.report_id)
})
watch(() => props.report_id, async (newReportId) => {
  if (newReportId) {
    console.log('Report ID changed:', newReportId)
    await fetchReport(newReportId)
  }
})

function getFileLink(filePath) {
      // console.log('File path:', filePath)
      return `https://dev.nuttynarwhal.com/${filePath}`
  }

</script>
<template>
    <div class="row">
        <div class="col">
            <span>Report Side</span>
            <div v-if="patientReport">
                <p>Report ID: {{ patientReport.ID }}</p>
                <p>Report Date: {{ patientReport.report_date }}</p>
                <p>Report Type: {{ patientReport.file_path }}</p>
                <a 
                    :href="getFileLink(patientReport.file_path)" 
                    target="_blank" 
                    rel="noopener noreferrer">
                    Open Report
                </a>
            </div>
            <div v-else>
                <p>No report data available.</p>
            </div>
        </div>
    </div>
</template>
