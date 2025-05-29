"use client"

import { useListSupportApplications } from "@/api/client"
import {
  HandlerListSupportApplicationsResponse,
  HandlerSupportApplicationResponse,
} from "@/api/model"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import Link from "next/link"

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

export default function ApplicationPage() {
  // API呼び出し
  const { data, isLoading, isError, error } = useListSupportApplications<{
    data: HandlerListSupportApplicationsResponse
  }>()

  const applications = data?.data?.support_applications || []

  // 金額をフォーマットする関数
  const formatAmount = (amount: number) => {
    return new Intl.NumberFormat("ja-JP").format(amount) + "円"
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                支援申請一覧
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                農業災害に関する支援申請の一覧です。申請者、申請日、申請金額、状況などの情報を確認できます。
              </p>
            </div>
            <div className="flex items-end">
              <Link
                href="/application/new"
                className="flex cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#111418] text-white gap-2 text-sm font-bold leading-normal tracking-[0.015em]"
              >
                <span>新規申請</span>
              </Link>
            </div>
          </div>
          <div className="px-4 py-3 @container">
            <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
              <table className="flex-1">
                <thead>
                  <tr className="bg-white">
                    <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                      申請ID
                    </th>
                    <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                      申請日
                    </th>
                    <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium leading-normal">
                      申請者
                    </th>
                    <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium leading-normal">
                      災害名
                    </th>
                    <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium leading-normal">
                      申請金額
                    </th>
                    <th className="table-column-720 px-4 py-3 text-left text-[#111418] w-[100px] text-sm font-medium leading-normal">
                      ステータス
                    </th>
                    <th className="table-column-840 px-4 py-3 text-left text-[#111418] w-[100px] text-[#637588] text-sm font-medium leading-normal">
                      アクション
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {isLoading ? (
                    <tr>
                      <td
                        colSpan={7}
                        className="h-[72px] px-4 py-2 text-center text-[#637588]"
                      >
                        読み込み中...
                      </td>
                    </tr>
                  ) : isError ? (
                    <tr>
                      <td
                        colSpan={7}
                        className="h-[72px] px-4 py-2 text-center text-red-500"
                      >
                        データの取得に失敗しました
                      </td>
                    </tr>
                  ) : applications && applications.length > 0 ? (
                    applications.map(
                      (item: HandlerSupportApplicationResponse) => (
                        <tr
                          key={item.application_id}
                          className="border-t border-t-[#dce0e5]"
                        >
                          <td className="table-column-120 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                            {item.application_id}
                          </td>
                          <td className="table-column-240 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                            {item.application_date}
                          </td>
                          <td className="table-column-360 h-[72px] px-4 py-2 w-[150px] text-[#637588] text-sm font-normal leading-normal">
                            {item.applicant_name}
                          </td>
                          <td className="table-column-480 h-[72px] px-4 py-2 w-[150px] text-[#637588] text-sm font-normal leading-normal">
                            {item.disaster_name}
                          </td>
                          <td className="table-column-600 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                            {item.requested_amount
                              ? formatAmount(item.requested_amount)
                              : ""}
                          </td>
                          <td className="table-column-720 h-[72px] px-4 py-2 w-[100px] text-sm font-normal leading-normal">
                            <button
                              type="button"
                              className={`flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-3 ${getStatusBadgeClass(
                                item.status || ""
                              )} text-sm font-medium leading-normal w-full`}
                            >
                              <span className="truncate">{item.status}</span>
                            </button>
                          </td>
                          <td className="table-column-840 h-[72px] px-4 py-2 w-[100px] text-sm font-bold leading-normal tracking-[0.015em]">
                            <Link
                              href={`/application/${item.application_id}`}
                              className="text-[#007bff] hover:underline"
                            >
                              詳細を表示
                            </Link>
                          </td>
                        </tr>
                      )
                    )
                  ) : (
                    <tr>
                      <td
                        colSpan={7}
                        className="h-[72px] px-4 py-2 text-center text-[#637588]"
                      >
                        申請データがありません
                      </td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
