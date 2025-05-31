export interface Disaster {
  id?: string
  name?: string
  disaster_type?: string
  status?: string
  occurred_at?: string
  // Add other fields as needed
}

export interface DisasterSearchParams {
  name: string
  disaster_type: string
  status: string
  date_from: string
  date_to: string
}
