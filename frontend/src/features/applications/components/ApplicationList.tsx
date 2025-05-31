import { Link } from "@tanstack/react-router"
import type { HandlerSupportApplicationResponse } from "../../../api/generated/model"
import { formatDate, formatAmount, getStatusBadgeClass, getStatusLabel } from "../utils/formatters"

interface ApplicationListProps {
  applications: HandlerSupportApplicationResponse[]
  isLoading: boolean
  error: unknown
}

export const ApplicationList = ({ applications, isLoading, error }: ApplicationListProps) => {
  if (isLoading) {
    return <div>読み込み中</div>
  }

  if (error) {
    return <div>申請情報の取得に失敗しました</div>
  }

  return (
    <div className="px-4 py-3">
      {applications.length === 0 ? (
        <div className="flex flex-col items-center justify-center p-8 bg-white rounded-lg border border-[#dce0e5]">
          <p className="text-[#637588] text-sm">申請情報が見つかりませんでした</p>
        </div>
      ) : (
        <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
          <table className="flex-1">
            <thead>
              <tr className="bg-white">
                <th className="px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                  申請日
                </th>
                <th className="px-4 py-3 text-left text-[#111418] w-[200px] text-sm font-medium leading-normal">
                  申請者名
                </th>
                <th className="px-4 py-3 text-left text-[#111418] w-[200px] text-sm font-medium leading-normal">
                  災害名
                </th>
                <th className="px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium leading-normal">
                  申請金額
                </th>
                <th className="px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                  ステータス
                </th>
                <th className="px-4 py-3 text-left text-[#111418] w-[120px] text-[#637588] text-sm font-medium leading-normal">
                  アクション
                </th>
              </tr>
            </thead>
            <tbody>
              {applications.map((application) => (
                <tr
                  key={application.application_id}
                  className="border-t border-t-[#dce0e5]"
                >
                  <td className="h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                    {formatDate(application.application_date || "")}
                  </td>
                  <td className="h-[72px] px-4 py-2 w-[200px] text-[#111418] text-sm font-normal leading-normal">
                    {application.applicant_name || "-"}
                  </td>
                  <td className="h-[72px] px-4 py-2 w-[200px] text-[#637588] text-sm font-normal leading-normal">
                    {application.disaster_name || "-"}
                  </td>
                  <td className="h-[72px] px-4 py-2 w-[150px] text-[#637588] text-sm font-normal leading-normal">
                    {formatAmount(application.requested_amount || 0)}
                  </td>
                  <td className="h-[72px] px-4 py-2 w-[120px] text-sm font-normal leading-normal">
                    <span
                      className={`inline-flex min-w-[84px] max-w-[480px] items-center justify-center overflow-hidden rounded-lg h-8 px-4 ${getStatusBadgeClass(application.status || "")} text-sm font-medium leading-normal`}
                    >
                      {getStatusLabel(application.status || "")}
                    </span>
                  </td>
                  <td className="h-[72px] px-4 py-2 w-[120px] text-sm font-bold leading-normal tracking-[0.015em]">
                    <Link
                      to={`/application/${application.application_id}`}
                      className="text-[#197fe5] hover:underline"
                    >
                      詳細を表示
                    </Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}