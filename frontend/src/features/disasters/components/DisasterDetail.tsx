import { GoogleMap, Marker, useJsApiLoader } from "@react-google-maps/api"
import { Link, useParams } from "@tanstack/react-router"
import { useState } from "react"
import type { HandlerDisasterResponse } from "../../../api/generated/model"
import { useDisaster } from "../../../api/hooks/useDisasters"
import { useGoogleAPI } from "../../../providers/GoogleAPIprovider"
import {
  formatDate,
  getStatusBadgeClass,
  getStatusLabel,
} from "../utils/formatters"
import { DisasterImageUpload } from "./DisasterImageUpload"

type TabType = "map" | "images" | "detail"

export const DisasterDetail = () => {
  // タブの状態管理
  const [activeTab, setActiveTab] = useState<TabType>("detail")

  // Get the disaster ID from the URL parameters
  const { disasterId } = useParams({ from: "/disasters/$disasterId" })

  // Get Google Maps API key from provider
  const { apiKey } = useGoogleAPI()

  // Load Google Maps API
  const { isLoaded } = useJsApiLoader({
    googleMapsApiKey: apiKey,
    id: "google-map-script",
  })

  // Fetch the disaster details
  const { data: disasterResponse, isLoading, error } = useDisaster(disasterId)

  // Format currency (damage amount)
  const formatCurrency = (value?: number) => {
    return new Intl.NumberFormat("ja-JP", {
      style: "currency",
      currency: "JPY",
      maximumFractionDigits: 0,
    }).format(value || 0)
  }

  const containerStyle = {
    width: "100%",
    height: "600px",
  }

  // Format area size
  const formatArea = (value?: number) => {
    return `${(value || 0).toLocaleString()} ha`
  }

  if (isLoading) {
    return <div className="p-4">読み込み中...</div>
  }

  if (error) {
    return <div className="p-4">災害情報の取得に失敗しました</div>
  }

  const disaster = disasterResponse?.data as HandlerDisasterResponse

  const center = {
    lat: disaster.latitude || 35.6895, // デフォルトは東京の緯度
    lng: disaster.longitude || 139.6917, // デフォルトは東京の経度
  }

  if (!disaster) {
    return <div className="p-4">災害情報が見つかりませんでした</div>
  }

  // タブのスタイル関数
  const getTabStyle = (tabType: TabType) => {
    const baseStyle =
      "px-4 py-2 rounded-t-lg font-medium transition-colors duration-200 border-b-2"
    if (activeTab === tabType) {
      return `${baseStyle} bg-white text-[#197fe5] border-[#197fe5]`
    }
    return `${baseStyle} bg-[#f8f9fa] text-[#637588] border-transparent hover:bg-[#e9ecef] hover:text-[#111418]`
  }

  return (
    <div className="p-4">
      <div className="mb-4">
        <Link
          to="/disasters"
          className="text-[#197fe5] hover:underline flex items-center"
        >
          ← 災害情報リストに戻る
        </Link>
      </div>

      <div className="flex flex-wrap justify-between gap-3 mb-6">
        <div className="flex min-w-72 flex-col gap-3">
          <h1 className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            {disaster.name || "無題の災害"}
          </h1>
          <div>
            <span
              className={`inline-flex min-w-[84px] max-w-[480px] items-center justify-center overflow-hidden rounded-lg h-8 px-4 ${getStatusBadgeClass(disaster.status || "")} text-sm font-medium leading-normal`}
            >
              {getStatusLabel(disaster.status || "")}
            </span>
          </div>
        </div>
      </div>

      {/* タブナビゲーション */}
      <div className="mb-6">
        <div className="border-b border-[#dce0e5] bg-[#f8f9fa] rounded-t-lg">
          <nav className="flex space-x-1 p-1">
            <button
              type="button"
              onClick={() => setActiveTab("detail")}
              className={getTabStyle("detail")}
            >
              ✍️ 詳細
            </button>
            <button
              type="button"
              onClick={() => setActiveTab("map")}
              className={getTabStyle("map")}
            >
              📍 地図
            </button>
            <button
              type="button"
              onClick={() => setActiveTab("images")}
              className={getTabStyle("images")}
            >
              📷 画像
            </button>
          </nav>
        </div>

        {/* タブコンテンツ */}
        <div className="border border-t-0 border-[#dce0e5] rounded-b-lg overflow-hidden bg-white">
          {activeTab === "detail" && (
            <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-6">
                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">
                    都道府県
                  </h3>
                  <p className="text-[#111418] text-base font-normal">
                    {disaster.municipality?.prefecture_name_kanji || "未分類"}
                  </p>
                </div>
                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">市町村</h3>
                  <p className="text-[#111418] text-base font-normal">
                    {disaster.municipality?.municipality_name_kanji || "未分類"}
                  </p>
                </div>
                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">
                    工種区分
                  </h3>
                  <p className="text-[#111418] text-base font-normal">
                    {disaster.work_category?.category_name || "未分類"}
                  </p>
                </div>

                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">発生日</h3>
                  <p className="text-[#111418] text-base font-normal">
                    {disaster.occurred_at
                      ? formatDate(disaster.occurred_at)
                      : "不明"}
                  </p>
                </div>

                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">被害額</h3>
                  <p className="text-[#111418] text-base font-normal">
                    {formatCurrency(disaster.estimated_damage_amount)}
                  </p>
                </div>

                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">
                    被害面積
                  </h3>
                  <p className="text-[#111418] text-base font-normal">
                    {formatArea(disaster.affected_area_size)}
                  </p>
                </div>

                <div className="flex flex-col gap-2">
                  <h3 className="text-[#637588] text-sm font-medium">
                    影響レベル
                  </h3>
                  <p className="text-[#111418] text-base font-normal">
                    {disaster.impact_level || "未設定"}
                  </p>
                </div>
              </div>

              {disaster.summary && (
                <div className="border-t border-[#dce0e5] p-6">
                  <h3 className="text-[#637588] text-sm font-medium mb-2">
                    災害の詳細
                  </h3>
                  <p className="text-[#111418] text-base font-normal whitespace-pre-wrap">
                    {disaster.summary}
                  </p>
                </div>
              )}
            </div>
          )}
          {activeTab === "map" && (
            <div>
              {isLoaded ? (
                <GoogleMap
                  mapContainerStyle={containerStyle}
                  center={center}
                  zoom={10}
                >
                  <Marker position={center} />
                </GoogleMap>
              ) : (
                <div
                  style={containerStyle}
                  className="flex items-center justify-center bg-gray-100"
                >
                  <p>地図を読み込み中...</p>
                </div>
              )}
            </div>
          )}

          {activeTab === "images" && (
            <div style={{ height: "600px" }} className="p-6">
              <DisasterImageUpload />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
