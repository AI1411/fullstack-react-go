export interface Disaster {
  id?: string
  name?: string
  disaster_type?: string
  status?: string
  occurred_at?: string
  latitude?: number
  longitude?: number
  address?: string
  affected_area_size?: number
  estimated_damage_amount?: number
  impact_level?: string
  place_id?: string
  prefecture?: {
    id?: string
    name?: string
  }
  summary?: string
}

export interface DisasterSearchParams {
  name: string
  disaster_type: string
  status: string
  date_from: string
  date_to: string
}
