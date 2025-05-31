import type { FacilityEquipment } from "../hooks/useFacilityEquipments"

type FacilityEquipmentListProps = {
  facilityEquipments: FacilityEquipment[] | undefined
  isLoading: boolean
  error: Error | null
}

export const FacilityEquipmentList = ({
  facilityEquipments,
  isLoading,
  error,
}: FacilityEquipmentListProps) => {
  if (isLoading) {
    return <div className="p-4">Loading...</div>
  }

  if (error) {
    return <div className="p-4 text-red-500">Error: {error.message}</div>
  }

  if (!facilityEquipments || facilityEquipments.length === 0) {
    return <div className="p-4">No facility equipment found.</div>
  }

  return (
    <div className="overflow-x-auto p-4">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              ID
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Description
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Created At
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Updated At
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {facilityEquipments.map((equipment) => (
            <tr key={equipment.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {equipment.id}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {equipment.name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {equipment.description}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {new Date(equipment.created_at).toLocaleString()}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {new Date(equipment.updated_at).toLocaleString()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
