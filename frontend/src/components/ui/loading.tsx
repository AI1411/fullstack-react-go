import React from 'react'

interface LoadingSpinnerProps {
  size?: 'sm' | 'md' | 'lg'
  message?: string
  className?: string
}

export const LoadingSpinner: React.FC<LoadingSpinnerProps> = ({
  size = 'md',
  message,
  className = ''
}) => {
  const sizeClasses = {
    sm: 'h-4 w-4',
    md: 'h-8 w-8',
    lg: 'h-12 w-12'
  }

  return (
    <div className={`flex items-center justify-center ${className}`}>
      <div
        className={`animate-spin rounded-full border-b-2 border-[#111418] ${sizeClasses[size]}`}
      />
      {message && (
        <span className="ml-2 text-[#637588] text-sm">{message}</span>
      )}
    </div>
  )
}

interface LoadingTableProps {
  rows?: number
  columns?: number
}

export const LoadingTable: React.FC<LoadingTableProps> = ({
  rows = 5,
  columns = 4
}) => {
  return (
    <div className="animate-pulse">
      {/* ヘッダー */}
      <div className="flex space-x-4 mb-4">
        {Array.from({ length: columns }, (_, i) => (
          <div key={i} className="bg-gray-200 h-6 rounded flex-1" />
        ))}
      </div>
      {/* 行 */}
      <div className="space-y-2">
        {Array.from({ length: rows }, (_, i) => (
          <div key={i} className="flex space-x-4">
            {Array.from({ length: columns }, (_, j) => (
              <div key={j} className="bg-gray-200 h-16 rounded flex-1" />
            ))}
          </div>
        ))}
      </div>
    </div>
  )
}

interface LoadingCardProps {
  count?: number
}

export const LoadingCard: React.FC<LoadingCardProps> = ({ count = 3 }) => {
  return (
    <div className="flex flex-wrap gap-4">
      {Array.from({ length: count }, (_, i) => (
        <div key={i} className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
          <div className="animate-pulse">
            <div className="bg-gray-200 h-4 rounded mb-2" />
            <div className="bg-gray-200 h-8 rounded" />
          </div>
        </div>
      ))}
    </div>
  )
}
