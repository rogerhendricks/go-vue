import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import DeviceIndex  from '../views/devices/Index.vue'
import LeadIndex  from '../views/leads/Index.vue'
import UserSettings from '../UserSettings.vue'
import DocorIndex from '../views/doctors/Index.vue'
import Doctor from '../views/doctors/Doctor.vue'
import PatientIndex from '../views/patients/Index.vue'
import Patient from '../views/patients/Patient.vue'
import ImplantedDevices from '../views/patients/partials/ImplantedDevices.vue'
import ImplantedLeads from '../views/patients/partials/ImplantedLeads.vue'
import UpdateForm from '../views/reports/partials/UpdateForm.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/settings',
    name: 'UserSettings',
    component: UserSettings,
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/devices',
    name: 'Devices',
    component: DeviceIndex,
    meta: { requiresAuth: true }
  },
  { 
    path: '/leads',
    name: 'Leads',
    component: LeadIndex,
    meta: { requiresAuth: true }
  },
  {
    path: '/doctors',
    name: 'Doctors',
    component: DocorIndex,
    meta: { requiresAuth: true }
  },
  {
    path: '/doctor/:id',
    name: 'Doctor',
    component: Doctor,
    props: true
  },
  {
    path: '/patients',
    name: 'Patients',
    component: PatientIndex,
    meta: { requiresAuth: true }
  },
  {
    path: '/patients/:id',
    name: 'Patient',
    component: Patient,
    meta: {requiresAuth: true},
  },
  {
    path: '/patients/:id/implanted-devices',
    name: 'ImplantedDevices',
    component: ImplantedDevices,
    meta: {requiresAuth: true},
  },
  {
    path: '/patients/:id/implanted-leads',
    name: 'ImplantedLeads',
    component: ImplantedLeads,
    meta: {requiresAuth: true},
  },
  {
    path: '/reports/:id/update/',
    name: 'UpdateReport',
    component: UpdateForm,
    meta: {requiresAuth: true},
  },
  {
    path: '/:catchAll(.*)',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!localStorage.getItem('user')) {
      next('/login')
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router