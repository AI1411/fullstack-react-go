import { useState } from "react"
import { useApplications } from "../hooks/useApplications"
import type { ApplicationSearchParams } from "../types"
import { ApplicationList } from "./ApplicationList"
import { ApplicationSearchForm } from "./ApplicationSearchForm"
import { ApplicationForm } from "./ApplicationForm"

export const ApplicationFeature = () => {
  // 検索パラメータの状態管理
  const [searchParams, setSearchParams] = useState<ApplicationSearchParams>({
    applicant_name: "",
    disaster_name: "",
    status: "",
    date_from: "",
    date_to: "",
  })

  // 検索実行時のハンドラ
  const handleSearch = (params: ApplicationSearchParams) => {
    setSearchParams(params)
  }

  // 申請情報を取得
  const { applications, isLoading, error } = useApplications(searchParams)

  // 申請フォーム送信成功時のハンドラ
  const handleApplicationSuccess = () => {
    // 検索を再実行して最新の申請情報を取得
    setSearchParams({ ...searchParams })
  }

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            支援申請
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            災害に対する支援申請を行い、申請状況を確認できます。新規申請は下部のフォームから行ってください。
          </p>
        </div>
      </div>

      {/* 検索フォーム */}
      <ApplicationSearchForm onSearch={handleSearch} />

      {/* 申請情報リスト */}
      <ApplicationList 
        applications={applications} 
        isLoading={isLoading} 
        error={error} 
      />

      {/* 申請フォーム */}
      <div className="mt-8">
        <div className="px-4 mb-4">
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em]">
            新規支援申請
          </h2>
          <p className="text-[#637588] text-sm font-normal leading-normal mt-1">
            以下のフォームに必要事項を入力して、支援申請を行ってください。
          </p>
        </div>
        <ApplicationForm onSuccess={handleApplicationSuccess} />
      </div>
    </>
  )
}