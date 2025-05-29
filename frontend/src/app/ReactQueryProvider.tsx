"use client"

import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactQueryStreamedHydration } from "@tanstack/react-query-next-experimental"
import React from "react"

export default function ReactQueryProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [queryClient] = React.useState(() => new QueryClient())

  return (
    <QueryClientProvider client={queryClient}>
      <ReactQueryStreamedHydration>{children}</ReactQueryStreamedHydration>
    </QueryClientProvider>
  )
}
