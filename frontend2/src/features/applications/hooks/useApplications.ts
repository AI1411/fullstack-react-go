import { useListSupportApplications } from "../../../api/generated/client"
import type { HandlerListSupportApplicationsResponse } from "../../../api/generated/model"
import type { ApplicationSearchParams } from "../types"

export const useApplications = (searchParams: ApplicationSearchParams) => {
  const {
    data: applicationsResponse,
    isLoading,
    error,
  } = useListSupportApplications<{ data: HandlerListSupportApplicationsResponse }>({
    query: {
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
      queryKey: ["applications", searchParams],
    },
    axios: {
      params: searchParams,
    },
  })

  const applications = applicationsResponse?.data?.support_applications || []

  return {
    applications,
    isLoading,
    error,
  }
}