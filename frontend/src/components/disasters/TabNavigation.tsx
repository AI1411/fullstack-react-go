"use client"

import { Dispatch, SetStateAction } from "react"

export type TabType = "overview" | "damages" | "documents" | "timeline"

type TabNavigationProps = {
  activeTab: TabType
  setActiveTab: Dispatch<SetStateAction<TabType>>
}

export default function TabNavigation({
  activeTab,
  setActiveTab,
}: TabNavigationProps) {
  return (
    <div className="pb-3">
      <div className="flex border-b border-[#dce0e5] px-4 gap-8">
        <TabButton
          label="概要"
          tabValue="overview"
          activeTab={activeTab}
          setActiveTab={setActiveTab}
        />
        <TabButton
          label="被害状況"
          tabValue="damages"
          activeTab={activeTab}
          setActiveTab={setActiveTab}
        />
        <TabButton
          label="関連書類"
          tabValue="documents"
          activeTab={activeTab}
          setActiveTab={setActiveTab}
        />
        <TabButton
          label="タイムライン"
          tabValue="timeline"
          activeTab={activeTab}
          setActiveTab={setActiveTab}
        />
      </div>
    </div>
  )
}

type TabButtonProps = {
  label: string
  tabValue: TabType
  activeTab: TabType
  setActiveTab: Dispatch<SetStateAction<TabType>>
}

function TabButton({
  label,
  tabValue,
  activeTab,
  setActiveTab,
}: TabButtonProps) {
  const isActive = activeTab === tabValue

  return (
    <button
      type="button"
      className={`flex flex-col items-center justify-center border-b-[3px] ${
        isActive
          ? "border-b-[#111418] text-[#111418]"
          : "border-b-transparent text-[#637588]"
      } pb-[13px] pt-4 cursor-pointer`}
      onClick={() => setActiveTab(tabValue)}
    >
      <p
        className={`${isActive ? "text-[#111418]" : "text-[#637588]"} text-sm font-bold leading-normal tracking-[0.015em]`}
      >
        {label}
      </p>
    </button>
  )
}
