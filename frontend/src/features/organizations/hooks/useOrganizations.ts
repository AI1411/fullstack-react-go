import { useQuery } from "@tanstack/react-query"
import axios from "axios"
import type { User } from "../../users/hooks/useUsers"

export type Organization = {
  id: string
  name: string
  description: string
  created_at: string
  updated_at: string
  users?: User[]
}

export type OrganizationSearchParams = {
  name?: string
}

export const useOrganizations = (params: OrganizationSearchParams = {}) => {
  return useQuery<Organization[]>({
    queryKey: ["organizations", params],
    queryFn: async () => {
      const { data } = await axios.get("/organizations", { params })
      return data
    },
  })
}
