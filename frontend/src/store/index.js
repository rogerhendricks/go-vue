import { createStore } from 'vuex'

import axios from '../axiosConfig'

export default createStore({
    state: {
        devices: [],
        leads: [],
        doctors: [],
        patient: {},
        reports: [],
        isLoading: false
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
        },
        setPatient(state, patient) {
          // console.log('Setting patient in store:', patient);
            state.patient = { ...patient}
        },
        setLoading(state, isLoading) {
          state.isLoading = isLoading
        },
        setReports(state, reports) {
          state.reports = reports
        },
        setReport(state, report) {
          state.report = report;
        },
    },
    actions: {
        async fetchDevices({ commit }) {
            try {
              const response = await axios.get('/api/deviceList')
              commit('setDevices', response.data)
            } catch (error) {
              console.error('Error fetching devices:', error)
            }
          },
          async fetchLeads({ commit }) {
            try {
              const response = await axios.get('/api/leadList')
              commit('setLeads', response.data)
            } catch (error) {
              console.error('Error fetching leads:', error)
            }
          },
          async fetchDoctors({ commit }) {
            try {
              const response = await axios.get('/api/doctorsList')
              commit('setDoctors', response.data)
            } catch (error) {
              console.error('Error fetching doctors:', error)
            }
          },
          async fetchPatient({ commit }, patientId) {
            try {
              const response = await axios.get(`/api/patients/${patientId}`)
              // Assuming the response contains a patient object
              commit('setPatient', response.data)
            } catch (error) {
              console.error('Error fetching patient:', error)
            }
          },
          async fetchReports({ commit }, patientId) {
            try {
              const response = await axios.get(`/api/${patientId}/reports`)
              commit('setReports', response.data)
            } catch (error) {
              console.error('Error fetching reports:', error)
            }
          },
          async fetchReport({ commit }, id) {
            try {
              const response = await axios.get(`/api/reports/${id}`)
              commit('setReport', response.data)
            } catch (error) {
              console.error('Error fetching report:', error)
            }
          }
          // async updatePatient({ commit }, patient) {
          //   try {
          //     await axios.put(`/api/patients/${patient.id}`, patient)
          //     commit('setPatient', patient)
          //   } catch (error) {
          //     console.error('Error updating patient:', error)
          //   }
          // }
    },
    modules: {}
})