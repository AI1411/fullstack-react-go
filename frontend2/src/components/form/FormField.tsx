import React from 'react'
import { Controller, Control, FieldValues, Path, FieldError } from 'react-hook-form'
import { Input } from '../ui/Input'

interface FormFieldProps<T extends FieldValues> {
  name: Path<T>
  control: Control<T>
  label: string
  error?: FieldError
  type?: string
  placeholder?: string
  options?: { value: string; label: string }[]
  disabled?: boolean
  className?: string
  rows?: number
}

export const FormField = <T extends FieldValues>({
  name,
  control,
  label,
  error,
  type = 'text',
  placeholder,
  options,
  disabled = false,
  className = '',
  rows,
}: FormFieldProps<T>) => {
  return (
    <div className={className}>
      <label htmlFor={name} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      <Controller
        name={name}
        control={control}
        render={({ field }) => {
          if (type === 'textarea') {
            return (
              <textarea
                id={name}
                placeholder={placeholder}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                rows={rows || 3}
                disabled={disabled}
                {...field}
              />
            )
          }
          
          if (type === 'select') {
            return (
              <select
                id={name}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                disabled={disabled}
                {...field}
              >
                <option value="">{placeholder || 'Select an option'}</option>
                {options?.map((option) => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            )
          }
          
          return (
            <Input
              id={name}
              type={type}
              placeholder={placeholder}
              className="mt-1"
              disabled={disabled}
              {...field}
            />
          )
        }}
      />
      {error && <p className="mt-1 text-sm text-red-600">{error.message}</p>}
    </div>
  )
}