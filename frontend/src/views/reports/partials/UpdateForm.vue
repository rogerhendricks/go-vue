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
  report_type: "",
  report_status: "",
  current_heart_rate: 0,
  current_rhythm: "",
  current_dependency: "",
  mdc_idc__stat_ataf_burden_percent: 0,
  mdc_idc_set_brady_mode: '',
  mdc_idc_set_brady_lowrate: 0,
  mdc_idc_set_brady_max_tracking_rate: 0,
  mdc_idc_set_brady_max_sensor_rate: 0,
  mdc_idc_dev_sav: "",
  mdc_idc_dev_pav: "",
  mdc_idc_stat_brady_ra_percent_paced: 0,
  mdc_idc_stat_brady_rv_percent_paced: 0,
  mdc_idc_stat_brady_lv_percent_paced: 0,
  mdc_idc_stat_brady_biv_percent_paced: 0,
  mdc_idc_batt_volt: 0,
  mdc_idc_batt_remaining: 0,
  mdc_idc_batt_status: "",
  mdc_idc_cap_charge_time: 0,
  mdc_idc_msmt_ra_impedance_mean: 0,
  mdc_idc_msmt_ra_sensing_mean: 0,
  mdc_idc_msmt_ra_threshold: 0,
  mdc_idc_msmt_ra_pw: 0,
  mdc_idc_msmt_rv_impedance_mean: 0,
  mdc_idc_msmt_rv_sensing_mean: 0,
  mdc_idc_msmt_rv_threshold: 0,
  mdc_idc_msmt_rv_pw: 0,
  mdc_idc_msmt_shock_impedance: 0,
  mdc_idc_msmt_lv_impedance_mean: 0,
  mdc_idc_msmt_lv_threshold: 0,
  mdc_idc_msmt_lv_pw: 0,
  comments: "",
  is_completed: false,
  author_id : currentUser.ID,
  file_path: '',
  patient: null,
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
    formData.value.report_type = newReport.report_type
    formData.value.report_status = newReport.report_status
    formData.value.current_heart_rate = newReport.current_heart_rate
    formData.value.current_rhythm = newReport.current_rhythm
    formData.value.current_dependency = newReport.current_dependency
    formData.value.mdc_idc__stat_ataf_burden_percent = newReport.mdc_idc__stat_ataf_burden_percent
    formData.value.mdc_idc_set_brady_mode = newReport.mdc_idc_set_brady_mode
    formData.value.mdc_idc_set_brady_lowrate = newReport.mdc_idc_set_brady_lowrate
    formData.value.mdc_idc_set_brady_max_tracking_rate = newReport.mdc_idc_set_brady_max_tracking_rate
    formData.value.mdc_idc_set_brady_max_sensor_rate = newReport.mdc_idc_set_brady_max_sensor_rate
    formData.value.mdc_idc_dev_sav = newReport.mdc_idc_dev_sav
    formData.value.mdc_idc_dev_pav = newReport.mdc_idc_dev_pav
    formData.value.mdc_idc_stat_brady_ra_percent_paced = newReport.mdc_idc_stat_brady_ra_percent_paced
    formData.value.mdc_idc_stat_brady_rv_percent_paced = newReport.mdc_idc_stat_brady_rv_percent_paced
    formData.value.mdc_idc_stat_brady_lv_percent_paced = newReport.mdc_idc_stat_brady_lv_percent_paced
    formData.value.mdc_idc_stat_brady_biv_percent_paced = newReport.mdc_idc_stat_brady_biv_percent_paced
    formData.value.mdc_idc_batt_volt = newReport.mdc_idc_batt_volt
    formData.value.mdc_idc_batt_remaining = newReport.mdc_idc_batt_remaining
    formData.value.mdc_idc_batt_status = newReport.mdc_idc_batt_status
    formData.value.mdc_idc_cap_charge_time = newReport.mdc_idc_cap_charge_time
    formData.value.mdc_idc_msmt_ra_impedance_mean = newReport.mdc_idc_msmt_ra_impedance_mean
    formData.value.mdc_idc_msmt_ra_sensing_mean = newReport.mdc_idc_msmt_ra_sensing_mean
    formData.value.mdc_idc_msmt_ra_threshold = newReport.mdc_idc_msmt_ra_threshold
    formData.value.mdc_idc_msmt_ra_pw = newReport.mdc_idc_msmt_ra_pw
    formData.value.mdc_idc_msmt_rv_impedance_mean = newReport.mdc_idc_msmt_rv_impedance_mean
    formData.value.mdc_idc_msmt_rv_sensing_mean = newReport.mdc_idc_msmt_rv_sensing_mean
    formData.value.mdc_idc_msmt_rv_threshold = newReport.mdc_idc_msmt_rv_threshold
    formData.value.mdc_idc_msmt_rv_pw = newReport.mdc_idc_msmt_rv_pw
    formData.value.mdc_idc_msmt_shock_impedance = newReport.mdc_idc_msmt_shock_impedance
    formData.value.mdc_idc_msmt_lv_impedance_mean = newReport.mdc_idc_msmt_lv_impedance_mean
    formData.value.mdc_idc_msmt_lv_threshold = newReport.mdc_idc_msmt_lv_threshold
    formData.value.mdc_idc_msmt_lv_pw = newReport.mdc_idc_msmt_lv_pw
    formData.value.comments = newReport.comments
    formData.value.is_completed = newReport.is_completed
    formData.value.patient = newReport.patient_id
    formData.value.file_path = newReport.file_path
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
    uploadData.append('report_type', formData.value.report_type)
    uploadData.append('report_status', formData.value.report_status)
    uploadData.append('current_heart_rate', formData.value.current_heart_rate)
    uploadData.append('current_rhythm', formData.value.current_rhythm)
    uploadData.append('current_dependency', formData.value.current_dependency)
    uploadData.append('mdc_idc__stat_ataf_burden_percent', formData.value.mdc_idc__stat_ataf_burden_percent)
    uploadData.append('mdc_idc_set_brady_mode', formData.value.mdc_idc_set_brady_mode)
    uploadData.append('mdc_idc_set_brady_lowrate', formData.value.mdc_idc_set_brady_lowrate)
    uploadData.append('mdc_idc_set_brady_max_tracking_rate', formData.value.mdc_idc_set_brady_max_tracking_rate)
    uploadData.append('mdc_idc_set_brady_max_sensor_rate', formData.value.mdc_idc_set_brady_max_sensor_rate)
    uploadData.append('mdc_idc_dev_sav', formData.value.mdc_idc_dev_sav)
    uploadData.append('mdc_idc_dev_pav', formData.value.mdc_idc_dev_pav)
    uploadData.append('mdc_idc_stat_brady_ra_percent_paced', formData.value.mdc_idc_stat_brady_ra_percent_paced)
    uploadData.append('mdc_idc_stat_brady_rv_percent_paced', formData.value.mdc_idc_stat_brady_rv_percent_paced)
    uploadData.append('mdc_idc_stat_brady_lv_percent_paced', formData.value.mdc_idc_stat_brady_lv_percent_paced)
    uploadData.append('mdc_idc_stat_brady_biv_percent_paced', formData.value.mdc_idc_stat_brady_biv_percent_paced)
    uploadData.append('mdc_idc_batt_volt', formData.value.mdc_idc_batt_volt)
    uploadData.append('mdc_idc_batt_remaining', formData.value.mdc_idc_batt_remaining)
    uploadData.append('mdc_idc_batt_status', formData.value.mdc_idc_batt_status)
    uploadData.append('mdc_idc_cap_charge_time', formData.value.mdc_idc_cap_charge_time)
    uploadData.append('mdc_idc_msmt_ra_impedance_mean', formData.value.mdc_idc_msmt_ra_impedance_mean)
    uploadData.append('mdc_idc_msmt_ra_sensing_mean', formData.value.mdc_idc_msmt_ra_sensing_mean)
    uploadData.append('mdc_idc_msmt_ra_threshold', formData.value.mdc_idc_msmt_ra_threshold)
    uploadData.append('mdc_idc_msmt_ra_pw', formData.value.mdc_idc_msmt_ra_pw)
    uploadData.append('mdc_idc_msmt_rv_impedance_mean', formData.value.mdc_idc_msmt_rv_impedance_mean)
    uploadData.append('mdc_idc_msmt_rv_sensing_mean', formData.value.mdc_idc_msmt_rv_sensing_mean)
    uploadData.append('mdc_idc_msmt_rv_threshold', formData.value.mdc_idc_msmt_rv_threshold)
    uploadData.append('mdc_idc_msmt_rv_pw', formData.value.mdc_idc_msmt_rv_pw)
    uploadData.append('mdc_idc_msmt_shock_impedance', formData.value.mdc_idc_msmt_shock_impedance)
    uploadData.append('mdc_idc_msmt_lv_impedance_mean', formData.value.mdc_idc_msmt_lv_impedance_mean)
    uploadData.append('mdc_idc_msmt_lv_threshold', formData.value.mdc_idc_msmt_lv_threshold)
    uploadData.append('mdc_idc_msmt_lv_pw', formData.value.mdc_idc_msmt_lv_pw)
    uploadData.append('comments', formData.value.comments)
    uploadData.append('is_completed', formData.value.is_completed)
    uploadData.append('author_id', formData.value.author_id)
    
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
    <div class="row mb-3">
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
        <!-- Report Data -->
        <div class="row border border-warning rounded mb-3">
            <div class="col p-2">
                <label for="report_date" class="form-label">Report Date:</label>
                <input
                    id="report_date"
                    class="form-control"
                    v-model="formData.report_date"
                    type="date"
                    required
                    @change="handleInput"
                />
            </div>
            <div class="col p-2">
                <label for="report_type" class="form-label">Report Type:</label>
                <select
                    class="form-select"
                    id="report_type"
                    v-model="formData.report_type"
                    required
                >
                    <option value="in_clinic">In Clinic</option>
                    <option value="hospital">Hospital</option>
                    <option value="remote">Remote</option>
                </select>
            </div>
            <div class="col p-2">
                <label for="report_status" class="form-label">Report Status:</label>
                <select
                    id="report_status"
                    class="form-select"
                    v-model="formData.report_status">
                    <option value="pending">Pending</option>
                    <option value="completed">Completed</option>
                    <option value="cancelled">Cancelled</option>
                    <option value="unknown">Unknown</option>
                </select>
            </div>
        </div>
        <!-- Patient Substrate -->
        <div class="row border border-danger-subtle rounded mb-3">
            <div class="col p-2">
                <label for="current_rhythm" class="form-label">Current Rhythm:</label>
                <input
                    type="text"
                    class="form-control"
                    v-model="formData.current_rhythm"
                    id="current_rhythm"
                />
            </div>
            <div class="col p-2">
                <label for="current_dependency" class="form-label">Dependent</label>
                <select
                    class="form-select"
                    v-model="formData.current_dependency"
                    name="current_dependency"
                    id="current_dependency"
                >
                    <option value="dependent">Dependent</option>
                    <option value="non_dependent">Non Dependent</option>
                    <option value="unknown">Unkown</option>
                </select>
            </div>
            <div class="col p-2">
                <label for="current_heart_rate" class="form-label">Heart Rate</label>
                <input
                    type="number"
                    class="form-control"
                    v-model="formData.current_heart_rate"
                    name="current_heart_rate"
                    id="current_heart_rate"
                />
            </div>
            <!-- Patient Arrhythmias -->
            <div class="col p-2">
                <label for="mdc_idc__stat_ataf_burden_percent" class="form-label">ATAF Burden Percent:</label>
                <input
                    type="number"
                    step="0.01"
                    class="form-control"
                    v-model="formData.mdc_idc__stat_ataf_burden_percent"
                    id="mdc_idc__stat_ataf_burden_percent"
                />
            </div>
        </div>
        <!-- Device Measurements -->
        <div class="row border border-success-subtle rounded mb-3">
            <div class="col">
                <div class="table-responsive">
                    <table class="table table-borderless">
                        <thead>
                            <tr>
                                <th></th>
                                <th>Impedance</th>
                                <th>Sensing</th>
                                <th>Threshold</th>
                                <th>Pulse Width</th>
                                <th>Paced %</th>
                                <!-- <th>Biv Paced</th> -->
                            </tr>
                        </thead>
                        <tbody class="table-group-divider">
                            <tr>
                                <th>RA</th>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_ra_impedance_mean"
                                        id="mdc_idc_msmt_ra_impedance_mean"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_ra_sensing_mean"
                                        id="mdc_idc_msmt_ra_sensing_mean"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_ra_threshold"
                                        id="mdc_idc_msmt_ra_threshold"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_ra_pw"
                                        id="mdc_idc_msmt_ra_pw"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_stat_brady_ra_percent_paced"
                                        id="mdc_idc_stat_brady_ra_percent_paced"
                                    />
                                </td>
                            </tr>
                            <tr>
                                <th>RV</th>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_rv_impedance_mean"
                                        id="mdc_idc_msmt_rv_impedance_mean"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_rv_sensing_mean"
                                        id="mdc_idc_msmt_rv_sensing_mean"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_rv_threshold"
                                        id="mdc_idc_msmt_rv_threshold"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_rv_pw"
                                        id="mdc_idc_msmt_rv_pw"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_stat_brady_rv_percent_paced"
                                        id="mdc_idc_stat_brady_rv_percent_paced"
                                    />
                                </td>
                            </tr>
                            <tr>
                                <th>LV</th>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_lv_impedance_mean"
                                        id="mdc_idc_msmt_lv_impedance_mean"
                                    />
                                </td>
                                <td></td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_lv_threshold"
                                        id="mdc_idc_msmt_lv_threshold"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_lv_pw"
                                        id="mdc_idc_msmt_lv_pw"
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_stat_brady_lv_percent_paced"
                                        id="mdc_idc_stat_brady_lv_percent_paced"
                                    />
                                </td>
                            </tr>
                            <tr>
                                <th>Shock</th>
                                <td colspan="4">
                                    <input
                                        type="number"
                                        step="0.01"
                                        class="form-control"
                                        v-model="formData.mdc_idc_msmt_shock_impedance"
                                        id="mdc_idc_msmt_shock_impedance"
                                    />
                                </td>
                                <td></td>
                            </tr>
                            <!-- <tr>
                                <th>BIV</th>
                                <td colspan="4"></td>
                                <td>
                                    <input
                                        type="number"
                                        class="form-control"
                                        v-model="formData.mdc_idc_stat_brady_biv_percent_paced"
                                        id="mdc_idc_stat_brady_biv_percent_paced"
                                    />
                                </td>
                            </tr> -->
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        <!-- Device Settings -->
        <div class="row border border-primary-subtle rounded mb-3">
            <div class="col p-2">
                <label for="mdc_idc_set_brady_mode" class="form-label"
                    >Brady Mode:</label
                >
                <input
                    type="text"
                    class="form-control"
                    v-model="formData.mdc_idc_set_brady_mode"
                    id="mdc_idc_set_brady_mode"
                />
            </div>
            <div class="col p-2">
                <label for="mdc_idc_set_brady_lowrate" class="form-label"
                    >Brady Lowrate:</label
                >
                <input
                    type="number"
                    class="form-control"
                    v-model="formData.mdc_idc_set_brady_lowrate"
                    id="mdc_idc_set_brady_lowrate"
                />
            </div>
            <div class="col p-2">
                <label
                    for="mdc_idc_set_brady_max_tracking_rate"
                    class="form-label"
                    >Max Tracking Rate:</label
                >
                <input
                    type="number"
                    class="form-control"
                    v-model="formData.mdc_idc_set_brady_max_tracking_rate"
                    id="mdc_idc_set_brady_max_tracking_rate"
                />
            </div>
            <div class="col p-2">
                <label
                    for="mdc_idc_set_brady_max_sensor_rate"
                    class="form-label"
                    >Max Sensor Rate:</label
                >
                <input
                    type="number"
                    class="form-control"
                    v-model="formData.mdc_idc_set_brady_max_sensor_rate"
                    id="mdc_idc_set_brady_max_sensor_rate"
                />
            </div>
        </div>
        <!-- Battery Status -->
        <div class="row border border-secondary-subtle rounded mb-3">
            <div class="col p-2">
                <label for="mdc_idc_batt_volt" class="form-label"
                    >Battery Voltage:</label
                >
                <input
                    type="number"
                    step="0.01"
                    class="form-control"
                    v-model="formData.mdc_idc_batt_volt"
                    id="mdc_idc_batt_volt"
                />
            </div>
            <div class="col p-2">
                <label for="mdc_idc_batt_remaining" class="form-label"
                    >Battery Remaining:</label
                >
                <input
                    type="number"
                    step="0.01"
                    class="form-control"
                    v-model="formData.mdc_idc_batt_remaining"
                    id="mdc_idc_batt_remaining"
                />
            </div>
            <div class="col p-2">
                <label for="mdc_idc_batt_status" class="form-label"
                    >Battery Status:</label
                >
                <input
                    type="text"
                    class="form-control"
                    v-model="formData.mdc_idc_batt_status"
                    id="mdc_idc_batt_status"
                />
            </div>
            <div class="col p-2">
                <label for="mdc_idc_cap_charge_time" class="form-label"
                    >Cap Charge Time:</label
                >
                <input
                    type="number"
                    step="0.001"
                    class="form-control"
                    v-model="formData.mdc_idc_cap_charge_time"
                    id="mdc_idc_cap_charge_time"
                />
            </div>
        </div>
        <!-- Comments -->
        <div class="mb-3">
            <label for="comments" class="form-label">Comments:</label>
            <textarea
                class="form-control"
                v-model="formData.comments"
                id="comments"
            ></textarea>
        </div>

        <!-- File upload -->
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
