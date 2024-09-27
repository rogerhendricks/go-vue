<script setup>

import { ref, onMounted, computed } from 'vue'
import { useStore } from 'vuex'
import axios from '../../../axiosConfig'
import { useRoute } from 'vue-router'
import {useToast} from 'vue-toastification';

const toast = useToast();

const route = useRoute()
const store = useStore()
const doctorsFromStore = computed(() => store.state.doctors.doctors || [])
const leadsFromStore = computed(() => store.state.leads.leads || [])
const patient = computed(() => store.state.patient.patient || {})
const loading = ref(false)
const newLead = ref({ 
  implant_date: '',
  explant_date: '',
  lead_id: '',
  serial: '',
  is_active:"true",
  patient_id: route.params.id,
  doctor_id: '', 
})
const editedLead = ref({
  ID: '',
  implant_date: '',
  explant_date: '',
  is_active: '',
  lead_id: '',
  serial: '',
  doctor_id: ''
});

const implantedLeads = ref([])
onMounted(async () => {
    getImplantedLeads()
    if (!leadsFromStore.value) {
        await store.dispatch('fetchLeads')
    }
})

async function getImplantedLeads() {
        loading.value = true
        const patientId = route.params.id
        try {
         const response = await axios.get(`/api/${patientId}/implantedLeads`)
         implantedLeads.value = response.data.implantedLeads
            // console.log('Implanted Leads:', implantedLeads.value)
        } catch (error) {
            console.error('Error fetching implanted leads:', error)
        } finally {
            loading.value = false
        }
}

async function createLead() {
    const patientId = Number(route.params.id)
    try {
      const response = await axios.post(`/api/${patientId}/implantedLeads`, {
            implant_date: newLead.value.implant_date,
            explant_date: newLead.value.explant_date || null, // Send null for optional fields
            is_active: newLead.value.is_active,
            lead_id: newLead.value.lead_id,
            serial: newLead.value.serial,
            patient_id: patientId, // Send patient_id as a number
            doctor_id: newLead.value.doctor_id
        })
        implantedLeads.value.push(response.data.implantedLead)
        console.log('New implanted leads:', implantedLeads.value)
        newLead.value = { 
            implant_date: '',
            explant_date: '',
            lead_id: '',
            serial: '',
            is_active:"true",
            patient_id: patientId,
            doctor_id: '', 
        }
        hideModal('createLeadModal')
        toast.success('Lead created successfully',{
            timeout: 2000,
        });
    } catch (error) {
        console.error('Error creating lead:', error)
        toast.error('Failed to create lead',{
            timeout: 2000,
        });
    }
}

function openEditModal(lead) {
    editedLead.value = { ...lead}
    console.log('Edited Lead:', editedLead.value)
    showModal('editLeadModal')
}

async function updateLead() {
    const patientId = Number(route.params.id)
    try {
        const response = await axios.put(`/api/implantedLeads/${editedLead.value.ID}`, editedLead.value)
        const editedLeadId = Number(editedLead.value.ID);
        const index = implantedLeads.value.findIndex(lead => lead.ID === editedLeadId)
        if (index !== -1) {
            implantedLeads.value[index] = response.data.implantedLead;
        } else {
            console.error('Lead not found:', editedLeadId)
        }
        hideModal('editLeadModal')
        toast.success('Lead updated successfully',{
            timeout: 2000,
        });
    } catch (error) {
        toast.error('Failed to update lead',{
            timeout: 2000,
        });
        console.error('Error updating lead:', error)
    }
}

function confirmDelete(leadId) {
      if (window.confirm('Do you want to delete this lead?')) {
        deleteLead(leadId)
      }
    }

async function deleteLead(leadId) {
    try {
        await axios.delete(`/api/implantedLeads/${leadId}`)
        implantedLeads.value = implantedLeads.value.filter(implantedLead => implantedLead.ID !== leadId)
        toast.success('Lead deleted successfully',{
            timeout: 2000,
        });
    } catch (error) {
        toast.error('Error deleting lead',{
            timeout: 2000,
        });
        console.error('Error deleting lead:', error)
    }
}

function showModal(modalId) {
      const modal = new bootstrap.Modal(document.getElementById(modalId))
      modal.show()
  }

function hideModal(modalId) {
    const modalEl = document.getElementById(modalId)
    const modal = bootstrap.Modal.getInstance(modalEl)
    modal.hide()
}
</script>
<template>
    <div class="p-2">
        <div class="row">
            <div class="col-md-2">
                <div class="d-grid d-md-flex justify-content-md-start"> 
                    <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#createLeadModal">Add Lead</button>
                </div>
            </div>
            <div v-if="loading" class="col">Loading...</div>
            <div class="col px-4" v-else>
                <div v-for="lead in implantedLeads" :key="lead.id">
                    <div class="row d-flex flex-wrap text-bg-secondary p-3 rounded-2 mb-3"> 
                        <div class="col">
                            <b>Implant Date:</b> {{ lead.implant_date }}
                        </div>
                        <div class="col">
                            <b>Serial:</b> {{ lead.serial }}
                        </div>
                        <div class="col">
                            <b>Lead:</b> {{ lead.Lead.manufacturer }} {{ lead.Lead.name }}
                        </div>
                        <div class="col">
                            <b>Active:</b> {{ lead.is_active}}
                        </div>
                        <div class="col">
                            <b>Implanter:</b> {{ lead.Doctor.name }}
                        </div>
                        <div class="col">
                            <b>Explant Date:</b> {{ lead.explant_date }}
                        </div>
                        <div class="col">
                            <div class="btn-group">
                                <button @click="openEditModal(lead)" class="btn btn-sm btn-primary">Edit</button>
                                <button @click="confirmDelete(lead.ID)" class="btn btn-sm btn-danger">Delete</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>


        <!-- Create Lead Modal -->
        <div class="modal fade" id="createLeadModal" tabindex="-1" aria-labelledby="createLeadModalLabel" aria-hidden="true">
            <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                <h5 class="modal-title" id="createLeadModalLabel">Create Lead</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                <form @submit.prevent="createLead">
                    <div class="row mb-3">
                    <div class="col">
                        <label for="implantDate" class="form-label">Implant Date</label>
                        <input type="date" v-model="newLead.implant_date" class="form-control" id="implantDate" required>
                    </div>
                    <div class="col">
                        <label for="serial" class="form-label">Serial</label>
                        <input type="text" v-model="newLead.serial" class="form-control" id="serial" required>
                    </div>
                    <div class="col">
                        <label for="explantDate" class="form-label">Explant Date</label>
                        <input type="date" v-model="newLead.explant_date" class="form-control border border-danger" id="explantDate">
                    </div>
                    <div class="col">
                        <label for="editLeadIsActive" class="form-label">Active</label>
                        <select v-model="editedLead.is_active" class="form-select" id="editLeadIsActive" required>
                        <option value="true">Yes</option>
                        <option value="false">No</option>
                        </select>
                    </div>
                    <div class="col">
                        <label for="lead" class="form-label">Lead</label>
                        <select v-model="newLead.lead_id" class="form-select" id="lead" required>
                        <option v-for="lead in leadsFromStore" :key="lead.ID" :value="lead.ID">{{ lead.name }}</option>
                        </select>
                    </div>
                    <div class="col">
                        <label for="doctor" class="form-label">Doctor</label>
                        <select v-model="newLead.doctor_id" class="form-select" id="doctor">
                        <option v-for="doctor in doctorsFromStore" :key="doctor.ID" :value="doctor.ID">{{ doctor.name }}</option>
                        </select>
                    </div>
                    </div>
                    <button type="submit" class="btn btn-primary">Save</button>
                </form>
                </div>
            </div>
            </div>
        </div>

    <!-- Edit Lead Modal -->
    <div class="modal fade" id="editLeadModal" tabindex="-1" aria-labelledby="editLeadModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="editLeadModalLabel">Edit Lead</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateLead">
              <div class="row mb-3">
                  <div class="col">
                    <label for="editLeadImplantDate" class="form-label">Implant Date</label>
                    <input type="date" v-model="editedLead.implant_date" class="form-control" id="editLeadImplantDate" required>
                  </div>
                  <div class="col">
                    <label for="editLeadExplantDate" class="form-label">Explant Date</label>
                    <input type="date" v-model="editedLead.explant_date" class="form-control border border-danger" id="editLeadExplantDate">
                  </div>
                    <div class="col">
                        <label for="editLeadSerial" class="form-label">Serial</label>
                        <input type="text" v-model="editedLead.serial" class="form-control" id="editLeadSerial" required>
                    </div>
                  <div class="col">
                    <label for="editLeadIsActive" class="form-label">Active</label>
                    <select v-model="editedLead.is_active" class="form-select" id="editLeadIsActive" required>
                      <option value="Yes">Yes</option>
                      <option value="No">No</option>
                    </select>
                  </div>
                  <div class="col">
                    <label for="editLeadLead" class="form-label">Lead</label>
                    <select v-model="editedLead.lead_id" class="form-select" id="editLeadLead" required>
                      <option v-for="lead in leadsFromStore" :key="lead.ID" :value="lead.ID">{{ lead.name }}</option>
                    </select>
                  </div>
                  <div class="col">
                    <label for="editLeadDoctor" class="form-label">Doctor</label>
                    <select v-model="editedLead.doctor_id" class="form-select" id="editLeadDoctor">
                      <option v-for="doctor in doctorsFromStore" :key="doctor.ID" :value="doctor.ID">{{ doctor.name }}</option>
                    </select>
                  </div>
              </div>
              <button type="submit" class="btn btn-primary">Update</button>
            </form>
          </div>
        </div>
      </div>
    </div>
</template>