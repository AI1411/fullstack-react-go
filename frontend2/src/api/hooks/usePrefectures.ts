// Custom React Query hooks for prefecture-related API endpoints
import { useQuery } from '@tanstack/react-query'
import { getPrefectures, getPrefectureById } from '../generated/client'

// Hook for fetching all prefectures
export const usePrefectures = () => {
  return useQuery({
    queryKey: ['prefectures'],
    queryFn: () => getPrefectures(),
  })
}

// Hook for fetching a single prefecture by ID
export const usePrefecture = (id: string) => {
  return useQuery({
    queryKey: ['prefecture', id],
    queryFn: () => getPrefectureById({ id }),
    enabled: !!id,
  })
}

// Hook for fetching prefectures with disasters
export const usePrefecturesWithDisasters = () => {
  return useQuery({
    queryKey: ['prefectures', 'with-disasters'],
    queryFn: () => getPrefectures({ withDisasters: true }),
  })
}