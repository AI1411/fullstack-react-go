import React, { useState } from 'react'
import { Button } from '../ui/Button'
import { Input } from '../ui/Input'

interface SearchFormProps {
  onSearch: (searchParams: Record<string, string>) => void
  fields: {
    id: string
    label: string
    placeholder?: string
    type?: string
    options?: { value: string; label: string }[]
  }[]
  initialValues?: Record<string, string>
  submitLabel?: string
  className?: string
}

export const SearchForm: React.FC<SearchFormProps> = ({
  onSearch,
  fields,
  initialValues = {},
  submitLabel = 'Search',
  className = '',
}) => {
  const [formValues, setFormValues] = useState<Record<string, string>>(initialValues)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target
    setFormValues((prev) => ({ ...prev, [name]: value }))
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Filter out empty values
    const filteredValues = Object.entries(formValues).reduce(
      (acc, [key, value]) => {
        if (value.trim() !== '') {
          acc[key] = value
        }
        return acc
      },
      {} as Record<string, string>
    )
    onSearch(filteredValues)
  }

  const handleReset = () => {
    setFormValues({})
    onSearch({})
  }

  return (
    <form onSubmit={handleSubmit} className={`space-y-4 ${className}`}>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {fields.map((field) => (
          <div key={field.id}>
            <label htmlFor={field.id} className="block text-sm font-medium text-gray-700">
              {field.label}
            </label>
            {field.options ? (
              <select
                id={field.id}
                name={field.id}
                value={formValues[field.id] || ''}
                onChange={handleChange}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              >
                <option value="">All</option>
                {field.options.map((option) => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            ) : (
              <Input
                id={field.id}
                name={field.id}
                type={field.type || 'text'}
                placeholder={field.placeholder}
                value={formValues[field.id] || ''}
                onChange={handleChange}
                className="mt-1"
              />
            )}
          </div>
        ))}
      </div>

      <div className="flex justify-end space-x-2">
        <Button
          type="button"
          variant="outline"
          onClick={handleReset}
        >
          Reset
        </Button>
        <Button type="submit">{submitLabel}</Button>
      </div>
    </form>
  )
}