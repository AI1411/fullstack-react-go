"use client"

import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import React, { useState } from "react"

const ReportGenerationPage: React.FC = () => {
  const [formData, setFormData] = useState({
    reportType: "",
    project: "",
    affectedArea: "",
    damageType: "",
    damageSeverity: "",
    reportDate: "",
    additionalNotes: "",
  })

  const handleSelectChange = (field: string, value: string) => {
    setFormData((prev) => ({
      ...prev,
      [field]: value,
    }))
  }

  const handleTextareaChange = (value: string) => {
    setFormData((prev) => ({
      ...prev,
      additionalNotes: value,
    }))
  }

  const handleGenerateReport = () => {
    console.log("Generating report with data:", formData)
    // ここにレポート生成のロジックを実装
  }

  return (
    <div
      className="relative flex size-full min-h-screen flex-col bg-white group/design-root overflow-x-hidden"
      style={
        {
          "--select-button-svg": `url('data:image/svg+xml,%3csvg xmlns=%27http://www.w3.org/2000/svg%27 width=%2724px%27 height=%2724px%27 fill=%27rgb(99,117,136)%27 viewBox=%270 0 256 256%27%3e%3cpath d=%27M181.66,170.34a8,8,0,0,1,0,11.32l-48,48a8,8,0,0,1-11.32,0l-48-48a8,8,0,0,1,11.32-11.32L128,212.69l42.34-42.35A8,8,0,0,1,181.66,170.34Zm-96-84.68L128,43.31l42.34,42.35a8,8,0,0,0,11.32-11.32l-48-48a8,8,0,0,0-11.32,0l-48,48A8,8,0,0,0,85.66,85.66Z%27%3e%3c/path%3e%3c/svg%3e')`,
          fontFamily: '"Public Sans", "Noto Sans", sans-serif',
        } as React.CSSProperties
      }
    >
      <div className="layout-container flex h-full grow flex-col">
        <Header />

        <div className="px-40 flex flex-1 justify-center py-5">
          <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
            <div className="flex flex-wrap justify-between gap-3 p-4">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72">
                レポート生成
              </p>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  レポート種別
                </p>
                <select
                  value={formData.reportType}
                  onChange={(e) =>
                    handleSelectChange("reportType", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="damage-assessment">被害査定レポート</option>
                  <option value="recovery-plan">復旧計画レポート</option>
                  <option value="financial-impact">経済影響レポート</option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  プロジェクト
                </p>
                <select
                  value={formData.project}
                  onChange={(e) =>
                    handleSelectChange("project", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="project-001">
                    プロジェクト001 - 台風被害対応
                  </option>
                  <option value="project-002">
                    プロジェクト002 - 水害復旧
                  </option>
                  <option value="project-003">
                    プロジェクト003 - 干ばつ対策
                  </option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  被災地域
                </p>
                <select
                  value={formData.affectedArea}
                  onChange={(e) =>
                    handleSelectChange("affectedArea", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="kyoto">京都府</option>
                  <option value="osaka">大阪府</option>
                  <option value="hyogo">兵庫県</option>
                  <option value="nara">奈良県</option>
                  <option value="shiga">滋賀県</option>
                  <option value="wakayama">和歌山県</option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  被害種別
                </p>
                <select
                  value={formData.damageType}
                  onChange={(e) =>
                    handleSelectChange("damageType", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="flood">洪水</option>
                  <option value="typhoon">台風</option>
                  <option value="hail">雹害</option>
                  <option value="drought">干ばつ</option>
                  <option value="landslide">地滑り</option>
                  <option value="strong-wind">強風</option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  被害程度
                </p>
                <select
                  value={formData.damageSeverity}
                  onChange={(e) =>
                    handleSelectChange("damageSeverity", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="minor">軽微</option>
                  <option value="moderate">中程度</option>
                  <option value="severe">深刻</option>
                  <option value="critical">甚大</option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  レポート日付
                </p>
                <select
                  value={formData.reportDate}
                  onChange={(e) =>
                    handleSelectChange("reportDate", e.target.value)
                  }
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                >
                  <option value="">選択してください</option>
                  <option value="2024-12-01">2024年12月1日</option>
                  <option value="2024-11-30">2024年11月30日</option>
                  <option value="2024-11-29">2024年11月29日</option>
                  <option value="custom">カスタム日付</option>
                </select>
              </label>
            </div>

            <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
              <label className="flex flex-col min-w-40 flex-1">
                <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                  追加メモ
                </p>
                <textarea
                  placeholder="追加のメモや観察事項を入力してください"
                  value={formData.additionalNotes}
                  onChange={(e) => handleTextareaChange(e.target.value)}
                  className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] min-h-36 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                />
              </label>
            </div>

            <div className="flex px-4 py-3 justify-end">
              <button
                onClick={handleGenerateReport}
                className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-[#197fe5] text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-[#1565c0] transition-colors"
              >
                <span className="truncate">レポート生成</span>
              </button>
            </div>
          </div>
        </div>

        <Footer />
      </div>
    </div>
  )
}

export default ReportGenerationPage
