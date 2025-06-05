import { useParams } from "@tanstack/react-router"
import { useState } from "react"

interface UploadResponse {
  filename: string
  path: string
  size: number
  uploaded: string
}

export const DisasterImageUpload = () => {
  // Get the disaster ID from the URL parameters
  const { disasterId } = useParams({ from: "/disasters/$disasterId" })

  // State for the selected file
  const [selectedFile, setSelectedFile] = useState<File | null>(null)

  // State for the upload status
  const [isUploading, setIsUploading] = useState(false)

  // State for the upload response
  const [uploadResponse, setUploadResponse] = useState<UploadResponse | null>(
    null
  )

  // State for error messages
  const [error, setError] = useState<string | null>(null)

  // Handle file selection
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files
    if (files && files.length > 0) {
      setSelectedFile(files[0])
      setError(null)
    }
  }

  // Handle file upload
  const handleUpload = async () => {
    if (!selectedFile) {
      setError("ファイルを選択してください")
      return
    }

    // Check if the file is an image
    if (!selectedFile.type.startsWith("image/")) {
      setError("画像ファイルを選択してください")
      return
    }

    setIsUploading(true)
    setError(null)

    try {
      const formData = new FormData()
      formData.append("image", selectedFile)

      const response = await fetch(
        `${import.meta.env.VITE_API_BASE_URL || ""}/disasters/${disasterId}/images`,
        {
          method: "POST",
          body: formData,
        }
      )

      if (!response.ok) {
        throw new Error(`アップロードに失敗しました: ${response.statusText}`)
      }

      const data = await response.json()
      setUploadResponse(data)
      setSelectedFile(null)

      // Reset the file input
      const fileInput = document.getElementById(
        "disaster-image"
      ) as HTMLInputElement
      if (fileInput) {
        fileInput.value = ""
      }
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "アップロードに失敗しました"
      )
    } finally {
      setIsUploading(false)
    }
  }

  return (
    <div className="bg-white rounded-lg border border-[#dce0e5] overflow-hidden mb-6 p-6">
      <h2 className="text-[#111418] text-xl font-bold mb-4">
        災害画像のアップロード
      </h2>

      <div className="mb-4">
        <label
          htmlFor="disaster-image"
          className="block text-[#637588] text-sm font-medium mb-2"
        >
          画像ファイルを選択
        </label>
        <input
          type="file"
          id="disaster-image"
          accept="image/*"
          onChange={handleFileChange}
          className="block w-full text-sm text-[#637588] file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-medium file:bg-[#f0f2f5] file:text-[#637588] hover:file:bg-[#e4e7eb]"
        />
        {selectedFile && (
          <p className="mt-2 text-sm text-[#637588]">
            選択されたファイル: {selectedFile.name} (
            {Math.round(selectedFile.size / 1024)} KB)
          </p>
        )}
      </div>

      <div className="flex items-center gap-4">
        <button
          onClick={handleUpload}
          disabled={!selectedFile || isUploading}
          className={`px-4 py-2 rounded-lg ${
            !selectedFile || isUploading
              ? "bg-gray-300 text-gray-500 cursor-not-allowed"
              : "bg-[#197fe5] text-white hover:bg-[#1668c3]"
          }`}
        >
          {isUploading ? "アップロード中..." : "アップロード"}
        </button>

        {error && <p className="text-red-500 text-sm">{error}</p>}
      </div>

      {uploadResponse && (
        <div className="mt-4 p-4 bg-green-50 border border-green-200 rounded-lg">
          <p className="text-green-700 font-medium">アップロード成功!</p>
          <p className="text-sm text-green-600">
            ファイル名: {uploadResponse.filename}
          </p>
          <p className="text-sm text-green-600">
            サイズ: {Math.round(uploadResponse.size / 1024)} KB
          </p>
        </div>
      )}
    </div>
  )
}
