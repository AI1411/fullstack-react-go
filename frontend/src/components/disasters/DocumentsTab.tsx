"use client"

import { useRef, useState } from "react"

// サンプルデータの型定義
interface ImageDocument {
  id: string
  name: string
  url: string
  uploadedAt: string
  size: number
  description?: string
}

// ファイルサイズをフォーマットする関数
const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return "0 Bytes"
  const k = 1024
  const sizes = ["Bytes", "KB", "MB", "GB"]
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Number.parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i]
}

// 日付フォーマット関数
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  })
}

// 画像プレビューモーダルコンポーネント
const ImageModal = ({
  image,
  onClose,
}: {
  image: ImageDocument | null
  onClose: () => void
}) => {
  if (!image) return null

  return (
    <div className="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div className="max-w-4xl max-h-full w-full">
        <div className="bg-white rounded-lg overflow-hidden">
          {/* ヘッダー */}
          <div className="flex justify-between items-center p-4 border-b border-gray-200">
            <div>
              <h3 className="text-lg font-bold text-[#111418]">{image.name}</h3>
              <p className="text-sm text-[#637588]">
                {formatFileSize(image.size)} • {formatDate(image.uploadedAt)}
              </p>
            </div>
            <button
              onClick={onClose}
              className="text-gray-500 hover:text-gray-700 text-2xl font-light w-8 h-8 flex items-center justify-center"
            >
              ×
            </button>
          </div>

          {/* 画像表示エリア */}
          <div className="p-4 max-h-[70vh] overflow-auto">
            <img
              src={image.url}
              alt={image.name}
              className="max-w-full h-auto mx-auto rounded"
              style={{ maxHeight: "60vh" }}
            />
          </div>

          {/* フッター */}
          <div className="flex justify-between items-center p-4 border-t border-gray-200 bg-gray-50">
            {image.description && (
              <p className="text-sm text-[#637588] flex-1 mr-4">
                {image.description}
              </p>
            )}
            <div className="flex gap-2">
              <a
                href={image.url}
                download={image.name}
                className="px-4 py-2 bg-[#111418] text-white text-sm rounded hover:bg-gray-800 transition-colors"
              >
                ダウンロード
              </a>
              <a
                href={image.url}
                target="_blank"
                rel="noopener noreferrer"
                className="px-4 py-2 border border-[#dce0e5] text-[#111418] text-sm rounded hover:bg-gray-50 transition-colors"
              >
                新しいタブで開く
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

// ファイルアップロードコンポーネント
const FileUpload = ({
  onUpload,
}: {
  onUpload: (files: FileList) => void
}) => {
  const [isDragging, setIsDragging] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault()
    setIsDragging(false)
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      onUpload(e.dataTransfer.files)
    }
  }

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      onUpload(e.target.files)
    }
  }

  return (
    <div
      className={`border-2 border-dashed rounded-lg p-6 text-center transition-colors cursor-pointer ${
        isDragging
          ? "border-blue-400 bg-blue-50"
          : "border-gray-300 hover:border-gray-400"
      }`}
      onDrop={handleDrop}
      onDragOver={(e) => e.preventDefault()}
      onDragEnter={() => setIsDragging(true)}
      onDragLeave={() => setIsDragging(false)}
      onClick={() => fileInputRef.current?.click()}
    >
      <input
        ref={fileInputRef}
        type="file"
        multiple
        accept="image/*"
        onChange={handleFileSelect}
        className="hidden"
      />

      <div className="text-6xl text-gray-400 mb-4">📁</div>
      <p className="text-[#111418] font-medium mb-2">
        画像ファイルをアップロード
      </p>
      <p className="text-[#637588] text-sm mb-4">
        ファイルをドラッグ&ドロップするか、クリックして選択してください
      </p>
      <div className="inline-flex items-center px-4 py-2 bg-[#111418] text-white text-sm rounded hover:bg-gray-800 transition-colors">
        ファイルを選択
      </div>
      <p className="text-xs text-[#637588] mt-2">
        対応形式: JPG, PNG, GIF, WebP
      </p>
    </div>
  )
}

export default function DocumentsTab() {
  // サンプルデータ（実際の実装では API から取得）
  const [images, setImages] = useState<ImageDocument[]>([
    {
      id: "1",
      name: "被害状況_現場写真1.jpg",
      url: "https://images.unsplash.com/photo-1500382017468-9049fed747ef?w=800&h=600&fit=crop",
      uploadedAt: "2025-05-29T10:30:00Z",
      size: 1024000,
      description: "災害発生直後の現場写真",
    },
    {
      id: "2",
      name: "復旧作業_進捗写真.jpg",
      url: "https://images.unsplash.com/photo-1581094794329-c8112a89af12?w=800&h=600&fit=crop",
      uploadedAt: "2025-05-29T14:15:00Z",
      size: 856000,
      description: "復旧作業の進捗状況",
    },
    {
      id: "3",
      name: "避難所_状況確認.png",
      url: "https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=800&h=600&fit=crop",
      uploadedAt: "2025-05-29T16:45:00Z",
      size: 720000,
      description: "避難所の設営状況",
    },
  ])

  const [selectedImage, setSelectedImage] = useState<ImageDocument | null>(null)
  const [viewMode, setViewMode] = useState<"grid" | "list">("grid")

  const handleUpload = (files: FileList) => {
    // 実際の実装では API にアップロード
    Array.from(files).forEach((file, index) => {
      if (file.type.startsWith("image/")) {
        const newImage: ImageDocument = {
          id: `new-${Date.now()}-${index}`,
          name: file.name,
          url: URL.createObjectURL(file),
          uploadedAt: new Date().toISOString(),
          size: file.size,
        }
        setImages((prev) => [...prev, newImage])
      }
    })
  }

  const handleDelete = (imageId: string) => {
    setImages((prev) => prev.filter((img) => img.id !== imageId))
  }

  return (
    <>
      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        関連書類
      </h2>

      <div className="px-4 py-3">
        {/* コントロールバー */}
        <div className="flex justify-between items-center mb-6">
          <div className="flex items-center gap-4">
            <span className="text-[#637588] text-sm">
              {images.length} 件の画像
            </span>
          </div>

          <div className="flex items-center gap-2">
            {/* 表示切替ボタン */}
            <button
              type="button"
              onClick={() => setViewMode("grid")}
              className={`p-2 rounded ${
                viewMode === "grid"
                  ? "bg-[#111418] text-white"
                  : "bg-gray-100 text-gray-600 hover:bg-gray-200"
              }`}
            >
              ⊞
            </button>
            <button
              type="button"
              onClick={() => setViewMode("list")}
              className={`p-2 rounded ${
                viewMode === "list"
                  ? "bg-[#111418] text-white"
                  : "bg-gray-100 text-gray-600 hover:bg-gray-200"
              }`}
            >
              ☰
            </button>
          </div>
        </div>

        {/* アップロードエリア */}
        <div className="mb-6">
          <FileUpload onUpload={handleUpload} />
        </div>

        {/* 画像一覧 */}
        {images.length > 0 ? (
          <div
            className={
              viewMode === "grid"
                ? "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
                : "space-y-4"
            }
          >
            {images.map((image) => (
              <div
                key={image.id}
                className={`border border-[#dce0e5] rounded-lg overflow-hidden hover:shadow-lg transition-shadow ${
                  viewMode === "list" ? "flex items-center" : ""
                }`}
              >
                {/* 画像部分 */}
                <div
                  className={`cursor-pointer ${
                    viewMode === "list"
                      ? "w-24 h-24 flex-shrink-0"
                      : "aspect-video"
                  }`}
                  onClick={() => setSelectedImage(image)}
                >
                  <img
                    src={image.url}
                    alt={image.name}
                    className="w-full h-full object-cover hover:opacity-90 transition-opacity"
                  />
                </div>

                {/* 情報部分 */}
                <div className={`p-4 ${viewMode === "list" ? "flex-1" : ""}`}>
                  <div className="flex justify-between items-start mb-2">
                    <h3 className="text-[#111418] font-medium text-sm truncate flex-1 mr-2">
                      {image.name}
                    </h3>
                    <button
                      onClick={() => handleDelete(image.id)}
                      className="text-red-500 hover:text-red-700 text-sm"
                    >
                      🗑
                    </button>
                  </div>

                  <div className="space-y-1">
                    <p className="text-[#637588] text-xs">
                      {formatFileSize(image.size)}
                    </p>
                    <p className="text-[#637588] text-xs">
                      {formatDate(image.uploadedAt)}
                    </p>
                    {image.description && (
                      <p className="text-[#637588] text-xs line-clamp-2">
                        {image.description}
                      </p>
                    )}
                  </div>

                  <div className="flex gap-2 mt-3">
                    <button
                      onClick={() => setSelectedImage(image)}
                      className="px-3 py-1 bg-[#111418] text-white text-xs rounded hover:bg-gray-800 transition-colors"
                    >
                      プレビュー
                    </button>
                    <a
                      href={image.url}
                      download={image.name}
                      className="px-3 py-1 border border-[#dce0e5] text-[#111418] text-xs rounded hover:bg-gray-50 transition-colors"
                    >
                      ダウンロード
                    </a>
                  </div>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="bg-gray-50 rounded-lg p-8 text-center">
            <div className="text-6xl text-gray-400 mb-4">🖼️</div>
            <p className="text-[#637588] text-base mb-4">
              まだ画像がアップロードされていません
            </p>
            <p className="text-[#637588] text-sm">
              上のアップロードエリアから画像を追加してください
            </p>
          </div>
        )}
      </div>

      {/* 画像プレビューモーダル */}
      <ImageModal
        image={selectedImage}
        onClose={() => setSelectedImage(null)}
      />
    </>
  )
}
