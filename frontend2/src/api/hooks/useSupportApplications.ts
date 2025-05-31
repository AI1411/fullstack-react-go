// Custom React Query hooks for support application-related API endpoints
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import {
  createSupportApplication,
  getSupportApplication,
  listSupportApplications,
} from "../generated/client"

// Hook for fetching all support applications
export const useSupportApplications = (params?: {
  page?: number
  limit?: number
  disasterId?: string
  status?: string
}) => {
  return useQuery({
    queryKey: ["supportApplications", params],
    queryFn: () => listSupportApplications({ params }),
  })
}

// Hook for fetching a single support application by ID
export const useSupportApplication = (id: string) => {
  return useQuery({
    queryKey: ["supportApplication", id],
    queryFn: () => getSupportApplication(id),
    enabled: !!id,
  })
}

// Hook for creating a new support application
export const useCreateSupportApplication = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: createSupportApplication,
    onSuccess: (data) => {
      // Invalidate relevant queries
      queryClient.invalidateQueries({ queryKey: ["supportApplications"] })
    },
  })
}
