// Custom React Query hooks for prefecture-related API endpoints
import { useQuery } from '@tanstack/react-query'
import { getPrefecture, listPrefectures } from '../generated/client'

// Hook for fetching all prefectures
export const usePrefectures = () => {
  return useQuery({
    queryKey: ['prefectures'],
    queryFn: () => listPrefectures(),
  })
}

// Hook for fetching a single prefecture by ID
export const usePrefecture = (id: number) => {
  return useQuery({
    queryKey: ['prefecture', id],
    queryFn: () => getPrefecture(id),
    enabled: !!id,
  })
}

// Hook for fetching prefectures with disasters
export const usePrefecturesWithDisasters = () => {
  return useQuery({
    queryKey: ['prefectures', 'with-disasters'],
    queryFn: () => listPrefectures({ params: { withDisasters: true } }),
  })
}