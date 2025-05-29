"use client"

import { HandlerListTimelinesResponse } from "@/api/model"
import { ErrorDisplay } from "@/components/ui/error-display"
import { LoadingSpinner } from "@/components/ui/loading"
import { formatDate } from "../../app/disasters/[id]/utils"

type TimelineTabProps = {
  timelinesResponse?: { data: HandlerListTimelinesResponse }
  isLoadingTimelines: boolean
  isErrorTimelines: boolean
  errorTimelines: unknown
  refetchTimelines: () => void
  handleError: (error: unknown) => { message: string }
}

export default function TimelineTab({
  timelinesResponse,
  isLoadingTimelines,
  isErrorTimelines,
  errorTimelines,
  refetchTimelines,
  handleError,
}: TimelineTabProps) {
  return (
    <>
      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        タイムライン
      </h2>
      <div className="px-4 py-3">
        {isLoadingTimelines ? (
          <div className="flex items-center justify-center min-h-[200px]">
            <LoadingSpinner
              size="md"
              message="タイムライン情報を読み込み中..."
            />
          </div>
        ) : isErrorTimelines ? (
          <ErrorDisplay
            title="タイムライン情報の取得に失敗しました"
            message={handleError(errorTimelines).message}
            onRetry={() => refetchTimelines()}
            className="mt-4"
          />
        ) : timelinesResponse?.data?.timelines &&
          timelinesResponse.data.timelines.length > 0 ? (
          <div className="space-y-4">
            {timelinesResponse.data.timelines.map((timeline) => (
              <div
                key={timeline.id}
                className="border border-[#dce0e5] rounded-lg p-4"
              >
                <div className="flex justify-between items-start mb-2">
                  <h3 className="text-[#111418] text-lg font-bold">
                    {timeline.event_name}
                  </h3>
                  {timeline.severity && (
                    <span
                      className={`px-2 py-1 text-xs font-medium rounded-full ${
                        timeline.severity === "高"
                          ? "bg-red-100 text-red-800"
                          : timeline.severity === "中"
                            ? "bg-yellow-100 text-yellow-800"
                            : "bg-blue-100 text-blue-800"
                      }`}
                    >
                      {timeline.severity}
                    </span>
                  )}
                </div>
                <p className="text-[#637588] text-sm mb-2">
                  {timeline.event_time ? formatDate(timeline.event_time) : "-"}
                </p>
                <p className="text-[#111418] text-sm whitespace-pre-wrap">
                  {timeline.description || "-"}
                </p>
              </div>
            ))}
          </div>
        ) : (
          <div className="bg-gray-50 rounded-lg p-8 text-center">
            <p className="text-[#637588] text-base">
              この災害に関連するタイムラインはありません。
            </p>
          </div>
        )}
      </div>
    </>
  )
}
