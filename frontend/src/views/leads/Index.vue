<script setup>
    import axios from '../../axiosConfig'
    import { ref, computed } from 'vue'
    import CreateForm from './partials/CreateForm.vue'
    import EditForm from './partials/EditForm.vue'

    const leads = ref([])
    const loading = ref(false)
    const selectedLead = ref(null)
    let searchTerm = ref('');
    let currentPage = ref(1);
    let pageSize = ref(10);

    const filteredLeads = computed(() => {
        if (!searchTerm.value) {
            return leads.value;
        }
        return leads.value.filter(lead =>
            lead.manufacturer.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
            lead.name.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
            lead.leadModel.toLowerCase().includes(searchTerm.value.toLowerCase())
        );
    });

    const paginatedLeads =computed(() => {
        const start = (currentPage.value - 1) * pageSize.value;
        const end = currentPage.value * pageSize.value;
        return filteredLeads.value.slice(start,end);
    });

    async function getLeads() {
        loading.value = true
        try {
            const response = await axios.get('/api/leads')
            leads.value = response.data.leads
            console.log('Leads fetched successfully:', leads.value)
        } catch (error) {
            console.error('Error fetching leads:', error)
        } finally {
            loading.value = false
        }
    }
    getLeads()
</script>
<template>
    <div class="row">
        <h1>Leads</h1>
    </div>
    <div class="row">
        <div class="col-md-3">
            <CreateForm />
        </div>
        <div class="col-md-9 mt-4">
            <div class="input-group mb-3">
                <input  type="text" class="form-control" placeholder="Search" v-model="searchTerm">
                <button class="btn btn-outline-secondary" type="button" id="button-addon2">Search</button>
            </div>
            <!-- <EditForm v-if="selectedLead" :lead="selectedLead" @cancel="selectedLead = null" /> -->

            <div class="my-2">
                <button class="btn btn-primary" @click="getLeads">Refresh</button>
            </div>
            <ol>
                <li v-for="lead in paginatedLeads">
                    <button type="button" class="btn" @click="selectedLead = lead" data-bs-toggle="modal" data-bs-target="#staticBackdrop">
                        {{ lead.manufacturer }} {{ lead.name }}
                    </button>
                    <!-- <button class="btn btn-danger btn-sm">Delete</button> -->
                </li>
            </ol>

            <nav aria-label="Page navigation example">
                <ul class="pagination
                justify-content-center">
                    <li class="page-item" :class="{ disabled
                    : currentPage === 1 }">
                        <a class="page-link" href="#" @click="currentPage = 1">First</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === 1 }">
                        <a class="page-link" href="#" @click="currentPage--">Previous</a>
                    </li>
                    <li class="page-item" v-for="page in Math.ceil(filteredLeads.length / pageSize)" :key="page" :class="{ active
                    : currentPage === page }">
                        <a class="page-link" href="#" @click="currentPage = page">{{ page }}</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === Math.ceil(filteredLeads.length / pageSize) }">
                        <a class="page-link" href="#" @click="currentPage++">Next</a>
                    </li>
                    <li class="page-item" :class="{ disabled
                    : currentPage === Math.ceil(filteredLeads.length / pageSize) }">
                        <a class="page-link" href="#" @click="currentPage = Math.ceil(filteredLeads.length / pageSize)">Last</a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>


    <!-- Modal -->
    <div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
        <div class="modal-header">
            <h1 class="modal-title fs-5" id="staticBackdropLabel">Edit Lead</h1>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
            <EditForm v-if="selectedLead" :lead="selectedLead" :key="selectedLead.ID" @cancel="selectedLead = null" />
        </div>
        <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="selectedLead = null" >Close</button>
        </div>
        </div>
    </div>
    </div>
</template>