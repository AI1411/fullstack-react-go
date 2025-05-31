import { useQuery } from "@tanstack/react-query"
import axios from "axios"
import type { User } from "./useUsers"

export const useUser = (userId: string) => {
  return useQuery<User>({
    queryKey: ["user", userId],
    queryFn: async () => {
      const { data } = await axios.get(`/users/${userId}`)
      return data
    },
    enabled: !!userId,
  })
}
