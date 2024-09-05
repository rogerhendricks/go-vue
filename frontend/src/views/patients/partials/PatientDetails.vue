<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import EditForm from './EditForm.vue'
const store = useStore()
const patient = computed(() => store.state.patient.patient)
</script>
<template>
    <nav>
    <div class="nav nav-tabs" id="nav-tab" role="tablist">
        <button class="nav-link active" id="nav-patient-tab" data-bs-toggle="tab" data-bs-target="#nav-patient" type="button" role="tab" aria-controls="nav-patient" aria-selected="true">Patient</button>
        <button class="nav-link" id="nav-devices-tab" data-bs-toggle="tab" data-bs-target="#nav-devices" type="button" role="tab" aria-controls="nav-devices" aria-selected="false">Devices</button>
        <button class="nav-link" id="nav-leads-tab" data-bs-toggle="tab" data-bs-target="#nav-leads" type="button" role="tab" aria-controls="nav-leads" aria-selected="false">Leads</button>
        <button class="nav-link" id="nav-doctors-tab" data-bs-toggle="tab" data-bs-target="#nav-doctors" type="button" role="tab" aria-controls="nav-doctors" aria-selected="false" >Doctors</button>
        <button type="button" class="btn" data-bs-toggle="modal" data-bs-target="#editPatientModal">Edit</button>
    </div>
    </nav>
    <div class="tab-content text-bg-light border border-top-0" id="nav-tabContent">
        <div class="tab-pane fade show active" id="nav-patient" role="tabpanel" aria-labelledby="nav-patient-tab" tabindex="0">
            <div class="p-2">
                <div class="row">
                    <div class="col">
                        <p>Name: {{ patient.name }}</p>
                    </div>
                    <div class="col">
                        <p>DOB: {{ patient.dob }}</p>
                    </div>
                    <div class="col">
                        <p>Phone: {{ patient.phone }}</p>
                    </div>
                    <div class="col">
                        <p>Email: {{ patient.email }}</p>
                    </div>  
                </div>
                <div class="row">
                    <div class="col">
                        <p>Address: {{ patient.address }}</p>
                    </div>
                    <div class="col">
                        <p>City: {{ patient.city }}</p>
                    </div>
                    <div class="col">
                        <p>State: {{ patient.state }}</p>
                    </div>
                    <div class="col">
                        <p>Zip: {{ patient.zip }}</p>
                    </div>
                    <div class="col">
                        <p>Country: {{ patient.country }}</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="tab-pane fade" id="nav-devices" role="tabpanel" aria-labelledby="nav-devices-tab" tabindex="0">...</div>
        <div class="tab-pane fade" id="nav-leads" role="tabpanel" aria-labelledby="nav-leads-tab" tabindex="0">...</div>
        <div class="tab-pane fade" id="nav-doctors" role="tabpanel" aria-labelledby="nav-doctors-tab" tabindex="0">...</div>
    </div>

    <!-- Modal dialog for editing patient details -->
    <div class="modal fade" id="editPatientModal" tabindex="-1" aria-labelledby="editPatientModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="editPatientModalLabel">Edit Patient</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <EditForm :patient="patient"/>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                    <button type="button" class="btn btn-primary">Save changes</button>
                </div>
            </div>
        </div>
    </div>

</template>