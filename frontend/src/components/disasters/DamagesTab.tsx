"use client"

import { HandlerDisasterResponse } from "@/api/model"
import { formatAmount, formatArea } from "../../app/disasters/[id]/utils"

type DamagesTabProps = {
  disaster: HandlerDisasterResponse
}

export default function DamagesTab({ disaster }: DamagesTabProps) {
  return (
    <>
      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        被害状況
      </h2>
      <div className="flex flex-wrap gap-4 p-4">
        <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
          <p className="text-[#111418] text-base font-medium leading-normal">
            被害レベル
          </p>
          <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
            {disaster.impact_level || "-"}
          </p>
        </div>
        <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
          <p className="text-[#111418] text-base font-medium leading-normal">
            被災面積
          </p>
          <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
            {formatArea(disaster.affected_area_size)}
          </p>
        </div>
        <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
          <p className="text-[#111418] text-base font-medium leading-normal">
            推定被害額
          </p>
          <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
            {formatAmount(disaster.estimated_damage_amount)}
          </p>
        </div>
      </div>

      {/* 被害詳細 */}
      <div className="px-4 py-3">
        <h3 className="text-[#111418] text-lg font-bold leading-tight mb-4">
          被害詳細
        </h3>
        <div className="bg-gray-50 rounded-lg p-4">
          <p className="text-[#111418] text-sm leading-relaxed whitespace-pre-wrap">
            {disaster.summary || "被害詳細情報がありません。"}
          </p>
        </div>
      </div>
    </>
  )
}
