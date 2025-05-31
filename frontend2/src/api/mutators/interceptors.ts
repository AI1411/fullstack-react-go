// Request and response interceptors for Axios
import { AxiosInstance, AxiosError, AxiosResponse } from 'axios'
import { refreshToken } from '../../features/auth/hooks/useAuth'

// Add request and response interceptors to an Axios instance
export const setupInterceptors = (axiosInstance: AxiosInstance): AxiosInstance => {
  // Request interceptor
  axiosInstance.interceptors.request.use(
    (config) => {
      // You can modify the request config here
      return config
    },
    (error) => {
      return Promise.reject(error)
    }
  )

  // Response interceptor
  axiosInstance.interceptors.response.use(
    (response: AxiosResponse) => {
      return response
    },
    async (error: AxiosError) => {
      const originalRequest = error.config
      
      // Handle 401 Unauthorized errors (token expired)
      if (error.response?.status === 401 && originalRequest && !originalRequest.headers._retry) {
        originalRequest.headers._retry = true
        
        try {
          // Attempt to refresh the token
          const newToken = await refreshToken()
          
          // Update the Authorization header with the new token
          originalRequest.headers.Authorization = `Bearer ${newToken}`
          
          // Retry the original request with the new token
          return axiosInstance(originalRequest)
        } catch (refreshError) {
          // If token refresh fails, redirect to login
          window.location.href = '/auth/login'
          return Promise.reject(refreshError)
        }
      }
      
      return Promise.reject(error)
    }
  )

  return axiosInstance
}