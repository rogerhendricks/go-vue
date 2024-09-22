<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import axios from '../../../axiosConfig'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'
import { PDFDocument } from 'pdf-lib'
import { useToast } from 'vue-toastification'

const toast = useToast()
const route = useRoute()
const router = useRouter()
const store = useStore()
const report = computed(() => store.state.report)

const formData = ref({
  report_date: '',
  file_path: '',
  patient: null,
  files: null,
})

// Watch for when the report is fetched and update formData accordingly
watch(report, (newReport) => {
  if (newReport) {
    formData.value.report_date = newReport.report_date
    formData.value.file_path = newReport.file_path
    formData.value.patient = newReport.patient_id
  }
})

onMounted(async () => {
  // Fetch the report if it's not already in the store
  if (!report.value) {
    await store.dispatch('fetchReport', route.params.id)
  } else {
    // Populate form data if report is already available
    formData.value.report_date = report.value.report_date
    formData.value.file_path = report.value.file_path
    formData.value.patient = report.value.patient_id
  }
})

const handleSubmit = async () => {
  try {
    // Create a new PDF document to merge files
    const mergedPdfDoc = await PDFDocument.create()

    // Loop through each selected file and merge them
    for (const file of formData.value.files) {
      const arrayBuffer = await file.arrayBuffer()
      const pdfDoc = await PDFDocument.load(arrayBuffer)
      const copiedPages = await mergedPdfDoc.copyPages(pdfDoc, pdfDoc.getPageIndices())
      copiedPages.forEach((page) => mergedPdfDoc.addPage(page))
    }

    // Serialize the merged PDF into a Blob
    const mergedPdfBytes = await mergedPdfDoc.save()
    const mergedPdfBlob = new Blob([mergedPdfBytes], { type: 'application/pdf' })

    // Create FormData and append the merged PDF Blob
    const uploadData = new FormData()
    uploadData.append('patient_id', formData.value.patient)
    uploadData.append('report_date', formData.value.report_date)
    uploadData.append('file', mergedPdfBlob, 'merged_report.pdf')

    // Send the FormData with merged PDF to the backend
    const response = await axios.put(`/api/reports/${route.params.id}/`, uploadData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    // Update the report in the Vuex store
    store.commit('setReport', response.data.report)
    console.log('Report updated:', response.data.report)
    toast.success('Report updated', { timeout: 2000 })
  } catch (error) {
    console.error('Error updating report:', error)
    toast.error('Error updating report', { timeout: 2000 })
  }
}

const handleFileUpload = (event) => {
  formData.value.files = event.target.files
}
function goBackUsingBack() {
      if (router) {
        router.back();
      }
    }

</script>

<template>
  <button class="btn" type="button" @click="goBackUsingBack">Go Back </button>
  <div class="row">
    <div class="col">
      <form @submit.prevent="handleSubmit">
        <div class="mb-3">
          <label for="report_date" class="form-label">Report Date:</label>
          <input
            id="report_date"
            class="form-control"
            v-model="formData.report_date"
            type="date"
            required
          />
        </div>
        <div class="mb-3">
          <label for="file_path" class="form-label">Files:</label>
          <input
            class="form-control"
            type="file"
            id="file_path"
            multiple
            @change="handleFileUpload"
          />
        </div>
        <div class="mb-3">
          <button class="btn btn-primary" type="submit">Submit</button>
        </div>
      </form>
    </div>
  </div>
</template>
