import { useState } from "react"
import { useFacilityEquipments } from "../hooks/useFacilityEquipments"
import type { FacilityEquipmentSearchParams } from "../hooks/useFacilityEquipments"
import { FacilityEquipmentList } from "./FacilityEquipmentList"
import { FacilityEquipmentSearchForm } from "./FacilityEquipmentSearchForm"

export const FacilityEquipmentFeature = () => {
  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState<FacilityEquipmentSearchParams>({
    name: "",
  })

  // 検索実行時のハンドラ
  const handleSearch = (params: FacilityEquipmentSearchParams) => {
    setSearchParams(params)
  }

  // 施設設備情報を取得
  const { data: facilityEquipments, isLoading, error } = useFacilityEquipments(searchParams)

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            施設設備マスタ
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            施設設備の情報を管理するためのマスタ画面です。施設設備の追加、編集、削除が行えます。
          </p>
        </div>
      </div>

      {/* 検索フォーム */}
      <FacilityEquipmentSearchForm onSearch={handleSearch} />

      {/* 施設設備リスト */}
      <FacilityEquipmentList 
        facilityEquipments={facilityEquipments} 
        isLoading={isLoading} 
        error={error} 
      />
    </>
  )
}