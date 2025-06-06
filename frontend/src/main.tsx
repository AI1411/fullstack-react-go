import { RouterProvider } from "@tanstack/react-router"
import { TanStackRouterDevtools } from "@tanstack/router-devtools"
import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import AxiosProvider from "./providers/AxiosProvider"
import ReactQueryProvider from "./providers/ReactQueryProvider"
import { router } from "./route"
import "./index.css"
import { GoogleAPIProvider } from "./providers/GoogleAPIprovider.tsx"

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
    <GoogleAPIProvider>
      <AxiosProvider>
        <ReactQueryProvider>
          <RouterProvider router={router} />
          <TanStackRouterDevtools router={router} />
        </ReactQueryProvider>
      </AxiosProvider>
    </GoogleAPIProvider>
  </StrictMode>
)
