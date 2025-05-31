import React, { createContext, useContext } from "react"
import axios from "axios"

// Create a context for the Axios instance
const AxiosContext = createContext(axios)

// Create a provider component
export const AxiosProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  // Configure axios defaults if needed
  axios.defaults.baseURL = import.meta.env.VITE_API_URL || "http://localhost:8080"

  return (
    <AxiosContext.Provider value={axios}>{children}</AxiosContext.Provider>
  )
}

// Create a hook to use the axios instance
export const useAxios = () => useContext(AxiosContext)

export default AxiosProvider