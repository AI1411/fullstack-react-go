import { useParams } from "@tanstack/react-router"
import { ImageUpload } from "../../../components/ImageUpload"

interface UploadResponse {
  filename: string
  path: string
  size: number
  uploaded: string
}

export const DisasterImageUpload = () => {
  // Get the disaster ID from the URL parameters
  const { disasterId } = useParams({ from: "/disasters/$disasterId" })

  // Handle file upload
  const handleUpload = async (formData: FormData): Promise<UploadResponse> => {
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

    return await response.json()
  }

  return <ImageUpload onUpload={handleUpload} title="災害画像のアップロード" />
}
