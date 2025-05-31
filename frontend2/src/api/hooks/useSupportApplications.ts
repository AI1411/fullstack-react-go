// Custom React Query hooks for support application-related API endpoints
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { 
  getSupportApplications, 
  getSupportApplicationById, 
  createSupportApplication, 
  updateSupportApplication, 
  deleteSupportApplication 
} from '../generated/client'

// Hook for fetching all support applications
export const useSupportApplications = (params?: { 
  page?: number; 
  limit?: number;
  disasterId?: string;
  status?: string;
}) => {
  return useQuery({
    queryKey: ['supportApplications', params],
    queryFn: () => getSupportApplications(params),
  })
}

// Hook for fetching a single support application by ID
export const useSupportApplication = (id: string) => {
  return useQuery({
    queryKey: ['supportApplication', id],
    queryFn: () => getSupportApplicationById({ id }),
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
      queryClient.invalidateQueries({ queryKey: ['supportApplications'] })
      if (data.disasterId) {
        queryClient.invalidateQueries({ 
          queryKey: ['supportApplications', { disasterId: data.disasterId }] 
        })
      }
    },
  })
}

// Hook for updating an existing support application
export const useUpdateSupportApplication = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: updateSupportApplication,
    onSuccess: (data) => {
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ['supportApplications'] })
      queryClient.invalidateQueries({ queryKey: ['supportApplication', data.id] })
      if (data.disasterId) {
        queryClient.invalidateQueries({ 
          queryKey: ['supportApplications', { disasterId: data.disasterId }] 
        })
      }
    },
  })
}

// Hook for deleting a support application
export const useDeleteSupportApplication = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: deleteSupportApplication,
    onSuccess: (_, variables) => {
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ['supportApplications'] })
      queryClient.invalidateQueries({ queryKey: ['supportApplication', variables.id] })
    },
  })
}