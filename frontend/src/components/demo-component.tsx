'use client';

import { useQuery } from '@tanstack/react-query';
import { fetchTodos } from '@/lib/api';
import { useCounterStore } from '@/lib/store';
import { Button } from '@/components/ui/button';

export function DemoComponent() {
  // Use zustand store
  const { count, increment, decrement, reset } = useCounterStore();

  // Use TanStack Query
  const { data: todos, isLoading, isError } = useQuery({
    queryKey: ['todos'],
    queryFn: fetchTodos,
  });

  return (
    <div className="space-y-8">
      <div className="p-6 bg-gray-100 dark:bg-gray-800 rounded-lg">
        <h2 className="text-xl font-bold mb-4">Zustand State Management</h2>
        <div className="flex items-center gap-4 mb-4">
          <Button onClick={decrement} variant="outline">-</Button>
          <span className="text-2xl font-bold">{count}</span>
          <Button onClick={increment} variant="outline">+</Button>
        </div>
        <Button onClick={reset} variant="secondary">Reset</Button>
      </div>

      <div className="p-6 bg-gray-100 dark:bg-gray-800 rounded-lg">
        <h2 className="text-xl font-bold mb-4">TanStack Query Data Fetching</h2>
        {isLoading ? (
          <p>Loading todos...</p>
        ) : isError ? (
          <p className="text-red-500">Error loading todos</p>
        ) : (
          <ul className="space-y-2">
            {todos?.map((todo) => (
              <li key={todo.id} className="flex items-center gap-2">
                <span className={todo.completed ? "line-through" : ""}>
                  {todo.title}
                </span>
                <span className="text-xs px-2 py-1 rounded-full bg-gray-200 dark:bg-gray-700">
                  {todo.completed ? "Completed" : "Pending"}
                </span>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}