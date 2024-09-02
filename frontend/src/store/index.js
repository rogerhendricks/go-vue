import { createStore } from 'vuex'

import axios from '../axiosConfig'

export default createStore({
    state: {
        devices: [],
        leads: [],
        doctors: []
    },
    getters:{
        doctors: state => state.doctors,
    },
    mutations: {
        setDevices(state, devices) {
            state.devices = devices
        },
        setLeads(state, leads) {
            state.leads = leads
        },
        setDoctors(state, doctors) {
            state.doctors = doctors
        }
    },
    actions: {
        async fetchDevices({ commit }) {
            // fetch devices from server and commit the mutation
            const response = await axios.get('/api/deviceList')
            commit('setDevices', response.data)
        },
        async fetchLeads({ commit }) {
            // fetch leads from server and commit the mutation
            const response = await axios.get('/api/leads')
            commit('setLeads', response.data)
        },
        async fetchDoctors({ commit }) {
            // fetch doctors from server and commit the mutation
            const response = await axios.get('/api/doctorsList')
            commit('setDoctors', response.data)
        }
    },
    modules: {}
})