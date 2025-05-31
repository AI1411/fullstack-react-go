import { Link, useParams } from "@tanstack/react-router"
import type { HandlerSupportApplicationResponse } from "../../../api/generated/model"
import { useSupportApplication } from "../../../api/hooks/useSupportApplications"
import {
  formatDate,
  formatAmount,
  getStatusBadgeClass,
  getStatusLabel,
} from "../utils/formatters"

export const ApplicationDetail = () => {
  // Get the application ID from the URL parameters
  const { applicationId } = useParams({ from: "/application/$applicationId" })

  // Fetch the application details
  const { data: applicationResponse, isLoading, error } = useSupportApplication(applicationId)

  if (isLoading) {
    return <div className="p-4">読み込み中...</div>
  }

  if (error) {
    return <div className="p-4">申請情報の取得に失敗しました</div>
  }

  const application = applicationResponse?.data as HandlerSupportApplicationResponse

  if (!application) {
    return <div className="p-4">申請情報が見つかりませんでした</div>
  }

  return (
    <div className="p-4">
      <div className="mb-4">
        <Link
          to="/applications"
          className="text-[#197fe5] hover:underline flex items-center"
        >
          ← 申請リストに戻る
        </Link>
      </div>

      <div className="flex flex-wrap justify-between gap-3 mb-6">
        <div className="flex min-w-72 flex-col gap-3">
          <h1 className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            {application.applicant_name || "無名の申請者"}の申請
          </h1>
          <div>
            <span
              className={`inline-flex min-w-[84px] max-w-[480px] items-center justify-center overflow-hidden rounded-lg h-8 px-4 ${getStatusBadgeClass(application.status || "")} text-sm font-medium leading-normal`}
            >
              {getStatusLabel(application.status || "")}
            </span>
          </div>
        </div>
      </div>

      <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-6">
          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">申請者名</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.applicant_name || "不明"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">申請日</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.application_date ? formatDate(application.application_date) : "不明"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">災害名</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.disaster_name || "不明"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">申請金額</h3>
            <p className="text-[#111418] text-base font-normal">
              {formatAmount(application.requested_amount || 0)}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">審査日</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.reviewed_at ? formatDate(application.reviewed_at) : "未審査"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">承認日</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.approved_at ? formatDate(application.approved_at) : "未承認"}
            </p>
          </div>

          <div className="flex flex-col gap-2">
            <h3 className="text-[#637588] text-sm font-medium">完了日</h3>
            <p className="text-[#111418] text-base font-normal">
              {application.completed_at ? formatDate(application.completed_at) : "未完了"}
            </p>
          </div>
        </div>

        {application.notes && (
          <div className="border-t border-[#dce0e5] p-6">
            <h3 className="text-[#637588] text-sm font-medium mb-2">
              備考
            </h3>
            <p className="text-[#111418] text-base font-normal whitespace-pre-wrap">
              {application.notes}
            </p>
          </div>
        )}
      </div>
    </div>
  )
}