"use client"

import Header from "@/components/layout/header/page"
import { useListDisasters } from "@/api/client"
import Link from "next/link"

export default function Home() {
  // 災害情報を取得
  const {
    data: disastersResponse,
    isLoading,
    isError,
  } = useListDisasters({
    query: {
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
    },
  })

  const disasters = disastersResponse?.data || []

  // 統計情報を計算
  const totalDisasters = disasters.length
  const pendingCount = disasters.filter(
    (d) => d.status === "pending" || d.status === "under_review"
  ).length
  const recentDisasters = disasters
    .sort(
      (a, b) =>
        new Date(b.occurred_at || "").getTime() -
        new Date(a.occurred_at || "").getTime()
    )
    .slice(0, 5)

  // 日付フォーマット関数
  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString("ja-JP", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
    })
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

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                ダッシュボード
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                お疲れさまです！システムの概要と最新の災害情報をご確認ください。
              </p>
            </div>
          </div>

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            システム概要
          </h2>
          <div className="flex flex-wrap gap-4 p-4">
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                総災害件数
              </p>
              {isLoading ? (
                <div className="animate-pulse bg-gray-200 h-8 rounded"></div>
              ) : (
                <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                  {totalDisasters.toLocaleString()}
                </p>
              )}
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                対応待ち
              </p>
              {isLoading ? (
                <div className="animate-pulse bg-gray-200 h-8 rounded"></div>
              ) : (
                <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                  {pendingCount.toLocaleString()}
                </p>
              )}
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                完了済み
              </p>
              {isLoading ? (
                <div className="animate-pulse bg-gray-200 h-8 rounded"></div>
              ) : (
                <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                  {disasters
                    .filter((d) => d.status === "completed")
                    .length.toLocaleString()}
                </p>
              )}
            </div>
          </div>

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            最新災害情報
          </h2>

          {isLoading ? (
            <div className="p-4">
              <div className="animate-pulse">
                <div className="bg-gray-200 h-6 rounded mb-2"></div>
                <div className="bg-gray-200 h-4 rounded mb-4"></div>
                <div className="bg-gray-200 h-32 rounded"></div>
              </div>
            </div>
          ) : isError ? (
            <div className="p-4">
              <div className="bg-red-50 border border-red-200 rounded-lg p-4">
                <p className="text-red-600">災害情報の取得に失敗しました</p>
              </div>
            </div>
          ) : recentDisasters.length > 0 ? (
            <div className="p-4">
              <div className="flex items-stretch justify-between gap-4 rounded-lg">
                <div className="flex flex-col gap-1 flex-[2_2_0px]">
                  <p className="text-[#111418] text-base font-bold leading-tight">
                    {recentDisasters[0].name}
                  </p>
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    {recentDisasters[0].summary}
                  </p>
                  <p className="text-[#637588] text-xs font-normal leading-normal mt-2">
                    発生日:{" "}
                    {recentDisasters[0].occurred_at
                      ? formatDate(recentDisasters[0].occurred_at)
                      : "-"}{" "}
                    | 状態: {getStatusLabel(recentDisasters[0].status || "")}
                  </p>
                </div>
                <div
                  className="w-full bg-center bg-no-repeat aspect-video bg-cover rounded-lg flex-1"
                  style={{
                    backgroundImage:
                      'url("https://lh3.googleusercontent.com/aida-public/AB6AXuAwL6_AdRbGqk3fz9oAyKgApsJ5lzCZr323vDQidQ9sUfYW8fL05o-F1utFzuhac0AdevlWakVlW9vzMCRB7o_50MQ7boxvgVAkfcpYppzmOj0ApvLQc-dIfIhILwFZEzbaAXyFtfO4opsZF3lJTppgRsbw5Bs-DkYdzhVUAOh0Azxj54F00OhUq-XNDvnuNK5PCypb7MbFNqq1njXjjA8Mze_JUSLaovZSz4hDcO_wGxaoLBjVYxHtVpnfSQdmZDYXQmn10XKZ7kY")',
                  }}
                ></div>
              </div>
            </div>
          ) : (
            <div className="p-4">
              <p className="text-[#637588] text-center">災害情報がありません</p>
            </div>
          )}

          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            最近の災害一覧
          </h2>
          <div className="px-4 py-3 @container">
            {isLoading ? (
              <div className="animate-pulse">
                <div className="bg-gray-200 h-12 rounded mb-2"></div>
                <div className="space-y-2">
                  {[...Array(5)].map((_, i) => (
                    <div key={i} className="bg-gray-200 h-16 rounded"></div>
                  ))}
                </div>
              </div>
            ) : isError ? (
              <div className="bg-red-50 border border-red-200 rounded-lg p-4">
                <p className="text-red-600">災害一覧の取得に失敗しました</p>
              </div>
            ) : recentDisasters.length > 0 ? (
              <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
                <table className="flex-1">
                  <thead>
                    <tr className="bg-white">
                      <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                        災害コード
                      </th>
                      <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                        災害名
                      </th>
                      <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                        災害種別
                      </th>
                      <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                        ステータス
                      </th>
                      <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                        発生日
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    {recentDisasters.map((disaster) => (
                      <tr
                        key={disaster.id}
                        className="border-t border-t-[#dce0e5]"
                      >
                        <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                          <Link
                            href={`/disasters/${disaster.id}`}
                            className="text-[#007bff] hover:underline"
                          >
                            {disaster.disaster_code}
                          </Link>
                        </td>
                        <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                          {disaster.name}
                        </td>
                        <td className="table-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                          <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                            <span className="truncate">
                              {disaster.disaster_type}
                            </span>
                          </button>
                        </td>
                        <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                          <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                            <span className="truncate">
                              {getStatusLabel(disaster.status || "")}
                            </span>
                          </button>
                        </td>
                        <td className="table-column-600 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                          {disaster.occurred_at
                            ? formatDate(disaster.occurred_at)
                            : "-"}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            ) : (
              <div className="text-center p-8">
                <p className="text-[#637588] text-lg">災害情報がありません</p>
              </div>
            )}
          </div>

          {/* 全件表示リンク */}
          <div className="p-4 text-center">
            <Link
              href="/disasters"
              className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#111418] hover:bg-[#333] transition-colors"
            >
              すべての災害情報を見る
            </Link>
          </div>
        </div>
      </main>
    </div>
  )
}
