import React from 'react'
import { useForm, Controller } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { Button } from '../ui/Button'
import { Input } from '../ui/Input'
import { usePrefectures } from '../../api/hooks/usePrefectures'

// Validation schema for disaster form
const disasterSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  description: z.string().min(1, 'Description is required'),
  disasterType: z.string().min(1, 'Disaster type is required'),
  prefectureId: z.string().min(1, 'Prefecture is required'),
  startDate: z.string().min(1, 'Start date is required'),
  endDate: z.string().optional(),
})

type DisasterFormData = z.infer<typeof disasterSchema>

interface DisasterFormProps {
  initialData?: Partial<DisasterFormData>
  onSubmit: (data: DisasterFormData) => void
  isSubmitting?: boolean
}

export const DisasterForm: React.FC<DisasterFormProps> = ({
  initialData,
  onSubmit,
  isSubmitting = false,
}) => {
  const { control, handleSubmit, formState: { errors } } = useForm<DisasterFormData>({
    resolver: zodResolver(disasterSchema),
    defaultValues: initialData || {
      name: '',
      description: '',
      disasterType: '',
      prefectureId: '',
      startDate: '',
      endDate: '',
    },
  })

  const { data: prefecturesData } = usePrefectures()

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label htmlFor="name" className="block text-sm font-medium text-gray-700">
          Disaster Name
        </label>
        <Controller
          name="name"
          control={control}
          render={({ field }) => (
            <Input
              id="name"
              placeholder="Enter disaster name"
              className="mt-1"
              {...field}
            />
          )}
        />
        {errors.name && (
          <p className="mt-1 text-sm text-red-600">{errors.name.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="description" className="block text-sm font-medium text-gray-700">
          Description
        </label>
        <Controller
          name="description"
          control={control}
          render={({ field }) => (
            <textarea
              id="description"
              placeholder="Enter disaster description"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              rows={3}
              {...field}
            />
          )}
        />
        {errors.description && (
          <p className="mt-1 text-sm text-red-600">{errors.description.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="disasterType" className="block text-sm font-medium text-gray-700">
          Disaster Type
        </label>
        <Controller
          name="disasterType"
          control={control}
          render={({ field }) => (
            <select
              id="disasterType"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              {...field}
            >
              <option value="">Select disaster type</option>
              <option value="EARTHQUAKE">Earthquake</option>
              <option value="FLOOD">Flood</option>
              <option value="TYPHOON">Typhoon</option>
              <option value="LANDSLIDE">Landslide</option>
              <option value="OTHER">Other</option>
            </select>
          )}
        />
        {errors.disasterType && (
          <p className="mt-1 text-sm text-red-600">{errors.disasterType.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="prefectureId" className="block text-sm font-medium text-gray-700">
          Prefecture
        </label>
        <Controller
          name="prefectureId"
          control={control}
          render={({ field }) => (
            <select
              id="prefectureId"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              {...field}
            >
              <option value="">Select prefecture</option>
              {prefecturesData?.data.map((prefecture) => (
                <option key={prefecture.id} value={prefecture.id}>
                  {prefecture.name}
                </option>
              ))}
            </select>
          )}
        />
        {errors.prefectureId && (
          <p className="mt-1 text-sm text-red-600">{errors.prefectureId.message}</p>
        )}
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div>
          <label htmlFor="startDate" className="block text-sm font-medium text-gray-700">
            Start Date
          </label>
          <Controller
            name="startDate"
            control={control}
            render={({ field }) => (
              <Input
                id="startDate"
                type="date"
                className="mt-1"
                {...field}
              />
            )}
          />
          {errors.startDate && (
            <p className="mt-1 text-sm text-red-600">{errors.startDate.message}</p>
          )}
        </div>

        <div>
          <label htmlFor="endDate" className="block text-sm font-medium text-gray-700">
            End Date (Optional)
          </label>
          <Controller
            name="endDate"
            control={control}
            render={({ field }) => (
              <Input
                id="endDate"
                type="date"
                className="mt-1"
                {...field}
              />
            )}
          />
        </div>
      </div>

      <div className="flex justify-end">
        <Button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Saving...' : initialData ? 'Update Disaster' : 'Create Disaster'}
        </Button>
      </div>
    </form>
  )
}