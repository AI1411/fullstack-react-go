import type { ReactNode } from "react"
import { createContext, useContext } from "react"

interface GoogleAPIContextType {
  apiKey: string
}

const GoogleAPIContext = createContext<GoogleAPIContextType | undefined>(
  undefined
)

interface GoogleAPIProviderProps {
  children: ReactNode
  apiKey?: string
}

export const GoogleAPIProvider = ({
  children,
  apiKey,
}: GoogleAPIProviderProps) => {
  // 環境変数からAPIキーを取得（デフォルト）
  const resolvedApiKey =
    apiKey || import.meta.env.VITE_GOOGLE_MAPS_API_KEY || ""

  if (!resolvedApiKey) {
    console.warn("Google Maps API キーが設定されていません")
  }

  return (
    <GoogleAPIContext.Provider value={{ apiKey: resolvedApiKey }}>
      {children}
    </GoogleAPIContext.Provider>
  )
}

export const useGoogleAPI = () => {
  const context = useContext(GoogleAPIContext)
  if (context === undefined) {
    throw new Error("useGoogleAPI は GoogleAPIProvider 内で使用してください")
  }
  return context
}
