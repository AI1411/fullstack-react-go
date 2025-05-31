import { useOrganizations } from "../hooks/useOrganizations"
import { OrganizationList } from "./OrganizationList"

export const OrganizationFeature = () => {
  // 組織情報を取得
  const { data: organizations, isLoading, error } = useOrganizations()

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            組織一覧
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            システムに登録されている組織の一覧です。
          </p>
        </div>
      </div>

      {/* 組織リスト */}
      <OrganizationList 
        organizations={organizations} 
        isLoading={isLoading} 
        error={error} 
      />
    </>
  )
}