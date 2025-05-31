import React from 'react'
import { useForm, Controller } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { Button } from '../ui/Button'
import { Input } from '../ui/Input'
import { useDisasters } from '../../api/hooks/useDisasters'

// Validation schema for support application form
const supportApplicationSchema = z.object({
  applicantName: z.string().min(1, 'Applicant name is required'),
  applicantContact: z.string().min(1, 'Contact information is required'),
  disasterId: z.string().min(1, 'Disaster is required'),
  supportType: z.string().min(1, 'Support type is required'),
  description: z.string().min(1, 'Description is required'),
  requestedAmount: z.string().optional(),
  address: z.string().min(1, 'Address is required'),
})

type SupportApplicationFormData = z.infer<typeof supportApplicationSchema>

interface SupportApplicationFormProps {
  initialData?: Partial<SupportApplicationFormData>
  onSubmit: (data: SupportApplicationFormData) => void
  isSubmitting?: boolean
  disasterId?: string
}

export const SupportApplicationForm: React.FC<SupportApplicationFormProps> = ({
  initialData,
  onSubmit,
  isSubmitting = false,
  disasterId,
}) => {
  const { control, handleSubmit, formState: { errors } } = useForm<SupportApplicationFormData>({
    resolver: zodResolver(supportApplicationSchema),
    defaultValues: initialData || {
      applicantName: '',
      applicantContact: '',
      disasterId: disasterId || '',
      supportType: '',
      description: '',
      requestedAmount: '',
      address: '',
    },
  })

  const { data: disastersData } = useDisasters()

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label htmlFor="applicantName" className="block text-sm font-medium text-gray-700">
          Applicant Name
        </label>
        <Controller
          name="applicantName"
          control={control}
          render={({ field }) => (
            <Input
              id="applicantName"
              placeholder="Enter your full name"
              className="mt-1"
              {...field}
            />
          )}
        />
        {errors.applicantName && (
          <p className="mt-1 text-sm text-red-600">{errors.applicantName.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="applicantContact" className="block text-sm font-medium text-gray-700">
          Contact Information
        </label>
        <Controller
          name="applicantContact"
          control={control}
          render={({ field }) => (
            <Input
              id="applicantContact"
              placeholder="Enter phone number or email"
              className="mt-1"
              {...field}
            />
          )}
        />
        {errors.applicantContact && (
          <p className="mt-1 text-sm text-red-600">{errors.applicantContact.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="address" className="block text-sm font-medium text-gray-700">
          Address
        </label>
        <Controller
          name="address"
          control={control}
          render={({ field }) => (
            <Input
              id="address"
              placeholder="Enter your address"
              className="mt-1"
              {...field}
            />
          )}
        />
        {errors.address && (
          <p className="mt-1 text-sm text-red-600">{errors.address.message}</p>
        )}
      </div>

      {!disasterId && (
        <div>
          <label htmlFor="disasterId" className="block text-sm font-medium text-gray-700">
            Disaster
          </label>
          <Controller
            name="disasterId"
            control={control}
            render={({ field }) => (
              <select
                id="disasterId"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                {...field}
              >
                <option value="">Select disaster</option>
                {disastersData?.data.map((disaster) => (
                  <option key={disaster.id} value={disaster.id}>
                    {disaster.name}
                  </option>
                ))}
              </select>
            )}
          />
          {errors.disasterId && (
            <p className="mt-1 text-sm text-red-600">{errors.disasterId.message}</p>
          )}
        </div>
      )}

      <div>
        <label htmlFor="supportType" className="block text-sm font-medium text-gray-700">
          Support Type
        </label>
        <Controller
          name="supportType"
          control={control}
          render={({ field }) => (
            <select
              id="supportType"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              {...field}
            >
              <option value="">Select support type</option>
              <option value="FINANCIAL">Financial Support</option>
              <option value="MATERIAL">Material Support</option>
              <option value="SHELTER">Shelter</option>
              <option value="MEDICAL">Medical Support</option>
              <option value="OTHER">Other</option>
            </select>
          )}
        />
        {errors.supportType && (
          <p className="mt-1 text-sm text-red-600">{errors.supportType.message}</p>
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
              placeholder="Describe your support needs"
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
        <label htmlFor="requestedAmount" className="block text-sm font-medium text-gray-700">
          Requested Amount (if applicable)
        </label>
        <Controller
          name="requestedAmount"
          control={control}
          render={({ field }) => (
            <Input
              id="requestedAmount"
              type="number"
              placeholder="Enter amount"
              className="mt-1"
              {...field}
            />
          )}
        />
      </div>

      <div className="flex justify-end">
        <Button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Submitting...' : initialData ? 'Update Application' : 'Submit Application'}
        </Button>
      </div>
    </form>
  )
}