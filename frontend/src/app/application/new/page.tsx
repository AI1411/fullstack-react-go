"use client"

import { useCreateSupportApplication } from "@/api/client"
import { HandlerCreateSupportApplicationRequest } from "@/api/model"
import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import { useRouter } from "next/navigation"
import React, { useState } from "react"

const SupportApplicationCreatePage: React.FC = () => {
  const router = useRouter()
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const createSupportApplicationMutation = useCreateSupportApplication({
    mutation: {
      onSuccess: () => {
        setIsSubmitting(false)
        router.push("/application")
      },
      onError: (error) => {
        setIsSubmitting(false)
        setError("申請の作成に失敗しました。入力内容を確認してください。")
        console.error("Error creating support application:", error)
      },
    },
  })

  const [formData, setFormData] =
    useState<HandlerCreateSupportApplicationRequest>({
      application_id: "",
      application_date: new Date().toISOString().split("T")[0], // 今日の日付をデフォルト値に
      applicant_name: "",
      disaster_name: "",
      requested_amount: 0,
      status: "審査中", // デフォルト値
      notes: "",
    })

  const handleInputChange = (
    field: keyof HandlerCreateSupportApplicationRequest,
    value: string | number
  ) => {
    setFormData((prev) => ({
      ...prev,
      [field]: value,
    }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    setError(null)

    try {
      const requestData = {
        ...formData,
        requested_amount: Number(formData.requested_amount),
      }

      console.log("送信データ:", requestData)

      createSupportApplicationMutation.mutate(requestData)
      // 削除: router.push("/application")
      // onSuccessコールバックで既にリダイレクトが処理されるため、ここでのリダイレクトは不要
    } catch (error) {
      console.error("エラー:", error)
      setError(`申請の作成に失敗しました: ${error.message}`)
    } finally {
      setIsSubmitting(false)
    }
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
                支援申請作成
              </p>
            </div>

            {error && (
              <div className="mx-4 mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                {error}
              </div>
            )}

            <form onSubmit={handleSubmit} className="flex flex-col">
              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    申請ID <span className="text-red-500">*</span>
                  </p>
                  <input
                    type="text"
                    value={formData.application_id}
                    onChange={(e) =>
                      handleInputChange("application_id", e.target.value)
                    }
                    placeholder="例: A001"
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                    required
                  />
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    申請日 <span className="text-red-500">*</span>
                  </p>
                  <input
                    type="date"
                    value={formData.application_date}
                    onChange={(e) =>
                      handleInputChange("application_date", e.target.value)
                    }
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                    required
                  />
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    申請者名 <span className="text-red-500">*</span>
                  </p>
                  <input
                    type="text"
                    value={formData.applicant_name}
                    onChange={(e) =>
                      handleInputChange("applicant_name", e.target.value)
                    }
                    placeholder="例: 山田農園"
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                    required
                  />
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    災害名 <span className="text-red-500">*</span>
                  </p>
                  <input
                    type="text"
                    value={formData.disaster_name}
                    onChange={(e) =>
                      handleInputChange("disaster_name", e.target.value)
                    }
                    placeholder="例: 京都府洪水被害"
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                    required
                  />
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    申請金額（円） <span className="text-red-500">*</span>
                  </p>
                  <input
                    type="number"
                    value={formData.requested_amount || ""}
                    onChange={(e) =>
                      handleInputChange(
                        "requested_amount",
                        Number.parseInt(e.target.value) || 0
                      )
                    }
                    placeholder="例: 2500000"
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                    required
                    min="0"
                  />
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    ステータス
                  </p>
                  <select
                    value={formData.status || "審査中"}
                    onChange={(e) =>
                      handleInputChange("status", e.target.value)
                    }
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-14 bg-[image:--select-button-svg] placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                  >
                    <option value="審査中">審査中</option>
                    <option value="書類確認中">書類確認中</option>
                    <option value="承認済">承認済</option>
                    <option value="支払処理中">支払処理中</option>
                    <option value="完了">完了</option>
                  </select>
                </label>
              </div>

              <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                <label className="flex flex-col min-w-40 flex-1">
                  <p className="text-[#111418] text-base font-medium leading-normal pb-2">
                    備考
                  </p>
                  <textarea
                    placeholder="追加の備考や詳細情報を入力してください"
                    value={formData.notes || ""}
                    onChange={(e) => handleInputChange("notes", e.target.value)}
                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] min-h-36 placeholder:text-[#637588] p-[15px] text-base font-normal leading-normal"
                  />
                </label>
              </div>

              <div className="flex px-4 py-3 justify-end gap-3">
                <button
                  type="button"
                  onClick={() => router.push("/application")}
                  className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-white border border-[#dce0e5] text-[#111418] text-base font-bold leading-normal tracking-[0.015em] hover:bg-gray-50 transition-colors"
                >
                  <span className="truncate">キャンセル</span>
                </button>
                <button
                  type="submit"
                  disabled={isSubmitting}
                  className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-[#197fe5] text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-[#1565c0] transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  <span className="truncate">
                    {isSubmitting ? "送信中..." : "申請を作成"}
                  </span>
                </button>
              </div>
            </form>
          </div>
        </div>

        <Footer />
      </div>
    </div>
  )
}

export default SupportApplicationCreatePage
