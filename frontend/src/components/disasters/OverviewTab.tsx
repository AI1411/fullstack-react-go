"use client"

import { HandlerDisasterResponse } from "@/api/model"
import {
  formatAmount,
  formatArea,
  formatDate,
} from "../../app/disasters/[id]/utils"

type OverviewTabProps = {
  disaster: HandlerDisasterResponse
}

export default function OverviewTab({ disaster }: OverviewTabProps) {
  return (
    <>
      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        災害概要
      </h2>
      <div className="p-4 grid grid-cols-[20%_1fr] gap-x-6">
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            災害コード
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.disaster_code || "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            災害種別
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.disaster_type || "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            発生場所
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.prefecture?.name || "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            発生日時
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.occurred_at ? formatDate(disaster.occurred_at) : "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            被害レベル
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.impact_level || "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            被災面積
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {formatArea(disaster.affected_area_size)}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            推定被害額
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {formatAmount(disaster.estimated_damage_amount)}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            ステータス
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal">
            {disaster.status || "-"}
          </p>
        </div>
        <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
          <p className="text-[#637588] text-sm font-normal leading-normal">
            被害概要
          </p>
          <p className="text-[#111418] text-sm font-normal leading-normal whitespace-pre-wrap">
            {disaster.summary || "-"}
          </p>
        </div>
      </div>
    </>
  )
}
