// Authentication related Axios configuration
import baseAxios from './base'
import { getAuthToken } from '../../utils/storage'

// Create an authenticated Axios instance that includes the auth token
const authAxios = baseAxios.create()

// Add auth token to requests
authAxios.interceptors.request.use(
  (config) => {
    const token = getAuthToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

export default authAxios