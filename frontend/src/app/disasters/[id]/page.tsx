import Header from "@/components/layout/header/page"
import Footer from "@/components/layout/footer/page"
import Link from "next/link"
import { notFound } from "next/navigation"

// サンプルデータ。IDに対応する詳細情報を含める
const disasterDetailsData: { [key: string]: any } = {
  "1": {
    id: "2023-001",
    name: "Typhoon Lan",
    date: "August 15, 2023",
    type: "Typhoon",
    location: "Kyoto Prefecture",
    affectedArea: "150 hectares",
    estimatedDamage: "500 million yen",
    damageSummary: {
      total: 120,
      farmlands: 80,
      facilities: 40,
    },
    documents: [
      {
        name: "Damage Report 2023-001",
        type: "Report",
        uploaded: "August 16, 2023",
      },
      {
        name: "Application Form 2023-001",
        type: "Form",
        uploaded: "August 17, 2023",
      },
      {
        name: "Assessment Report 2023-001",
        type: "Report",
        uploaded: "August 20, 2023",
      },
    ],
  },
  // 他のIDのデータもここに追加...
}

const getDisasterById = (id: string) => {
  // 実際のIDは "2023-001" のような形式かもしれませんが、
  // URLの "1" とマッピングさせるため、ここでは "1" をキーとします。
  return disasterDetailsData[id]
}

type Props = {
  params: { id: string }
}

export default function DisasterDetailPage({ params }: Props) {
  const disaster = getDisasterById(params.id)

  if (!disaster) {
    notFound()
  }

  return (
    <div className="layout-container flex h-full grow flex-col">
      <Header />
      <main className="px-40 flex flex-1 justify-center py-5">
        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
          {/* Breadcrumbs */}
          <div className="flex flex-wrap gap-2 p-4">
            <Link
              href="/disasters"
              className="text-[#637588] text-base font-medium leading-normal hover:underline"
            >
              Disasters
            </Link>
            <span className="text-[#637588] text-base font-medium leading-normal">
              /
            </span>
            <Link
              href="/disasters"
              className="text-[#637588] text-base font-medium leading-normal hover:underline"
            >
              Disaster List
            </Link>
            <span className="text-[#637588] text-base font-medium leading-normal">
              /
            </span>
            <span className="text-[#111418] text-base font-medium leading-normal">
              Disaster Details
            </span>
          </div>

          {/* Page Header */}
          <div className="flex flex-wrap justify-between gap-3 p-4">
            <div className="flex min-w-72 flex-col gap-3">
              <p className="text-[#111418] tracking-light text-[32px] font-bold leading-tight">
                Disaster #{disaster.id}
              </p>
              <p className="text-[#637588] text-sm font-normal leading-normal">
                {disaster.name}, {disaster.date}
              </p>
            </div>
          </div>

          {/* Tab Navigation */}
          {/* TODO: タブの動的切り替えには 'use client' と useState が必要です */}
          <div className="pb-3">
            <div className="flex border-b border-[#dce0e5] px-4 gap-8">
              <Link
                className="flex flex-col items-center justify-center border-b-[3px] border-b-[#111418] text-[#111418] pb-[13px] pt-4"
                href="#"
              >
                <p className="text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]">
                  Overview
                </p>
              </Link>
              <Link
                className="flex flex-col items-center justify-center border-b-[3px] border-b-transparent text-[#637588] pb-[13px] pt-4"
                href="#"
              >
                <p className="text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                  Damages
                </p>
              </Link>
              <Link
                className="flex flex-col items-center justify-center border-b-[3px] border-b-transparent text-[#637588] pb-[13px] pt-4"
                href="#"
              >
                <p className="text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                  Documents
                </p>
              </Link>
              <Link
                className="flex flex-col items-center justify-center border-b-[3px] border-b-transparent text-[#637588] pb-[13px] pt-4"
                href="#"
              >
                <p className="text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                  Timeline
                </p>
              </Link>
            </div>
          </div>

          {/* Disaster Overview Section */}
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Disaster Overview
          </h2>
          <div className="p-4 grid grid-cols-[20%_1fr] gap-x-6">
            <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Disaster Type
              </p>
              <p className="text-[#111418] text-sm font-normal leading-normal">
                {disaster.type}
              </p>
            </div>
            <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Location
              </p>
              <p className="text-[#111418] text-sm font-normal leading-normal">
                {disaster.location}
              </p>
            </div>
            <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Date
              </p>
              <p className="text-[#111418] text-sm font-normal leading-normal">
                {disaster.date}
              </p>
            </div>
            <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Affected Area
              </p>
              <p className="text-[#111418] text-sm font-normal leading-normal">
                {disaster.affectedArea}
              </p>
            </div>
            <div className="col-span-2 grid grid-cols-subgrid border-t border-t-[#dce0e5] py-5">
              <p className="text-[#637588] text-sm font-normal leading-normal">
                Estimated Damage
              </p>
              <p className="text-[#111418] text-sm font-normal leading-normal">
                {disaster.estimatedDamage}
              </p>
            </div>
          </div>

          {/* Damage Summary Section */}
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Damage Summary
          </h2>
          <div className="flex flex-wrap gap-4 p-4">
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Total Damages
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                {disaster.damageSummary.total}
              </p>
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Damaged Farmlands
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                {disaster.damageSummary.farmlands}
              </p>
            </div>
            <div className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-[#dce0e5]">
              <p className="text-[#111418] text-base font-medium leading-normal">
                Damaged Facilities
              </p>
              <p className="text-[#111418] tracking-light text-2xl font-bold leading-tight">
                {disaster.damageSummary.facilities}
              </p>
            </div>
          </div>

          {/* Related Documents Section */}
          <h2 className="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5">
            Related Documents
          </h2>
          <div className="px-4 py-3 @container">
            <div className="flex overflow-hidden rounded-lg border border-[#dce0e5] bg-white">
              <table className="flex-1">
                <thead>
                  <tr className="bg-white">
                    <th className="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Document Name
                    </th>
                    <th className="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Type
                    </th>
                    <th className="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal">
                      Uploaded Date
                    </th>
                    <th className="table-column-480 px-4 py-3 text-left text-[#111418] w-60 text-[#637588] text-sm font-medium leading-normal">
                      Action
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {disaster.documents.map((doc: any, index: number) => (
                    <tr key={index} className="border-t border-t-[#dce0e5]">
                      <td className="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal">
                        {doc.name}
                      </td>
                      <td className="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                        {doc.type}
                      </td>
                      <td className="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#637588] text-sm font-normal leading-normal">
                        {doc.uploaded}
                      </td>
                      <td className="table-column-480 h-[72px] px-4 py-2 w-60 text-[#637588] text-sm font-bold leading-normal tracking-[0.015em]">
                        <Link
                          href="#"
                          className="text-blue-600 hover:underline"
                        >
                          View
                        </Link>
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
