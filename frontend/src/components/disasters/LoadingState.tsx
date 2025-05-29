"use client"

import Footer from "@/components/layout/footer/page"
import Header from "@/components/layout/header/page"
import { LoadingSpinner } from "@/components/ui/loading"

export default function LoadingState() {
  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex items-center justify-center min-h-[400px]">
            <LoadingSpinner size="lg" message="災害情報を読み込み中..." />
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}