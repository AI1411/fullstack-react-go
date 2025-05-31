import { useQuery } from "@tanstack/react-query"
import axios from "axios"
import type { Organization } from "./useOrganizations"

export const useOrganization = (organizationId: string) => {
  return useQuery<Organization>({
    queryKey: ["organization", organizationId],
    queryFn: async () => {
      const { data } = await axios.get(`/organizations/${organizationId}`)
      return data
    },
    enabled: !!organizationId,
  })
}
