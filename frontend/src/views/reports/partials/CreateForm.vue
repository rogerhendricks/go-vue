<script setup>

import { ref, onMounted, computed, watch, defineProps } from 'vue'
import axios from '../../../axiosConfig'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import {useToast} from 'vue-toastification';
import { PDFDocument } from 'pdf-lib'

const store = useStore()
const toast = useToast();
const props = defineProps(['patient'])

const initialFormData = ref({
      report_date: '',
      file_path: '',
      patient_id: props.patient.ID
    })

const formData = ref({ ...initialFormData.value })
const files = ref([])

const handleFileUpload = (e) => {
      files.value = Array.from(e.target.files)
      // formData.value.file_path = file
    }

  const mergePDFs = async (pdfFiles) => {
  const pdfDoc = await PDFDocument.create()
  
  for (const file of pdfFiles) {
    const fileData = await file.arrayBuffer()
    const doc = await PDFDocument.load(fileData)
    const [copiedPage] = await pdfDoc.copyPages(doc, doc.getPageIndices())
    copiedPage.forEach(page => pdfDoc.addPage(page))
  }

  const mergedPdfBytes = await pdfDoc.save()
  return new Blob([mergedPdfBytes], { type: 'application/pdf' })
}


const handleSubmit = async () => {
      try {
        const mergedPdfBlob = await mergePDFs(files.value)
        const form = new FormData()
        form.append('report_date', formData.value.report_date)
        form.append('patient_id', formData.value.patient_id)
        form.append('file', mergedPdfBlob, 'merged_report.pdf')

        const response = await axios.post(`/api/${props.patient.ID}/reports`, form, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        console.log('Report created:', response.data)
        toast.success('Report created successfully',{
            timeout: 2000,
        })
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

</script>
<template>
      <form @submit.prevent="handleSubmit">
        <div class="mb-3">
            <label for="report_date" class=""form-label>Report Date:</label>
            <input id="report_date" class="form-control" v-model="formData.report_date" type="date" required input="handleInput">
        </div>
        <div class="mb-3">
          <label for="file_path" class="form-label">Files:</label>
          <input class="form-control" type="file" id="file_path"  multiple required @change="handleFileUpload">
        </div>
        <div class="mb-3">
          <button class="btn btn-primary" type="submit">Submit</button>
        </div>
    </form>

</template>