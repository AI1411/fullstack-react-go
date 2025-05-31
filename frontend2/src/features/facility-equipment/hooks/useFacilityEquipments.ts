import { useQuery } from "@tanstack/react-query"
import axios from "axios"

export type FacilityEquipment = {
  id: string
  name: string
  description: string
  created_at: string
  updated_at: string
}

export type FacilityEquipmentSearchParams = {
  name?: string
}

export const useFacilityEquipments = (params: FacilityEquipmentSearchParams = {}) => {
  return useQuery<FacilityEquipment[]>({
    queryKey: ["facilityEquipments", params],
    queryFn: async () => {
      const { data } = await axios.get("/facility-equipment", { params })
      return data
    },
  })
}