"use client"

import { useGetDisaster, useGetDisastersIdTimelines } from "@/api/client"
import {
  HandlerDisasterResponse,
  HandlerListTimelinesResponse,
} from "@/api/model"
import { useErrorHandler } from "@/hooks/useErrorHandler"
import { notFound } from "next/navigation"
import { useState } from "react"

// Import components
import DamagesTab from "@/components/disasters/DamagesTab"
import DocumentsTab from "@/components/disasters/DocumentsTab"
import ErrorState from "@/components/disasters/ErrorState"
import LoadingState from "@/components/disasters/LoadingState"
import OverviewTab from "@/components/disasters/OverviewTab"
import PageHeader from "@/components/disasters/PageHeader"
import TabNavigation, { TabType } from "@/components/disasters/TabNavigation"
import TimelineTab from "@/components/disasters/TimelineTab"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"

type Props = {
  params: { id: string }
}

export default function DisasterDetailPage({ params }: Props) {
  const { handleError } = useErrorHandler()

  // APIから災害詳細を取得
  const {
    data: disasterResponse,
    isLoading,
    isError,
    error,
    refetch,
  } = useGetDisaster<{ data: HandlerDisasterResponse }>(
    params.id, // pathパラメータのidを直接渡す
    {
      query: {
        staleTime: 5 * 60 * 1000, // 5分間キャッシュ
        enabled: !!params.id, // idが存在する場合のみクエリを実行
      },
    }
  )

  // APIから災害に紐づくタイムライン一覧を取得
  const {
    data: timelinesResponse,
    isLoading: isLoadingTimelines,
    isError: isErrorTimelines,
    error: errorTimelines,
    refetch: refetchTimelines,
  } = useGetDisastersIdTimelines<{ data: HandlerListTimelinesResponse }>(
    params.id, // pathパラメータのidを直接渡す
    {
      query: {
        staleTime: 5 * 60 * 1000, // 5分間キャッシュ
        enabled: !!params.id, // idが存在する場合のみクエリを実行
      },
    }
  )

  const [activeTab, setActiveTab] = useState<TabType>("overview")

  // ローディング状態
  if (isLoading) {
    return <LoadingState />
  }

  // エラー状態
  if (isError) {
    const apiError = handleError(error)
    return <ErrorState message={apiError.message} onRetry={() => refetch()} />
  }

  // レスポンスからデータを取得
  const disaster = disasterResponse?.data

  if (!disaster) {
    notFound()
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          {/* Page Header with Breadcrumbs */}
          <PageHeader disaster={disaster} />

          {/* Tab Navigation */}
          <TabNavigation activeTab={activeTab} setActiveTab={setActiveTab} />

          {/* Tab Content */}
          {activeTab === "overview" && <OverviewTab disaster={disaster} />}

          {activeTab === "damages" && <DamagesTab disaster={disaster} />}

          {activeTab === "documents" && <DocumentsTab />}

          {activeTab === "timeline" && (
            <TimelineTab
              timelinesResponse={timelinesResponse}
              isLoadingTimelines={isLoadingTimelines}
              isErrorTimelines={isErrorTimelines}
              errorTimelines={errorTimelines}
              refetchTimelines={refetchTimelines}
              handleError={handleError}
            />
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
