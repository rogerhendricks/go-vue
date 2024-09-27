<script setup>
import { ref, onMounted, computed, watch, defineProps, defineEmits } from 'vue'
import axios from '../../../axiosConfig'
import { useStore } from 'vuex'
import { useRoute, useRouter} from 'vue-router'
import { PDFDocument, StandardFonts, rgb } from 'pdf-lib'
import { useToast } from 'vue-toastification'

const props = defineProps(['patient', 'report_id'])

const currentUser = JSON.parse(localStorage.getItem('user') || 'null')
const toast = useToast()
const route = useRoute()
const router = useRouter()
const store = useStore()

const emit = defineEmits(['report-updated', 'report-deleted'])
const report = computed(() => store.state.report)
const patientReport = ref([])
const formData = ref({
  report_date: '',
  file_path: '',
  patient: null,
  author_id : currentUser.ID,
  files: null,
})

// Create refs for the selected doctor and address
const selectedDoctor = ref(props.patient.Doctors[0] || {})
const selectedAddress = ref(selectedDoctor.value.Addresses[0] || {})

// Watch for changes in the selected doctor and update the selected address
watch(selectedDoctor, (newDoctor) => {
  selectedAddress.value = newDoctor.Addresses[0]
})

// Watch for when the report is fetched and update formData accordingly
watch(report, (newReport) => {
  if (newReport) {
    formData.value.report_date = newReport.report_date
    formData.value.file_path = newReport.file_path
    formData.value.patient = newReport.patient_id
  }
})
watch(() => props.report_id, async (newReportId) => {
  if (newReportId) {
    console.log('Report ID changed:', newReportId)
    await fetchReport(newReportId)
  }
})
const fetchReport = async (reportId) => {
  try {
    const response = await axios.get(`/api/reports/${reportId}/`)
    patientReport.value = response.data.report
    store.commit('setReport', response.data.report);
  } catch (error) {
    console.error('Error fetching report:', error)
    toast.error('Error fetching report',{
        timeout: 2000
    })
  }
}

onMounted(async () => {
  await fetchReport(props.report_id)
})

const handleSubmit = async () => {
  try {
    const uploadData = new FormData()
    uploadData.append('patient_id', formData.value.patient)
    uploadData.append('report_date', formData.value.report_date)

    // Check if new files are selected
    if (formData.value.files && formData.value.files.length > 0) {
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

      // Append the merged PDF Blob to FormData
      uploadData.append('file', mergedPdfBlob, 'merged_report.pdf')
    }

    // Send the FormData with or without the merged PDF to the backend
    const response = await axios.put(`/api/reports/${props.report_id}/`, uploadData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    // Update the report in the Vuex store
    store.commit('setReport', response.data.report)
    console.log('Report updated:', response.data.report)
    toast.success('Report updated successfully', { timeout: 2000 })
    emit('report-updated')
  } catch (error) {
    console.error('Error updating report:', error)
    toast.error('Error updating report', { timeout: 2000 })
  }
}

// Handle file upload
const handleFileUpload = (event) => {
  formData.value.files = event.target.files
}

function getFileLink(filePath) {
      // console.log('File path:', filePath)
      return `https://dev.nuttynarwhal.com/${filePath}`
}


// Create a PDF from the form data
const createPDF = async () => {
  try {
    // Fetch the 'report.pdf' template
    // const templateResponse = await axios.get('/report.pdf', { responseType: 'arraybuffer' })
    // const templateBytes = new Uint8Array(templateResponse.data)
    // Load the template into a PDFDocument
    // const pdfDoc = await PDFDocument.load(templateBytes)

    const pdfDoc = await PDFDocument.create()
    const page = pdfDoc.addPage()

    // Get the form fields
    const { report_date, file_path, patient, files } = formData.value

    // Add the form fields to the PDF
    // const page = pdfDoc.getPages()[0]
    const { width, height } = page.getSize()
    const font = await pdfDoc.embedFont(StandardFonts.Helvetica)
    const color = rgb(0, 0, 0)
    page.drawText(report_date, { x: 50, y: height - 70, size: 50, font, color })
    page.drawText(`Doctor: ${selectedDoctor.value.name}`, { x: 50, y: height - 100, size: 50, font, color })
    page.drawText(`Address: ${selectedAddress.value.street}`, { x: 50, y: height - 200, size: 50, font, color })

    // Serialize the PDFDocument to bytes
    const pdfBytes = await pdfDoc.save()

    // Create a Blob from the bytes
    const blob = new Blob([pdfBytes], { type: 'application/pdf' })

    // Create a URL for the Blob
    const url = URL.createObjectURL(blob)

    // Return the URL
    window.open(url, '_blank')
    return url
  } catch (error) {
    console.error('Error creating PDF:', error)
    toast.error('Error creating PDF', {
      timeout: 2000
    })
  }
}


// Deleting report
function confirmDelete() {
      if (window.confirm('Do you want to delete this report?')) {
        deleteReport()
      }
    }

const deleteReport = async () => {
  try {
    await axios.delete(`/api/reports/${props.report_id}/`)
    await store.dispatch('deleteReport', props.report_id)
    toast.success('Report deleted', { timeout: 2000 })
    emit('report-deleted', props.report_id)
    router.push(`/patients/${props.patient.ID}`)
  } catch (error) {
    console.error('Error deleting report:', error)
    toast.error('Error deleting report', { timeout: 2000 })
  }
}
</script>

<template>
    <div class="row">
    <div class="col">
      <!-- Doctor selection -->
      <select class="form-select" v-model="selectedDoctor">
        <option v-for="doctor in props.patient.Doctors" :key="doctor.id" :value="doctor">
          {{ doctor.name }}
        </option>
      </select>
    </div>
    <div class="col">
      <!-- Address selection -->
      <select class="form-select" v-model="selectedAddress">
        <option v-for="address in selectedDoctor.Addresses" :key="address.id" :value="address">
          {{ address.street }} {{ address.city }} {{ address.state }} {{ address.zip }}
        </option>
      </select>
    </div>
    <div class="col">
      <button type="button" class="btn btn-primary" @click="createPDF">Create PDF</button>
    </div>
  </div>
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
          <label for="file_path" class="form-label">Files:
                <a v-if = "formData.file_path"
                  :href="getFileLink(formData.file_path)" 
                  target="_blank" 
                  rel="noopener noreferrer">
                  Open Report
                </a>
          </label>
          <input
            class="form-control"
            type="file"
            id="file_path"
            multiple
            @change="handleFileUpload"
          />
        </div>
        <div class="mb-3">
          <div class="btn-group">
            <button class="btn btn-primary" type="submit">Update</button>
            <button class="btn btn-danger" type="button" @click="confirmDelete">Delete</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>
