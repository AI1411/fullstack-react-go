import { useListDisasters } from "../../../api/generated/client"
import type { HandlerListDisastersResponse } from "../../../api/generated/model"
import type { DisasterSearchParams } from "../types"

export const useDisasters = (searchParams: DisasterSearchParams) => {
  const {
    data: disastersResponse,
    isLoading,
    error,
  } = useListDisasters<{ data: HandlerListDisastersResponse }>({
    queryKey: ["disasters", searchParams],
    axios: {
      params: searchParams,
    },
  })

  const disasters = disastersResponse?.data?.disasters || []

  return {
    disasters,
    isLoading,
    error,
  }
}
