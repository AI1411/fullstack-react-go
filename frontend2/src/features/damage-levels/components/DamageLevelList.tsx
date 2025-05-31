import type { DamageLevel } from "../hooks/useDamageLevels"

type DamageLevelListProps = {
  damageLevels: DamageLevel[] | undefined
  isLoading: boolean
  error: Error | null
}

export const DamageLevelList = ({
  damageLevels,
  isLoading,
  error,
}: DamageLevelListProps) => {
  if (isLoading) {
    return <div className="p-4">Loading...</div>
  }

  if (error) {
    return <div className="p-4 text-red-500">Error: {error.message}</div>
  }

  if (!damageLevels || damageLevels.length === 0) {
    return <div className="p-4">No damage levels found.</div>
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
          {damageLevels.map((level) => (
            <tr key={level.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {level.id}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {level.name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {level.description}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {new Date(level.created_at).toLocaleString()}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {new Date(level.updated_at).toLocaleString()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
