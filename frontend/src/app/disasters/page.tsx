"use client"

import { useListDisasters } from "@/api/client"
import {
  type HandlerDisasterResponse,
  type HandlerListDisastersResponse,
} from "@/api/model"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import { EmptyState, ErrorDisplay } from "@/components/ui/error-display"
import { LoadingSpinner, LoadingTable } from "@/components/ui/loading"
import { useErrorHandler } from "@/hooks/useErrorHandler"
import { FileX } from "lucide-react"
import Link from "next/link"

// ステータスに応じたバッジの色を定義
const getStatusBadgeClass = (status: string) => {
  switch (status) {
    case "pending":
      return "bg-[#fff7e6] text-[#ff8b00]"
    case "under_review":
      return "bg-[#edf5ff] text-[#0055cc]"
    case "in_progress":
      return "bg-[#f0f2f4] text-[#111418]"
    case "completed":
      return "bg-[#e6fcf5] text-[#00a3bf]"
    default:
      return "bg-[#f0f2f4] text-[#111418]"
  }
}

// ステータスの日本語表示
const getStatusLabel = (status: string) => {
  switch (status) {
    case "pending":
      return "未着手"
    case "under_review":
      return "審査中"
    case "in_progress":
      return "対応中"
    case "completed":
      return "完了"
    default:
      return status
  }
}

// 日付フォーマット関数
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  })
}

export default function DisasterInfoPage() {
  const { handleError } = useErrorHandler()

  // API呼び出し
  const {
    data: disastersResponse,
    isLoading,
    isError,
    error,
    refetch,
  } = useListDisasters<{ data: HandlerListDisastersResponse }>({
    query: {
      retry: 3,
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
    },
  })

  if (isLoading) {
    return <div>読み込み中</div>
  }

  if (error) {
    return <div>データの取得に失敗しました</div>
  }

  const disasters = disastersResponse?.data?.disasters || []

  // エラーハンドリング
  const apiError = isError ? handleError(error) : null

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                災害情報リスト
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                報告されたすべての農業災害イベントを閲覧・管理できます。各項目には、
                発生日、被災地域、被害の概要に関する詳細情報が記載されています。
              </p>
            </div>
          </div>

          <div className="px-4 py-3 @container">
            {isLoading && (
              <div className="p-8">
                <LoadingSpinner
                  message="災害情報を読み込み中..."
                  className="mb-4"
                />
                <LoadingTable rows={10} columns={5} />
              </div>
            )}

            {isError && apiError && (
              <ErrorDisplay
                title="災害情報の取得に失敗しました"
                message={apiError.message}
                onRetry={() => refetch()}
                className="mb-4"
              />
            )}
            {!isLoading &&
              !isError &&
              (disasters.length === 0 ? (
                <EmptyState
                  title="災害情報が見つかりませんでした"
                  description="現在登録されている災害情報がありません。"
                  icon={<FileX className="h-12 w-12 text-gray-400" />}
                />
              ) : (
                <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
                  <table className="flex-1">
                    <thead>
                      <tr className="bg-white">
                        <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          発生日
                        </th>
                        <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          災害名
                        </th>
                        <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          災害種別
                        </th>
                        <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                          ステータス
                        </th>
                        <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-60 text-[#637588] text-sm font-medium leading-normal">
                          アクション
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {disasters.map((disaster: HandlerDisasterResponse) => (
                        <tr
                          key={disaster.id}
                          className="border-t border-t-[#dce0e5]"
                        >
                          <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                            {disaster.occurred_at
                              ? formatDate(disaster.occurred_at)
                              : "-"}
                          </td>
                          <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                            {disaster.name || "-"}
                          </td>
                          <td className="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                            {disaster.disaster_type || "-"}
                          </td>
                          <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                            <button
                              type="button"
                              className={`flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 ${getStatusBadgeClass(disaster.status || "")} text-sm font-medium leading-normal w-full`}
                            >
                              <span className="truncate">
                                {getStatusLabel(disaster.status || "")}
                              </span>
                            </button>
                          </td>
                          <td className="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-bold leading-normal tracking-[0.015em]">
                            <Link
                              href={`/disasters/${disaster.id}`}
                              className="text-[#007bff] hover:underline"
                            >
                              詳細を表示
                            </Link>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              ))}
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
