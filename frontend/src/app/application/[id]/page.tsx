"use client"

import { useGetSupportApplication } from "@/api/client"
import { HandlerSupportApplicationResponse } from "@/api/model"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import Link from "next/link"
import { useParams } from "next/navigation"

// ステータスに応じたバッジの色を定義
const getStatusBadgeClass = (status: string) => {
  switch (status) {
    case "審査中":
      return "bg-[#f0f2f4] text-[#111418]"
    case "書類確認中":
      return "bg-[#edf5ff] text-[#0055cc]"
    case "承認済":
      return "bg-[#e3fcef] text-[#006644]"
    case "支払処理中":
      return "bg-[#fff7e6] text-[#ff8b00]"
    case "完了":
      return "bg-[#e6fcf5] text-[#00a3bf]"
    default:
      return "bg-[#f0f2f4] text-[#111418]"
  }
}

export default function ApplicationDetailPage() {
  const params = useParams()
  const applicationId = params.id as string

  // API呼び出し
  const { data, isLoading, isError, error } = useGetSupportApplication<{
    data: HandlerSupportApplicationResponse
  }>(applicationId)

  const application = data?.data

  // 金額をフォーマットする関数
  const formatAmount = (amount: number) => {
    return new Intl.NumberFormat("ja-JP").format(amount) + "円"
  }

  // 日付をフォーマットする関数
  const formatDate = (dateString: string | undefined) => {
    if (!dateString) return "-"
    return dateString
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                支援申請詳細
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                農業災害に関する支援申請の詳細情報です。
              </p>
            </div>
            <div className="flex items-end">
              <Link
                href="/application"
                className="flex cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f4] text-[#111418] gap-2 text-sm font-bold leading-normal tracking-[0.015em]"
              >
                <span>一覧に戻る</span>
              </Link>
            </div>
          </div>

          {isLoading ? (
            <div className="p-4 text-center">読み込み中...</div>
          ) : isError ? (
            <div className="p-4 text-center text-red-500">
              データの取得に失敗しました
            </div>
          ) : application ? (
            <div className="px-4 py-3">
              <div className="overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
                <div className="p-6">
                  <div className="mb-6 flex justify-between items-center">
                    <h2 className="text-xl font-bold text-[#111418]">
                      申請 #{application.application_id}
                    </h2>
                    <div className="flex items-center">
                      <span className="mr-2 text-sm text-[#637588]">
                        ステータス:
                      </span>
                      <button
                        type="button"
                        className={`flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-3 ${getStatusBadgeClass(
                          application.status || ""
                        )} text-sm font-medium leading-normal`}
                      >
                        <span className="truncate">{application.status}</span>
                      </button>
                    </div>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div className="space-y-4">
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          申請者
                        </h3>
                        <p className="text-base text-[#111418]">
                          {application.applicant_name}
                        </p>
                      </div>
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          災害名
                        </h3>
                        <p className="text-base text-[#111418]">
                          {application.disaster_name}
                        </p>
                      </div>
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          申請金額
                        </h3>
                        <p className="text-base text-[#111418]">
                          {application.requested_amount
                            ? formatAmount(application.requested_amount)
                            : "-"}
                        </p>
                      </div>
                    </div>
                    <div className="space-y-4">
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          申請日
                        </h3>
                        <p className="text-base text-[#111418]">
                          {formatDate(application.application_date)}
                        </p>
                      </div>
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          審査日
                        </h3>
                        <p className="text-base text-[#111418]">
                          {formatDate(application.reviewed_at)}
                        </p>
                      </div>
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          承認日
                        </h3>
                        <p className="text-base text-[#111418]">
                          {formatDate(application.approved_at)}
                        </p>
                      </div>
                      <div>
                        <h3 className="text-sm font-medium text-[#637588]">
                          完了日
                        </h3>
                        <p className="text-base text-[#111418]">
                          {formatDate(application.completed_at)}
                        </p>
                      </div>
                    </div>
                  </div>

                  {application.notes && (
                    <div className="mt-6">
                      <h3 className="text-sm font-medium text-[#637588]">
                        備考
                      </h3>
                      <p className="text-base text-[#111418] whitespace-pre-line">
                        {application.notes}
                      </p>
                    </div>
                  )}

                  <div className="mt-6 pt-6 border-t border-[#dce0e5]">
                    <div className="flex justify-between text-sm text-[#637588]">
                      <div>
                        <span>作成日時: </span>
                        <span>{formatDate(application.created_at)}</span>
                      </div>
                      <div>
                        <span>更新日時: </span>
                        <span>{formatDate(application.updated_at)}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <div className="p-4 text-center text-[#637588]">
              申請データが見つかりません
            </div>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
