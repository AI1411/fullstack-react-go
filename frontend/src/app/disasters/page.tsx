import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import Link from "next/link" // next/link をインポート

// サンプルデータ (実際にはAPIなどから取得します)
const disasterData = [
  {
    id: "1", // 各アイテムに一意のIDを追加 (例)
    date: "2024-07-15",
    region: "京都府",
    summary: "洪水による広範囲な農作物被害",
    status: "審査中",
  },
  {
    id: "2",
    date: "2024-07-10",
    region: "大阪府",
    summary: "地滑りによる灌漑システムへの影響",
    status: "対応中",
  },
  {
    id: "3",
    date: "2024-07-05",
    region: "兵庫県",
    summary: "果樹園への深刻な雹害",
    status: "完了",
  },
  {
    id: "4",
    date: "2024-06-28",
    region: "奈良県",
    summary: "水田に影響を与える干ばつ",
    status: "審査中",
  },
  {
    id: "5",
    date: "2024-06-20",
    region: "滋賀県",
    summary: "温室への風害",
    status: "対応中",
  },
  {
    id: "6",
    date: "2024-06-12",
    region: "和歌山県",
    summary: "農地の冠水被害",
    status: "完了",
  },
  {
    id: "7",
    date: "2024-06-05",
    region: "三重県",
    summary: "アクセス道路への地滑り被害",
    status: "審査中",
  },
  {
    id: "8",
    date: "2024-05-28",
    region: "愛知県",
    summary: "野菜作物への雹害",
    status: "対応中",
  },
  {
    id: "9",
    date: "2024-05-20",
    region: "岐阜県",
    summary: "家畜に影響を与える干ばつ状況",
    status: "完了",
  },
  {
    id: "10",
    date: "2024-05-12",
    region: "静岡県",
    summary: "果樹への風害",
    status: "審査中",
  },
]

export default function DisasterInfoPage() {
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
            <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
              <table className="flex-1">
                <thead>
                  <tr className="bg-white">
                    <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      発生日
                    </th>
                    <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      地域
                    </th>
                    <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      被害概要
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
                  {disasterData.map(
                    (
                      item // index は不要になったため削除
                    ) => (
                      <tr key={item.id} className="border-t border-t-[#dce0e5]">
                        {" "}
                        {/* key に item.id を使用 */}
                        <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                          {item.date}
                        </td>
                        <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                          {item.region}
                        </td>
                        <td className="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                          {item.summary}
                        </td>
                        <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                          <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                            <span className="truncate">{item.status}</span>
                          </button>
                        </td>
                        <td className="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-bold leading-normal tracking-[0.015em]">
                          {/* Link コンポーネントでラップ */}
                          <Link
                            href={`/disasters/${item.id}`}
                            className="text-[#007bff] hover:underline"
                          >
                            詳細を表示
                          </Link>
                        </td>
                      </tr>
                    )
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
