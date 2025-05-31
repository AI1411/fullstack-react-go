// Custom React Query hooks for disaster-related API endpoints
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import {
  createDisaster,
  deleteDisaster,
  getDisaster,
  listDisasters,
  updateDisaster,
} from "../generated/client"
import type { HandlerDisasterResponse, HandlerUpdateDisasterRequest, ListDisastersParams } from "../generated/model"

// Hook for fetching all disasters
export const useDisasters = (params?: ListDisastersParams) => {
  return useQuery({
    queryKey: ["disasters", params],
    queryFn: () => listDisasters(params),
  })
}

// Hook for fetching a single disaster by ID
export const useDisaster = (id: string) => {
  return useQuery({
    queryKey: ["disaster", id],
    queryFn: () => getDisaster(id),
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
      queryClient.invalidateQueries({ queryKey: ["disasters"] })
    },
  })
}

// Hook for updating an existing disaster
export const useUpdateDisaster = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: HandlerUpdateDisasterRequest }) => 
      updateDisaster(id, data),
    onSuccess: (response) => {
      const disaster = response.data as HandlerDisasterResponse
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ["disasters"] })
      queryClient.invalidateQueries({ queryKey: ["disaster", disaster.id] })
    },
  })
}

// Hook for deleting a disaster
export const useDeleteDisaster = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id }: { id: string }) => deleteDisaster(id),
    onSuccess: (_, { id }) => {
      // Invalidate specific queries
      queryClient.invalidateQueries({ queryKey: ["disasters"] })
      queryClient.invalidateQueries({ queryKey: ["disaster", id] })
    },
  })
}
