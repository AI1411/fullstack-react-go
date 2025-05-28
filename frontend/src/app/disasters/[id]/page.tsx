"use client"

import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import Link from "next/link"
import { notFound } from "next/navigation"
import { useState } from "react"

// サンプルデータ。IDに対応する詳細情報を含める
const disasterDetailsData: { [key: string]: any } = {
  "1": {
    id: "2023-001",
    name: "台風ラン",
    date: "2023年8月15日",
    type: "台風",
    location: "京都府",
    affectedArea: "150ヘクタール",
    estimatedDamage: "5億円",
    damageSummary: {
      total: 120,
      farmlands: 80,
      facilities: 40,
    },
    timeline: [
      {
        date: "2023年8月14日 09:00",
        event: "気象庁より台風接近の警報発令",
        status: "警報",
        details: "台風12号が近畿地方に接近中。京都府南部に暴風警報発令。",
      },
      {
        date: "2023年8月15日 06:30",
        event: "災害対策本部設置",
        status: "対応",
        details: "京都府農林水産部に災害対策本部を設置。被害状況の収集を開始。",
      },
      {
        date: "2023年8月15日 13:45",
        event: "最初の被害報告",
        status: "被害",
        details: "南丹市から農地冠水の第一報。約20ヘクタールの水田が被害。",
      },
      {
        date: "2023年8月15日 19:20",
        event: "被害拡大",
        status: "被害",
        details:
          "亀岡市、京田辺市からも被害報告。被害面積が100ヘクタールを超える見込み。",
      },
      {
        date: "2023年8月16日 08:00",
        event: "現地調査開始",
        status: "調査",
        details: "被害調査班3チームが現地調査を開始。",
      },
      {
        date: "2023年8月17日 10:30",
        event: "緊急支援決定",
        status: "支援",
        details:
          "農林水産省と連携し、緊急支援策を決定。復旧作業の資金支援を開始。",
      },
      {
        date: "2023年8月20日 15:00",
        event: "被害確定報告書提出",
        status: "報告",
        details: "最終的な被害状況を取りまとめ、農林水産省へ報告書を提出。",
      },
    ],
    documents: [
      {
        name: "被害報告書 2023-001",
        type: "報告書",
        uploaded: "2023年8月16日",
      },
      {
        name: "申請書 2023-001",
        type: "申請書",
        uploaded: "2023年8月17日",
      },
      {
        name: "評価レポート 2023-001",
        type: "レポート",
        uploaded: "2023年8月20日",
      },
    ],
  },
  "2": {
    id: "2023-002",
    name: "大阪水害",
    date: "2024年7月10日",
    type: "洪水",
    location: "大阪府",
    affectedArea: "120ヘクタール",
    estimatedDamage: "3億5000万円",
    damageSummary: {
      total: 95,
      farmlands: 65,
      facilities: 30,
    },
    timeline: [
      {
        date: "2024年7月9日 18:00",
        event: "大雨特別警報発令",
        status: "警報",
        details: "気象庁より大阪府北部に大雨特別警報が発令。",
      },
      {
        date: "2024年7月10日 02:30",
        event: "河川氾濫発生",
        status: "被害",
        details: "淀川水系の支流が複数箇所で氾濫。農地への浸水被害が発生。",
      },
      {
        date: "2024年7月10日 07:00",
        event: "災害対策本部設置",
        status: "対応",
        details: "大阪府農政部に災害対策本部を設置。",
      },
      {
        date: "2024年7月11日 09:15",
        event: "現地調査開始",
        status: "調査",
        details: "被害状況の調査を開始。ドローンによる空撮も実施。",
      },
      {
        date: "2024年7月12日 14:00",
        event: "支援対策決定",
        status: "支援",
        details: "被災農家への支援策を決定。復旧資金の緊急融資を開始。",
      },
    ],
    documents: [
      {
        name: "被害報告書 2023-002",
        type: "報告書",
        uploaded: "2024年7月11日",
      },
      {
        name: "申請書 2023-002",
        type: "申請書",
        uploaded: "2024年7月12日",
      },
    ],
  },
  "3": {
    id: "2023-003",
    name: "兵庫県雹害",
    date: "2024年7月5日",
    type: "雹害",
    location: "兵庫県",
    affectedArea: "85ヘクタール",
    estimatedDamage: "2億2000万円",
    damageSummary: {
      total: 70,
      farmlands: 55,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-003",
        type: "報告書",
        uploaded: "2024年7月6日",
      },
    ],
  },
  "4": {
    id: "2023-004",
    name: "奈良県干ばつ",
    date: "2024年6月28日",
    type: "干ばつ",
    location: "奈良県",
    affectedArea: "200ヘクタール",
    estimatedDamage: "4億円",
    damageSummary: {
      total: 110,
      farmlands: 95,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-004",
        type: "報告書",
        uploaded: "2024年6月29日",
      },
    ],
  },
  "5": {
    id: "2023-005",
    name: "滋賀県強風被害",
    date: "2024年6月20日",
    type: "強風",
    location: "滋賀県",
    affectedArea: "70ヘクタール",
    estimatedDamage: "1億8000万円",
    damageSummary: {
      total: 60,
      farmlands: 35,
      facilities: 25,
    },
    documents: [
      {
        name: "被害報告書 2023-005",
        type: "報告書",
        uploaded: "2024年6月21日",
      },
    ],
  },
  "6": {
    id: "2023-006",
    name: "和歌山県水害",
    date: "2024年6月12日",
    type: "洪水",
    location: "和歌山県",
    affectedArea: "110ヘクタール",
    estimatedDamage: "2億8000万円",
    damageSummary: {
      total: 85,
      farmlands: 65,
      facilities: 20,
    },
    documents: [
      {
        name: "被害報告書 2023-006",
        type: "報告書",
        uploaded: "2024年6月13日",
      },
    ],
  },
  "7": {
    id: "2023-007",
    name: "三重県地滑り",
    date: "2024年6月5日",
    type: "地滑り",
    location: "三重県",
    affectedArea: "60ヘクタール",
    estimatedDamage: "1億5000万円",
    damageSummary: {
      total: 45,
      farmlands: 30,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-007",
        type: "報告書",
        uploaded: "2024年6月6日",
      },
    ],
  },
  "8": {
    id: "2023-008",
    name: "愛知県雹害",
    date: "2024年5月28日",
    type: "雹害",
    location: "愛知県",
    affectedArea: "75ヘクタール",
    estimatedDamage: "1億9000万円",
    damageSummary: {
      total: 65,
      farmlands: 50,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-008",
        type: "報告書",
        uploaded: "2024年5月29日",
      },
    ],
  },
  "9": {
    id: "2023-009",
    name: "岐阜県干ばつ",
    date: "2024年5月20日",
    type: "干ばつ",
    location: "岐阜県",
    affectedArea: "180ヘクタール",
    estimatedDamage: "3億2000万円",
    damageSummary: {
      total: 90,
      farmlands: 75,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-009",
        type: "報告書",
        uploaded: "2024年5月21日",
      },
    ],
  },
  "10": {
    id: "2023-010",
    name: "静岡県強風被害",
    date: "2024年5月12日",
    type: "強風",
    location: "静岡県",
    affectedArea: "65ヘクタール",
    estimatedDamage: "1億7000万円",
    damageSummary: {
      total: 55,
      farmlands: 40,
      facilities: 15,
    },
    documents: [
      {
        name: "被害報告書 2023-010",
        type: "報告書",
        uploaded: "2024年5月13日",
      },
    ],
  },
}

const getDisasterById = (id: string) => {
  // 実際のIDは "2023-001" のような形式かもしれませんが、
  // URLの "1" とマッピングさせるため、ここでは "1" をキーとします。
  return disasterDetailsData[id]
}

type Props = {
  params: { id: string }
}

export default function DisasterDetailPage({ params }: Props) {
  const disaster = getDisasterById(params.id)
  const [activeTab, setActiveTab] = useState<
    "overview" | "damages" | "documents" | "timeline"
  >("overview")

  if (!disaster) {
    notFound()
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
                災害 #{disaster.id}
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                {disaster.name}, {disaster.date}
              </p>
            </div>
          </div>

          <div className="pb-3">
            <div className="flex border-b border-[#dce0e5] px-4 gap-8">
              <button
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
                    災害種別
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.type}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    発生場所
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.location}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    発生日
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.date}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    被災面積
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.affectedArea}
                  </p>
                </div>
                <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
                  <p className="text-[#637588] text-sm font-normal leading-normal">
                    推定被害額
                  </p>
                  <p className="text-[#111418] text-sm font-normal leading-normal">
                    {disaster.estimatedDamage}
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
                    総被害件数
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {disaster.damageSummary.total}
                  </p>
                </div>
                <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
                  <p className="text-[#111418] text-base font-medium leading-normal">
                    農地被害
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {disaster.damageSummary.farmlands}
                  </p>
                </div>
                <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
                  <p className="text-[#111418] text-base font-medium leading-normal">
                    施設被害
                  </p>
                  <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                    {disaster.damageSummary.facilities}
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
                <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
                  <table className="flex-1">
                    <thead>
                      <tr className="bg-white">
                        <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          書類名
                        </th>
                        <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          種類
                        </th>
                        <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                          アップロード日
                        </th>
                        <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-[#637588] text-sm font-medium leading-normal">
                          アクション
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {disaster.documents.map((doc: any, index: number) => (
                        <tr key={index} className="border-t border-t-[#dce0e5]">
                          <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                            {doc.name}
                          </td>
                          <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                            {doc.type}
                          </td>
                          <td className="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                            {doc.uploaded}
                          </td>
                          <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                            <Link
                              href="#"
                              className="text-blue-600 hover:underline"
                            >
                              表示
                            </Link>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
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
              <div className="p-4">
                <div className="flex flex-col gap-4 border border-[#dce0e5] rounded-lg p-6">
                  {disaster.timeline ? (
                    <div className="relative">
                      {disaster.timeline.map((item: any, index: number) => (
                        <div key={index} className="mb-8 ml-6 relative">
                          {/* 縦線 */}
                          {index < disaster.timeline.length - 1 && (
                            <div className="absolute left-[-24px] top-[24px] bottom-[-32px] w-[2px] bg-[#dce0e5]"></div>
                          )}

                          {/* タイムラインの丸いポイント */}
                          <div className="absolute left-[-30px] top-0 h-[12px] w-[12px] rounded-full border-4 border-white bg-blue-500"></div>

                          <div className="flex flex-col gap-1">
                            {/* 日付 */}
                            <p className="text-[#637588] text-sm font-medium leading-normal">
                              {item.date}
                            </p>

                            {/* イベント名とステータス */}
                            <div className="flex items-center gap-2">
                              <h3 className="text-[#111418] text-base font-bold leading-normal">
                                {item.event}
                              </h3>
                              <span
                                className={`px-2 py-1 rounded-full text-xs font-medium ${
                                  item.status === "警報"
                                    ? "bg-red-100 text-red-800"
                                    : item.status === "被害"
                                      ? "bg-orange-100 text-orange-800"
                                      : item.status === "対応"
                                        ? "bg-blue-100 text-blue-800"
                                        : item.status === "調査"
                                          ? "bg-purple-100 text-purple-800"
                                          : item.status === "支援"
                                            ? "bg-green-100 text-green-800"
                                            : item.status === "報告"
                                              ? "bg-gray-100 text-gray-800"
                                              : "bg-gray-100 text-gray-800"
                                }`}
                              >
                                {item.status}
                              </span>
                            </div>

                            {/* 詳細 */}
                            <p className="text-[#111418] text-sm font-normal leading-normal">
                              {item.details}
                            </p>
                          </div>
                        </div>
                      ))}
                    </div>
                  ) : (
                    <p className="text-[#111418] text-base font-medium leading-normal">
                      このリストの災害にはタイムライン情報がありません。
                    </p>
                  )}
                </div>
              </div>
            </>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
