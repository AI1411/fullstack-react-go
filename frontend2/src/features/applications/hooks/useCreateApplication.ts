import { useCreateSupportApplication } from "../../../api/generated/client"
import type { ApplicationFormData } from "../types"

export const useCreateApplication = () => {
  const { mutate, isLoading, isSuccess, error } = useCreateSupportApplication()

  const createApplication = (formData: ApplicationFormData) => {
    mutate({
      data: {
        applicant_name: formData.applicant_name,
        application_date: formData.application_date,
        application_id: formData.application_id,
        disaster_name: formData.disaster_name,
        notes: formData.notes,
        requested_amount: formData.requested_amount,
        status: formData.status,
      },
    })
  }

  return {
    createApplication,
    isLoading,
    isSuccess,
    error,
  }
}