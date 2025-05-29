"use client"

import { useGetDisaster, useGetDisastersIdTimelines } from "@/api/client"
import {
  HandlerDisasterResponse,
  HandlerListTimelinesResponse,
} from "@/api/model"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import { ErrorDisplay } from "@/components/ui/error-display"
import { LoadingSpinner } from "@/components/ui/loading"
import { useErrorHandler } from "@/hooks/useErrorHandler"
import Link from "next/link"
import { notFound } from "next/navigation"
import { useState } from "react"

type Props = {
  params: { id: string }
}

export default function DisasterDetailPage({ params }: Props) {
  const { handleError } = useErrorHandler()

  // APIから災害詳細を取得
  const {
    data: disasterResponse,
    isLoading,
    isError,
    error,
    refetch,
  } = useGetDisaster<{ data: HandlerDisasterResponse }>(
    params.id, // pathパラメータのidを直接渡す
    {
      query: {
        staleTime: 5 * 60 * 1000, // 5分間キャッシュ
        enabled: !!params.id, // idが存在する場合のみクエリを実行
      },
    }
  )

  // APIから災害に紐づくタイムライン一覧を取得
  const {
    data: timelinesResponse,
    isLoading: isLoadingTimelines,
    isError: isErrorTimelines,
    error: errorTimelines,
    refetch: refetchTimelines,
  } = useGetDisastersIdTimelines<{ data: HandlerListTimelinesResponse }>(
    params.id, // pathパラメータのidを直接渡す
    {
      query: {
        staleTime: 5 * 60 * 1000, // 5分間キャッシュ
        enabled: !!params.id, // idが存在する場合のみクエリを実行
      },
    }
  )

  const [activeTab, setActiveTab] = useState<
    "overview" | "damages" | "documents" | "timeline"
  >("overview")

  // ローディング状態
  if (isLoading) {
    return (
      <div className="layout-container flex h-full grow flex-col">
        <Header />
        <main className="px-40 flex flex-1 justify-center py-5">
          <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
            <div className="flex items-center justify-center min-h-[400px]">
              <LoadingSpinner size="lg" message="災害情報を読み込み中..." />
            </div>
          </div>
        </main>
        <Footer />
      </div>
    )
  }

  // エラー状態
  if (isError) {
    const apiError = handleError(error)
    return (
      <div className="layout-container flex h-full grow flex-col">
        <Header />
        <main className="px-40 flex flex-1 justify-center py-5">
          <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
            <ErrorDisplay
              title="災害情報の取得に失敗しました"
              message={apiError.message}
              onRetry={() => refetch()}
              className="mt-8"
            />
          </div>
        </main>
        <Footer />
      </div>
    )
  }

  // レスポンスからデータを取得
  const disaster = disasterResponse?.data

  if (!disaster) {
    notFound()
  }

  // 日付フォーマット関数
  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString("ja-JP", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
    })
  }

  // 数値フォーマット関数
  const formatNumber = (num: number | null | undefined): string => {
    if (num === null || num === undefined) return "-"
    return num.toLocaleString("ja-JP")
  }

  // 面積フォーマット関数
  const formatArea = (area: number | null | undefined): string => {
    if (area === null || area === undefined) return "-"
    return `${area.toLocaleString("ja-JP")} ha`
  }

  // 金額フォーマット関数
  const formatAmount = (amount: number | null | undefined): string => {
    if (amount === null || amount === undefined) return "-"
    return `¥${amount.toLocaleString("ja-JP")}`
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          {/* Breadcrumbs */}
          <div className="flex flex-wrap gap-2 p-4">
            <Link
              href="/"
              className="text-[#637588] text-base font-medium leading-normal hover:underline"
            >
              TOP
            </Link>
            <span className="text-[#637588] text-base font-medium leading-normal">
              /
            </span>
            <Link
              href="/disasters"
              className="text-[#637588] text-base font-medium leading-normal hover:underline"
            >
              災害リスト
            </Link>
            <span className="text-[#637588] text-base font-medium leading-normal">
              /
            </span>
            <span className="text-[#111418] text-base font-medium leading-normal">
              災害詳細
            </span>
          </div>

          {/* Page Header */}
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                災害 #{disaster.disaster_code || disaster.id}
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                {disaster.name} /{" "}
                {disaster.occurred_at ? formatDate(disaster.occurred_at) : "-"}
              </p>
            </div>
          </div>

          {/* タブナビゲーション */}
          <div className="pb-3">
            <div className="flex border-b border-[#dce0e5] px-4 gap-8">
              <button
                type="button"
                className={`flex flex-col items-center justify-center border-b-[3px] ${
                  activeTab === "overview"
                    ? "border-b-[#111418] text-[#111418]"
                    : "border-b-transparent text-[#637588]"
                } pb-[13px] pt-4 cursor-pointer`}
                onClick={() => setActiveTab("overview")}
              >
                <p
                  className={`${activeTab === "overview" ? "text-[#111418]" : "text-[#637588]"} text-sm font-bold leading-normal tracking-[0.015em]`}
                >
                  概要
                </p>
              </button>
              <button
                type="button"
                className={`flex flex-col items-center justify-center border-b-[3px] ${
                  activeTab === "damages"
                    ? "border-b-[#111418] text-[#111418]"
                    : "border-b-transparent text-[#637588]"
                } pb-[13px] pt-4 cursor-pointer`}
                onClick={() => setActiveTab("damages")}
              >
                <p
                  className={`${activeTab === "damages" ? "text-[#111418]" : "text-[#637588]"} text-sm font-bold leading-normal tracking-[0.015em]`}
                >
                  被害状況
                </p>
              </button>
              <button
                type="button"
                className={`flex flex-col items-center justify-center border-b-[3px] ${
                  activeTab === "documents"
                    ? "border-b-[#111418] text-[#111418]"
                    : "border-b-transparent text-[#637588]"
                } pb-[13px] pt-4 cursor-pointer`}
                onClick={() => setActiveTab("documents")}
              >
                <p
                  className={`${activeTab === "documents" ? "text-[#111418]" : "text-[#637588]"} text-sm font-bold leading-normal tracking-[0.015em]`}
                >
                  関連書類
                </p>
              </button>
              <button
                type="button"
                className={`flex flex-col items-center justify-center border-b-[3px] ${
                  activeTab === "timeline"
                    ? "border-b-[#111418] text-[#111418]"
                    : "border-b-transparent text-[#637588]"
                } pb-[13px] pt-4 cursor-pointer`}
                onClick={() => setActiveTab("timeline")}
              >
                <p
                  className={`${activeTab === "timeline" ? "text-[#111418]" : "text-[#637588]"} text-sm font-bold leading-normal tracking-[0.015em]`}
                >
                  タイムライン
                </p>
              </button>
            </div>
          </div>

          {/* Disaster Overview Section */}
          {activeTab === "overview" && (
            <>
              <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
                災害概要
              </h2>
              <div className="p-4 grid grid-cols-[20%_1fr] gap-x-6">
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    災害コード
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.disaster_code || "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    災害種別
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.disaster_type || "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    発生場所
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.prefecture?.name || "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    発生日時
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.occurred_at
                      ? formatDate(disaster.occurred_at)
                      : "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    被害レベル
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.impact_level || "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    被災面積
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {formatArea(disaster.affected_area_size)}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    推定被害額
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {formatAmount(disaster.estimated_damage_amount)}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    ステータス
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.status || "-"}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    被害概要
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal whitespace-pre-wrap">
                    {disaster.summary || "-"}
                  </p>
                </div>
              </div>
            </>
          )}

          {/* Damage Summary Section */}
          {activeTab === "damages" && (
            <>
              <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
                被害状況
              </h2>
              <div className="flex flex-wrap gap-4 p-4">
                <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
                  <p className="text-[#111418] text-base font-medium leading-normal">
                    被害レベル
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {disaster.impact_level || "-"}
                  </p>
                </div>
                <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
                  <p className="text-[#111418] text-base font-medium leading-normal">
                    被災面積
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {formatArea(disaster.affected_area_size)}
                  </p>
                </div>
                <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
                  <p className="text-[#111418] text-base font-medium leading-normal">
                    推定被害額
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {formatAmount(disaster.estimated_damage_amount)}
                  </p>
                </div>
              </div>

              {/* 被害詳細 */}
              <div className="px-4 py-3">
                <h3 className="text-[#111418] text-lg font-bold leading-tight mb-4">
                  被害詳細
                </h3>
                <div className="bg-gray-50 rounded-lg p-4">
                  <p className="text-[#111418] text-sm leading-relaxed whitespace-pre-wrap">
                    {disaster.summary || "被害詳細情報がありません。"}
                  </p>
                </div>
              </div>
            </>
          )}

          {/* Related Documents Section */}
          {activeTab === "documents" && (
            <>
              <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
                関連書類
              </h2>
              <div className="px-4 py-3 @container">
                <div className="bg-gray-50 rounded-lg p-8 text-center">
                  <p className="text-[#637588] text-base">
                    関連書類の機能は現在開発中です。
                  </p>
                </div>
              </div>
            </>
          )}

          {/* Timeline Section */}
          {activeTab === "timeline" && (
            <>
              <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
                タイムライン
              </h2>
              <div className="px-4 py-3">
                {isLoadingTimelines ? (
                  <div className="flex items-center justify-center min-h-[200px]">
                    <LoadingSpinner
                      size="md"
                      message="タイムライン情報を読み込み中..."
                    />
                  </div>
                ) : isErrorTimelines ? (
                  <ErrorDisplay
                    title="タイムライン情報の取得に失敗しました"
                    message={handleError(errorTimelines).message}
                    onRetry={() => refetchTimelines()}
                    className="mt-4"
                  />
                ) : timelinesResponse?.data?.timelines &&
                  timelinesResponse.data.timelines.length > 0 ? (
                  <div className="space-y-4">
                    {timelinesResponse.data.timelines.map((timeline) => (
                      <div
                        key={timeline.id}
                        className="border border-[#dce0e5] rounded-lg p-4"
                      >
                        <div className="flex justify-between items-start mb-2">
                          <h3 className="text-[#111418] text-lg font-bold">
                            {timeline.event_name}
                          </h3>
                          {timeline.severity && (
                            <span
                              className={`px-2 py-1 text-xs font-medium rounded-full ${
                                timeline.severity === "高"
                                  ? "bg-red-100 text-red-800"
                                  : timeline.severity === "中"
                                    ? "bg-yellow-100 text-yellow-800"
                                    : "bg-blue-100 text-blue-800"
                              }`}
                            >
                              {timeline.severity}
                            </span>
                          )}
                        </div>
                        <p className="text-[#637588] text-sm mb-2">
                          {timeline.event_time
                            ? formatDate(timeline.event_time)
                            : "-"}
                        </p>
                        <p className="text-[#111418] text-sm whitespace-pre-wrap">
                          {timeline.description || "-"}
                        </p>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="bg-gray-50 rounded-lg p-8 text-center">
                    <p className="text-[#637588] text-base">
                      この災害に関連するタイムラインはありません。
                    </p>
                  </div>
                )}
              </div>
            </>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
