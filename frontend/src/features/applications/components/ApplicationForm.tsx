import { useState } from "react"
import type { ApplicationFormData } from "../types"
import { useCreateApplication } from "../hooks"

interface ApplicationFormProps {
  onSuccess?: () => void
}

export const ApplicationForm = ({ onSuccess }: ApplicationFormProps) => {
  const [formData, setFormData] = useState<ApplicationFormData>({
    applicant_name: "",
    application_date: new Date().toISOString().split("T")[0],
    application_id: `APP-${Date.now()}`,
    disaster_name: "",
    notes: "",
    requested_amount: 0,
    status: "pending",
  })

  const { createApplication, isLoading, isSuccess, error } = useCreateApplication()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target
    setFormData((prev) => ({
      ...prev,
      [name]: name === "requested_amount" ? Number(value) : value,
    }))
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    createApplication(formData)
    if (onSuccess && isSuccess) {
      onSuccess()
    }
  }

  return (
    <div className="mx-4 mb-4 p-4 bg-white rounded-lg border border-[#dce0e5]">
      <h3 className="text-lg font-medium mb-4">新規支援申請</h3>

      {error && (
        <div className="mb-4 p-3 bg-[#ffebe6] text-[#ff5630] rounded-lg">
          申請の送信に失敗しました。入力内容を確認してください。
        </div>
      )}

      {isSuccess && (
        <div className="mb-4 p-3 bg-[#e6fcf5] text-[#00875a] rounded-lg">
          申請が正常に送信されました。
        </div>
      )}

      <form onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
          {/* 申請者名 */}
          <div>
            <label
              htmlFor="applicant_name"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              申請者名 <span className="text-[#ff5630]">*</span>
            </label>
            <input
              type="text"
              id="applicant_name"
              name="applicant_name"
              value={formData.applicant_name}
              onChange={handleChange}
              required
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
              placeholder="申請者の氏名を入力"
            />
          </div>

          {/* 申請日 */}
          <div>
            <label
              htmlFor="application_date"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              申請日 <span className="text-[#ff5630]">*</span>
            </label>
            <input
              type="date"
              id="application_date"
              name="application_date"
              value={formData.application_date}
              onChange={handleChange}
              required
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
            />
          </div>

          {/* 災害名 */}
          <div>
            <label
              htmlFor="disaster_name"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              災害名 <span className="text-[#ff5630]">*</span>
            </label>
            <input
              type="text"
              id="disaster_name"
              name="disaster_name"
              value={formData.disaster_name}
              onChange={handleChange}
              required
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
              placeholder="関連する災害名を入力"
            />
          </div>

          {/* 申請金額 */}
          <div>
            <label
              htmlFor="requested_amount"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              申請金額 (円) <span className="text-[#ff5630]">*</span>
            </label>
            <input
              type="number"
              id="requested_amount"
              name="requested_amount"
              value={formData.requested_amount}
              onChange={handleChange}
              required
              min="0"
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
              placeholder="申請金額を入力"
            />
          </div>
        </div>

        {/* 備考 */}
        <div className="mb-4">
          <label
            htmlFor="notes"
            className="block text-sm font-medium text-[#111418] mb-1"
          >
            備考
          </label>
          <textarea
            id="notes"
            name="notes"
            value={formData.notes}
            onChange={handleChange}
            rows={4}
            className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
            placeholder="申請に関する補足情報があれば入力してください"
          />
        </div>

        {/* 送信ボタン */}
        <div className="flex justify-end">
          <button
            type="submit"
            disabled={isLoading}
            className="inline-flex items-center justify-center px-6 py-2 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-[#197fe5] hover:bg-[#1565c0] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#197fe5] transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isLoading ? "送信中..." : "申請を送信"}
          </button>
        </div>
      </form>
    </div>
  )
}