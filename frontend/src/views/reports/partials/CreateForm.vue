<script setup>

import { ref, onMounted, computed, watch, defineProps, toRefs, defineEmits } from 'vue'
import axios from '../../../axiosConfig'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import {useToast} from 'vue-toastification';
import { PDFDocument, StandardFonts, rgb } from 'pdf-lib'
import fileImport from '../../../utils/fileImport';


const store = useStore()
const deviceFromStore = computed(() => store.state.device || {})
console.log('Device from store:', deviceFromStore.value)
const currentUser = JSON.parse(localStorage.getItem('user') || 'null')
const toast = useToast();
const props = defineProps(['patient'])
const emit = defineEmits(['report-created'])

const importFile = fileImport

// Create refs for the selected doctor and address
const selectedDoctor = ref(props.patient.Doctors[0])
const selectedAddress = ref(selectedDoctor.value.Addresses[0])

// Watch for changes in the selected doctor and update the selected address
watch(selectedDoctor, (newDoctor) => {
  selectedAddress.value = newDoctor.Addresses[0]
})

const initialFormData = ref({
      report_date: '',
      report_type: '',
      current_dependency: '',
      current_heart_rate: '',
      author_id : currentUser.ID,
      file_path: '',
      patient_id: props.patient.ID,
      files: [],
    })

const formData = ref({ ...initialFormData.value })
const files = ref([])

const handleFileUpload = (e) => {
      formData.value.files = e.target.files
    }

const mergePDFs = async (pdfFiles) => {
  const pdfDoc = await PDFDocument.create()
  
  for (const file of pdfFiles) {
    const fileData = await file.arrayBuffer()
    const doc = await PDFDocument.load(fileData)
    const copiedPages = await pdfDoc.copyPages(doc, doc.getPageIndices())
    copiedPages.forEach(page => pdfDoc.addPage(page))
  }

  const mergedPdfBytes = await pdfDoc.save()
  return new Blob([mergedPdfBytes], { type: 'application/pdf' })
}


const handleSubmit = async () => {
      try {
        const form = new FormData()
        form.append('author_id', currentUser.ID)
        form.append('report_date', formData.value.report_date)
        form.append('patient_id', formData.value.patient_id)
        form.append('report_type', formData.value.report_type)
        form.append('current_dependency', formData.value.current_dependency)
        form.append('current_heart_rate', formData.value.current_heart_rate)
        
        if (formData.value.files && formData.value.files.length > 0){
          const mergedPdfBlob = await mergePDFs(files.value)
          form.append('file', mergedPdfBlob, 'merged_report.pdf')
        } else {
          // If no files are selected, this will still send the form data without a file
          console.log("No files selected, submitting form without file.")
        }

        const response = await axios.post(`/api/${props.patient.ID}/reports`, form, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        console.log('Report created:', response.data)
        toast.success('Report created successfully',{
            timeout: 2000,
        })
        emit('report-created')
        formData.value = { ...initialFormData };
        files.value = []
      } catch (error) {
        console.error('Error creating report:', error)
        toast.error('Error creating report',{
            timeout: 2000,
        })
      }
    }

function handleInput(event) {
        formData.value[event.target.name] = event.target.value
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
    page.drawText(`Doctor: ${selectedDoctor.value.name}`, { x: 50, y: height - 100, size: 50, font, color })
    page.drawText(`Address: ${selectedAddress.value.street}`, { x: 50, y: height - 200, size: 50, font, color })
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

const handleFileChange = async (event) => {
    const deviceSerial = deviceFromStore.value.serial
    const target = event.target;
    const file = target.files[0];
try{
    const data = await fileImport(file, deviceSerial)
    console.log('Parsed dataResult:', data);
    for (const key in data) {
      // formData.value[key] = data[key]
      console.log('Key:', key, 'Value:', data[key]);
      // if (formData.value.hasOwnProperty(key)) {
      //   formData.value[key] = dataResult[key];
      // }
    }
    return data;
  } catch (error) {
    console.error('Error importing file:', error)
    toast.error('Error importing file', {
      timeout: 2000
    })
  }
}
</script>
<template>
  <div class="row">
    <div class="col">
      <!-- <button type="button" class="btn btn-secondary" @click="importFile">Import</button> -->
      <input type="file" @change="handleFileChange($event)" class="form-control">
    </div>
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

      <form @submit.prevent="handleSubmit">
        <div class="mb-3">
            <label for="report_date" class=""form-label>Report Date:</label>
            <input id="report_date" class="form-control" v-model="formData.report_date" type="date" required input="handleInput">
        </div>
        <div class="mb-3">
          <label for="report_type" class="form-label">Report Type:</label>
          <select class="form-select" id="report_type" v-model="formData.report_type" required @change="handleInput">
            <option value="in_clinic">In Clinic</option>
            <option value="hospital">Hospital</option>
            <option value="remote">Remote</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="current_dependency" class="form-label">Dependent</label>
          <select class="form-select" v-model="formData.current_dependency" name="current_dependency" id="current_dependency">
            <option value="dependent">Dependent</option>
            <option value="non_dependent">Non Dependent</option>
            <option value="unknown">Unkown</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="current_heart_rate" class="form-label">Heart Rate</label>
          <input type="number" class="form-control" v-model="formData.current_heart_rate" name="current_heart_rate" id="current_heart_rate">
        </div>
        <div class="mb-3">
          <label for="file_path" class="form-label">Files:</label>
          <input class="form-control" type="file" id="file_path"  multiple @change="handleFileUpload">
        </div>
        <div class="mb-3">
          <button class="btn btn-primary" type="submit">Submit</button>
        </div>
    </form>

</template>