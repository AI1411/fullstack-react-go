// Custom React Query hooks for disaster-related API endpoints
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { 
  getDisasters, 
  getDisasterById, 
  createDisaster, 
  updateDisaster, 
  deleteDisaster 
} from '../generated/client'

// Hook for fetching all disasters
export const useDisasters = (params?: { page?: number; limit?: number }) => {
  return useQuery({
    queryKey: ['disasters', params],
    queryFn: () => getDisasters(params),
  })
}

// Hook for fetching a single disaster by ID
export const useDisaster = (id: string) => {
  return useQuery({
    queryKey: ['disaster', id],
    queryFn: () => getDisasterById({ id }),
    enabled: !!id,
  })
}

// Hook for creating a new disaster
export const useCreateDisaster = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: createDisaster,
    onSuccess: () => {
      // Invalidate the disasters list query to refetch
      queryClient.invalidateQueries({ queryKey: ['disasters'] })
    },
  })
}

// Hook for updating an existing disaster
export const useUpdateDisaster = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: updateDisaster,
    onSuccess: (data) => {
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ['disasters'] })
      queryClient.invalidateQueries({ queryKey: ['disaster', data.id] })
    },
  })
}

// Hook for deleting a disaster
export const useDeleteDisaster = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: deleteDisaster,
    onSuccess: (_, variables) => {
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ['disasters'] })
      queryClient.invalidateQueries({ queryKey: ['disaster', variables.id] })
    },
  })
}