<script setup>
import { ref, onMounted, computed, watch, defineProps } from 'vue'
import axios from '../../axiosConfig'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import {useToast} from 'vue-toastification';
import CreateForm from './partials/CreateForm.vue'
import UpdateForm from './partials/UpdateForm.vue'
import Report from './Report.vue'


const props = defineProps(['patient'])
const toast = useToast();
const route = useRoute()
const store = useStore()
const patient = computed(() => store.state.patient.patient || {})
const reportList = computed(() => store.state.reports.reports || [])
const isLoading = computed(() => store.state.isLoading)
let selectedReportId = ref(null)
const patientReports = ref([])

onMounted(async () => {
    await store.dispatch('fetchReports', props.patient.ID)
})

let showForm = ref(false)
const toggleForm = ()=> {
    showForm.value = !showForm.value
}
const selectReport = (id) => {
    if(showForm.value) {
        showForm.value = false
    }
    selectedReportId.value = id
    console.log('Selected report:', selectedReportId.value)
}

const handleReportCreated = async () => {
  // Fetch the updated reports after creating a new one
  await store.dispatch('fetchReports', props.patient.ID)
  showForm.value = false // Close the form after creation
}

const handleReportUpdated = async (updatedReport) => {
  await store.dispatch('fetchReports', props.patient.ID) // Refresh the reports list
  selectedReportId.value = null // Reset the selected report, closing the form
}

const handleReportDeleted = async (deletedReportId) => {
  await store.dispatch('fetchReports', props.patient.ID) // Refresh the reports list
  selectedReportId.value = null // Reset the selected report, closing the form
}
</script>
<template>
    <div class="row">
        <div class="col-md-2">
            <button class="btn btn-secondary mb-2" @click="toggleForm">Create</button>
            <div v-for="report in reportList" :key="report.ID">
                <button type="button" class="btn" @click="selectReport(report.ID)">{{ report.report_date }}</button>
            </div>
        </div>
        <div class="col-md-10">
            <div v-if="showForm">
                <CreateForm  :patient="props.patient" @report-created="handleReportCreated"/>
            </div>
            <div v-else>
                <!-- <Report v-if="selectedReportId" :report_id="selectedReportId" :patient="props.patient"/> -->
                <UpdateForm v-if="selectedReportId" :report_id="selectedReportId" :patient="props.patient" @report-deleted="handleReportDeleted" @report-updated="handleReportUpdated"/>
            </div>
        </div>
    </div>
</template>