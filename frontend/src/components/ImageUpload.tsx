import { useState } from "react"

interface UploadResponse {
  filename: string
  path: string
  size: number
  uploaded: string
}

interface FileWithPreview extends File {
  preview?: string;
}

interface ImageUploadProps {
  onUpload: (formData: FormData) => Promise<UploadResponse>
  title?: string
}

export const ImageUpload = ({ onUpload, title }: ImageUploadProps) => {
  // State for the selected files
  const [selectedFiles, setSelectedFiles] = useState<FileWithPreview[]>([])

  // State for the upload status
  const [isUploading, setIsUploading] = useState(false)

  // State for the upload responses
  const [uploadResponses, setUploadResponses] = useState<UploadResponse[]>([])

  // State for error messages
  const [error, setError] = useState<string | null>(null)

  // State to track which files have been uploaded
  const [uploadedFiles, setUploadedFiles] = useState<Set<string>>(new Set())

  // Handle file selection
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files
    if (files && files.length > 0) {
      const fileArray = Array.from(files).map(file => {
        return Object.assign(file, {
          preview: URL.createObjectURL(file)
        });
      });
      // Append new files to existing selection instead of replacing
      setSelectedFiles(prevFiles => [...prevFiles, ...fileArray]);
      setError(null);

      // Reset the file input to allow selecting the same file again
      event.target.value = "";
    }
  }

  // Handle removing a file from selection
  const handleRemoveFile = (index: number) => {
    setSelectedFiles(prevFiles => {
      const newFiles = [...prevFiles];
      // Revoke the object URL to avoid memory leaks
      if (newFiles[index].preview) {
        URL.revokeObjectURL(newFiles[index].preview as string);
      }
      newFiles.splice(index, 1);
      return newFiles;
    });
  }

  // Handle clearing all files
  const handleClearAll = () => {
    // Revoke all object URLs to avoid memory leaks
    for (const file of selectedFiles) {
      if (file.preview) {
        URL.revokeObjectURL(file.preview as string);
      }
    }
    setSelectedFiles([]);
    setUploadedFiles(new Set());
    setUploadResponses([]);
    setError(null);
  }

  // Check if a file has been uploaded
  const isFileUploaded = (file: FileWithPreview) => {
    const fileId = `${file.name}-${file.size}-${file.lastModified}`;
    return uploadedFiles.has(fileId);
  }

  // Handle file upload
  const handleUpload = async () => {
    if (selectedFiles.length === 0) {
      setError("ファイルを選択してください")
      return
    }

    // Check if all files are images
    const nonImageFiles = selectedFiles.filter(file => !file.type.startsWith("image/"))
    if (nonImageFiles.length > 0) {
      setError("すべてのファイルが画像である必要があります")
      return
    }

    // Filter out files that have already been uploaded
    const filesToUpload = selectedFiles.filter(file => {
      const fileId = `${file.name}-${file.size}-${file.lastModified}`;
      return !uploadedFiles.has(fileId);
    });

    if (filesToUpload.length === 0) {
      setError("すべてのファイルはすでにアップロードされています")
      return
    }

    setIsUploading(true)
    setError(null)
    setUploadResponses([])

    try {
      const responses: UploadResponse[] = []
      const newUploadedFiles = new Set(uploadedFiles);

      // Upload each file
      for (const file of filesToUpload) {
        const formData = new FormData()
        formData.append("image", file)

        const data = await onUpload(formData)
        responses.push(data)

        // Mark this file as uploaded
        const fileId = `${file.name}-${file.size}-${file.lastModified}`;
        newUploadedFiles.add(fileId);
      }

      setUploadResponses(responses)
      setUploadedFiles(newUploadedFiles);

      // Don't clear selected files, so users can add more
      // Don't revoke object URLs, as we're still displaying the images
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
      {title && (
        <h2 className="text-[#111418] text-xl font-bold mb-4">
          {title}
        </h2>
      )}

      <div className="mb-4">
        <div className="flex justify-between items-center mb-2">
          <label
            htmlFor="image-upload"
            className="block text-[#637588] text-sm font-medium"
          >
            {selectedFiles.length > 0 ? '画像をさらに追加する' : '画像ファイルを選択（複数可）'}
          </label>
          {selectedFiles.length > 0 && (
            <button
              type="button"
              onClick={handleClearAll}
              className="text-sm text-red-500 hover:text-red-700"
            >
              すべてクリア
            </button>
          )}
        </div>
        <input
          type="file"
          id="image-upload"
          accept="image/*"
          multiple
          onChange={handleFileChange}
          className="block w-full text-sm text-[#637588] file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-medium file:bg-[#f0f2f5] file:text-[#637588] hover:file:bg-[#e4e7eb]"
        />
        {selectedFiles.length > 0 && (
          <div className="mt-2">
            <p className="text-sm text-[#637588] font-medium">
              選択されたファイル: {selectedFiles.length}件
            </p>
            <div className="mt-2 flex flex-wrap gap-2">
              {selectedFiles.map((file) => (
                <div key={`${file.name}-${file.size}-${file.lastModified}`} className="text-sm text-[#637588]">
                  {file.name} ({Math.round(file.size / 1024)} KB)
                </div>
              ))}
            </div>

            {/* Image previews */}
            <div className="mt-4 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
              {selectedFiles.map((file, index) => (
                file.preview && (
                  <div key={`preview-${file.name}-${file.size}-${file.lastModified}`} 
                       className={`relative rounded-lg overflow-hidden border ${isFileUploaded(file) ? 'border-green-500' : 'border-gray-200'}`}>
                    <img 
                      src={file.preview} 
                      alt={`Preview of ${file.name}`}
                      className="w-full h-32 object-cover"
                    />
                    <div className="absolute bottom-0 left-0 right-0 bg-black bg-opacity-50 text-white text-xs p-1 truncate">
                      {file.name}
                    </div>
                    {isFileUploaded(file) && (
                      <div className="absolute top-1 left-1 bg-green-500 text-white text-xs px-2 py-1 rounded-md">
                        アップロード済み
                      </div>
                    )}
                    <button 
                      type="button"
                      onClick={() => handleRemoveFile(index)}
                      className="absolute top-1 right-1 bg-red-500 text-white rounded-full w-6 h-6 flex items-center justify-center hover:bg-red-600"
                      title="削除"
                    >
                      ×
                    </button>
                  </div>
                )
              ))}
            </div>
          </div>
        )}
      </div>

      <div className="flex items-center gap-4">
        <button
          type="button"
          onClick={handleUpload}
          disabled={selectedFiles.length === 0 || isUploading}
          className={`px-4 py-2 rounded-lg ${
            selectedFiles.length === 0 || isUploading
              ? "bg-gray-300 text-gray-500 cursor-not-allowed"
              : "bg-[#197fe5] text-white hover:bg-[#1668c3]"
          }`}
        >
          {isUploading ? "アップロード中..." : "未アップロードの画像をアップロード"}
        </button>

        {error && <p className="text-red-500 text-sm">{error}</p>}
      </div>

      {uploadResponses.length > 0 && (
        <div className="mt-4 p-4 bg-green-50 border border-green-200 rounded-lg">
          <p className="text-green-700 font-medium">アップロード成功!</p>
          <div className="mt-2">
            {uploadResponses.map((response) => (
              <div key={`${response.filename}-${response.size}-${response.uploaded}`} className="mb-2 pb-2 border-b border-green-100 last:border-b-0 last:mb-0 last:pb-0">
                <p className="text-sm text-green-600">
                  ファイル名: {response.filename}
                </p>
                <p className="text-sm text-green-600">
                  サイズ: {Math.round(response.size / 1024)} KB
                </p>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  )
}