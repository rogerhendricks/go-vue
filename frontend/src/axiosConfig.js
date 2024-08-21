import axios from 'axios'

axios.defaults.withCredentials = true
axios.interceptors.request.use(function (config) {
  const token = document.cookie.split('; ').find(row => row.startsWith('csrf_'))
  if (token) {
    config.headers['X-CSRF-Token'] = token.split('=')[1]
  }
  return config
})

export default axios