import { Link, useParams } from "@tanstack/react-router"
import { useOrganization } from "../hooks/useOrganization"
import { formatDate } from "../../../utils/formatters"

export const OrganizationDetail = () => {
  // Get the organization ID from the URL parameters
  const { organizationId } = useParams({ from: "/organizations/$organizationId" })

  // Fetch the organization details
  const { data: organization, isLoading, error } = useOrganization(organizationId)

  if (isLoading) {
    return <div className="p-4">読み込み中...</div>
  }

  if (error) {
    return <div className="p-4">組織情報の取得に失敗しました</div>
  }

  if (!organization) {
    return <div className="p-4">組織が見つかりませんでした</div>
  }

  return (
    <div className="p-4">
      <div className="mb-4">
        <Link
          to="/organizations"
          className="text-[#197fe5] hover:underline flex items-center"
        >
          ← 組織リストに戻る
        </Link>
      </div>

      <div className="flex flex-wrap justify-between gap-3 mb-6">
        <div className="flex min-w-72 flex-col gap-3">
          <h1 className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            {organization.name}
          </h1>
        </div>
      </div>

      <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-6">
          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">ID</h3>
            <p className="text-[#111418] text-base font-normal">
              {organization.id}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">名前</h3>
            <p className="text-[#111418] text-base font-normal">
              {organization.name}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">説明</h3>
            <p className="text-[#111418] text-base font-normal">
              {organization.description || "説明なし"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">作成日</h3>
            <p className="text-[#111418] text-base font-normal">
              {organization.created_at ? formatDate(organization.created_at) : "不明"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">更新日</h3>
            <p className="text-[#111418] text-base font-normal">
              {organization.updated_at ? formatDate(organization.updated_at) : "不明"}
            </p>
          </div>
        </div>
      </div>

      {/* ユーザー一覧 */}
      <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
        <div className="p-6">
          <h2 className="text-[#111418] text-xl font-semibold mb-4">所属ユーザー一覧</h2>

          {organization.users && organization.users.length > 0 ? (
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
                    <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">名前</th>
                    <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">メールアドレス</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {organization.users.map((user) => (
                    <tr key={user.id}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{user.id}</td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{user.name}</td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{user.email}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p className="text-gray-500">所属ユーザーはいません</p>
          )}
        </div>
      </div>
    </div>
  )
}
