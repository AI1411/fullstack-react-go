// Simple API utility functions for demonstration purposes

// Simulate a fetch request with a delay
const simulateFetch = <T>(data: T, delay: number = 1000): Promise<T> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(data)
    }, delay)
  })
}

// Example API endpoints
export const fetchTodos = async (): Promise<
  { id: number; title: string; completed: boolean }[]
> => {
  // Simulate API call
  return simulateFetch([
    { id: 1, title: "Learn React", completed: true },
    { id: 2, title: "Learn Next.js", completed: true },
    { id: 3, title: "Learn Zustand", completed: false },
    { id: 4, title: "Learn TanStack Query", completed: false },
  ])
}
