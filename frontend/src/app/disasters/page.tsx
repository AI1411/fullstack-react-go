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
import { FileX, Search, X } from "lucide-react"
import Link from "next/link"
import { useState } from "react"

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

  // 検索条件の状態管理
  const [searchName, setSearchName] = useState("")
  const [searchDisasterType, setSearchDisasterType] = useState("")
  const [searchStatus, setSearchStatus] = useState("")
  const [searchDateFrom, setSearchDateFrom] = useState("")
  const [searchDateTo, setSearchDateTo] = useState("")

  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState({
    name: "",
    disaster_type: "",
    status: "",
    date_from: "",
    date_to: "",
  })

  // 検索実行関数
  const handleSearch = () => {
    setSearchParams({
      name: searchName,
      disaster_type: searchDisasterType,
      status: searchStatus,
      date_from: searchDateFrom,
      date_to: searchDateTo,
    })
  }

  // 検索条件クリア関数
  const handleClearSearch = () => {
    setSearchName("")
    setSearchDisasterType("")
    setSearchStatus("")
    setSearchDateFrom("")
    setSearchDateTo("")
    setSearchParams({
      name: "",
      disaster_type: "",
      status: "",
      date_from: "",
      date_to: "",
    })
  }

  // API呼び出し
  const {
    data: disastersResponse,
    isLoading,
    isError,
    error,
    refetch,
  } = useListDisasters<{ data: HandlerListDisastersResponse }>({
    axios: {
      params: searchParams,
    },
    query: {
      retry: 3,
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
      queryKey: ["disasters", searchParams],
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

          {/* 検索フォーム - テーブル幅に合わせて調整 */}
          <div className="mx-4 mb-4 p-4 bg-white rounded-lg border border-[#dce0e5]">
            <div className="flex items-center justify-between mb-3">
              <h3 className="text-lg font-medium">検索条件</h3>
              <button
                type="button"
                onClick={handleClearSearch}
                className="inline-flex items-center px-3 py-1.5 text-sm font-medium text-[#637588] hover:text-[#111418] transition-colors"
              >
                <X className="h-4 w-4 mr-1" />
                クリア
              </button>
            </div>

            {/* 1行目: 発生日の期間指定 */}
            <div className="flex gap-4 mb-4">
              <div className="flex items-center gap-2 flex-1">
                <div className="flex-1 max-w-[200px]">
                  <label
                    htmlFor="date_from"
                    className="block text-sm font-medium text-[#111418] mb-1"
                  >
                    発生日（開始）
                  </label>
                  <input
                    type="date"
                    id="date_from"
                    value={searchDateFrom}
                    onChange={(e) => setSearchDateFrom(e.target.value)}
                    className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
                  />
                </div>
                <div className="flex items-center pt-6">
                  <span className="text-[#637588] text-sm">〜</span>
                </div>
                <div className="flex-1 max-w-[200px]">
                  <label
                    htmlFor="date_to"
                    className="block text-sm font-medium text-[#111418] mb-1"
                  >
                    発生日（終了）
                  </label>
                  <input
                    type="date"
                    id="date_to"
                    value={searchDateTo}
                    onChange={(e) => setSearchDateTo(e.target.value)}
                    className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
                  />
                </div>
              </div>
            </div>

            {/* 2行目: その他の検索条件 */}
            <div className="flex gap-4 mb-4">
              {/* 災害名 - テーブルの発生日+災害名列の幅に対応 */}
              <div className="flex-1 max-w-[350px]">
                <label
                  htmlFor="name"
                  className="block text-sm font-medium text-[#111418] mb-1"
                >
                  災害名
                </label>
                <input
                  type="text"
                  id="name"
                  value={searchName}
                  onChange={(e) => setSearchName(e.target.value)}
                  className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
                  placeholder="災害名で検索"
                />
              </div>

              {/* 災害種別 - テーブルの災害種別列の幅に対応 */}
              <div className="flex-1 max-w-[200px]">
                <label
                  htmlFor="disaster_type"
                  className="block text-sm font-medium text-[#111418] mb-1"
                >
                  災害種別
                </label>
                <select
                  id="disaster_type"
                  value={searchDisasterType}
                  onChange={(e) => setSearchDisasterType(e.target.value)}
                  className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm bg-white"
                >
                  <option value="">すべて</option>
                  <option value="洪水">洪水</option>
                  <option value="地滑り">地滑り</option>
                  <option value="雹害">雹害</option>
                  <option value="干ばつ">干ばつ</option>
                  <option value="風害">風害</option>
                  <option value="地震">地震</option>
                  <option value="霜害">霜害</option>
                  <option value="病害虫">病害虫</option>
                </select>
              </div>

              {/* ステータス - テーブルのステータス列の幅に対応 */}
              <div className="flex-1 max-w-[150px]">
                <label
                  htmlFor="status"
                  className="block text-sm font-medium text-[#111418] mb-1"
                >
                  ステータス
                </label>
                <select
                  id="status"
                  value={searchStatus}
                  onChange={(e) => setSearchStatus(e.target.value)}
                  className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm bg-white"
                >
                  <option value="">すべて</option>
                  <option value="pending">未着手</option>
                  <option value="under_review">審査中</option>
                  <option value="in_progress">対応中</option>
                  <option value="completed">完了</option>
                </select>
              </div>

              {/* 検索ボタン - アクション列の幅に対応 */}
              <div className="flex items-end max-w-[120px]">
                <button
                  type="button"
                  onClick={handleSearch}
                  className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-[#197fe5] hover:bg-[#1565c0] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#197fe5] transition-colors"
                >
                  <Search className="h-4 w-4 mr-2" />
                  検索
                </button>
              </div>
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
                        <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                          発生日
                        </th>
                        <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[350px] text-sm font-medium leading-normal">
                          災害名
                        </th>
                        <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[200px] text-sm font-medium leading-normal">
                          災害種別
                        </th>
                        <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium leading-normal">
                          ステータス
                        </th>
                        <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-[120px] text-[#637588] text-sm font-medium leading-normal">
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
                          <td className="table-column-120 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                            {disaster.occurred_at
                              ? formatDate(disaster.occurred_at)
                              : "-"}
                          </td>
                          <td className="table-column-240 h-[72px] px-4 py-2 w-[350px] text-[#111418] text-sm font-normal leading-normal">
                            {disaster.name || "-"}
                          </td>
                          <td className="table-column-360 h-[72px] px-4 py-2 w-[200px] text-[#637588] text-sm font-normal leading-normal">
                            {disaster.disaster_type || "-"}
                          </td>
                          <td className="table-column-480 h-[72px] px-4 py-2 w-[150px] text-sm font-normal leading-normal">
                            <button
                              type="button"
                              className={`flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 ${getStatusBadgeClass(disaster.status || "")} text-sm font-medium leading-normal w-full`}
                            >
                              <span className="truncate">
                                {getStatusLabel(disaster.status || "")}
                              </span>
                            </button>
                          </td>
                          <td className="table-column-600 h-[72px] px-4 py-2 w-[120px] text-sm font-bold leading-normal tracking-[0.015em]">
                            <Link
                              href={`/disasters/${disaster.id}`}
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
              ))}
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
