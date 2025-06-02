import { useState } from "react"
import { GoogleMap, Marker, useJsApiLoader } from "@react-google-maps/api"
import { useDisasters } from "../hooks/useDisasters"
import type { DisasterSearchParams } from "../types"
import { DisasterList } from "./DisasterList"
import { DisasterSearchForm } from "./DisasterSearchForm"
import { useGoogleAPI } from "../../../providers/GoogleAPIprovider"

export const DisastersFeature = () => {
  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState<DisasterSearchParams>({
    name: "",
    disaster_type: "",
    status: "",
    date_from: "",
    date_to: "",
  })

  // 表示モード（リスト/マップ）の状態管理
  const [viewMode, setViewMode] = useState<"list" | "map">("list")

  // Get Google Maps API key from provider
  const { apiKey } = useGoogleAPI()

  // Load Google Maps API
  const { isLoaded } = useJsApiLoader({
    googleMapsApiKey: apiKey,
    id: 'google-map-script'
  })

  // 検索実行時のハンドラ
  const handleSearch = (params: DisasterSearchParams) => {
    setSearchParams(params)
  }

  // 災害情報を取得
  const { disasters, isLoading, error } = useDisasters(searchParams)

  // マップのスタイル
  const containerStyle = {
    width: "100%",
    height: "600px",
  }

  // マップの中心位置（日本の中心あたり）
  const center = {
    lat: 36.2048,
    lng: 138.2529,
  }

  return (
    <>
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
        <div className="flex items-center gap-2">
          <button
            onClick={() => setViewMode("list")}
            className={`px-4 py-2 rounded-lg ${
              viewMode === "list"
                ? "bg-[#197fe5] text-white"
                : "bg-gray-100 text-[#637588]"
            }`}
          >
            リスト表示
          </button>
          <button
            onClick={() => setViewMode("map")}
            className={`px-4 py-2 rounded-lg ${
              viewMode === "map"
                ? "bg-[#197fe5] text-white"
                : "bg-gray-100 text-[#637588]"
            }`}
          >
            マップ表示
          </button>
        </div>
      </div>

      {/* 検索フォーム */}
      <DisasterSearchForm onSearch={handleSearch} />

      {viewMode === "list" ? (
        /* 災害情報リスト */
        <DisasterList 
          disasters={disasters} 
          isLoading={isLoading} 
          error={error} 
        />
      ) : (
        /* 災害情報マップ */
        <div className="p-4">
          <div className="border rounded-lg overflow-hidden">
            {isLoaded ? (
              <GoogleMap mapContainerStyle={containerStyle} center={center} zoom={6}>
                {disasters.map((disaster) => 
                  disaster.latitude && disaster.longitude ? (
                    <Marker
                      key={disaster.id}
                      position={{
                        lat: disaster.latitude,
                        lng: disaster.longitude,
                      }}
                      onClick={() => {
                        if (disaster.id) {
                          // Use the correct URL format based on the project's routing
                          window.location.href = `/disasters/${disaster.id}`;
                        }
                      }}
                      title={disaster.name || "無題の災害"}
                    />
                  ) : null
                )}
              </GoogleMap>
            ) : (
              <div style={containerStyle} className="flex items-center justify-center bg-gray-100">
                <p>地図を読み込み中...</p>
              </div>
            )}
          </div>
        </div>
      )}
    </>
  )
}
