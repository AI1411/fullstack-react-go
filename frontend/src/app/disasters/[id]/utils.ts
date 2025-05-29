/**
 * Utility functions for formatting data in the disaster detail page
 */

/**
 * Format a date string to a localized date and time string
 */
export const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  })
}

/**
 * Format a number to a localized string with commas
 */
export const formatNumber = (num: number | null | undefined): string => {
  if (num === null || num === undefined) return "-"
  return num.toLocaleString("ja-JP")
}

/**
 * Format an area value to a localized string with unit
 */
export const formatArea = (area: number | null | undefined): string => {
  if (area === null || area === undefined) return "-"
  return `${area.toLocaleString("ja-JP")} ha`
}

/**
 * Format an amount value to a localized currency string
 */
export const formatAmount = (amount: number | null | undefined): string => {
  if (amount === null || amount === undefined) return "-"
  return `¥${amount.toLocaleString("ja-JP")}`
}
