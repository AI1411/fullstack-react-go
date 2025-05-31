import { useState } from "react"
import { useDisasters } from "../hooks/useDisasters"
import type { DisasterSearchParams } from "../types"
import { DisasterList } from "./DisasterList"
import { DisasterSearchForm } from "./DisasterSearchForm"

export const DisastersFeature = () => {
  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState<DisasterSearchParams>({
    name: "",
    disaster_type: "",
    status: "",
    date_from: "",
    date_to: "",
  })

  // 検索実行時のハンドラ
  const handleSearch = (params: DisasterSearchParams) => {
    setSearchParams(params)
  }

  // 災害情報を取得
  const { disasters, isLoading, error } = useDisasters(searchParams)

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
      </div>

      {/* 検索フォーム */}
      <DisasterSearchForm onSearch={handleSearch} />

      {/* 災害情報リスト */}
      <DisasterList 
        disasters={disasters} 
        isLoading={isLoading} 
        error={error} 
      />
    </>
  )
}