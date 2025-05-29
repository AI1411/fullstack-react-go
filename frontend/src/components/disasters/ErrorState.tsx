"use client"

import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import { ErrorDisplay } from "@/components/ui/error-display"

type ErrorStateProps = {
  message: string
  onRetry: () => void
}

export default function ErrorState({ message, onRetry }: ErrorStateProps) {
  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <ErrorDisplay
            title="災害情報の取得に失敗しました"
            message={message}
            onRetry={onRetry}
            className="mt-8"
          />
        </div>
      </main>
      <Footer />
    </div>
  )
}