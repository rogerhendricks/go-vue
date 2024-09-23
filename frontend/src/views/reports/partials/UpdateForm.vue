<script setup>
import { ref, onMounted, computed, watch, defineProps } from 'vue'
import axios from '../../../axiosConfig'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'
import { PDFDocument, StandardFonts, rgb } from 'pdf-lib'
import { useToast } from 'vue-toastification'

const props = defineProps(['patient', 'report_id'])

const toast = useToast()
const route = useRoute()
const router = useRouter()
const store = useStore()
const report = computed(() => store.state.report)
const patientReport = ref({})
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

// function goBackUsingBack() {
//       if (router) {
//         router.back();
//       }
// }

function getFileLink(filePath) {
      // console.log('File path:', filePath)
      return `https://dev.nuttynarwhal.com/${filePath}`
}

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
    // Add more fields as needed...

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


</script>

<template>
  <button type="button" class="btn btn-primary" @click="createPDF">Create PDF</button>
  <!-- <button class="btn" type="button" @click="goBackUsingBack">Go Back </button> -->
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
                            <a 
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
          <button class="btn btn-primary" type="submit">Submit</button>
        </div>
      </form>
    </div>
  </div>
</template>
