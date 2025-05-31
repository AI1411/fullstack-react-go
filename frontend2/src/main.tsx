import { RouterProvider } from "@tanstack/react-router"
import { TanStackRouterDevtools } from "@tanstack/router-devtools"
import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { router } from "./route"

function assertElement(
  element: HTMLElement | null
): asserts element is HTMLElement {
  if (!element) {
    throw new Error("Root element not found")
  }
}

const rootElement = document.getElementById("root")
assertElement(rootElement)

createRoot(rootElement).render(
  <StrictMode>
    <RouterProvider router={router} />
    <TanStackRouterDevtools router={router} />
  </StrictMode>
)
