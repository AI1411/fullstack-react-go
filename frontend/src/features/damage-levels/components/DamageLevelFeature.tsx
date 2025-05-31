import { useState } from "react"
import { useDamageLevels } from "../hooks/useDamageLevels"
import type { DamageLevelSearchParams } from "../hooks/useDamageLevels"
import { DamageLevelList } from "./DamageLevelList"
import { DamageLevelSearchForm } from "./DamageLevelSearchForm"

export const DamageLevelFeature = () => {
  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState<DamageLevelSearchParams>({
    name: "",
  })

  // 検索実行時のハンドラ
  const handleSearch = (params: DamageLevelSearchParams) => {
    setSearchParams(params)
  }

  // 被害程度情報を取得
  const { data: damageLevels, isLoading, error } = useDamageLevels(searchParams)

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            被害程度マスタ
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            被害程度の情報を管理するためのマスタ画面です。被害程度の追加、編集、削除が行えます。
          </p>
        </div>
      </div>

      {/* 検索フォーム */}
      <DamageLevelSearchForm onSearch={handleSearch} />

      {/* 被害程度リスト */}
      <DamageLevelList 
        damageLevels={damageLevels} 
        isLoading={isLoading} 
        error={error} 
      />
    </>
  )
}