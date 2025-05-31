// ステータスに応じたバッジの色を定義
export const getStatusBadgeClass = (status: string) => {
  switch (status) {
    case "pending":
      return "bg-[#fff7e6] text-[#ff8b00]"
    case "under_review":
      return "bg-[#edf5ff] text-[#0055cc]"
    case "in_progress":
      return "bg-[#f0f2f4] text-[#111418]"
    case "completed":
      return "bg-[#e6fcf5] text-[#00a3bf]"
    default:
      return "bg-[#f0f2f4] text-[#111418]"
  }
}

// ステータスの日本語表示
export const getStatusLabel = (status: string) => {
  switch (status) {
    case "pending":
      return "未着手"
    case "under_review":
      return "審査中"
    case "in_progress":
      return "対応中"
    case "completed":
      return "完了"
    default:
      return status
  }
}

// 日付フォーマット関数
export const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  })
}