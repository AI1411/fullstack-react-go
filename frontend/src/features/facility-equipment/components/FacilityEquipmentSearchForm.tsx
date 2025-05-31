import { useState } from "react"
import type { FacilityEquipmentSearchParams } from "../hooks/useFacilityEquipments"

type FacilityEquipmentSearchFormProps = {
  onSearch: (params: FacilityEquipmentSearchParams) => void
}

export const FacilityEquipmentSearchForm = ({
  onSearch,
}: FacilityEquipmentSearchFormProps) => {
  const [name, setName] = useState("")

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    onSearch({ name: name || undefined })
  }

  return (
    <div className="bg-white p-4 mb-4 rounded-lg shadow">
      <h2 className="text-lg font-semibold mb-4">検索条件</h2>
      <form onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label
              htmlFor="name"
              className="block text-sm font-medium text-gray-700"
            >
              名前
            </label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            />
          </div>
        </div>
        <div className="mt-4 flex justify-end">
          <button
            type="submit"
            className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            検索
          </button>
        </div>
      </form>
    </div>
  )
}