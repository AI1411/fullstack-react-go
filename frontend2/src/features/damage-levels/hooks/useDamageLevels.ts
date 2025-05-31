import { useQuery } from "@tanstack/react-query"
import axios from "axios"

export type DamageLevel = {
  id: string
  name: string
  description: string
  created_at: string
  updated_at: string
}

export type DamageLevelSearchParams = {
  name?: string
}

export const useDamageLevels = (params: DamageLevelSearchParams = {}) => {
  return useQuery<DamageLevel[]>({
    queryKey: ["damageLevels", params],
    queryFn: async () => {
      const { data } = await axios.get("/damage-levels", { params })
      return data
    },
  })
}