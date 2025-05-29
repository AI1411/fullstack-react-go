"use client"

import axios from "axios"
import React from "react"

export default function AxiosProvider({
  children,
}: {
  children: React.ReactNode
}) {
  axios.defaults.baseURL = "http://localhost:8080"

  return children
}
