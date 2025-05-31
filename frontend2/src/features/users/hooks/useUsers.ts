import { useQuery } from "@tanstack/react-query"
import axios from "axios"

export type User = {
  id: number
  name: string
  email: string
  created_at: string
  updated_at: string
}

export type UserSearchParams = {
  name?: string
  email?: string
}

export const useUsers = (params: UserSearchParams = {}) => {
  return useQuery<User[]>({
    queryKey: ["users", params],
    queryFn: async () => {
      const { data } = await axios.get("/users", { params })
      return data
    },
  })
}