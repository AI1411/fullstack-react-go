export interface ApplicationSearchParams {
  applicant_name: string
  disaster_name: string
  status: string
  date_from: string
  date_to: string
}

export interface ApplicationFormData {
  applicant_name: string
  application_date: string
  application_id: string
  disaster_name: string
  notes: string
  requested_amount: number
  status: string
}