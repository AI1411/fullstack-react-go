import type { Organization } from "../hooks/useOrganizations"

type OrganizationListProps = {
  organizations: Organization[] | undefined
  isLoading: boolean
  error: Error | null
}

export const OrganizationList = ({
  organizations,
  isLoading,
  error,
}: OrganizationListProps) => {
  if (isLoading) {
    return <div className="p-4">Loading...</div>
  }

  if (error) {
    return <div className="p-4 text-red-500">Error: {error.message}</div>
  }

  if (!organizations || organizations.length === 0) {
    return <div className="p-4">No organizations found.</div>
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
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {organizations.map((organization) => (
            <tr key={organization.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {organization.id}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {organization.name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {organization.description}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}