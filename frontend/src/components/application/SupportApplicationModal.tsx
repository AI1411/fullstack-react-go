"use client"

import React, { useState, useEffect } from "react"
import { useCreateSupportApplication } from "@/api/client"
import { HandlerCreateSupportApplicationRequest } from "@/api/model"

interface SupportApplicationModalProps {
  isOpen: boolean
  onClose: () => void
  onApplicationCreated: () => void
}

const SupportApplicationModal: React.FC<SupportApplicationModalProps> = ({
  isOpen,
  onClose,
  onApplicationCreated,
}) => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [formData, setFormData] =
    useState<HandlerCreateSupportApplicationRequest>({
      application_id: "",
      application_date: new Date().toISOString().split("T")[0],
      applicant_name: "",
      disaster_name: "",
      requested_amount: 0,
      status: "審査中",
      notes: "",
    })

  useEffect(() => {
    if (isOpen) {
      setError(null)
      setFormData({
        application_id: "",
        application_date: new Date().toISOString().split("T")[0],
        applicant_name: "",
        disaster_name: "",
        requested_amount: 0,
        status: "審査中",
        notes: "",
      })
    }
  }, [isOpen])

  const createSupportApplicationMutation = useCreateSupportApplication({
    mutation: {
      onSuccess: () => {
        setIsSubmitting(false)
        onApplicationCreated()
        onClose()
      },
      onError: (error: Error) => {
        setIsSubmitting(false)
        setError(
          `Failed to create application. Please check your input: ${error.message || "Unknown error"}`
        )
        console.error("Error creating support application:", error)
      },
    },
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
        requested_amount: Number(formData.requested_amount) || 0,
      }
      console.log("Submitting data (Modal):", requestData)
      createSupportApplicationMutation.mutate(requestData)
    } catch (err: any) {
      console.error("Error submitting form (Modal):", err)
      setError(`Failed to create application: ${err.message || "Unknown error"}`)
      setIsSubmitting(false)
    }
  }

  if (!isOpen) {
    return null
  }

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center p-4 z-50">
      <div className="bg-white p-6 rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold">支援申請作成</h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700 text-2xl"
            aria-label="Close"
          >
            &times;
          </button>
        </div>

        {error && (
          <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="flex flex-col space-y-4">
          <div>
            <label htmlFor="application_id_modal" className="block text-sm font-medium text-gray-700 mb-1">
              申請ID <span className="text-red-500">*</span>
            </label>
            <input
              id="application_id_modal"
              type="text"
              value={formData.application_id}
              onChange={(e) =>
                handleInputChange("application_id", e.target.value)
              }
              placeholder="例: A001"
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              required
            />
          </div>

          <div>
            <label htmlFor="application_date_modal" className="block text-sm font-medium text-gray-700 mb-1">
              申請日 <span className="text-red-500">*</span>
            </label>
            <input
              id="application_date_modal"
              type="date"
              value={formData.application_date}
              onChange={(e) =>
                handleInputChange("application_date", e.target.value)
              }
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              required
            />
          </div>

          <div>
            <label htmlFor="applicant_name_modal" className="block text-sm font-medium text-gray-700 mb-1">
              申請者名 <span className="text-red-500">*</span>
            </label>
            <input
              id="applicant_name_modal"
              type="text"
              value={formData.applicant_name}
              onChange={(e) =>
                handleInputChange("applicant_name", e.target.value)
              }
              placeholder="例: 山田農園"
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              required
            />
          </div>

          <div>
            <label htmlFor="disaster_name_modal" className="block text-sm font-medium text-gray-700 mb-1">
              災害名 <span className="text-red-500">*</span>
            </label>
            <input
              id="disaster_name_modal"
              type="text"
              value={formData.disaster_name}
              onChange={(e) =>
                handleInputChange("disaster_name", e.target.value)
              }
              placeholder="例: 京都府洪水被害"
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              required
            />
          </div>

          <div>
            <label htmlFor="requested_amount_modal" className="block text-sm font-medium text-gray-700 mb-1">
              申請金額（円） <span className="text-red-500">*</span>
            </label>
            <input
              id="requested_amount_modal"
              type="number"
              value={formData.requested_amount || ""}
              onChange={(e) =>
                handleInputChange(
                  "requested_amount",
                  Number.parseInt(e.target.value) || 0
                )
              }
              placeholder="例: 2500000"
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              required
              min="0"
            />
          </div>

          <div>
            <label htmlFor="status_modal" className="block text-sm font-medium text-gray-700 mb-1">
              ステータス
            </label>
            <select
              id="status_modal"
              value={formData.status || "審査中"}
              onChange={(e) => handleInputChange("status", e.target.value)}
              className="block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            >
              <option value="審査中">審査中</option>
              <option value="書類確認中">書類確認中</option>
              <option value="承認済">承認済</option>
              <option value="支払処理中">支払処理中</option>
              <option value="完了">完了</option>
            </select>
          </div>

          <div>
            <label htmlFor="notes_modal" className="block text-sm font-medium text-gray-700 mb-1">
              備考
            </label>
            <textarea
              id="notes_modal"
              placeholder="追加の備考や詳細情報を入力してください"
              value={formData.notes || ""}
              onChange={(e) => handleInputChange("notes", e.target.value)}
              className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              rows={3}
            />
          </div>

          <div className="flex justify-end space-x-3 pt-4">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors"
            >
              キャンセル
            </button>
            <button
              type="submit"
              disabled={isSubmitting}
              className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {isSubmitting ? "Submitting..." : "申請を作成"}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default SupportApplicationModal
