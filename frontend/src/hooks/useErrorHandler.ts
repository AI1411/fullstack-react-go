import { useCallback } from "react"
import { AxiosError } from "axios"

export interface ApiError {
  message: string
  status?: number
  code?: string
}

export const useErrorHandler = () => {
  const handleError = useCallback((error: unknown): ApiError => {
    if (error instanceof AxiosError) {
      // Axiosエラーの場合
      if (error.response) {
        // サーバーからのレスポンスがある場合
        return {
          message:
            error.response.data?.message ||
            error.message ||
            "サーバーエラーが発生しました",
          status: error.response.status,
          code: error.response.data?.code || error.code,
        }
      } else if (error.request) {
        // リクエストは送信されたがレスポンスがない場合
        return {
          message: "サーバーに接続できませんでした",
          status: undefined,
          code: "NETWORK_ERROR",
        }
      } else {
        // リクエスト設定でエラーが発生した場合
        return {
          message: error.message || "不明なエラーが発生しました",
          status: undefined,
          code: "REQUEST_ERROR",
        }
      }
    }

    // その他のエラー
    if (error instanceof Error) {
      return {
        message: error.message,
        status: undefined,
        code: "UNKNOWN_ERROR",
      }
    }

    return {
      message: "不明なエラーが発生しました",
      status: undefined,
      code: "UNKNOWN_ERROR",
    }
  }, [])

  return { handleError }
}
