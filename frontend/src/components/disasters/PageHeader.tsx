"use client"

import { HandlerDisasterResponse } from "@/api/model"
import Link from "next/link"
import { formatDate } from "../../app/disasters/[id]/utils"

type PageHeaderProps = {
  disaster: HandlerDisasterResponse
}

export default function PageHeader({ disaster }: PageHeaderProps) {
  return (
    <>
      {/* Breadcrumbs */}
      <div className="flex flex-wrap gap-2 p-4">
        <Link
          href="/frontend/public"
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
            災害 #{disaster.disaster_code || disaster.id}
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            {disaster.name} /{" "}
            {disaster.occurred_at ? formatDate(disaster.occurred_at) : "-"}
          </p>
        </div>
      </div>
    </>
  )
}
