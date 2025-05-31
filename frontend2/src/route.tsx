import {
  createRootRoute,
  createRoute,
  createRouter,
  Outlet,
} from "@tanstack/react-router"
import { Layout } from "./components/layouts/Header"
import { Home } from "./routes/Home"
import { Disasters } from "./routes/Disasters"
import { DisasterDetail } from "./routes/DisasterDetail"
import { Application } from "./routes/Application"
import { Hello } from "./App"

// Root route with layout
const rootRoute = createRootRoute({
  component: () => <Layout><Outlet /></Layout>,
})

// Home route
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Home,
})

// Disasters route
const disastersRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/disasters",
  component: Disasters,
})

// Disaster detail route
const disasterDetailRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/disasters/$disasterId",
  component: DisasterDetail,
})

// Application route
const applicationRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/application",
  component: Application,
})

// Legacy hello route
const helloRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/$helloId",
  component: () => <Hello />,
})

const routeTree = rootRoute.addChildren([
  indexRoute,
  disastersRoute,
  disasterDetailRoute,
  applicationRoute,
  helloRoute,
])

export const router = createRouter({ routeTree })

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router
  }
}
