import Header from "@/components/layout/header/page"
import Footer from "@/components/layout/footer/page"

// サンプルデータ (実際にはAPIなどから取得します)
const disasterData = [
  {
    date: "2024-07-15",
    region: "Kyoto Prefecture",
    summary: "Extensive crop damage due to flooding",
    status: "Pending",
  },
  {
    date: "2024-07-10",
    region: "Osaka Prefecture",
    summary: "Landslide impacting irrigation systems",
    status: "In Progress",
  },
  {
    date: "2024-07-05",
    region: "Hyogo Prefecture",
    summary: "Severe hail damage to orchards",
    status: "Completed",
  },
  {
    date: "2024-06-28",
    region: "Nara Prefecture",
    summary: "Drought affecting rice paddies",
    status: "Pending",
  },
  {
    date: "2024-06-20",
    region: "Shiga Prefecture",
    summary: "Wind damage to greenhouses",
    status: "In Progress",
  },
  {
    date: "2024-06-12",
    region: "Wakayama Prefecture",
    summary: "Flooding of agricultural lands",
    status: "Completed",
  },
  {
    date: "2024-06-05",
    region: "Mie Prefecture",
    summary: "Landslide impacting access roads",
    status: "Pending",
  },
  {
    date: "2024-05-28",
    region: "Aichi Prefecture",
    summary: "Hail damage to vegetable crops",
    status: "In Progress",
  },
  {
    date: "2024-05-20",
    region: "Gifu Prefecture",
    summary: "Drought conditions affecting livestock",
    status: "Completed",
  },
  {
    date: "2024-05-12",
    region: "Shizuoka Prefecture",
    summary: "Wind damage to fruit trees",
    status: "Pending",
  },
]

export default function DisasterInfoPage() {
  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                Disaster Information List
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                View and manage all reported agricultural disaster events. Each
                entry provides details on the event's date, affected area, and a
                summary of damages.
              </p>
            </div>
          </div>
          <div className="px-4 py-3 @container">
            <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
              <table className="flex-1">
                <thead>
                  <tr className="bg-white">
                    <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Date of Occurrence
                    </th>
                    <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Region
                    </th>
                    <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Damage Summary
                    </th>
                    <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal">
                      Status
                    </th>
                    <th className="table-column-600 px-4 py-3 text-left text-[#111418] w-60 text-[#637588] text-sm font-medium leading-normal">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {disasterData.map((item, index) => (
                    <tr key={index} className="border-t border-t-[#dce0e5]">
                      <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                        {item.date}
                      </td>
                      <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                        {item.region}
                      </td>
                      <td className="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                        {item.summary}
                      </td>
                      <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal">
                        <button className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f4] text-[#111418] text-sm font-medium leading-normal w-full">
                          <span className="truncate">{item.status}</span>
                        </button>
                      </td>
                      <td className="table-column-600 h-[72px] px-4 py-2 w-60 text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                        View Details
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
