"use client"

import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import Link from "next/link"

// サンプルデータ (実際にはAPIなどから取得します)
const applicationData = [
  {
    id: "A001",
    date: "2024-07-18",
    applicant: "山田農園",
    disasterName: "京都府洪水被害",
    disasterId: "1",
    requestAmount: "2,500,000円",
    status: "審査中",
    category: "施設復旧",
  },
  {
    id: "A002",
    date: "2024-07-16",
    applicant: "大阪みかん生産組合",
    disasterName: "大阪府地滑り被害",
    disasterId: "2",
    requestAmount: "1,800,000円",
    status: "書類確認中",
    category: "機械購入",
  },
  {
    id: "A003",
    date: "2024-07-12",
    applicant: "兵庫果樹園",
    disasterName: "兵庫県雹害",
    disasterId: "3",
    requestAmount: "3,200,000円",
    status: "承認済",
    category: "農地復旧",
  },
  {
    id: "A004",
    date: "2024-07-10",
    applicant: "奈良田んぼ組合",
    disasterName: "奈良県干ばつ被害",
    disasterId: "4",
    requestAmount: "1,500,000円",
    status: "審査中",
    category: "水利施設",
  },
  {
    id: "A005",
    date: "2024-07-05",
    applicant: "滋賀グリーンハウス",
    disasterName: "滋賀県風害",
    disasterId: "5",
    requestAmount: "4,200,000円",
    status: "支払処理中",
    category: "施設復旧",
  },
  {
    id: "A006",
    date: "2024-07-01",
    applicant: "和歌山みかん農園",
    disasterName: "和歌山県水害",
    disasterId: "6",
    requestAmount: "1,700,000円",
    status: "完了",
    category: "農地復旧",
  },
  {
    id: "A007",
    date: "2024-06-28",
    applicant: "三重農業協同組合",
    disasterName: "三重県地滑り被害",
    disasterId: "7",
    requestAmount: "2,900,000円",
    status: "審査中",
    category: "道路復旧",
  },
  {
    id: "A008",
    date: "2024-06-22",
    applicant: "愛知野菜生産者",
    disasterName: "愛知県雹害",
    disasterId: "8",
    requestAmount: "1,200,000円",
    status: "書類確認中",
    category: "農地復旧",
  },
  {
    id: "A009",
    date: "2024-06-18",
    applicant: "岐阜畜産農家",
    disasterName: "岐阜県干ばつ被害",
    disasterId: "9",
    requestAmount: "3,500,000円",
    status: "完了",
    category: "給水設備",
  },
  {
    id: "A010",
    date: "2024-06-15",
    applicant: "静岡果樹園",
    disasterName: "静岡県風害",
    disasterId: "10",
    requestAmount: "1,900,000円",
    status: "承認済",
    category: "農地復旧",
  },
]

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
                  {applicationData.map((item) => (
                    <tr key={item.id} className="border-t border-t-[#dce0e5]">
                      <td className="table-column-120 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                        {item.id}
                      </td>
                      <td className="table-column-240 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                        {item.date}
                      </td>
                      <td className="table-column-360 h-[72px] px-4 py-2 w-[150px] text-[#637588] text-sm font-normal leading-normal">
                        {item.applicant}
                      </td>
                      <td className="table-column-480 h-[72px] px-4 py-2 w-[150px] text-[#637588] text-sm font-normal leading-normal">
                        <Link
                          href={`/disasters/${item.disasterId}`}
                          className="text-[#007bff] hover:underline"
                        >
                          {item.disasterName}
                        </Link>
                      </td>
                      <td className="table-column-600 h-[72px] px-4 py-2 w-[120px] text-[#637588] text-sm font-normal leading-normal">
                        {item.requestAmount}
                      </td>
                      <td className="table-column-720 h-[72px] px-4 py-2 w-[100px] text-sm font-normal leading-normal">
                        <button
                          className={`flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-3 ${getStatusBadgeClass(
                            item.status
                          )} text-sm font-medium leading-normal w-full`}
                        >
                          <span className="truncate">{item.status}</span>
                        </button>
                      </td>
                      <td className="table-column-840 h-[72px] px-4 py-2 w-[100px] text-sm font-bold leading-normal tracking-[0.015em]">
                        <Link
                          href={`/application/${item.id}`}
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
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
