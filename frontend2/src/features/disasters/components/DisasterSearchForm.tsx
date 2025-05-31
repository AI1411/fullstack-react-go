import { useState } from "react"
import type { DisasterSearchParams } from "../types"

interface DisasterSearchFormProps {
  onSearch: (params: DisasterSearchParams) => void
}

export const DisasterSearchForm = ({ onSearch }: DisasterSearchFormProps) => {
  // 検索条件の状態管理
  const [searchName, setSearchName] = useState("")
  const [searchDisasterType, setSearchDisasterType] = useState("")
  const [searchStatus, setSearchStatus] = useState("")
  const [searchDateFrom, setSearchDateFrom] = useState("")
  const [searchDateTo, setSearchDateTo] = useState("")

  // 検索実行関数
  const handleSearch = () => {
    onSearch({
      name: searchName,
      disaster_type: searchDisasterType,
      status: searchStatus,
      date_from: searchDateFrom,
      date_to: searchDateTo,
    })
  }

  // 検索条件クリア関数
  const handleClearSearch = () => {
    setSearchName("")
    setSearchDisasterType("")
    setSearchStatus("")
    setSearchDateFrom("")
    setSearchDateTo("")
    onSearch({
      name: "",
      disaster_type: "",
      status: "",
      date_from: "",
      date_to: "",
    })
  }

  return (
    <div className="mx-4 mb-4 p-4 bg-white rounded-lg border border-[#dce0e5]">
      <div className="flex items-center justify-between mb-3">
        <h3 className="text-lg font-medium">検索条件</h3>
        <button
          type="button"
          onClick={handleClearSearch}
          className="inline-flex items-center px-3 py-1.5 text-sm font-medium text-[#637588] hover:text-[#111418] transition-colors"
        >
          クリア
        </button>
      </div>

      {/* 1行目: 発生日の期間指定 */}
      <div className="flex gap-4 mb-4">
        <div className="flex items-center gap-2 flex-1">
          <div className="flex-1 max-w-[200px]">
            <label
              htmlFor="date_from"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              発生日（開始）
            </label>
            <input
              type="date"
              id="date_from"
              value={searchDateFrom}
              onChange={(e) => setSearchDateFrom(e.target.value)}
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
            />
          </div>
          <div className="flex items-center pt-6">
            <span className="text-[#637588] text-sm">〜</span>
          </div>
          <div className="flex-1 max-w-[200px]">
            <label
              htmlFor="date_to"
              className="block text-sm font-medium text-[#111418] mb-1"
            >
              発生日（終了）
            </label>
            <input
              type="date"
              id="date_to"
              value={searchDateTo}
              onChange={(e) => setSearchDateTo(e.target.value)}
              className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
            />
          </div>
        </div>
      </div>

      {/* 2行目: その他の検索条件 */}
      <div className="flex gap-4 mb-4">
        {/* 災害名 */}
        <div className="flex-1 max-w-[350px]">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-[#111418] mb-1"
          >
            災害名
          </label>
          <input
            type="text"
            id="name"
            value={searchName}
            onChange={(e) => setSearchName(e.target.value)}
            className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm"
            placeholder="災害名で検索"
          />
        </div>

        {/* 災害種別 */}
        <div className="flex-1 max-w-[200px]">
          <label
            htmlFor="disaster_type"
            className="block text-sm font-medium text-[#111418] mb-1"
          >
            災害種別
          </label>
          <select
            id="disaster_type"
            value={searchDisasterType}
            onChange={(e) => setSearchDisasterType(e.target.value)}
            className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm bg-white"
          >
            <option value="">すべて</option>
            <option value="洪水">洪水</option>
            <option value="地滑り">地滑り</option>
            <option value="雹害">雹害</option>
            <option value="干ばつ">干ばつ</option>
            <option value="風害">風害</option>
            <option value="地震">地震</option>
            <option value="霜害">霜害</option>
            <option value="病害虫">病害虫</option>
          </select>
        </div>

        {/* ステータス */}
        <div className="flex-1 max-w-[150px]">
          <label
            htmlFor="status"
            className="block text-sm font-medium text-[#111418] mb-1"
          >
            ステータス
          </label>
          <select
            id="status"
            value={searchStatus}
            onChange={(e) => setSearchStatus(e.target.value)}
            className="w-full px-3 py-2 border border-[#dce0e5] rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-[#197fe5] focus:border-[#197fe5] text-sm bg-white"
          >
            <option value="">すべて</option>
            <option value="pending">未着手</option>
            <option value="under_review">審査中</option>
            <option value="in_progress">対応中</option>
            <option value="completed">完了</option>
          </select>
        </div>

        {/* 検索ボタン */}
        <div className="flex items-end max-w-[120px]">
          <button
            type="button"
            onClick={handleSearch}
            className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-lg shadow-sm text-sm font-medium text-white bg-[#197fe5] hover:bg-[#1565c0] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#197fe5] transition-colors"
          >
            検索
          </button>
        </div>
      </div>
    </div>
  )
}
