import { Link } from "@tanstack/react-router"
import { useEffect, useState } from "react"
import {
  Bar,
  BarChart,
  CartesianGrid,
  Cell,
  Legend,
  Line,
  LineChart,
  Pie,
  PieChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts"
import { useListDisasters } from "../api/generated/client"
import type { HandlerListDisastersResponse } from "../api/generated/model"

// 注意: このコンポーネントを使用するには、以下のコマンドでRechartsをインストールする必要があります:
// pnpm add recharts

export const Home = () => {
  // 災害情報を取得
  const {
    data: disastersResponse,
    isLoading,
    error,
  } = useListDisasters<{ data: HandlerListDisastersResponse }>({
    query: {
      staleTime: 5 * 60 * 1000, // 5分間キャッシュ
    },
  })

  // 集計データの状態
  const [summaryStats, setSummaryStats] = useState({
    totalDisasters: 0,
    totalDamageAmount: 0,
    totalAffectedArea: 0,
    byType: [],
    byMonth: [],
    byImpactLevel: [],
  })

  // データが変更されたら集計データを更新
  useEffect(() => {
    if (disastersResponse?.data?.disasters) {
      const disasters = disastersResponse.data.disasters

      // 基本統計情報の計算
      const totalDisasters = disasters.length
      const totalDamageAmount = disasters.reduce(
        (sum, d) => sum + (d.estimated_damage_amount || 0),
        0
      )
      const totalAffectedArea = disasters.reduce(
        (sum, d) => sum + (d.affected_area_size || 0),
        0
      )

      // 災害タイプ別の集計
      const typeCount = {}
      for (const d of disasters) {
        const type = d.disaster_type || "未分類"
        typeCount[type] = (typeCount[type] || 0) + 1
      }
      const byType = Object.entries(typeCount).map(([name, value]) => ({
        name,
        value,
      }))

      // 月別の災害発生数
      const monthCount = {}
      for (const d of disasters) {
        if (d.occurred_at) {
          const date = new Date(d.occurred_at)
          const monthKey = `${date.getFullYear()}-${date.getMonth() + 1}`
          monthCount[monthKey] = (monthCount[monthKey] || 0) + 1
        }
      }
      const byMonth = Object.entries(monthCount)
        .map(([month, count]) => ({ month, count }))
        .sort((a, b) => a.month.localeCompare(b.month))

      // 影響レベル別の集計
      const impactCount = {}
      for (const d of disasters) {
        const level = d.impact_level || "未分類"
        impactCount[level] = (impactCount[level] || 0) + 1
      }
      const byImpactLevel = Object.entries(impactCount).map(
        ([name, value]) => ({ name, value })
      )

      setSummaryStats({
        totalDisasters,
        totalDamageAmount,
        totalAffectedArea,
        byType,
        byMonth,
        byImpactLevel,
      })
    }
  }, [disastersResponse])

  // 円グラフの色
  const COLORS = [
    "#0088FE",
    "#00C49F",
    "#FFBB28",
    "#FF8042",
    "#8884d8",
    "#82ca9d",
  ]

  if (isLoading) {
    return <div>読み込み中</div>
  }

  if (error) {
    return <div>データの取得に失敗しました</div>
  }

  // 金額のフォーマット
  const formatCurrency = (value) => {
    return new Intl.NumberFormat("ja-JP", {
      style: "currency",
      currency: "JPY",
      maximumFractionDigits: 0,
    }).format(value)
  }

  // 面積のフォーマット
  const formatArea = (value) => {
    return `${value.toLocaleString()} ha`
  }

  return (
    <>
      <div className="flex flex-wrap justify-between gap-3 p-4">
        <div className="flex min-w-72 flex-col gap-3">
          <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
            ダッシュボード
          </p>
          <p className="text-[#637588] text-sm font-normal leading-normal">
            お疲れさまです！システムの概要と最新の災害情報をご確認ください。
          </p>
        </div>
      </div>

      {/* 統計カード */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 p-4">
        <div className="bg-white rounded-lg shadow p-6 border-l-4 border-blue-500">
          <h3 className="text-gray-500 text-sm font-medium">災害総数</h3>
          <p className="text-3xl font-bold">{summaryStats.totalDisasters}</p>
          <p className="text-gray-600 text-xs mt-2">登録されている災害の総数</p>
        </div>
        <div className="bg-white rounded-lg shadow p-6 border-l-4 border-green-500">
          <h3 className="text-gray-500 text-sm font-medium">被害総額</h3>
          <p className="text-3xl font-bold">
            {formatCurrency(summaryStats.totalDamageAmount)}
          </p>
          <p className="text-gray-600 text-xs mt-2">推定被害金額の合計</p>
        </div>
        <div className="bg-white rounded-lg shadow p-6 border-l-4 border-yellow-500">
          <h3 className="text-gray-500 text-sm font-medium">被害面積</h3>
          <p className="text-3xl font-bold">
            {formatArea(summaryStats.totalAffectedArea)}
          </p>
          <p className="text-gray-600 text-xs mt-2">影響を受けた地域の総面積</p>
        </div>
      </div>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        災害タイプ別分布
      </h2>

      {/* 災害タイプ別の円グラフ */}
      <div className="bg-white rounded-lg shadow p-4 mx-4 mb-6">
        <ResponsiveContainer width="100%" height={300}>
          <PieChart>
            <Pie
              data={summaryStats.byType}
              cx="50%"
              cy="50%"
              labelLine={true}
              label={({ name, percent }) =>
                `${name} (${(percent * 100).toFixed(0)}%)`
              }
              outerRadius={100}
              fill="#8884d8"
              dataKey="value"
            >
              {summaryStats.byType.map((entry, index) => (
                <Cell
                  key={`cell-${entry.name}`}
                  fill={COLORS[index % COLORS.length]}
                />
              ))}
            </Pie>
            <Tooltip formatter={(value) => [`${value}件`, "件数"]} />
            <Legend />
          </PieChart>
        </ResponsiveContainer>
      </div>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        月別災害発生数
      </h2>

      {/* 月別の折れ線グラフ */}
      <div className="bg-white rounded-lg shadow p-4 mx-4 mb-6">
        <ResponsiveContainer width="100%" height={300}>
          <LineChart
            data={summaryStats.byMonth}
            margin={{ top: 5, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="month" />
            <YAxis />
            <Tooltip formatter={(value) => [`${value}件`, "発生数"]} />
            <Legend />
            <Line
              type="monotone"
              dataKey="count"
              stroke="#8884d8"
              name="災害発生数"
            />
          </LineChart>
        </ResponsiveContainer>
      </div>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        影響レベル別分布
      </h2>

      {/* 影響レベル別の棒グラフ */}
      <div className="bg-white rounded-lg shadow p-4 mx-4 mb-6">
        <ResponsiveContainer width="100%" height={300}>
          <BarChart
            data={summaryStats.byImpactLevel}
            margin={{ top: 5, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip formatter={(value) => [`${value}件`, "件数"]} />
            <Legend />
            <Bar dataKey="value" name="災害数" fill="#82ca9d" />
          </BarChart>
        </ResponsiveContainer>
      </div>

      <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
        最近の災害一覧
      </h2>

      {/* 最新の災害リスト */}
      <div className="bg-white rounded-lg shadow mx-4 mb-6 overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                名称
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                種類
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                発生日
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                状態
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {disastersResponse?.data?.disasters?.slice(0, 5).map((disaster) => (
              <tr key={disaster.id}>
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {disaster.name}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {disaster.disaster_type}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {disaster.occurred_at
                    ? new Date(disaster.occurred_at).toLocaleDateString("ja-JP")
                    : "不明"}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <span
                    className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full 
                    ${
                      disaster.status === "対応中"
                        ? "bg-yellow-100 text-yellow-800"
                        : disaster.status === "完了"
                          ? "bg-green-100 text-green-800"
                          : "bg-gray-100 text-gray-800"
                    }`}
                  >
                    {disaster.status || "未設定"}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* 全件表示リンク */}
      <div className="p-4 text-center">
        <Link
          to="/disasters"
          className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#111418] hover:bg-[#333] transition-colors mr-4"
        >
          すべての災害情報を見る
        </Link>
        <Link
          to="/application"
          className="inline-flex items-center justify-center px-6 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-[#197fe5] hover:bg-[#1565c0] transition-colors"
        >
          支援申請を行う
        </Link>
      </div>
    </>
  )
}
